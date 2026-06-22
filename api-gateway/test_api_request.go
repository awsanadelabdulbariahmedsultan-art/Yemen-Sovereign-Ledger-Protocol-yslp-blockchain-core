/*
===================================================================================================
                         PROPRIETARY SOURCE CODE & INTELLECTUAL PROPERTY
===================================================================================================
MODULE NAME:    CBOYD API Gateway - Test API Request Script / سكريبت محاكاة واختبار البوابة
VERSION:        1.0.0 (Automated QA & Test Client)
AUTHOR/CREATOR: Eng. Awsan Adel Abdulbari Ahmed Sultan
COUNTRY:        Yemen (الجمهورية اليمنية)
CONTACT PHONE:  +967 777852433 / +967 776633003
EMAIL:          awsandew@outlook.com
LINKEDIN:       https://linkedin.com

LEGAL NOTICE:
Copyright (c) 2026 Eng. Awsan Adel Abdulbari Ahmed Sultan. All Rights Reserved.
===================================================================================================
*/

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// WalletTransferPayload هيكل البيانات المطابق تماماً للبوابة لغرض الإرسال
type WalletTransferPayload struct {
	WalletType    string  `json:"wallet_type"`
	SenderPhone   string  `json:"sender_phone"`
	ReceiverPhone string  `json:"receiver_phone"`
	Amount        float64 `json:"amount"`
	AssetSymbol   string  `json:"asset_symbol"`
}

func main() {
	fmt.Println("=====================================================================")
	fmt.Println("    CBOYD API GATEWAY - AUTOMATED WALLET TRANSACTION SIMULATOR       ")
	fmt.Println("    DEVELOPER: ENG. AWSAN ADEL ABDULBARI AHMED SULTAN                ")
	fmt.Println("=====================================================================")
	fmt.Println("[*] جاري تجهيز حزمة البيانات واختبار الربط الميداني لبوابات المحافظ...\n")

	// 1. رابط البوابة البرمجية الموحدة للمشروع (API Gateway Endpoint)
	apiURL := "http://localhost:8080/api/v1/wallet/transfer"

	// 2. محاكاة بيانات عملية دفع حقيقية (مثال: تحويل من محفظة أم فلوس إلى تاجر بالريال الرقمي)
	payload := WalletTransferPayload{
		WalletType:    "M-Floos (Kurimi)",
		SenderPhone:   "777785243",
		ReceiverPhone: "776633003",
		Amount:        125000.00, // مائة وخمسة وعشرون ألف ريال يمني
		AssetSymbol:   "YER",
	}

	// 3. تحويل البيانات برمجياً إلى صيغة JSON المشفرة
	jsonData, err := json.Marshal(payload)
	if err != nil {
		fmt.Printf("[❌] خطأ في معالجة وتحويل البيانات: %v\n", err)
		return
	}

	// 4. إنشاء وإرسال طلب الـ HTTP POST الفوري إلى خادم البوابة
	startTime := time.Now()
	resp, err := http.Post(apiURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("[❌] فشل الاتصال بالبوابة البرمجية! تأكد من أن ملف (main.go) الخاص بالـ api-gateway يعمل حالياً.")
		return
	}
	defer resp.Body.Close()

	// 5. قراءة الرد المستلم من البوابة البرمجية للمشروع
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("[❌] خطأ في قراءة حزمة الرد المستلمة: %v\n", err)
		return
	}

	duration := time.Since(startTime)

	// 6. تحليل وعرض النتيجة النهائية في سطر الأوامر لتأكيد الربط
	fmt.Println("[✅] تم استلام الرد من خادم البوابة البرمجية الموحدة:")
	fmt.Printf("[⏱️] زمن استجابة الميكروسيرفيس: %v\n", duration)
	fmt.Println("[📦] محتوى الرد الرقمي (JSON Response):")
	
	// تجميل طباعة الرد
	var prettyJSON bytes.Buffer
	errorPretty := json.Indent(&prettyJSON, body, "", "    ")
	if errorPretty == nil {
		fmt.Println(prettyJSON.String())
	} else {
		fmt.Println(string(body))
	}
	fmt.Println("=====================================================================")
}
