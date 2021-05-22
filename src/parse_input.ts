type Stdin = typeof Deno.stdin;

export function parseInput(
  input: Stdin,
): { [key: string]: string | number; data: string } {
  const buffer = new Uint8Array(128);
  const readSize = input.readSync(buffer);
  if (readSize === null) {
    throw new Error("pass JSON to stdin");
  }

  const readData = new TextDecoder().decode(buffer);
  const json = JSON.parse(readData.substring(0, readSize));
  return {
    ...json,
    date: new Date().toJSON(),
  };
}
