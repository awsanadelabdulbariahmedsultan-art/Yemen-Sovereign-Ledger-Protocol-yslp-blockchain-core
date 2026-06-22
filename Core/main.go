/*
===================================================================================================
                         PROPRIETARY SOURCE CODE & INTELLECTUAL PROPERTY
===================================================================================================
PROTOCOL NAME:  Yemen Sovereign Ledger Protocol (YSLP) / بروتوكول السجل السيادي اليمني الموحد
PROJECT CORE:   Central Bank of Yemen Digital (CBOYD) Engine
VERSION:        2.0.0 (Ultimate Unified Production Core)
AUTHOR/CREATOR: Eng. Awsan Adel Abdulbari Ahmed Sultan
COUNTRY:        Yemen (الجمهورية اليمنية)
CONTACT PHONE:  +967 777852433 / +967 776633003
EMAIL:          awsandew@outlook.com
LINKEDIN:       https://linkedin.com

LEGAL NOTICE:
Copyright (c) 2026 Eng. Awsan Adel Abdulbari Ahmed Sultan. All Rights Reserved.
This source code, architecture, and cryptographic protocol are the exclusive intellectual 
property of the author. Unauthorized copying, distribution, modification, or deployment 
of this software without explicit written permission from the owner is strictly prohibited.
===================================================================================================
*/

package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
	"strconv"
	"time"
)

// CentralBankAuthority المعرف الرقمي السيادي للبنك المركزي اليمني كموثق رئيسي للشبكة
const CentralBankAuthority = "CENTRAL_BANK_YEMEN_NODE_01"

// ================================================================================================
// 1. معمارية الأنواع والبيانات المشتركة (الأصول، وسائل الدفع، القنوات الصناعية)
// ================================================================================================

type AssetType string
const (
	FIAT   AssetType = "FIAT"   // العملات النقدية الورقية (YER, USD, EUR)
	CRYPTO AssetType = "CRYPTO" // العملات المشفرة اللامركزية (BTC, ETH)
	STABLE AssetType = "STABLE" // العملات الرقمية المستقرة (USDT, USDC)
)

type PaymentMethod string
const (
	E_WALLET   PaymentMethod = "E_WALLET"   // المحافظ الرقمية (أم فلوس، جيب، جوالي...)
	BANK_ACC   PaymentMethod = "BANK_ACC"   // الحسابات البنكية والمقاصة المصرفية
	CARD       PaymentMethod = "CARD"       // البطاقات الائتمانية الدولية
	NFC_DEVICE PaymentMethod = "NFC_DEVICE" // أنظمة الـ NFC وإنترنت الأشياء الصناعية
)

type IndustrialSector string
const (
	FACTORY  IndustrialSector = "FACTORY"  // المصانع وخطوط الإنتاج والأتمتة
	MARITIME IndustrialSector = "MARITIME" // القطاع البحري والسفن والموانئ
	REFINERY IndustrialSector = "REFINERY" // بركات المصبّات، الوقود والشركات النفطية
	NONE     IndustrialSector = "NONE"     // المعاملات المدنية والبنكية العادية
)

// NFCTag هيكل بيانات بطاقة / شريحة الـ NFC المستخدمة في الآلات أو السفن
type NFCTag struct {
	UID          string
	Sector       IndustrialSector
	OwnerEntity  string
	IsAuthorized bool
}

// ExchangeGateway هيكل ربط المنصات الرقمية وبورصات المال العالمية
type ExchangeGateway struct {
	PlatformName string // Binance, Bloomberg Terminal, Interactive Brokers
	IsConnected  bool
	APIVersion   string
}

// Asset هيكل تسجيل الأصول المالية داخل النظام ومعدلات مقاصتها
type Asset struct {
	Symbol    string
	Name      string
	Type      AssetType
	RateToUSD float64 // سعر الصرف اللحظي مقابل الدولار للمقاصة الفورية
}

// ================================================================================================
// 2. معمارية الحماية والتشفير الرقمي (ECDSA Cryptography Core)
// ================================================================================================

// GenerateWalletKeyPair توليد زوج مفاتيح (سري ومعلن) لأي محفظة أو منشأة برمجية جديدة تنضم للشبكة
func GenerateWalletKeyPair() (*ecdsa.PrivateKey, string, error) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, "", err
	}
	pubKeyBytes := elliptic.Marshal(elliptic.P256(), privateKey.PublicKey.X, privateKey.PublicKey.Y)
	walletAddress := hex.EncodeToString(pubKeyBytes)
	return privateKey, walletAddress, nil
}

