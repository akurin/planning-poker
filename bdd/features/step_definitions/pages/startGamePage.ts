import { By, ThenableWebDriver, until } from "selenium-webdriver";

export class StartGamePage {
  private _driver: ThenableWebDriver;

  constructor(driver: ThenableWebDriver) {
    this._driver = driver;
  }

  async setGameName(gameName: string): Promise<void> {
    const gameNameInput = await this._driver.findElement(
      By.css("[data-qa='game-name']")
    );
    await gameNameInput.sendKeys(gameName);
  }

  async startGame(): Promise<void> {
    const startGameButton = await this._driver.findElement(
      By.css("[data-qa='start-game']")
    );
    await startGameButton.click();
  }
}
