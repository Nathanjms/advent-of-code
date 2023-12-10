import fs from "fs";

const inputPath = "./day4/input";

export function partOne() {
  var input = fs.readFileSync(inputPath, "utf8");
  var inputArray = input.trim().split("\n");

  let sum = 0;

  inputArray.forEach((line) => {
    const { winningNumbers, userNumbers } = extractLineInfo(line);

    // Get the number of matching numbers:
    const matchingNumbers = new Set([...winningNumbers].filter((x) => userNumbers.has(x)));

    if (matchingNumbers.size > 0) {
      // multiply the sum by 2 for each matching number, except the first one. 2^0 = 1
      sum += 2 ** (matchingNumbers.size - 1);
    }
  });

  console.log({ day: 4, part: 1, value: sum });
}

export function partTwo() {
  var input = fs.readFileSync(inputPath, "utf8");
  var inputArray = input.trim().split("\n");

  // Build quantity for each line, each starting at 1:
  let quantity = inputArray.map((line) => 1);

  for (let i = 0; i < inputArray.length; i++) {
    const { winningNumbers, userNumbers } = extractLineInfo(inputArray[i]);

    // Get the number of matching numbers. We only need to do this once per line
    const matchingNumberQuantity = new Set([...winningNumbers].filter((x) => userNumbers.has(x))).size;

    // For the quantity of this line, we need to add cards to the next x, where x is the matching number quantity
    for (let j = 1; j <= matchingNumberQuantity; j++) {
      // Add the quantity of this line to the next line(s)
      quantity[i + j] += quantity[i];
    }
  }

  console.log({ day: 4, part: 2, value: quantity.reduce((a, b) => a + b, 0) });
}

function extractLineInfo(line) {
  const regex = /Card\s+(\d+):(.*)\|(.*)/;
  const match = line.match(regex);
  return {
    id: match[1],
    winningNumbers: new Set(
      match[2]
        .trim()
        .split(" ")
        .filter((number) => number !== "")
        .map((number) => Number(number))
    ),
    userNumbers: new Set(
      match[3]
        .trim()
        .split(" ")
        .filter((number) => number !== "")
        .map((number) => Number(number))
    ),
  };
}
