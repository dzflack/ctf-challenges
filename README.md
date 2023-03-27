# CTF-Challenges

This repo contains a collection of the CTF challenges I have written over the last couple years.

Challenge files are in the [challenges](challenges) folder, and their corresponding solutions in [solutions](solutions)

## Requirements

- docker
- docker-compose

## Start

To get the challenges up and running:

1. Clone this repo `git clone https://github.com/dzflack/ctf-challenges`
2. Run docker compose `cd ctf-challenges/challenges && docker-compose up`
3. Navigate to the challenges as listed below

## Challenges

The challenge descriptions and their associated URLs are below

### gitouttahere

Just another day in the land of poorly deployed websites. (basic web enumeration is required for this challenge, please don't DoS our server though).

<https://127.0.0.1:9000>

### monitctf

Here at ${insert_hip_and_cool_company} we use this nifty little application to monitor our processes. Can you please do the haxing of it?

<http://127.0.0.1:9100>

### cccrypto

Can you decrypt the thing?

<http://127.0.0.1:9200>

### jarjarbinks

Please look at this new website, the Yeewinator

<http://127.0.0.1:9300>

### ggrf

Check out this crackin HTML5 web site. Our GRC team told us that they run a highly sensitive flag service on this server...

<http://127.0.0.1:9400>

### obfuscatego

Something has gone wrong with our code! Please navigate to the site for more info

<http://127.0.0.1:9500>

### badifi

Simply find the 0 day and login as root. "www-data:www-data" should get you started.

Filedrop: [ctf-badifi-source.zip](https://pub-6103938fe2954b398b3c27265496ccf4.r2.dev/ctf-badifi-source.zip)

<http://127.0.0.1:9600>

### zippitydoodah

While trying to hax our brand new Yelsa Model 99, we've intercepted a firmware blob which we believe contains a secret file that will help us. Can you retrieve this secret file and show us the way? Oh right, and there is some Crypto of type Zip involved somehow.

> NOTE - There is no interactive service for this chal. Filedrop is below

Filedrop: [alpine-standard-3.16.2-aarch64.iso.bin](https://pub-6103938fe2954b398b3c27265496ccf4.r2.dev/alpine-standard-3.16.2-aarch64.iso.bin)

### Cache em All

We retrieved this Enterprise Application which seems to cache credentials. Can you retrieve the credentials of the application's last user?

> NOTE - There is no interactive service for this chal. Filedrop is below

Filedrop: [cache-credentials.zip](https://pub-6103938fe2954b398b3c27265496ccf4.r2.dev/cache-credentials.zip)

## Solutions

- [gitouttahere](solutions/gitouttahere.md)
- [monitctf](solutions/monitctf.md)
- [obfuscatego](solutions/obfuscatego.md)
- [cccrypto](solutions/cccrypto/README.md)
- [jarjarbinks](solutions/jarjarbinks/README.md)
- [ggrf](solutions/ggrf/README.md)
- [badifi](solutions/badifi/README.md)
- [zippitydoodah](solutions/zippitydoodah.md)
- [cache-em-all](solutions/cache-em-all.md)

## Flag Format

`CTF{something_here}`
