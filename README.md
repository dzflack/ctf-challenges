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

`https://127.0.0.1:9000`

### monitctf

Here at ${insert_hip_and_cool_company} we use this nifty little application to monitor our processes. Can you please do the haxing of it?

`http://127.0.0.1:9100`

### cccrypto

Can you decrypt the thing?

`http://127.0.0.1:9200`

### jarjarbinks

Please look at this new website, the Yeewinator

`http://127.0.0.1:9300`

### ggrf

Check out this crackin HTML5 web site. Our GRC team told us that they run a highly sensitive flag service on this server...

`http://127.0.0.1:9400`

### obfuscatego

Something has gone wrong with our code! Please navigate to the site for more info

`http://127.0.0.1:9500`

## Solutions

- [gitouttahere](solutions/gitouttahere.md)
- [monitctf](solutions/monitctf.md)
- [obfuscatego](solutions/obfuscatego.md)
- [cccrypto](solutions/cccrypto/README.md)
- [jarjarbinks](solutions/jarjarbinks/README.md)
- [ggrf](solutions/ggrf/README.md)

## Flag Format

`CTF{something_here}`
