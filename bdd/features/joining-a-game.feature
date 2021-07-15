Feature: Joining a game
  As a planning poker player
  In order take part in planning
  I want to be able to join a game

  Scenario: Joining a game
    Given There is a game "Sprint planning"
    When I navigate to the game page
    And enter the name "John Doe"
    Then I should see the name "John Doe" in the list of players