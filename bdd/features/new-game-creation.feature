Feature: New game starting
  As a planning poker moderator
  In order to start planning
  I want to be able to start a new game

  Background:
    Given I navigated to the "Start game" page

  Scenario: Start a new game
    Given I entered the game name "Sprint planning"
    When I start the game
    Then I should be redirected to the game "Sprint planning"