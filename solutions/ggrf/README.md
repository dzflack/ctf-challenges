# Solution

- XSS in PDF generation which leads to SSRF, which can be used to enumerate local ports and retrieve the flag 
- The example domain used in `solve.py` is `ni.mba`, this should be changed to a domain you control that is less than 8 chars
- Create a DNS record pointing your domain to your LAN IP (i.e. `ni.mba` -> `192.168.1.3`)
- `cd` to `./solutions/ggrf/`
- Serve `b.js` on port 80, via `sudo python3 -m http.server`
- Run `python3 solve.py`, wait ~20 seconds, open the retrieved PDF, and view the flag.