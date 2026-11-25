# Automated Vuln-Scanner CLI

A custom, concurrent port scanner written in Go designed to quickly identify open critical ports during initial penetration testing reconnaissance.

## 📖 Complete Brief
This tool is a lightweight, high-speed port scanner built in Go. By utilizing Go concurrency (goroutines), it rapidly enumerates a target's open ports without the overhead of bulkier tools like Nmap. It defaults to a list of the most common high-risk ports (such as SSH, HTTP, HTTPS, SMB, and RDP), making it an excellent utility for quick reconnaissance during initial penetration testing or internal network audits.

## ✨ Features
- **Concurrency**: Built heavily on goroutines to scan multiple ports simultaneously, drastically lowering scan times.
- **Port Profiling**: Defaults to scanning high-value targets (ports 21, 22, 23, 80, 443, 8080, etc.).
- **Cross-Platform Compatibility**: Easily complies to native executables for Windows, Linux, and macOS.
- **Speed & Simplicity**: Provides instant, easy-to-read terminal output.

## 🛠 Prerequisites
- Go 1.18 or higher installed on your system.

## 🚀 How to Run

1. **Navigate to the core directory:**
   ```bash
   cd ./automated-vuln-scanner/
   ```

2. **Build the binary executable:**
   ```bash
   go build -o vuln-scanner main.go
   ```

3. **Run the scanner against a target:**
   ```bash
   ./vuln-scanner <target>
   ```
   *Example:*
   ```bash
   ./vuln-scanner scanme.nmap.org
   ```

*Disclaimer: Only use this tool against targets you have explicit permission to test.*
