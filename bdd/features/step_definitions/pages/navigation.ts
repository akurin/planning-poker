import { ThenableWebDriver } from "selenium-webdriver";

export class Navigation {
  private _driver: ThenableWebDriver;

  constructor(driver: ThenableWebDriver) {
    this._driver = driver;
  }

  async navigateTo(url: string): Promise<void> {
    await this._driver.get(url);
  }
}