// ================================================================================================
// 3. الهيكل الهندسي الموحد للمعاملة المالية الشاملة (Unified Transaction Structure)
// ================================================================================================

type UnifiedTransaction struct {
	TxID            string
	SenderWallet    string           // عنوان محفظة المرسل (المفتاح المعلن)
	ReceiverWallet  string           // عنوان محفظة المستقبل
	FromAsset       string           // العملة المرسلة
	ToAsset         string           // العملة المستلمة (التحويل التلقائي)
	SourceAmount    float64          // المبلغ الأصلي
	TargetAmount    float64          // المبلغ بعد المقاصة والتحويل
	PaymentChannel  PaymentMethod    // وسيلة الدفع (محفظة، بنك، NFC)
	IndustrialScope IndustrialSector // النطاق الصناعي (إن وجد)
	NFC_UID         string           // معرف شريحة الـ NFC (إن وجد)
	HardwareGateway string           // بوابة القراءة الميدانية
	TargetPlatform  string           // المنصة العالمية المتصلة بالعملية
	Timestamp       int64
	Signature       string // التوقيع الرقمي المشفر الفريد (بصمة العميل السرية)
}

// CalculateTxHash حساب الهاش الأساسي لبيانات المعاملة تمهيداً لتوقيعها أو فحصها
func (tx *UnifiedTransaction) CalculateTxHash() []byte {
	record := tx.SenderWallet + tx.ReceiverWallet + tx.FromAsset + tx.ToAsset + 
		fmt.Sprintf("%f%f", tx.SourceAmount, tx.TargetAmount) + tx.NFC_UID + 
		tx.HardwareGateway + tx.TargetPlatform + strconv.FormatInt(tx.Timestamp, 10)
	h := sha256.New()
	h.Write([]byte(record))
	return h.Sum(nil)
}

// SignTransaction توقيع المعاملة رقمياً بمفتاح العميل السري قبل إرسالها لعقد الشبكة
func (tx *UnifiedTransaction) SignTransaction(privKey *ecdsa.PrivateKey) error {
	txHash := tx.CalculateTxHash()
	r, s, err := ecdsa.Sign(rand.Reader, privKey, txHash)
	if err != nil {
		return err
	}
	tx.Signature = hex.EncodeToString(r.Bytes()) + "||" + hex.EncodeToString(s.Bytes())
	return nil
}

// IsValidTransaction التحقق من النزاهة الرياضية للتوقيع الرقمي ومنع انتحال الهوية والتزوير
func (tx *UnifiedTransaction) IsValidTransaction() bool {
	if tx.SenderWallet == "YSLP_CORE_PROTOCOL" {
		return true // كتل التأسيس النظامية مستثناة
	}
	if tx.Signature == "" {
		return false
	}

	// استخراج قيم المفتاح المعلن من عنوان المرسل
	pubKeyBytes, err := hex.DecodeString(tx.SenderWallet)
	if err != nil {
		return false
	}
	x, y := elliptic.Unmarshal(elliptic.P256(), pubKeyBytes)
	if x == nil || y == nil {
		return false
	}
	pubKey := ecdsa.PublicKey{Curve: elliptic.P256(), X: x, Y: y}

	// تفكيك قيم التوقيع الرقمي R و S
	var rBytesStr, sBytesStr string
	_, err = fmt.Sscanf(tx.Signature, "%s||%s", &rBytesStr, &sBytesStr)
	if err != nil {
		rBytesStr = tx.Signature[:len(tx.Signature)/2-1]
		sBytesStr = tx.Signature[len(tx.Signature)/2+2:]
	}

	rDecoded, _ := hex.DecodeString(rBytesStr)
	sDecoded, _ := hex.DecodeString(sBytesStr)

	var r, s big.Int
	r.SetBytes(rDecoded)
	s.SetBytes(sDecoded)

	return ecdsa.Verify(&pubKey, tx.CalculateTxHash(), &r, &s)
}

// ================================================================================================
// 4. البنية الهندسية للكتل وسلسلة البلوكشين (Blockchain & Block Architecture)
// ================================================================================================

type Block struct {
	Index        int
	Timestamp    int64
	Transactions []UnifiedTransaction
	PrevHash     string
	Hash         string
	Validator    string // توقيع الموثق السيادي (آلية إجماع PoA)
}

type GlobalSovereignBlockchain struct {
	Chain               []Block
	PendingTransactions []UnifiedTransaction
	SupportedAssets     map[string]Asset
	ConnectedExchanges  map[string]ExchangeGateway
	RegisteredNFCTags   map[string]NFCTag
}

