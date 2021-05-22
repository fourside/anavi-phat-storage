import { parseInput } from "./parse_input.ts";

function main(): void {
  try {
    const json = parseInput(Deno.stdin);
    console.log(json);
  } catch (error) {
    console.error(error);
    Deno.exit(-1);
  }
}

if (import.meta.main) {
  main();
}
