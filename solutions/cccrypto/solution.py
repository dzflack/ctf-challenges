#!/bin/python3
from Crypto.Cipher import DES3
from Crypto import Random
from hashlib import md5
import base64

for username in range(1, 999998):
    initialKey = ("Jsg#kkdf*777" + str(username))

    hash = md5(initialKey.encode()).digest()
    finalKey = (hash + hash[:8])

    iv = bytes([93, 3, 22, 99, 4, 82, 162, 34])

    cipher = DES3.new(finalKey, DES3.MODE_CBC, iv)

    encrypted_text = base64.b64decode("ODIOBqJDWHAtvLIv8Zk51WcfZFRKDxJ+")
    plaint = cipher.decrypt(encrypted_text)

    try:
        plaint.decode("utf-8")
        print(username)
        print(plaint.decode("utf-8"))
    except:
        continue