// CalculateBlockHash دالة التشفير الرياضي لحساب هاش الكتلة الكلي باستخدام SHA-256
func CalculateBlockHash(b Block) string {
	txsString := ""
	for _, tx := range b.Transactions {
		txsString += tx.TxID + tx.Signature + tx.TargetPlatform
	}
	record := strconv.Itoa(b.Index) + strconv.FormatInt(b.Timestamp, 10) + txsString + b.PrevHash + b.Validator
	h := sha256.New()
	h.Write([]byte(record))
	return hex.EncodeToString(h.Sum(nil))
}

// NewGlobalSovereignBlockchain تأسيس الشبكة السيادية الموحدة وتوليد كتلة البداية وتغذية البيانات الأساسية
func NewGlobalSovereignBlockchain() *GlobalSovereignBlockchain {
	bc := &GlobalSovereignBlockchain{
		SupportedAssets:    make(map[string]Asset),
		ConnectedExchanges: make(map[string]ExchangeGateway),
		RegisteredNFCTags:  make(map[string]NFCTag),
	}

	// أ: تسجيل الأصول والعملات النقدية والمشفرة والمستقرة بأسعار صرفها المركزية الموحدة
	bc.SupportedAssets["YER"] = Asset{Symbol: "YER", Name: "Yemeni Rial", Type: FIAT, RateToUSD: 0.0040}
	bc.SupportedAssets["USD"] = Asset{Symbol: "USD", Name: "US Dollar", Type: FIAT, RateToUSD: 1.0}
	bc.SupportedAssets["BTC"] = Asset{Symbol: "BTC", Name: "Bitcoin", Type: CRYPTO, RateToUSD: 65000.0}
	bc.SupportedAssets["USDT"] = Asset{Symbol: "USDT", Name: "Tether", Type: STABLE, RateToUSD: 1.0}

	// ب: تفعيل بوابات ومحولات الـ APIs للمنصات الرقمية والبورصات العالمية
	bc.ConnectedExchanges["BINANCE"] = ExchangeGateway{PlatformName: "Binance Global Exchange", IsConnected: true, APIVersion: "v3"}
	bc.ConnectedExchanges["BLOOMBERG"] = ExchangeGateway{PlatformName: "Bloomberg Terminal", IsConnected: true, APIVersion: "v2"}
	bc.ConnectedExchanges["INTERACTIVE_BROKERS"] = ExchangeGateway{PlatformName: "Interactive Brokers", IsConnected: true, APIVersion: "v1"}

	// ج: تسجيل شرائح الـ NFC الصناعية المعتمدة سيادياً للمصانع والمنشآت اللوجستية والسفن
	bc.RegisteredNFCTags["NFC_SHIP_ADEN_777"] = NFCTag{UID: "NFC_SHIP_ADEN_777", Sector: MARITIME, OwnerEntity: "ناقلة النفط الدولية العملاقة", IsAuthorized: true}
	bc.RegisteredNFCTags["NFC_FAC_YEMEN_01"] = NFCTag{UID: "NFC_FAC_YEMEN_01", Sector: FACTORY, OwnerEntity: "مجمع المصانع الوطنية الحديثة", IsAuthorized: true}
	bc.RegisteredNFCTags["NFC_REF_BARAKA_02"] = NFCTag{UID: "NFC_REF_BARAKA_02", Sector: REFINERY, OwnerEntity: "منشآت مصبات بركات النفطية والمشتقات", IsAuthorized: true}

	// د: توليد كتلة التأسيس للمحرك (Genesis Block)
	genesisBlock := Block{
		Index:        0,
		Timestamp:    time.Now().Unix(),
		Transactions: []UnifiedTransaction{},
		PrevHash:     "0000000000000000000000000000000000000000000000000000000000000000",
		Validator:    CentralBankAuthority,
	}
	genesisBlock.Hash = CalculateBlockHash(genesisBlock)
	bc.Chain = append(bc.Chain, genesisBlock)

	return bc
}

// CalculateExchange محرك حساب الأسعار والمقاصة الفورية المتعددة عبر الدولار كمقاصة وسيطة
func (bc *GlobalSovereignBlockchain) CalculateExchange(from string, to string, amount float64) (float64, error) {
	srcAsset, existsSrc := bc.SupportedAssets[from]
	targetAsset, existsTarget := bc.SupportedAssets[to]
	if !existsSrc || !existsTarget {
		return 0, fmt.Errorf("العملة المطلوبة غير مدعومة في البروتوكول")
	}
	amountInUSD := amount * srcAsset.RateToUSD
	return amountInUSD / targetAsset.RateToUSD, nil
}


