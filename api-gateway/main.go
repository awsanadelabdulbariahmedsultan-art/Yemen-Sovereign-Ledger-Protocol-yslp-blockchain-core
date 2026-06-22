/*
===================================================================================================
                         PROPRIETARY SOURCE CODE & INTELLECTUAL PROPERTY
===================================================================================================
MODULE NAME:    CBOYD API Gateway & Wallet Microservices / بوابة المحافظ الرقمية
VERSION:        1.0.0 (Production Routing Core)
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
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// APIResponse الهيكل القياسي للرد على طلبات تطبيقات المحافظ
type APIResponse struct {
	Success   bool   `json:"success"`
	Message   string `json:"message"`
	TxID      string `json:"tx_id,omitempty"`
	Timestamp int64  `json:"timestamp"`
}

// WalletTransferRequest هيكل البيانات المستلمة من تطبيقات الـ Mobile API لطلب التحويل
type WalletTransferRequest struct {
	WalletType   string  `json:"wallet_type"`   // مثال: M-Floos, Jaib, Jawali
	SenderPhone  string  `json:"sender_phone"`  // رقم حساب المرسل
	ReceiverPhone string `json:"receiver_phone"`// رقم حساب المستلم
	Amount       float64 `json:"amount"`         // المبلغ بالريال اليمني
	AssetSymbol  string  `json:"asset_symbol"`  // YER, USD, USDT
}

// handleWalletTransfer معالج طلبات التحويل الفوري بين المحافظ والمقاصة الرقمية
func handleWalletTransfer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(APIResponse{Success: false, Message: "طريقة الطلب غير مدعومة، يجب استخدام POST", Timestamp: time.Now().Unix()})
		return
	}

	var req WalletTransferRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || req.Amount <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(APIResponse{Success: false, Message: "بيانات الطلب غير صالحة أو غير مكتملة", Timestamp: time.Now().Unix()})
		return
	}

	// محاكاة التوجيه البرمجي والتوقيع الرقمي في نوات البلوكشين (YSLP Core Integration)
	log.Printf("[API GATEWAY] طلب تحويل مستلم من محفظة [%s] بمبلغ %.2f %s\n", req.WalletType, req.Amount, req.AssetSymbol)
	
	mockTxID := "cboyd_tx_" + fmt.Sprintf("%d", time.Now().UnixNano())[:8]

	response := APIResponse{
		Success:   true,
		Message:   fmt.Sprintf("تم استلام الطلب من محفظة %s وتوجيهه لنواة البلوكشين للتوقيع والتعميد", req.WalletType),
		TxID:      mockTxID,
		Timestamp: time.Now().Unix(),
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func main() {
	// تعيين المسارات (Routes) للبوابة البرمجية للمحافظ ومزودي الخدمة
	http.HandleFunc("/api/v1/wallet/transfer", handleWalletTransfer)

	fmt.Println("=======================================================================================")
	fmt.Println("       CBOYD UNIFIED API GATEWAY & MICROSERVICES INTERFACE                             ")
	fmt.Println("       DEVELOPER: ENG. AWSAN ADEL ABDULBARI AHMED SULTAN                               ")
	fmt.Println("=======================================================================================")
	fmt.Println("[*] خادم البوابة البرمجية للمحافظ يعمل الآن بنجاح على المنفذ المحلي: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
