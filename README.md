# CloudFront Project — How the Internet Reaches My App

## GitHub Actions OIDC setup

The deployment role must trust pushes to `main` from this repository. Apply the
repository's trust policy with:

```bash
aws iam update-assume-role-policy \
  --role-name GitHubActionsDeployRole \
  --policy-document file://aws/github-actions-trust-policy.json
```