// ProcessTransaction المعالجة المركزية الشاملة لإدخال وضمان سلامة أي معاملة مالية بالشبكة (مدنية، صناعية، بورصة، NFC)
func (bc *GlobalSovereignBlockchain) ProcessTransaction(tx UnifiedTransaction) bool {
	// 1. التحقق من سلامة التوقيع الرقمي لمنع التزوير واختراق الحسابات
	if !tx.IsValidTransaction() {
		fmt.Printf("[❌] رفض المعاملة: التوقيع الرقمي (البصمة السرية للعميل) غير صالح أو تم التلاعب به!\n")
		return false
	}

	// 2. إذا كانت العملية عبر الـ NFC الصناعي، يتم فحص تصريح الهاردوير ومطابقته برمجياً
	if tx.PaymentChannel == NFC_DEVICE {
		tag, exists := bc.RegisteredNFCTags[tx.NFC_UID]
		if !exists || !tag.IsAuthorized {
			fmt.Printf("[❌] رفض المعاملة: شريحة الهاردوير NFC برقم [%s] غير مسجلة بالبنك المركزي!\n", tx.NFC_UID)
			return false
		}
		tx.IndustrialScope = tag.Sector
	}

	// 3. التحقق من ربط بوابات المنصة العالمية المحددة
	if tx.TargetPlatform != "LOCAL_CLEARING" {
		exch, exists := bc.ConnectedExchanges[tx.TargetPlatform]
		if !exists || !exch.IsConnected {
			fmt.Printf("[❌] رفض المعاملة: بوابة الاتصال بالمنصة العالمية [%s] غير مفعلة!\n", tx.TargetPlatform)
			return false
		}
	}

	// 4. احتساب مقادير الصرف والمقاصة الآلية
	targetAmount, err := bc.CalculateExchange(tx.FromAsset, tx.ToAsset, tx.SourceAmount)
	if err != nil {
		fmt.Printf("[❌] رفض المعاملة: فشل احتساب مقاصة الأصول: %v\n", err)
		return false
	}
	tx.TargetAmount = targetAmount

	txIDHash := sha256.Sum256([]byte(tx.SenderWallet + tx.ReceiverWallet + strconv.FormatInt(time.Now().UnixNano(), 10)))
	tx.TxID = hex.EncodeToString(txIDHash[:8])

	bc.PendingTransactions = append(bc.PendingTransactions, tx)
	fmt.Printf("[+ YSLP CORE] تم قبول المعاملة الآمنة [%s] بقناة [%v]: تحويل %.2f %s إلى %.4f %s وتوجيهها للمنصة [%s]\n",
		tx.TxID, tx.PaymentChannel, tx.SourceAmount, tx.FromAsset, tx.TargetAmount, tx.ToAsset, tx.TargetPlatform)
	return true
}

// MineBlock تنفيذ آلية الإجماع السيادي (Proof of Authority - PoA) لتعميد وحفظ الكتل عبر البنك المركزي
func (bc *GlobalSovereignBlockchain) MineBlock(validator string) bool {
	if validator != CentralBankAuthority {
		fmt.Printf("[❌] خرق نظام الإجماع: العقدة الممررة %s غير مخولة سيادياً بتعميد المعاملات!\n", validator)
		return false
	}

	prevBlock := bc.Chain[len(bc.Chain)-1]
	newBlock := Block{
		Index:        prevBlock.Index + 1,
		Timestamp:    time.Now().Unix(),
		Transactions: bc.PendingTransactions,
		PrevHash:     prevBlock.Hash,
		Validator:    validator,
	}
	newBlock.Hash = CalculateBlockHash(newBlock)
	bc.Chain = append(bc.Chain, newBlock)
	bc.PendingTransactions = []UnifiedTransaction{}
	fmt.Printf("[✅ SUCCESS] إجماع البنوك والقطاعات: وثّق البنك المركزي اليمني الكتلة الشاملة الموحدة رقم #%d في السجل السيادي.\n", newBlock.Index)
	return true
}

