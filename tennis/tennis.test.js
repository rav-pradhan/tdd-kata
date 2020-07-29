import {
  assertEquals,
} from "https://deno.land/std/testing/asserts.ts";

import TennisMatch from './tennis.js'

const setScore = (match, p1, p2) => {
  for (let i = 0; i < p1; i++) {
    match.playerOneWinsAPoint()
  }

  for (let i = 0; i < p2; i++) {
    match.playerTwoWinsAPoint()
  }
}

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
  setScore(match, 1, 0);
  const score = match.getScore();
  assertEquals(score, "15-0");
})

Deno.test("when player two wins a point, the score is 0-15", () => {
  const match = new TennisMatch("Roger Federer", "Rafael Nadal");
  setScore(match, 0, 1);
  const score = match.getScore();
  assertEquals(score, "0-15");
})

Deno.test("when both players win a point, the score is 15-15", () => {
  const match = new TennisMatch("Roger Federer", "Rafael Nadal");
  setScore(match, 1, 1);
  const score = match.getScore();
  assertEquals(score, "15-15");
})

Deno.test("that the score is 30-0 when player one wins two points", () => {
  const match = new TennisMatch("Roger Federer", "Rafael Nadal");
  setScore(match, 2, 0);
  const score = match.getScore();
  assertEquals(score, "30-0");
})


Deno.test("that the score is 0-40 when player two wins three points", () => {
  const match = new TennisMatch("Roger Federer", "Rafael Nadal");
  setScore(match, 0, 3);
  const score = match.getScore();
  assertEquals(score, "0-40");
})

Deno.test("when player is on game point, and they win the next point, then player wins the game", () => {
  const match = new TennisMatch("Roger Federer", "Rafael Nadal");
  setScore(match, 4, 0);
  const score = match.getScore();
  assertEquals(score, "Game won by Roger Federer")
})

Deno.test("when both players are on 40-40, and a player wins a point, the score is A to that player", () => {
  const match = new TennisMatch("Roger Federer", "Rafael Nadal");
  setScore(match, 3, 3);
  match.playerOneWinsAPoint();

  const score = match.getScore();
  assertEquals(score, "A-40")
})

Deno.test("when both players are on 40-40, and a player wins the next two consecutive points, that player is deemed the winner", () => {
  const match = new TennisMatch("Roger Federer", "Rafael Nadal");
  setScore(match, 3, 5);

  const score = match.getScore();
  assertEquals(score, "Game won by Rafael Nadal")
})

Deno.test("when it's advantage, and the opposing player wins the point, score returns to 40-40", () => {
  const match = new TennisMatch("Roger Federer", "Rafael Nadal");
  setScore(match, 3, 3);
  match.playerOneWinsAPoint();
  match.playerTwoWinsAPoint();

  const score = match.getScore();
  assertEquals(score, "40-40")
})