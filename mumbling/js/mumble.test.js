import {
  assertEquals,
} from "https://deno.land/std/testing/asserts.ts";

import mumbleLetters from './mumble.js'

Deno.test("an empty string is returned when no parameter is passed", () => {
  const result = mumbleLetters()
  assertEquals(result, "")
})

Deno.test("that A is returned when a is the parameter passed", () => {
  const result = mumbleLetters("a");
  assertEquals(result, "A");
})

Deno.test("that A-Bb is returned when ab is the parameter passed", () => {
  const result = mumbleLetters("ab");
  assertEquals(result, "A-Bb");
})

Deno.test("that A-Bb-Ccc is returned when abC is the parameter passed", () => {
  const result = mumbleLetters("abc");
  assertEquals(result, "A-Bb-Ccc");
})

Deno.test("that mixed casing is handled: A-Bb-Ccc-Dddd is returned when abCD is the parameter passed", () => {
  const result = mumbleLetters("abCD");
  assertEquals(result, "A-Bb-Ccc-Dddd");
})