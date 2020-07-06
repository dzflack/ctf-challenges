# Solution

- Dir bust `https://127.0.0.1:9000/.git/` directory
- Get 403 response
- Look for juicy git files
- Find `https://127.0.0.1:9000/.git/config` with creds
- Used creds to authenticate to web root and get flag
