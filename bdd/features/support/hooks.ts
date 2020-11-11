import { After } from "cucumber";

After(async function () {
  await this.quit();
});
