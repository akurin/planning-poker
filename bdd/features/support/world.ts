import { setWorldConstructor, World } from "cucumber";
import { Builder, ThenableWebDriver, WebDriver } from "selenium-webdriver";
import chrome from "selenium-webdriver/chrome";
import { Level, Preferences, Type } from "selenium-webdriver/lib/logging";

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
    const builder = new Builder()
      .forBrowser("chrome")
      .setChromeOptions(new chrome.Options().headless());

    if (process.env.CI)
      builder.setChromeOptions(new chrome.Options().headless());

    const p = new Preferences();
    p.setLevel(Type.BROWSER, Level.ALL);
    p.setLevel(Type.SERVER, Level.ALL);
    builder.setLoggingPrefs(p);
    return builder.build();
  }

  async taskDelay(seconds: number): Promise<void> {
    await this._driver.sleep(seconds * 1000);
  }
}

setWorldConstructor(CustomWorld);
