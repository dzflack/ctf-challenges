# Solution

- SSRF (I mean it's contrived but yeah) + Zip Slip leading to WAR deployment in webapps folder
- Run [solution.sh](solution.sh)
- This will generate app.jar, which has a directory traversal WAR filename, with a webshell in it
- Serve app.jar with `python3 -m http.server 8000`
- Send a POST to the app as instructed by the script, which triggers the SSRF to download and extract app.jar
- Browse to the webshell as instructed by the script, and cat the flag to win
