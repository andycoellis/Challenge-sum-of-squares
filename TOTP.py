import pyotp
import hashlib
import base64


secret_key = 'REDACTED'
key = base64.b32encode(bytearray(secret_key, 'ascii'))


pw = pyotp.TOTP(key, digest=hashlib.sha512, digits=10, interval=30)
output = pw.now()

print(output)