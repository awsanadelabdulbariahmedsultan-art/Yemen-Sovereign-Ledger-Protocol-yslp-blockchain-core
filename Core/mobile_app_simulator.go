/*
===================================================================================================
                         PROPRIETARY SOURCE CODE & INTELLECTUAL PROPERTY
===================================================================================================
PROTOCOL NAME:  Yemen Sovereign Ledger Protocol (YSLP) - Mobile UI App Simulator
VERSION:        2.6.5 (Cloud Deployable Production Core)
AUTHOR/CREATOR: Eng. Awsan Adel Abdulbari Ahmed Sultan
COUNTRY:        Yemen (الجمهورية اليمنية)
CONTACT PHONE:  +967 777852433 / +967 776633003
EMAIL:          awsandew@gmail.com
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
	"fmt"
	"log"
	"net/http"
	"os" // حزمة قراءة متغيرات البيئة السحابية للمنافذ الديناميكية
)

func appSimulatorHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	
	htmlContent := `
	<!DOCTYPE html>
	<html lang="ar" dir="rtl">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>CBOYD - محاكي تطبيق الهاتف المحمول</title>
		<style>
			:root {
				--phone-bg: #0b132b;
				--app-primary: #1c2541;
				--app-secondary: #3a506b;
				--app-accent: #5bc0be;
				--neon-green: #64dfdf;
				--card-gradient: linear-gradient(135deg, #11998e, #38ef7d);
				--text-light: #ffffff;
			}
			* { box-sizing: border-box; margin: 0; padding: 0; font-family: 'Segoe UI', sans-serif; }
			body { background-color: #111; display: flex; justify-content: center; align-items: center; min-height: 100vh; padding: 10px; }
			
			/* هيكل محاكاة جسم الهاتف الذكي */
			.smartphone {
				width: 360px;
				height: 740px;
				background-color: #000;
				border: 12px solid #333;
				border-radius: 40px;
				box-shadow: 0 25px 50px rgba(0,0,0,0.5);
				position: relative;
				overflow: hidden;
				display: flex;
				flex-direction: column;
			}
			/* الكاميرا العلوية للهاتف (النوتش) */
			.smartphone::before {
				content: '';
				position: absolute;
				top: 0;
				left: 50%;
				transform: translateX(-50%);
				width: 140px;
				height: 25px;
				background-color: #333;
				border-bottom-left-radius: 15px;
				border-bottom-right-radius: 15px;
				z-index: 10;
			}
			
			.app-container {
				background-color: var(--phone-bg);
				width: 100%;
				height: 100%;
				padding: 25px 15px 15px 15px;
				display: flex;
				flex-direction: column;
				overflow-y: auto;
				color: var(--text-light);
			}
			
			/* شريط الحالة العلوي للهاتف */
			.status-bar { display: flex; justify-content: space-between; font-size: 11px; opacity: 0.8; margin-bottom: 15px; padding: 0 5px; }
			
			/* ترويسة التطبيق الداخلي */
			.app-header { display: flex; justify-content: space-between; align-items: center; margin-bottom: 20px; }
			.app-header h2 { font-size: 18px; font-weight: bold; color: var(--text-light); }
			.profile-badge { width: 35px; height: 35px; background-color: var(--app-secondary); border-radius: 50%; display: flex; justify-content: center; align-items: center; font-size: 12px; border: 2px solid var(--app-accent); font-weight: bold; }
			
			/* بطاقة الرصيد الرقمية المتحركة */
			.balance-card { background: var(--card-gradient); padding: 20px; border-radius: 20px; box-shadow: 0 10px 20px rgba(56,239,125,0.2); margin-bottom: 20px; position: relative; overflow: hidden; }
			.balance-card::after { content: ''; position: absolute; top: -50%; left: -50%; width: 200%; height: 200%; background: linear-gradient(45deg, transparent, rgba(255,255,255,0.1), transparent); transform: rotate(45deg); animation: shine 4s infinite; }
			.balance-card p { font-size: 12px; opacity: 0.9; margin-bottom: 5px; }
			.balance-card h1 { font-size: 24px; font-weight: bold; margin-bottom: 15px; letter-spacing: 0.5px; }
			.card-meta { display: flex; justify-content: space-between; font-size: 11px; opacity: 0.8; }
			
			/* أزرار العمليات السريعة بالتطبيق */
			.action-grid { display: grid; grid-template-columns: repeat(3, 1fr); gap: 12px; margin-bottom: 25px; }
			.action-btn { background-color: var(--app-primary); border: 1px solid rgba(255,255,255,0.05); padding: 12px; border-radius: 15px; text-align: center; cursor: pointer; transition: all 0.2s; }
			.action-btn:hover { background-color: var(--app-secondary); border-color: var(--app-accent); }
			.action-btn .icon { font-size: 20px; margin-bottom: 5px; }
			.action-btn span { font-size: 11px; display: block; font-weight: 500; }
			
			/* أقسام الخدمات المتطورة للمنظومة */
			.section-title { font-size: 14px; font-weight: bold; margin-bottom: 12px; color: var(--app-accent); display: flex; justify-content: space-between; align-items: center; }
			
			.service-list { display: flex; flex-direction: column; gap: 10px; margin-bottom: 20px; }
			.service-item { background-color: var(--app-primary); padding: 15px; border-radius: 15px; display: flex; justify-content: space-between; align-items: center; border-right: 4px solid var(--app-accent); }
			.service-details h4 { font-size: 13px; margin-bottom: 3px; }
			.service-details p { font-size: 11px; color: #a0aec0; }
			.status-tag { font-size: 10px; background-color: rgba(100, 223, 223, 0.1); color: var(--neon-green); padding: 3px 8px; border-radius: 10px; font-weight: bold; }
			
			/* زر التمرير الصناعي واللوجستي الحصري */
			.nfc-pulse-btn { background: radial-gradient(circle, var(--app-secondary), var(--phone-bg)); border: 2px dashed var(--app-accent); padding: 20px; border-radius: 20px; text-align: center; cursor: pointer; margin-top: 10px; position: relative; animation: borderPulse 2s infinite; }
			.nfc-pulse-btn h4 { color: var(--neon-green); font-size: 14px; margin-bottom: 5px; }
			.nfc-pulse-btn p { font-size: 11px; opacity: 0.8; }
			
			/* وثيقة الحماية بأسفل الهاتف */
			.app-footer { text-align: center; font-size: 10px; color: #a0aec0; margin-top: auto; padding-top: 20px; border-top: 1px solid rgba(255,255,255,0.05); }
			
			@keyframes shine { 100% { left: 125%; } }
			@keyframes borderPulse { 0% { border-color: rgba(91, 192, 190, 0.4); } 50% { border-color: rgba(100, 223, 223, 1); } 100% { border-color: rgba(91, 192, 190, 0.4); } }
		</style>
	</head>
	<body>

		<div class="smartphone">
			<div class="app-container">
				
				<!-- شريط الهاتف الافتراضي -->
				<div class="status-bar">
					<span>22:36 📱</span>
					<span>CBOYD 5G • 🔋 100%</span>
				</div>
				
				<!-- ترويسة محفظة أوسان سلطان -->
				<div class="app-header">
					<div>
						<p style="font-size: 11px; opacity: 0.7;">أهلاً بك يا مهندس</p>
						<h2>أوسان عادل سلطان</h2>
					</div>
					<div class="profile-badge">ENG</div>
				</div>
				
				<!-- بطاقة الأصول الفعالة بالريال الرقمي -->
				<div class="balance-card">
					<p>الحساب السيادي الفعال (YERD)</p>
					<h1>75,450,000.00 YER</h1>
					<div class="card-meta">
						<span>ID: 01010305468</span>
						<span>تأمين ECDSA P-256 🔒</span>
					</div>
				</div>
				
				<!-- شبكة أزرار العمليات الميدانية السريعة -->
				<div class="action-grid">
					<div class="action-btn">
						<div class="icon">📤</div>
						<span>إرسال أموال</span>
					</div>
					<div class="action-btn">
						<div class="icon">📥</div>
						<span>استلام أموال</span>
					</div>
					<div class="action-btn">
						<div class="icon">🔲</div>
						<span>مسح الـ QR</span>
					</div>
				</div>
				
				<!-- أقسام الربط والتبادل والمقاصة الفورية -->
				<div class="section-title">
					<span>القنوات والربط والتبادل الدولي</span>
					<span style="font-size: 10px; color:#64dfdf;">متصل حياً</span>
				</div>
				
				<div class="service-list">
					<div class="service-item">
						<div class="service-details">
							<h4>مقاصة وتحويل متعدد الأصول</h4>
							<p>تبادل فوري: YER ⇄ USD ⇄ USDT</p>
						</div>
						<span class="status-tag">Binance/Bloomberg</span>
					</div>
					
					<div class="service-item">
						<div class="service-details">
							<h4>ربط المحافظ الإلكترونية اليمنية</h4>
							<p>أم فلوس، جيب، جوالي، ون كاش...</p>
						</div>
						<span class="status-tag">API Gateway</span>
					</div>
				</div>
				
				<!-- زر التمرير الميداني للصناعات الثقيلة والسفن والبركات النفطية -->
				<div class="nfc-pulse-btn">
					<h4>تفعيل دفع وتمرير الهاردوير NFC</h4>
					<p>للمصانع • أرصفة Mوانئ وسفن الشحن • تفريغ بركات النفط</p>
				</div>
				
				<!-- ميتاداتا حقوق النشر والابتكار بأسفل شاشة الهاتف -->
				<div class="app-footer">
					<p>بروتوكول YSLP السيادي لليمن v2.6.5</p>
					<p style="margin-top: 3px; font-size:9px; color: #5bc0be;">حقوق الملكية مسجلة © 2026 م. أوسان سلطان</p>
				</div>

			</div>
		</div>

	</body>
	</html>
	`
	fmt.Fprint(w, htmlContent)
}

func main() {
	http.HandleFunc("/", appSimulatorHandler)
	
	// قراءة المنفذ الديناميكي من خوادم السحابة لتجنب أخطاء حجز الـ Port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // البورت الافتراضي في حال التشغيل المحلي أو في Cloud Shell
	}

	fmt.Println("=======================================================================================")
	fmt.Println("       [⚡ YSLP PRO] NEW DEPLOYABLE MOBILE APP SIMULATOR IS ONLINE...                  ")
	fmt.Println("       DESIGN ARCHITECT: ENG. AWSAN ADEL ABDULBARI AHMED SULTAN                       ")
	fmt.Println("=======================================================================================")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
