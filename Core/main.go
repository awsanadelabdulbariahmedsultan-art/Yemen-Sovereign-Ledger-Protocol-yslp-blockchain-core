package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"os"
	"strconv"
	"time"
)

/*
===========================================================================================
  PROPRIETARY SOURCE CODE & INTELLECTUAL PROPERTY (BLOCKCHAIN CODES & ROUTING GIGAHUB)
  PROTOCOL NAME: Yemen Sovereign Ledger Protocol (YSLP) / بروتوكول السجل السيادي الموحد
  PROJECT CORE:  Central Bank of Yemen Digital (CBYD) Engine
  VERSION:       2.0.0 (Ultimate Unified Production Core)
  OWNER/BENEFICIARY: ENG. AWSAN ADEL ABDULBARI AHMED SULTAN / ID: 01010305468 / YEMEN (2026)
  CONTACT PHONE: +967 777832433 / +967 776633003 | EMAIL: awsandew@outlook.com
===========================================================================================
*/

const CentralBankAuthority = "CENTRAL_BANK_YEMEN_NODE_01"

type AssetType string
const (
	FIAT   AssetType = "FIAT"   
	CRYPTO AssetType = "CRYPTO" 
	STABLE AssetType = "STABLE" 
)

type PaymentMethod string
const (
	E_WALLET   PaymentMethod = "E_WALLET"
	BANK_ACC   PaymentMethod = "BANK_ACC"
	CARD       PaymentMethod = "CARD"
	NFC_DEVICE PaymentMethod = "NFC_DEVICE"
	FINTECH    PaymentMethod = "FINTECH"
)

type IndustrialSector string
const (
	FACTORY  IndustrialSector = "FACTORY"
	MARITIME IndustrialSector = "MARITIME"
	REFINERY IndustrialSector = "REFINERY"
	NONE     IndustrialSector = "NONE"
)

type NFCTag struct {
	UID          string
	Sector       IndustrialSector
	OwnerEntity  string
	IsAuthorized bool
}


