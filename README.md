# **you-are-l** 🚀  
*A CLI tool for domain & URL intelligence*  

---

## **🔹 Overview**  
`you-are-l` is a powerful command-line tool that helps you analyze domains, URLs, and SSL certificates. It provides:  

✔️ **WHOIS lookups** (full & brief)  
✔️ **DNS record retrieval** (A, AAAA, SOA, etc.)  
✔️ **Short URL unredirection** (trace the final URL)  
✔️ **Webpage metadata & headers fetching**  
✔️ **SSL certificate inspection**  

---

## **🔹 Installation**  

### **1️⃣ Build from Source (Go Required)**
```sh
git clone https://github.com/yourusername/you-are-l.git
cd you-are-l
go build -o you-are-l
```

### **2️⃣ Run the CLI**
```sh
./you-are-l --help
```

---

## **🔹 Usage**  

### 📌 **WHOIS Lookup**  
Retrieve WHOIS information about a domain.  
```sh
./you-are-l whois full example.com
./you-are-l whois brief example.com
```

### 🔗 **Unredirect a Shortened URL**  
Find the final destination of a shortened link.  
```sh
./you-are-l unredirect https://bit.ly/xyz
```

### 🌐 **Get DNS Records**  
Fetch DNS records like A, AAAA, SOA, etc.  
```sh
./you-are-l dns example.com
```

### 📜 **Fetch Webpage Metadata & Headers**  
Retrieve title, description, and HTTP headers of a webpage.  
```sh
./you-are-l fetch https://example.com
```

### 🔒 **Check SSL Certificate Details**  
Retrieve SSL certificate details, including issuer, expiration, serial number, and fingerprint.  
```sh
./you-are-l ssl example.com
```

---

## **🔹 Example Output**  

### ✅ **WHOIS Lookup**
```sh
./you-are-l whois brief example.com
```
```plaintext
Domain Name: example.com
Registrar: Internet Assigned Numbers Authority (IANA)
Updated Date: 2024-01-01
Creation Date: 1995-08-14
Expiration Date: 2030-08-14
```

### ✅ **DNS Lookup**
```sh
./you-are-l dns example.com
```
```plaintext
A Records:
93.184.216.34
AAAA Records:
2606:2800:220:1:248:1893:25c8:1946
SOA Record:
ns.icann.org
```

### ✅ **Unredirect a URL**
```sh
./you-are-l unredirect https://bit.ly/xyz
```
```plaintext
Final URL: https://example.com/articles/some-article
```

### ✅ **Fetch Webpage Content**
```sh
./you-are-l fetch https://example.com
```
```plaintext
Fetched URL: https://example.com
Status Code: 200
Title: Example Domain
Description: This is an example domain used for testing purposes.
Headers:
Content-Type: text/html; charset=UTF-8
```

### ✅ **SSL Certificate Details**
```sh
./you-are-l ssl example.com
```
```plaintext
Subject: example.com
Issuer: Let's Encrypt
Expires: 15 June 2025
Renewed: 10 March 2025
Serial Num: 03A9F3E4A6C3D89B1E71C5F29A
Fingerprint: 1A:2B:3C:4D:5E:6F:7G:8H:9I:10J:11K:12L:13M:14N:15O
Extended Key Usage: TLS Web Server Authentication
```

---

## **🔹 Features Roadmap 🛠️**  
- ✅ Add WHOIS Lookup  
- ✅ Add URL Unredirection  
- ✅ Fetch DNS Records  
- ✅ Fetch Webpage Metadata  
- ✅ SSL Certificate Inspection  
- 🔜 Add Reverse IP Lookup  
- 🔜 Add Subdomain Enumeration  

---

## **🔹 Contributing**  
Feel free to contribute! Open a PR or create an issue.  

1. Fork the repository  
2. Create a new branch (`feature-xyz`)  
3. Commit your changes  
4. Open a pull request  


📢 **Made with ❤️ in Go** | 🌍 **For Ethical Use Only**