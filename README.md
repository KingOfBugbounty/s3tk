<div align="center">
  <img width="768" alt="ChatGPT Image 12 de set de 2025, 21_02_01" src="https://github.com/user-attachments/assets/40674e6f-914b-4cd4-a7a9-421657631756" />
</div>


<h1 align="center">
  S3Scan - S3 Bucket Security Scanner / <a href="https://x.com/OFJAAAH" target="_blank" rel="noopener">@✖️OFJAAAH</a>
</h1>

<p align="center">
  <strong>A powerful S3 bucket security scanner designed for penetration testing and bug bounty hunting</strong>
</p>

<p align="center">
  <img src="https://img.shields.io/badge/Go-1.21+-blue.svg" alt="Go Version">
  <img src="https://img.shields.io/badge/License-MIT-green.svg" alt="License">
  <img src="https://img.shields.io/badge/Platform-Linux%20%7C%20Windows%20%7C%20macOS-lightgrey.svg" alt="Platform">
  <img src="https://img.shields.io/github/release/user/s3scan.svg" alt="Release">
</p>

<p align="center">
  This tool automatically detects misconfigurations and security vulnerabilities in AWS S3 buckets.
</p>

## Features

- **Comprehensive Security Testing**: Tests for multiple S3 bucket vulnerabilities
- **Multiple Input Formats**: Supports various S3 URL formats and bucket names
- **Permission Testing**: Checks LIST, UPLOAD, DELETE, and TAKEOVER capabilities
- **Batch Processing**: Scan multiple buckets from stdin input
- **Colorized Output**: Easy-to-read results with color-coded vulnerability status
- **Detailed Reporting**: Comprehensive summary of findings

## Security Tests Performed

### 1. Bucket Existence Check
- Verifies if the S3 bucket exists
- Identifies non-existent buckets that could be claimed

### 2. List Permissions (READ)
- Tests if bucket contents can be enumerated
- Detects publicly readable buckets
- **Risk**: Information disclosure, data exposure

### 3. Upload Permissions (WRITE)
- Tests if files can be uploaded to the bucket
- Identifies writable buckets
- **Risk**: Malicious file upload, defacement, hosting malware

### 4. Delete Permissions (DELETE)
- Tests if objects can be deleted from the bucket
- Detects buckets with delete permissions
- **Risk**: Data destruction, denial of service

### 5. Bucket Takeover Detection
- Identifies non-existent buckets that can be claimed
- Tests multiple AWS regions for takeover opportunities
- **Risk**: Subdomain takeover, brand impersonation

## Installation

### Prerequisites
- Go 1.21 or higher
- Internet connection for S3 API testing

### Building from Source

1. Clone or download the source code
2. Navigate to the project directory
3. Build the binary:

```bash
go build -o s3scan .
```

### Quick Setup

```bash
# Make the binary executable
chmod +x s3scan

# Optional: Move to system PATH
sudo mv s3scan /usr/local/bin/
```

## Usage

### Basic Usage

The tool reads S3 bucket URLs or names from stdin:

```bash
# Single bucket
echo "my-test-bucket" | ./s3scan

# Multiple buckets from file
cat buckets.txt | ./s3scan

# Multiple buckets inline
echo -e "bucket1\nbucket2\nbucket3" | ./s3scan
```

### Supported Input Formats

The scanner accepts various S3 URL formats:

- Bucket name: `my-bucket-name`
- Virtual-hosted style: `https://my-bucket.s3.amazonaws.com`
- Path-style: `https://s3.amazonaws.com/my-bucket`
- S3 URI: `s3://my-bucket-name`

### Example Commands

```bash
# Scan from a list of domains/subdomains
subfinder -d example.com | grep s3 | ./s3scan

# Scan buckets found during reconnaissance
echo "company-backups" | ./s3scan
echo "app-uploads" | ./s3scan
echo "static-assets" | ./s3scan

# Batch scan from file
cat << EOF | ./s3scan
company-data
backup-bucket
public-assets
user-uploads
EOF
```

### Integration with Other Tools

```bash
# With subfinder and grep
subfinder -d target.com | grep -i s3 | ./s3scan

# With amass
amass enum -d target.com | grep s3 | ./s3scan

# With waybackurls
echo "target.com" | waybackurls | grep s3 | ./s3scan
```

## Output Interpretation

### Vulnerability Status

- **[VULNERABLE]** - Red: Misconfiguration detected
- **[SECURE]** - Green: No vulnerabilities found
- **[EXISTS]** - Green: Bucket exists and accessible
- **[NOT FOUND]** - Yellow: Bucket doesn't exist
- **[TAKEOVER POSSIBLE]** - Magenta: Bucket can be claimed

### Permission Types

- **[LIST]** - Can enumerate bucket contents
- **[UPLOAD]** - Can upload files to bucket
- **[DELETE]** - Can delete objects from bucket
- **[TAKEOVER]** - Bucket doesn't exist and can be claimed

### Example Output

```
[MISCONFIGURED] company-backups
  └─ [LIST] Public read access - can enumerate bucket contents
  └─ [UPLOAD] Public write access - can upload malicious files

[TAKEOVER POSSIBLE] old-app-assets
  └─ [TAKEOVER] Bucket doesn't exist - can be claimed for subdomain takeover

[SECURE] private-data
```

## Legal and Ethical Usage

⚠️ **IMPORTANT**: This tool is designed for:
- Authorized penetration testing
- Bug bounty programs
- Security assessments on systems you own
- Educational purposes

**DO NOT USE** for:
- Unauthorized testing of third-party systems
- Malicious activities
- Illegal access attempts

Always ensure you have proper authorization before testing any S3 buckets.

## Defensive Recommendations

If vulnerabilities are found:

1. **For LIST vulnerabilities**: Configure bucket policies to deny public read access
2. **For UPLOAD vulnerabilities**: Remove public write permissions, implement proper IAM policies
3. **For DELETE vulnerabilities**: Restrict delete permissions to authorized users only
4. **For TAKEOVER opportunities**: Claim unused buckets or ensure they're not referenced in applications

## Contributing

This tool is designed for security professionals and researchers. Contributions that improve detection capabilities or add new security tests are welcome.

## License

This tool is provided for educational and authorized testing purposes only.
