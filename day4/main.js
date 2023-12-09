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

  console.log("sum :", sum);
}

export function partTwo() {
  var input = fs.readFileSync(inputPath, "utf8");
  var inputArray = input.trim().split("\n");

  let sum = 0;

  console.log("sum :", sum);
}
