import { ThenableWebDriver } from "selenium-webdriver";

export class GamePage {
  private _driver: ThenableWebDriver;

  constructor(driver: ThenableWebDriver) {
    this._driver = driver;
  }

  async getGameName(): Promise<string> {
    return await this._driver.getTitle();
  }
}
