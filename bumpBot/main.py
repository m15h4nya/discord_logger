import requests
import time
import json
import random

def get_conf() -> dict:
    cfg = json.load(open("config.json"))
    return cfg

cfg = get_conf()

bumps = cfg["bumps_body"]
header = cfg["bumps_header"]
baseUrl = "https://discord.com/api/v9"

def send_interaction():
    for bump_body in bumps:
        response = requests.post(f"{baseUrl}/interactions", json=bump_body, headers=header)
        time.sleep(random.randint(10, 20))
        if response.status_code != 204:
            requests.post(f"{baseUrl}/channels/521272424302641163/messages", json={"data": f"Error bumping: {response.status_code}"}, headers=header)
            time.sleep(random.randint(10, 20))


while True:
    send_interaction()
    time.sleep(60 * 10 * 3) # 30 minutes
 