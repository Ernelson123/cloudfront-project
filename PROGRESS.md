# Project Progress: How the Internet Reaches My App

**Domain:** `dodon.me`  
**Repository:** `github.com/Ernelson123/cloudfront-project`  
**Goal:** Learn enterprise-grade cloud infrastructure (AWS S3, CloudFront, Route 53, ACM) by building a secure, automated static website deployment pipeline.

---

## ✅ Phase 1: Foundations and a Secure Account
- [x] Enabled MFA on the AWS Root account.
- [x] Created a least-privilege IAM user (`ernelson-admin`) for daily work.
- [x] Set up a €5 Billing Budget alert in AWS.
- [x] Acquired the domain `dodon.me` via GitHub Student Developer Pack (Namecheap).
- [x] Created a Route 53 Public Hosted Zone for `dodon.me`.
- [x] Delegated DNS: Updated Namecheap Custom Nameservers to point to the 4 AWS Route 53 nameservers.
- [x] Initialized Git repository and Go module (`tools/infra-check`).
- **Checkpoint Passed:** `dig +trace dodon.me` confirms AWS is the authoritative DNS provider.

## ✅ Phase 2: The Website (Origin Content)
- [x] Created `site/index.html`, `site/404.html`, and `site/style.css`.
- [x] Ensured no inline scripts are used (preparing for strict CSP in Phase 7).
- [x] Built a local Go server (`tools/local-server/main.go`) using `http.FileServer` to test the site at `http://localhost:8080`.
- **Checkpoint Passed:** Site renders correctly locally. Understood that "static" means the server serves raw files from disk without executing code or querying a database per request.
- [x] Committed and tagged as `v0.2.0`.

## ✅ Phase 3: The Private Origin (S3)
- [x] Created S3 bucket: `dodon-me-website` in `eu-north-1` (Stockholm).
- [x] Enabled **Block All Public Access** (all 4 settings).
- [x] Enabled **Bucket Versioning** (for easy rollbacks).
- [x] Enabled **Default Encryption** (SSE-S3).
- [x] Uploaded `index.html`, `404.html`, and `style.css`.
- **Checkpoint Passed:** Accessing the S3 Object URL directly returns an `AccessDenied` XML error. The vault is securely locked.

---

## 🚧 Up Next: Phase 4
- Configure Route 53 DNS records (A, AAAA, CAA).
- Begin building the `infra-check dns` Go command to programmatically verify DNS resolution.

## ✅ Phase 4: DNS, the internet’s phone book
- [x] Added CAA record (`0 issue "amazon.com"`) to Route 53 to restrict certificate issuance to AWS.
- [x] Built and tested `tools/infra-check/dns.go` to programmatically verify NS and A records.
- [x] Fixed Go module structure (moved `go.mod` to project root).
- **Checkpoint Passed:** `go run ./tools/infra-check dns dodon.me` successfully returns the 4 AWS nameservers.
