import { fromFileUrl, dirname } from "jsr:@std/path";

// --- SETUP ---
const text = await readFile();
const [leftList, rightList] = parseInput(text);

leftList.sort((a, b) => a - b);
rightList.sort((a, b) => a - b);

// --- PART 1 ---
let difference = 0;

for (let i = 0; i < leftList.length; i++) {
  difference += Math.abs(leftList[i] - rightList[i]);
}

console.log({ difference });

// --- PART 2 ---
let similarityScore = 0;

leftList.forEach((val) => {
  let qty = 0;

  for (let j = 0; j < rightList.length; j++) {
    if (val === rightList[j]) {
      qty++;
    }
    if (rightList[j] > val) {
      break;
    }
  }

  similarityScore += qty * val;
});

console.log({ similarityScore });

async function readFile(): string {
  let path = dirname(fromFileUrl(import.meta.url)) + "/example-input";
  if (Deno.args.length > 0) {
    // If we pass an argument, use this as the input
    path = Deno.args[0];
  }

  return await Deno.readTextFile(path);
}

function parseInput(input: string): [number[], number[]] {
  let leftList = [] as number[];
  let rightList = [] as number[];
  input.split("\n").forEach((line) => {
    let [left, right] = line.split("   ").map(Number);

    leftList.push(left);
    rightList.push(right);
  });

  return [leftList, rightList];
}
