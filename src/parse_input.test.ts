import {
  assertEquals,
  assertExists,
  assertThrows,
} from "https://deno.land/std@0.97.0/testing/asserts.ts";
import { parseInput } from "./parse_input.ts";

Deno.test({
  name: "parseInput parse JSON adding date field",
  fn: () => {
    // arrange
    const reader = mockStdin('{"hoge": 123}');

    // act
    const json = parseInput(reader);

    // assert
    assertEquals(json.hoge, 123);
    assertExists(json.date); // should be added date field
  },
});

Deno.test({
  name: "parseInput throw Error if not given json string",
  fn: () => {
    // arrange
    const reader = mockStdin('"hoge": 123');

    // act && assert
    assertThrows(() => parseInput(reader), Error);
  },
});

function mockStdin(input: string): typeof Deno.stdin {
  return {
    close: () => {},
    read: async (u: Uint8Array): Promise<number | null> => {
      new TextEncoder().encodeInto(input, u);
      return await Promise.resolve(input.length);
    },
    readSync: (u: Uint8Array): number | null => {
      new TextEncoder().encodeInto(input, u);
      return input.length;
    },
    rid: Math.random(),
  };
}
