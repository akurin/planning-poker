import { setWorldConstructor, World } from "cucumber";
import { Builder, ThenableWebDriver } from "selenium-webdriver";

import { GamePage } from "../step_definitions/pages/gamePage";
import { Navigation } from "../step_definitions/pages/navigation";
import { StartGamePage } from "../step_definitions/pages/startGamePage";

class CustomWorld implements World {
  private readonly _driver: ThenableWebDriver;
  readonly navigation: Navigation;
  readonly startGamePage: StartGamePage;
  readonly gamePage: GamePage;

  constructor() {
    this._driver = this.createDriver();
    this.navigation = new Navigation(this._driver);
    this.startGamePage = new StartGamePage(this._driver);
    this.gamePage = new GamePage(this._driver);
  }

  async quit() {
    await this._driver.quit();
  }

  private createDriver() {
    const seleniumUrl = process.env.SELENIUM_URL;

    const builder = new Builder().forBrowser("chrome");
    if (seleniumUrl) builder.usingServer(seleniumUrl);

    return builder.build();
  }

  async taskDelay(seconds: number): Promise<void> {
    await this._driver.sleep(seconds * 1000);
  }
}

setWorldConstructor(CustomWorld);
