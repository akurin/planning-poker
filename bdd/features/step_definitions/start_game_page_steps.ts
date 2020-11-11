import { expect } from "chai";
import { Then, When } from "cucumber";
import { Given } from "cucumber";

Given('I navigated to the "Start game" page', async function () {
  await this.navigation.navigateTo("http://testhost:5000");
});

Given("I entered the game name {string}", async function (gameName: string) {
  await this.startGamePage.setGameName(gameName);
});

When("I start the game", async function () {
  await this.startGamePage.startGame();
});

Then("I should be redirected to the game {string}", async function (
  expectedGameName: string
) {
  const actualGameName = await this.gamePage.getGameName();
  expect(actualGameName).to.be.equal(expectedGameName);
});
