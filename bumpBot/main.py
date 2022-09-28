import time

from selenium import webdriver
from selenium.webdriver.common.keys import Keys
from selenium.common.exceptions import NoSuchElementException
from locators import Locators


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
    time.sleep(5)
    return channel_name == page.find_element(Locators.CHANNELNAME).text()


def skip_modal(page):
    try:
        time.sleep(5)
        page.find_element(*Locators.SKIP_MODAL).click()
    except NoSuchElementException:
        pass


def login(page):
    page.get("https://discord.com/login")
    skip_modal(page)
    page.find_element(*Locators.EMAIL).send_keys("misha@idwte.ru")
    time.sleep(5)
    page.find_element(*Locators.PASSWORD).send_keys("GanterStudios")
    time.sleep(5)
    page.find_element(*Locators.LOGIN).click()
    time.sleep(5)


def open_page(page):
    page.get("https://discord.com/channels/465780328611708937/521272424302641163")
    time.sleep(5)


def send_bump(page, text: str):
    page.find_element(*Locators.INPUTSPAN).send_keys(text)
    time.sleep(1)
    page.find_element(*Locators.INPUTSPAN).send_keys(Keys.ENTER)
    time.sleep(1)
    page.find_element(*Locators.INPUTSPANDELTA).send_keys(Keys.ENTER)
    time.sleep(1)


commands: [str] = ["/bump", "/like"]

driver_ = init_driver_remote()
login(driver_)
open_page(driver_)
while True:
    for i in commands:
        send_bump(driver_, i)
    time.sleep(60 * 10) #10 minutes