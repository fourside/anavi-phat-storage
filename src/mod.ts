import { parseArgs } from "./cli.ts";
import { parseInput } from "./parse_input.ts";
import { insert } from "./mongo.ts";

async function main(): Promise<void> {
  const parsed = parseArgs(Deno.args);
  if (parsed.exit) {
    Deno.exit(parsed.exitCode);
  }
  const { collection } = parsed;

  try {
    const json = parseInput(Deno.stdin);
    await insert(json, collection);
  } catch (error) {
    console.error(error);
    Deno.exit(-1);
  }
}

if (import.meta.main) {
  await main();
}
