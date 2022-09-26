from selenium.webdriver.common.by import By

class Locators:
    CHANNELNAME = By.XPATH, "//div/div/div[2]/section/div[1]/div[2]/h3"
    INPUTSPAN = By.XPATH, "//div/div[3]/div/div/div/span/span/span"
    INPUTSPANDELTA = By.XPATH, "//form/div/div[2]/div/div[2]/div/div/div/span/span[1]/span"
    EMAIL = By.XPATH, '//input[@name="email"]'
    PASSWORD = By.XPATH, '//input[@name="password"]'
    LOGIN = By.XPATH, '//button[@type="submit"]'
    SKIP_MODAL = By.XPATH, '//div[contains(text(), "Продолжить")]'