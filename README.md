# Automated Vuln-Scanner CLI

A custom, concurrent port scanner written in Go designed to quickly identify open critical ports during initial penetration testing reconnaissance.

## Features
- **Concurrency**: Uses Go routines to scan multiple ports concurrently, significantly speeding up the enumeration process compared to sequential scanning.
- **Port Profiling**: Defaults to a specific profile of common high-risk ports.
- **Cross-Platform**: Compiles easily for any OS.

## Setup & Usage
Ensure you have Go installed on your system.

```bash
# Build the binary
go build -o vuln-scanner main.go

# Run the scanner
./vuln-scanner scanme.nmap.org
```

*Disclaimer: Only use this tool against targets you have explicit permission to test.*
