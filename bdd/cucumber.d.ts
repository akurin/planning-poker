import { GamePage } from "./features/step_definitions/pages/gamePage";
import { Navigation } from "./features/step_definitions/pages/navigation";
import { StartGamePage } from "./features/step_definitions/pages/startGamePage";

declare module "cucumber" {
  export interface World {
    readonly navigation: Navigation;
    readonly startGamePage: StartGamePage;
    readonly gamePage: GamePage;

    quit(): Promise<void>;
  }
}