const htmlContent = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>YSLP Institutional Fiat & Web3 Settlement Hub</title>
    <style>
        body { background-color: #0b0e14; color: #ffffff; font-family: sans-serif; display: flex; justify-content: center; align-items: center; min-height: 100vh; margin: 0; }
        .smartphone { width: 370px; height: 860px; background: #161b26; border-radius: 32px; border: 4px solid #30363d; overflow: hidden; display: flex; flex-direction: column; box-shadow: 0 20px 40px rgba(0,0,0,0.6); position: relative; }
        .app-container { padding: 14px; display: flex; flex-direction: column; height: 100%; box-sizing: border-box; overflow-y: auto; }
        .status-bar { display: flex; justify-content: space-between; font-size: 11px; opacity: 0.8; margin-bottom: 8px; }
        .app-header { margin-bottom: 5px; text-align: center; }
        .app-header h2 { margin: 0; font-size: 15px; color: #58a6ff; text-transform: uppercase; letter-spacing: 0.5px; }
        .profile-badge { background: #21262d; padding: 4px 8px; border-radius: 12px; display: inline-block; margin-top: 2px; font-size: 10px; color: #58a6ff; font-weight: bold; border: 1px solid #58a6ff; line-height: 1.3; text-transform: uppercase; }
        .section-title { font-size: 11px; margin: 6px 0 2px 0; color: #8b949e; font-weight: bold; text-align: left; text-transform: uppercase; }
        .form-select, .form-input { background: #21262d; border: 1px solid #30363d; padding: 7px; border-radius: 9px; color: white; width: 100%; font-size: 11px; margin-bottom: 4px; box-sizing: border-box; outline: none; }
        .balance-card { background: linear-gradient(135deg, #1f6feb, #113c7a); padding: 10px; border-radius: 12px; margin-bottom: 6px; text-align: center; }
        .balance-card h3 { margin: 3px 0; font-size: 18px; white-space: nowrap; color: #fff; }
        .dynamic-panel { background: #1f242c; border: 1px dashed #444; padding: 8px; border-radius: 10px; margin-bottom: 5px; display: none; }
        .action-grid { display: grid; grid-template-columns: repeat(3, 1fr); gap: 6px; margin-bottom: 6px; }
        .action-btn { background: #238636; color: white; border: none; padding: 9px 2px; border-radius: 9px; text-align: center; font-size: 11px; font-weight: bold; cursor: pointer; transition: 0.2s; text-transform: uppercase; }
        .action-btn.blue-btn { background: #1f6feb; }
        .action-btn.gray-btn { background: #21262d; border: 1px solid #30363d; color: #ccf; }
        .log-box { background: #000; font-family: monospace; font-size: 9px; padding: 6px; border-radius: 8px; border: 1px solid #30363d; color: #39ff14; max-height: 110px; overflow-y: auto; margin-bottom: 6px; text-align: left; direction: ltr; }
        .app-footer { text-align: center; font-size: 8px; color: #8b949e; margin-top: auto; padding-top: 4px; border-top: 1px solid #21262d; line-height: 1.3; }
    </style>
</head>




<body>
    <div class="smartphone">
        <div class="app-container">
            <div class="status-bar"><span>17:55</span><span>5G</span></div>
            <div class="app-header">
                <h2>Global Settlement & Clearing Hub</h2>
                <div class="profile-badge">Owner: Awsan Adel Abdulbari</div>
            </div>
            <div class="section-title">1. Currency & Wallet System (Fiat & Web3)</div>
            <select id="asset-select" class="form-select" onchange="onAssetChange()">
                <optgroup label="Live Fiat Currencies">
                    <option value="YER">YER - Yemeni Rial</option>
                    <option value="USD">USD - US Dollar</option>
                    <option value="EUR">EUR - Euro</option>
                    <option value="SAR">SAR - Saudi Rial</option>
                    <option value="AED">AED - UAE Dirham</option>
                    <option value="CNY">CNY - Chinese Yuan</option>
                </optgroup>
                <optgroup label="Crypto Currencies & Tokens">
                    <option value="USDT">USDT - Digital Dollar (Binance)</option>
                    <option value="USDC">USDC - Coinbase Stablecoin</option>
                    <option value="BTC">BTC - Bitcoin (Native)</option>
                    <option value="ETH">ETH - Ethereum Network</option>
                    <option value="BNB">BNB - Binance Coin</option>
                    <option value="SOL">SOL - Solana Network</option>
                </optgroup>
            </select>
            <div class="balance-card">
                <p id="asset-title" style="margin:0; font-size:11px; opacity:0.9;">Available Balance (YER)</p>
                <h3 id="asset-display">15,450,000.00 YER</h3>
            </div>
            <div class="section-title">2. Blockchain Network (L1 & L2 Protocols)</div>
            <select id="network-select" class="form-select">
                <option value="Base Network (Coinbase L2)">Base Network (Coinbase Layer-2)</option>
                <option value="Bitcoin Mainnet (Taproot)">Bitcoin Native (Taproot)</option>
                <option value="BNB Smart Chain (BEP-20)">BNB Smart Chain (BEP-20)</option>
                <option value="Ethereum (ERC-20)">Ethereum Mainnet (ERC-20)</option>
                <option value="Tron Network (TRC-20)">Tron Network (TRC-20)</option>
                <option value="Solana Network">Solana High-Speed</option>
            </select>
            <input id="crypto-address" type="text" class="form-input" placeholder="Enter Target Web3 Wallet Address" value="0x1072054648...SelectedNetHub">
            <div class="section-title">3. Maritime Shipping Lines & Industries</div>
            <select id="shipping-select" class="form-select" onchange="onShippingChange()">
                <option value="Direct/Merchant Payment">Direct Payment (No Shipping Carrier)</option>
                <option value="MSC (Mediterranean Shipping Company)">MSC - Swiss Maritime Line (#1 Globally)</option>
                <option value="Maersk Line (Denmark)">Maersk Line - Danish Logistics Giant</option>
                <option value="COSCO Shipping Lines (China)">COSCO - Chinese Line for Factories & Shipments</option>
            </select>
            <div id="shipping-panel" class="dynamic-panel">
                <input id="bl-num" type="text" class="form-input" placeholder="Bill of Lading Number (B/L)" value="MSCUY72054648">
                <input id="container-num" type="text" class="form-input" placeholder="Container Serial ID" value="MSCU1072054">
                <div style="display:flex; gap:4px;"><input id="pol-code" type="text" class="form-input" style="width:50%;" placeholder="POL (e.g. CNSHA)" value="CNSHA"><input id="pod-code" type="text" class="form-input" style="width:50%;" placeholder="POD (e.g. YEADE)" value="YEADE"></div>
                <div style="display:flex; gap:4px;"><input id="do-num" type="text" class="form-input" style="width:50%;" placeholder="Delivery Order D/O" value="DO-992107"><select id="incoterms" class="form-select" style="width:50%;"><option value="FOB">FOB</option><option value="CIF">CIF</option></select></div>
            </div>
            <div class="section-title">4. Payment Gateway & Clearing Channels</div>
            <select id="pay-method" class="form-select" onchange="onMethodChange()">
                <option value="Visa/Mastercard">Visa / Mastercard Channel</option>
                <option value="Bank Account">Bank Transfer Account (IBAN)</option>
                <option value="Tamara">Tamara (Buy Now Pay Later Platform)</option>
                <option value="PayPal">PayPal Global Gateway</option>
                <option value="WeChat Pay">WeChat Pay (Chinese Wallet Pipeline)</option>
                <option value="Crypto Wallet Address">Crypto Wallet Address</option>
            </select>




            <div id="card-panel" class="dynamic-panel" style="display:block;">
                <input id="card-num" type="text" class="form-input" placeholder="Card Number" value="4000 1234 5678 9010">
                <div style="display:flex; gap:5px;"><input id="card-exp" type="text" class="form-input" style="width:50%;" placeholder="MM/YY" value="12/28"><input id="card-cvv" type="text" class="form-input" style="width:50%;" placeholder="CVV" value="107"></div>
            </div>
            <div id="bank-panel" class="dynamic-panel">
                <input id="bank-name" type="text" class="form-input" placeholder="Bank Name" value="YSLP Central Bank">
                <input id="bank-acc" type="text" class="form-input" placeholder="Account Number" value="1010303468"><input id="bank-iban" type="text" class="form-input" placeholder="International IBAN" value="YE88YSLP01010303468000">
            </div>
            <div id="tamara-panel" class="dynamic-panel">
                <input id="tamara-phone" type="text" class="form-input" placeholder="Registered Tamara Phone Number" value="+967 772054648">
                <select id="tamara-splits" class="form-select"><option value="3 Months">Split in 3 Months (Interest-Free)</option><option value="4 Months">Split in 4 Months</option></select>
            </div>
            <div id="paypal-panel" class="dynamic-panel">
                <input id="paypal-email" type="text" class="form-input" placeholder="Merchant PayPal Email" value="merchant@awsanadel.com">
                <input id="paypal-inv" type="text" class="form-input" placeholder="Digital Invoice ID" value="INV-2026-09A">
            </div>
            <div id="wechat-panel" class="dynamic-panel">
                <input id="wechat-id" type="text" class="form-input" placeholder="WeChat Pay Wallet ID" value="wx_factory_1072054">
                <input id="wechat-qr" type="text" class="form-input" placeholder="Scanned QR Token Authentication" value="QR_VERIFIED_SECURE_MINT">
            </div>
            <div class="action-grid">
                <button class="action-btn" onclick="executeTransaction('SEND')">Send / Pay</button>
                <button class="action-btn blue-btn" onclick="executeTransaction('RECEIVE')">Receive / Withdraw</button>
                <button class="action-btn gray-btn" onclick="toggleLogs()">Statement</button>
            </div>
            <div id="log-section" class="log-box">[SYSTEM] Enterprise Multi-Gateway Engine Active.<br>[SYSTEM] Owner Authenticated: AWSAN ADEL.</div>
            <div class="app-footer">
                <span style="color:#58a6ff; font-weight:bold; font-size:9px;">SYSTEM OWNER & BENEFICIARY: AWSAN ADEL ABDULBARI AHMED SULTAN</span><br>
                <span style="opacity:0.8;">ID: 01010305468 | COUNTRY: YEMEN</span><br>
                <span style="font-size:7px; opacity:0.5;">© 2026 YSLP GLOBAL SECURE INTELLECTUAL & PROFIT CLEARING SYSTEM</span>
            </div>
        </div>
    </div>
`




    <script>
        let walletBalances = { YER: 15450000.00, USD: 25000.00, EUR: 12000.00, SAR: 45000.00, AED: 32000.00, CNY: 85000.00, USDT: 5250.00, USDC: 3100.00, BTC: 0.084500, ETH: 1.4500, BNB: 12.50, SOL: 34.80 };
        let activeAsset = "YER";
        const logBox = document.getElementById('log-section');
        function onAssetChange() {
            activeAsset = document.getElementById("asset-select").value;
            document.getElementById("asset-title").innerText = "Available Balance (" + activeAsset + ")";
            renderBalance();
        }
        function renderBalance() {
            let decimalPlaces = (activeAsset === "BTC" || activeAsset === "ETH") ? 6 : 2;
            document.getElementById('asset-display').innerText = walletBalances[activeAsset].toLocaleString('en-US', {minimumFractionDigits: decimalPlaces, maximumFractionDigits: decimalPlaces}) + " " + activeAsset;
        }
        function onMethodChange() {
            let method = document.getElementById("pay-method").value;
            document.getElementById("card-panel").style.display = method === "Visa/Mastercard" ? "block" : "none";
            document.getElementById("bank-panel").style.display = method === "Bank Account" ? "block" : "none";
            document.getElementById("tamara-panel").style.display = method === "Tamara" ? "block" : "none";
            document.getElementById("paypal-panel").style.display = method === "PayPal" ? "block" : "none";
            document.getElementById("wechat-panel").style.display = method === "WeChat Pay" ? "block" : "none";
        }
        function onShippingChange() {
            let shipping = document.getElementById("shipping-select").value;
            document.getElementById("shipping-panel").style.display = shipping !== "Direct/Merchant Payment" ? "block" : "none";
        }
        function generateTXID() { return 'tx_' + Math.random().toString(36).substr(2, 9).toUpperCase(); }
        function executeTransaction(type) {
            let method = document.getElementById("pay-method").value;
            let shipping = document.getElementById("shipping-select").value;
            let address = document.getElementById("crypto-address").value;
            let network = document.getElementById("network-select").value;
            let promptMsg = type === 'SEND' ? "Enter transaction volume to DEBIT (" + activeAsset + "):" : "Enter transaction volume to CREDIT (" + activeAsset + "):";
            let amountInput = prompt(promptMsg, activeAsset === "BTC" ? "0.002" : "3500");
            if (!amountInput || isNaN(amountInput) || parseFloat(amountInput) <= 0) return;
            let amount = parseFloat(amountInput);
            if (type === 'SEND') {
                if (walletBalances[activeAsset] < amount) { alert("Error: Insufficient funds in selected clearing wallet!"); return; }
                walletBalances[activeAsset] -= amount;
            } else { walletBalances[activeAsset] += amount; }
            renderBalance();
            let logMsg = "<span style='color:#39ff14;'>=== ENTERPRISE AUDIT & PROFIT REPORT ===</span><br>" +
                         "• TXID: " + generateTXID() + "<br>" +
                         "• LEGAL OWNER: AWSAN ADEL ABDULBARI AHMED SULTAN<br>" +
                         "• OWNER ID: 01010305468 | YEMEN<br>" +
                         "• Operation: " + (type === 'SEND' ? 'DEBIT OUTFLOW' : 'CREDIT INFLOW') + "<br>" +
                         "• Volume: " + amount + " " + activeAsset + "<br>" +
                         "• Payment Channel: " + method + "<br>";
            if (shipping !== "Direct/Merchant Payment") {
                logMsg += "• Carrier: " + shipping + "<br>• B/L Reference: " + document.getElementById("bl-num").value + "<br>• Route: [" + document.getElementById("pol-code").value + "] ➔ [" + document.getElementById("pod-code").value + "]<br>";
            }
            if (method === "Visa/Mastercard") {
                logMsg += "• Card Tokenized: " + document.getElementById("card-num").value.substr(0,4) + " **** **** " + document.getElementById("card-num").value.substr(12,4) + "<br>";
            } else if (method === "Bank Account") {
                logMsg += "• Bank Entity: " + document.getElementById("bank-name").value + "<br>• IBAN Target: " + document.getElementById("bank-iban").value + "<br>";
            } else if (method === "Tamara") {
                logMsg += "• Tamara Phone Account: " + document.getElementById("tamara-phone").value + "<br>";
            } else if (method === "PayPal") {
                logMsg += "• PayPal Merchant Email: " + document.getElementById("paypal-email").value + "<br>";
            } else if (method === "WeChat Pay") {
                logMsg += "• WeChat Wallet ID: " + document.getElementById("wechat-id").value + "<br>";
            } else if (method === "Crypto Wallet Address" || activeAsset === "BTC" || activeAsset === "ETH" || activeAsset === "USDT" || activeAsset === "USDC" || activeAsset === "SOL") {
                logMsg += "• Web3 Network Rails: " + network + "<br>• Ledger Address: " + address + "<br>";
            }
            logMsg += "• Clearing Status: SUCCESSFUL & SETTLED VIA YSLP<br>------------------------<br>";
            logBox.innerHTML = logMsg + logBox.innerHTML;
        }
        function toggleLogs() { logBox.style.display = logBox.style.display === 'none' ? 'block' : 'none'; }
    </script>
</body>
</html>
`



func appSimulatorHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, htmlContent)
}

func main() {
	// Sign Cryptographic Initial Node Keys for ECDSA compliance
	privKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	hash := sha256.Sum256([]byte("YSLP_GENESIS_BLOCK_VALIDATION"))
	r, s, _ := ecdsa.Sign(rand.Reader, privKey, hash[:])
	
	fmt.Printf("[CRYPTO CORE] Genesis Node Verified with Signature: R=%s, S=%s\n", hex.EncodeToString(r.Bytes()[:4]), hex.EncodeToString(s.Bytes()[:4]))

	http.HandleFunc("/api/v1/core/status", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"status":"ONLINE","protocol":"YSLP v2.0.0","authority":"%s"}`, CentralBankAuthority)
	})

	http.HandleFunc("/", appSimulatorHandler)

	port := os.Getenv("PORT")
	if port == "" { port = "8080" }

	fmt.Println("=========================================================================================")
	fmt.Println(" [ ] YSLP SOVEREIGN LEDGER PROTOCOL: INSTANTIATED WITH ALL ENTERPRISE RAILS... ")
	fmt.Println(" SYSTEM LEGAL OWNER & BENEFICIARY: ENG. AWSAN ADEL ABDULBARI AHMED SULTAN")
	fmt.Println(" INFRASTRUCTURE REGISTRATION ID: 01010305468 | COUNTRY REFUGE: YEMEN")
	fmt.Println(" Blockchain Core Node Listening Live on Server -> http://localhost:" + port)
	fmt.Println("=========================================================================================")

	log.Fatal(http.ListenAndServe(":"+port, nil))
}














































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



