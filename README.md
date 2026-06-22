# 🏦 Central Bank of Yemen Digital (CBOYD) Engine
### Powered by: Yemen Sovereign Ledger Protocol (YSLP) v2.0.0

---

## 📄 وثيقة الملكية الفكرية والبرمجية / Intellectual Property Notice
* **المطور والمؤسس / Author & Creator:** Eng. Awsan Adel Abdulbari Ahmed Sultan
* **بلد المنشأ / Country:** Yemen (الجمهورية اليمنية)
* **وسائل الاتصال / Contact Phone:** +967 777852433 / +967 776633003
* **البريد الإلكتروني / Email:** awsandew@outlook.com
* **لينكد إن / LinkedIn:** [Eng. Awsan Adel](https://linkedin.com)

> **⚠️ LEGAL NOTICE:** Copyright (c) 2026 Eng. Awsan Adel Abdulbari Ahmed Sultan. All Rights Reserved. Unauthorized copying, distribution, or deployment of this protocol without explicit written permission from the owner is strictly prohibited.

---

## 🌍 نظرة عامة على المشروع (Arabic)
مشروع **CBOYD** هو الإطار العملي الموحد والمحرك السيادي القائم على تكنولوجيا البلوكشين لربط المحافظ الإلكترونية اليمنية المحلية (مثل أم فلوس، جيب، جوالي، ون كاش) وتسهيل التكامل المالي الرقمي مع المنصات والبورصات العالمية (Binance, Bloomberg Terminal) والمنظمات الدولية (البنك الدولي وصندوق النقد الدولي).

يعتمد المشروع في نواته على بروتوكول **YSLP** المكتوب بلغة **Go (Golang)**، والذي يدمج المعاملات البنكية المدنية، والأنظمة اللوجستية وإنترنت الأشياء (IoT / NFC) للمصانع، الموانئ، والبركات النفطية في سجل مشفر وموحد خاضع لرقابة وإجماع البنك المركزي اليمني.

### 🛠️ المميزات التقنية الأساسية:
1. **آلية إجماع بنكية (Proof of Authority - PoA):** تصفير رسوم الغاز العشوائية، ومعالجة آلاف المعاملات في الثانية مع حصر التوثيق بعقدة البنك المركزي السيادية.
2. **نظام تشفير المفاتيح (ECDSA P-256):** حماية المحافظ والحسابات؛ حيث لا تُقبل أي معاملة مالم تُوقع رقمياً بالبصمة البيومترية والمفتاح السري الخاص بالعميل لمنع التزوير.
3. **محرك المقاصة والتحويل آلي متعدد الأصول:** دمج الفيات (الريال اليمني، الدولار) مع الكريبتو والعملات المستقرة (USDT) والتحويل الفوري بناءً على أسعار البورصة اللحظية.
4. **أتمتة إنترنت الأشياء (NFC / IoT Core):** ربط أتمتة خطوط إنتاج المصانع، تفريغ بركات ومصبات الوقود النفطية، ورسوم تفريغ السفن بالموانئ مالياً وتلقائياً عبر الهاردوير الميداني.

---

## 🌍 Project Overview (English)
The **CBOYD** engine is a blockchain-driven sovereign framework designed to unify local Yemeni e-wallets (e.g., M-Floos, Jaib, Jawali, OneCash) and bridge them with global liquidity platforms, digital exchanges (Binance, Bloomberg), and international financial institutions (World Bank, IMF).

At its core runs the **Yemen Sovereign Ledger Protocol (YSLP)** built from scratch using **Go (Golang)**. This protocol consolidates civilian banking, industrial automation, maritime logistics, and refinery IoT systems into a singular, highly secure immutable ledger governed by the Central Bank of Yemen.

### 🛠️ Key Architectural Features:
1. **Sovereign Consensus (Proof of Authority - PoA):** Zero random gas fees, high throughput (thousands of TPS), and validator node exclusivity restricted to the Central Bank.
2. **Asymmetric Cryptography (ECDSA P-256):** Complete wallet protection utilizing private/public key pairs. Transactions require explicit user signature verification (biometric/private key) to prevent tampering.
3. **Multi-Asset Swap & Clearing Engine:** Instantaneous fiat-to-crypto (YER/USD to USDT/BTC) automated clearing using real-time global exchange rates.
4. **Industrial IoT & NFC Integration:** Automated payment settlements triggered by hardware taps on factory production lines, refinery valves/tanks, and port maritime shipping berths.

---

## 📂 هيكلية مجلدات النظام / Project Directory Structure
```text
CBOYD-Project/
├── core/
│   └── main.go           # The entire unified YSLP Blockchain Core Engine (Go)
├── api-gateway/          # Wallet API Connectors & Microservices
├── assets/               # System Architecture Diagrams & UI Logos (CBOYD Graphic Identity)
└── README.md             # This comprehensive production documentation file
```

---

## 💻 طريقة التشغيل والاختبار / Deployment & Testing
لإقلاع المحرك واختبار السيناريوهات الميدانية الأربعة (التحويل المشفر، حظر الهجمات السيبرانية، مدفوعات الموانئ اللوجستية، وأتمتة بركات النفط)، تتبع الخطوات التالية:

1. تأكد من تثبيت بيئة لغة **Go** على جهازك (نسخة 1.18 فما فوق).
2. ادخل إلى مجلد النواة:
   ```bash
   cd core
   ```
3. قم بتشغيل البروتوكول مباشرة عبر الأمر التالي:
   ```bash
   go run main.go
   ```

---
*This protocol stands as a living proof of Yemen's engineering capacity to build next-generation national FinTech infrastructures.*
