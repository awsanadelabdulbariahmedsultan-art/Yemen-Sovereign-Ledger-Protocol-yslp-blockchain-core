package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

/*
===========================================================================================
  PROPRIETARY SOURCE CODE & INTELLECTUAL PROPERTY (AUTOMATED TESTING SUITE)
  MODULE NAME:   CBYD API Gateway - Test API Request Script / سكربت محاكاة واختبار قنوات الدفع
  VERSION:       3.8.0 (Enterprise Automation Node)
  AUTHOR/OWNER:  ENG. AWSAN ADEL ABDULBARI AHMED SULTAN / ID: 01010305468 / YEMEN (2026)
  CONTACT PHONE: +967 777832433 / +967 776633003 | EMAIL: awsandew@outlook.com
===========================================================================================
*/

type WalletTransferPayload struct {
	WalletType      string  `json:"wallet_type"`
	SenderPhone     string  `json:"sender_phone"`
	ReceiverPhone   string  `json:"receiver_phone"`
	Amount          float64 `json:"amount"`
	AssetSymbol     string  `json:"asset_symbol"`
	PaymentMethod   string  `json:"payment_method"`
	ShippingCarrier string  `json:"shipping_carrier"`
	ContainerID     string  `json:"container_id"`
	BankIban        string  `json:"bank_iban"`
}

func main() {
	fmt.Println("=========================================================================================")
	fmt.Println(" [ ] CBYD API GATEWAY - AUTOMATED WALLET TRANSACTION SIMULATOR... ")
	fmt.Println(" SYSTEM LEGAL OWNER: ENG. AWSAN ADEL ABDULBARI AHMED SULTAN")
	fmt.Println("=========================================================================================")
	fmt.Println("[*] Initiating cross-border clearing payload simulation...")

	apiURL := "http://localhost:8080/api/v1/wallet/transfer"

	// Mocking a complete institutional multi-asset payment through WeChat/Shipping Line
	payload := WalletTransferPayload{
		WalletType:      "M-Floos (Kurimi)",
		SenderPhone:     "+967772054648",
		ReceiverPhone:   "+967776633003",
		Amount:          2500.00,
		AssetSymbol:     "USDT",
		PaymentMethod:   "WeChat Pay",
		ShippingCarrier: "COSCO Shipping Lines (China)",
		ContainerID:     "MSCU1072054",
		BankIban:        "YE88YSLP01010303468000",
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		fmt.Printf("[X] JSON Marshalling Error: %v\n", err)
		return
	}

	startTime := time.Now()
	resp, err := http.Post(apiURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("[X] Connectivity Error: Ensure api-gateway server is live on 8080. Error: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("[X] Read Response Payload Error: %v\n", err)
		return
	}

	duration := time.Since(startTime)

	fmt.Println("[✓] Network Clearing Endpoint Hit Successfully.")
	fmt.Printf("[✓] Settlement Gateway Response Latency: %v\n", duration)
	fmt.Println("[✓] Signed Node Payload (JSON Response Output):")

	var prettyJSON bytes.Buffer
	errorPretty := json.Indent(&prettyJSON, body, "", "  ")
	if errorPretty != nil {
		fmt.Println(string(body))
	} else {
		fmt.Println(prettyJSON.String())
	}
	fmt.Println("=========================================================================================")
}