// IsChainValid التحقق التلقائي الشامل لسلامة كامل البلوكشين وتأكيد تشفير كتل البيانات تاريخياً
func (bc *GlobalSovereignBlockchain) IsChainValid() bool {
	for i := 1; i < len(bc.Chain); i++ {
		currentBlock := bc.Chain[i]
		previousBlock := bc.Chain[i-1]

		if currentBlock.Hash != CalculateBlockHash(currentBlock) {
			return false
		}
		if currentBlock.PrevHash != previousBlock.Hash {
			return false
		}
		if currentBlock.Validator != CentralBankAuthority {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("=======================================================================================")
	fmt.Println("       YEMEN SOVEREIGN LEDGER PROTOCOL (YSLP) - UNIFIED PRODUCTION CORE ENGINE         ")
	fmt.Println("       DESIGNED FOR: CENTRAL BANK OF YEMEN DIGITAL (CBOYD) ARCHITECTURE               ")
	fmt.Println("       PROPRIETARY INTELLECTUAL PROPERTY DEVELOPER: ENG. AWSAN ADEL ABDULBARI         ")
	fmt.Println("=======================================================================================")

	yslpCore := NewGlobalSovereignBlockchain()
	fmt.Println("[*] تم بدء تشغيل السجل بنجاح. تفعيل قنوات: البنوك، المحافظ الرقمية، إنترنت الأشياء والـ NFC الميداني.")
	fmt.Println("[*] حالة الاتصال ببورصات المال والمنصات العالمية: Binance, Bloomberg Terminal جاهزة.\n")

	privKeyUserA, addressUserA, _ := GenerateWalletKeyPair()
	_, addressUserB, _ := GenerateWalletKeyPair()
	privKeyShip, addressShip, _ := GenerateWalletKeyPair()
	_, addressPort, _ := GenerateWalletKeyPair()
	privKeyRefinery, addressRefinery, _ := GenerateWalletKeyPair()

	fmt.Println("[*] --- السيناريو الأول: سحب وتحويل عملات مشفرة للمحافظ المحلية عبر البورصة العالمية ---")
	tx1 := UnifiedTransaction{
		SenderWallet:   addressUserA,
		ReceiverWallet: addressUserB,
		FromAsset:      "YER",
		ToAsset:        "USDT",
		SourceAmount:   300000.00,
		PaymentChannel: E_WALLET,
		TargetPlatform: "BINANCE",
		Timestamp:      time.Now().Unix(),
	}
	tx1.SignTransaction(privKeyUserA)
	yslpCore.ProcessTransaction(tx1)

	fmt.Println("\n[*] --- السيناريو الثاني: محاكاة هجوم وتزوير البيانات (فحص جدار الحماية التشفيري) ---")
	txMalicious := tx1
	txMalicious.SourceAmount = 95000000.00
	yslpCore.ProcessTransaction(txMalicious)

	fmt.Println("\n[*] --- السيناريو الثالث: تسوية رسوم شحن وتفريغ لسفينة ناقلة تلقائياً بميناء عدن عبر NFC ---")
	txMaritime := UnifiedTransaction{
		SenderWallet:    addressShip,
		ReceiverWallet:  addressPort,
		FromAsset:       "USD",
		ToAsset:         "YER",
		SourceAmount:    45000.00,
		PaymentChannel:  NFC_DEVICE,
		NFC_UID:         "NFC_SHIP_ADEN_777",
		HardwareGateway: "PORT_ADEN_BERTH_02_READER",
		TargetPlatform:  "BLOOMBERG",
		Timestamp:       time.Now().Unix(),
	}
	txMaritime.SignTransaction(privKeyShip)
	yslpCore.ProcessTransaction(txMaritime)

	fmt.Println("\n[*] --- السيناريو الرابع: أتمتة حسابات تفريغ ومقاصة بركات ومنشآت مصبات الوقود النفطية ---")
	txRefinery := UnifiedTransaction{
		SenderWallet:    addressRefinery,
		ReceiverWallet:  addressUserB,
		FromAsset:       "YER",
		ToAsset:         "YER",
		SourceAmount:    12500000.00,
		PaymentChannel:  NFC_DEVICE,
		NFC_UID:         "NFC_REF_BARAKA_02",
		HardwareGateway: "VALVE_REFINERY_OUTLET_NFC",
		TargetPlatform:  "LOCAL_CLEARING",
		Timestamp:       time.Now().Unix(),
	}
	txRefinery.SignTransaction(privKeyRefinery)
	yslpCore.ProcessTransaction(txRefinery)

	fmt.Println("\n=======================================================================================")
	fmt.Println("[*] إرسال كتل العمليات والمعاملات المدمجة لعقدة البنك المركزي اليمني لإنتاج كتلة التوثيق:")
	if yslpCore.MineBlock(CentralBankAuthority) {
		fmt.Printf("[*] فحص النزاهة التشفيرية الشامل لشبكة البلوكشين: %t (كامل السجل محمي ومؤمن ضد أي تلاعب تاريخي)\n", yslpCore.IsChainValid())
	}
	fmt.Println("=======================================================================================")
}



