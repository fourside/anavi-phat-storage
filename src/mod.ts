import { parseInput } from "./parse_input.ts";
import { insert } from "./mongo.ts";

async function main(): Promise<void> {
  try {
    const json = parseInput(Deno.stdin);
    console.log(json);
    await insert(json);
  } catch (error) {
    console.error(error);
    Deno.exit(-1);
  }
}

if (import.meta.main) {
  await main();
}
