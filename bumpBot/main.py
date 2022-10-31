import json
from time import sleep

from typing import Tuple, List
from selenium import webdriver
from selenium.webdriver.common.keys import Keys
from selenium.common.exceptions import NoSuchElementException
from locators import Locators


def get_conf() -> Tuple[str, str]:
    cfg = json.load(open("config.json"))
    return cfg


def init_driver():
    options = webdriver.ChromeOptions()
    options.add_argument('--ignore-ssl-errors=yes')
    options.add_argument('--ignore-certificate-errors')
    return webdriver.Chrome(
        executable_path="driver/chromedriver.exe",
        options=options
    )


def init_driver_remote():
    options = webdriver.ChromeOptions()
    options.add_argument('--ignore-ssl-errors=yes')
    options.add_argument('--ignore-certificate-errors')
    return webdriver.Remote(
        command_executor="http://172.17.0.2:4444/wd/hub",
        options=options
    )


def check_for_valid_channel(page: webdriver, channel_name: str):
    sleep(5)
    return channel_name == page.find_element(Locators.CHANNELNAME).text()


def skip_modal(page):
    try:
        sleep(5)
        page.find_element(*Locators.SKIP_MODAL).click()
    except NoSuchElementException:
        pass


def login(page, cfg: dict):
    page.get("https://discord.com/login")
    skip_modal(page)
    page.find_element(*Locators.EMAIL).send_keys(cfg['login'])
    sleep(5)
    page.find_element(*Locators.PASSWORD).send_keys(cfg['password'])
    sleep(5)
    page.find_element(*Locators.LOGIN).click()
    sleep(5)


def open_page(page):
    page.get("https://discord.com/channels/465780328611708937/521272424302641163")
    sleep(5)


def send_bump(page, text: str):
    page.find_element(*Locators.INPUTSPAN).send_keys(text)
    sleep(1)
    page.find_element(*Locators.INPUTSPAN).send_keys(Keys.ENTER)
    sleep(1)
    page.find_element(*Locators.INPUTSPANDELTA).send_keys(Keys.ENTER)
    sleep(1)


commands: List[str] = ["/bump", "/like", "/up"]


driver_ = init_driver_remote()
login(driver_, get_conf())
open_page(driver_)
while True:
    for i in commands:
        send_bump(driver_, i)
    sleep(60 * 10) #10 minutes
