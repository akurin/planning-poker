import { By, Condition, ThenableWebDriver, until } from "selenium-webdriver";

export class GamePage {
  private _driver: ThenableWebDriver;

  constructor(driver: ThenableWebDriver) {
    this._driver = driver;
  }

  async waitUntilLoaded(): Promise<void> {
    await this._driver.wait(until.elementLocated(By.css('[data-qa="game"]')));
    await this._driver.wait(
      until.elementLocated(By.css('[data-qa="game-title"]'))
    );
  }

  async getGameName(): Promise<string> {
    const titleElement = await this._driver.findElement(
      By.css('[data-qa="game-title"]')
    );
    return await titleElement.getText();
  }
}
