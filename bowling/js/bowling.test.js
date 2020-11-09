import { assertEquals } from "https://deno.land/std/testing/asserts.ts";
import bowlingScore from './bowling.js'

Deno.test("that a game with no pins fallen amounts to a score of 0", () => {
  const got = bowlingScore("00 00 00");
  const want = 0;
  assertEquals(got, want);
});

Deno.test("that a game calculates scores when all rolls are one", () => {
  const got = bowlingScore("11 11 11 11 11");
  const want = 10;
  assertEquals(got, want);
});

Deno.test("that spare rolls and their bonus points are accounted for", () => {
  const got = bowlingScore("8/ 54 3/ 90");
  const want = 52;
  assertEquals(got, want);
});

Deno.test("that a strike and bonus points are accounted for", () => {
  const got = bowlingScore("X 54 X 90");
  const want = 56;
  assertEquals(got, want);
});

Deno.test(
  "that two subsequent strikes calculates bonus points correctly",
  () => {
    const got = bowlingScore("X X 54");
    const want = 53;
    assertEquals(got, want);
  }
);

Deno.test("that a 300 points is given for a perfect score", () => {
  const got = bowlingScore("X X X X X X X X X XXX");
  const want = 300;
  assertEquals(got, want);
});

Deno.test("that 100 points is given for a game of spares", () => {
  const got = bowlingScore("0/ 0/ 0/ 0/ 0/ 0/ 0/ 0/ 0/ 0/0");
  const want = 100;
  assertEquals(got, want);
})
