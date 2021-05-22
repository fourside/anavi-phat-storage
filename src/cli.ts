import { parse } from "https://deno.land/std@0.97.0/flags/mod.ts";

type ReturnValue = {
  exit: true;
  exitCode: 0 | -1;
} | {
  exit: false;
  collection: string;
};

export function parseArgs(args: string[]): ReturnValue {
  const parsed = parse(args);
  if (!parsed.collection) {
    console.log("pass --collection");
    return { exit: true, exitCode: -1 };
  }
  return { exit: false, collection: parsed.collection };
}
