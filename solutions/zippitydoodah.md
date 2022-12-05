# Solution

* brew install libomp
* Download a release of bkcrack from here <https://github.com/kimci86/bkcrack/releases>
* Create a file `plain.txt` with the contents "alpine-standard-3.16.2 220706" (no space or newlines. this is from the .alpine-release file as retrieved from <https://dl-cdn.alpinelinux.org/alpine/v3.16/releases/aarch64/alpine-standard-3.16.2-aarch64.iso>)
* ./bkcrack -C alpine-standard-3.16.2-aarch64.iso.bin -c .alpine-release -p plain.txt -e
* Check the keys identified are "1f31450a c85ff17c cdd78d6a"
* ../bkcrack -C alpine-standard-3.16.2-aarch64.iso.bin -k 1f31450a c85ff17c cdd78d6a -U decrypted.zip password1
* Open decrypted.zip, using password1 as the password
* Flag is in secrets.txt
