import {
  assertEquals,
} from "https://deno.land/std/testing/asserts.ts";

import TennisMatch from './tennis.js'

Deno.test("that the tennis match is created with two players", () => {
  const match = new TennisMatch("Roger Federer", "Rafael Nadal");
  assertEquals(match.playerOne.name, "Roger Federer");
  assertEquals(match.playerTwo.name, "Rafael Nadal");
})

Deno.test("when the match is started, score is 0-0", () => {
  const match = new TennisMatch("Roger Federer", "Rafael Nadal");
  const score = match.getScore();
  assertEquals(score, "0-0");
})

Deno.test("when player one wins a point, the score is 15-0", () => {
  const match = new TennisMatch("Roger Federer", "Rafael Nadal");
  match.playerOneWinsAPoint();
  const score = match.getScore();
  assertEquals(score, "15-0");
})

Deno.test("when player two wins a point, the score is 0-15", () => {
  const match = new TennisMatch("Roger Federer", "Rafael Nadal");
  match.playerTwoWinsAPoint();
  const score = match.getScore();
  assertEquals(score, "0-15");
})

Deno.test("when both players win a point, the score is 15-15", () => {
  const match = new TennisMatch("Roger Federer", "Rafael Nadal");
  match.playerTwoWinsAPoint();
  match.playerOneWinsAPoint();
  const score = match.getScore();
  assertEquals(score, "15-15");
})

Deno.test("when player is on game point, and they win the next point, then player wins the game", () => {
  const match = new TennisMatch("Roger Federer", "Rafael Nadal");
  match.playerOneWinsAPoint();
  match.playerOneWinsAPoint();
  match.playerOneWinsAPoint();
  match.playerOneWinsAPoint();
  const score = match.getScore();
  assertEquals(score, "Game won by Roger Federer")
})

Deno.test("when both players are on 40-40, and a player wins a point, the score is A to that player", () => {
  const match = new TennisMatch("Roger Federer", "Rafael Nadal");
  match.playerOneWinsAPoint();
  match.playerOneWinsAPoint();
  match.playerOneWinsAPoint();
  match.playerTwoWinsAPoint();
  match.playerTwoWinsAPoint();
  match.playerTwoWinsAPoint();

  match.playerOneWinsAPoint();

  const score = match.getScore();
  assertEquals(score, "A-40")
})

Deno.test("when it's advantage, and the opposing player wins the point, score returns to 40-40", () => {
  const match = new TennisMatch("Roger Federer", "Rafael Nadal");
  match.playerOneWinsAPoint();
  match.playerOneWinsAPoint();
  match.playerOneWinsAPoint();
  match.playerTwoWinsAPoint();
  match.playerTwoWinsAPoint();
  match.playerTwoWinsAPoint();

  match.playerOneWinsAPoint();
  match.playerTwoWinsAPoint();

  const score = match.getScore();
  assertEquals(score, "40-40")
})

