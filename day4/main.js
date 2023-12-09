import fs from "fs";
import CardLine from "./CardLine.js";

const inputPath = "./day4/input";

export function partOne() {
  var input = fs.readFileSync(inputPath, "utf8");
  var inputArray = input.trim().split("\n");

  let sum = 0;

  inputArray.forEach((line) => {
    const cardLine = new CardLine(line);
    const cardId = cardLine.getId();
    const winningNumbers = cardLine.getWinningNumbers();
    const userNumbers = cardLine.getUserNumbers();

    // Get the number of matching numbers:
    const matchingNumbers = new Set([...winningNumbers].filter((x) => userNumbers.has(x)));

    if (matchingNumbers.size > 0) {
      sum += 2 ** (matchingNumbers.size - 1);
    }
  });

  console.log("sum :", sum);
}

export function partTwo() {
  var input = fs.readFileSync(inputPath, "utf8");
  var inputArray = input.trim().split("\n");

  let sum = 0;

  console.log("sum :", sum);
}
