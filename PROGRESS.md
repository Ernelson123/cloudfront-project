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

## ✅ Phase 5: HTTPS, Certificates, and Trust
- [x] Requested free public ACM certificate in `us-east-1` for `dodon.me` and `www.dodon.me`.
- [x] Validated certificate via DNS (auto-created CNAME records in Route 53).
- [x] Certificate status is now **Issued**.
- [x] Built `tools/infra-check/tlscheck.go` to programmatically verify TLS version, cipher suite, and certificate expiry.
- **Checkpoint Passed:** Tool correctly identifies that `dodon.me` has no A record yet (expected), and successfully validates TLS on known domains like `google.com`.

## ✅ Phase 6: The CDN (CloudFront)
- [x] Created CloudFront distribution with S3 as the origin.
- [x] Configured Origin Access Control (OAC) to securely grant CloudFront read access.
- [x] Updated S3 Bucket Policy to allow only this specific CloudFront distribution.
- [x] Set Default Root Object to `index.html` and configured custom 404 error response.
- [x] Attached ACM SSL certificate and added Alternate Domain Names (`dodon.me`, `www.dodon.me`).
- [x] Set Price Class to "North America and Europe" to protect the €5 budget.
- **Checkpoint Passed:** `curl -I` returns `HTTP/2 200`, `x-cache: Miss/Hit`, and confirms S3 encryption/versioning headers.

## ✅ Phase 7: Connect everything and harden security
- [x] Route 53 A + AAAA ALIAS records for dodon.me and www.dodon.me -> CloudFront.
- [x] CloudFront Function (JS) cleanly redirects www -> apex (301) without querystring artifacts.
- [x] Response Headers Policy: HSTS, CSP (no inline scripts), nosniff, X-Frame-Options, Referrer-Policy.
- [x] DNSSEC intentionally OFF (~$20/mo would break the €5 budget).
- **Checkpoint:** Go tls tool validates dodon.me with TLS 1.3. www correctly 301s to apex with clean URL. Security headers verified.

## ✅ Phase 8: Automation & Caching (Grand Finale)
- [x] Built `tools/infra-check/cache.go` to programmatically verify CloudFront `X-Cache` (Hit/Miss/RefreshHit) behavior.
- [x] Built `tools/infra-check/headers.go` to validate the presence of HSTS, CSP, X-Frame-Options, and more.
- [x] `infra-check` CLI is now a complete, multi-command validation tool for the entire architecture.
- [x] External validation: securityheaders.com grade = A/A+.
- **Checkpoint:** Full end-to-end verification of DNS, TLS, CDN caching, and security headers via custom Go tooling.

## 🌟 Bonus: securityheaders.com A -> A+
- [x] Added Permissions-Policy custom header to response headers policy (disables camera/mic/geo/pay/usb/sensors).
- **Result:** Independent scan = A+; all 6 scored headers green.

## ✅ Phase 9: CI/CD Automation (GitHub Actions)
- [x] Created `.github/workflows/deploy.yml` to trigger deployments on pushes to `main`.
- [x] Configured AWS IAM OIDC Identity Provider for GitHub (`token.actions.githubusercontent.com`).
- [x] Created IAM Role (`GitHubActionsDeployRole`) with a strict trust policy limiting access to `repo:Ernelson123/cloudfront-project:*`.
- [x] Attached a least-privilege IAM Policy granting only `s3:PutObject`, `s3:DeleteObject`, and `s3:ListBucket` on `dodon-me-website`.
- [x] **Checkpoint Passed:** Pushing a change to `site/index.html` automatically triggered the GitHub Action, which securely assumed the IAM role via OIDC (no hardcoded AWS secrets) and synced the files to S3.
