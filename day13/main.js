import fs from "fs";

const inputPath = "./day13/example-input";

export function partOne(input = null) {
  var input = input || inputPath;
  input = fs.readFileSync(input, "utf8");

  let puzzles = input
    .trim()
    .split("\n\n")
    .map((line) => line.split("\n"));

  let sum = 0;
  for (let puzzle of puzzles) {
    const { isVertical, count } = findVerticalOrHorizontalReflectionLine(puzzle);
    console.log({ isVertical, count });
    sum += isVertical ? count : 100 * count;
  }

  console.log({ day: 13, part: 1, value: sum });
}

export function partTwo(input = null) {
  var input = input || inputPath;
  input = fs.readFileSync(input, "utf8");
  var inputArray = input
    .trim()
    .split("\n")
    .map((line) => line.split(""));

  console.log({ day: 13, part: 2, value: "todo" });
}

function findVerticalOrHorizontalReflectionLine(puzzle) {
  let reflectionPoint = 0.5; // In-between 2 points.

  while (reflectionPoint > 0 && (reflectionPoint < puzzle[0].length - 1 || reflectionPoint < puzzle.length - 1)) {
    const lowIndex = Math.floor(reflectionPoint); // left for vertical, top for horizontal
    const highIndex = Math.ceil(reflectionPoint); // right for vertical, bottom for vertical
    // Horizontally...
    if (reflectionPoint < puzzle[0].length) {
      if (expandAndCheckAllLinesMatchVertically(lowIndex, highIndex, puzzle)) {
        // We've found the line! So we know it's (a) vertical and (b) at reflectionPoint. This means there are highIndex columns to the left of it
        return { isVertical: true, count: highIndex };
      }
    }

    // Vertically...
    if (reflectionPoint < puzzle.length) {
      if (expandAndCheckAllLinesMatchHorizontally(lowIndex, highIndex, puzzle)) {
        // We've found the line! So we know it's (a) horizontal and (b) at reflectionPoint. This means there are highIndex columns above of it
        return { isVertical: false, count: highIndex };
      }
    }
    reflectionPoint++;
  }

  throw Error("No Reflection found :c");
}

function expandAndCheckAllLinesMatchVertically(leftIndex, rightIndex, puzzle) {
  while (leftIndex >= 0 && rightIndex < puzzle[0].length) {
    let topIndex = 0;
    let bottomIndex = puzzle.length - 1;
    while (topIndex <= bottomIndex) {
      // Check all on the left index at the same height match all on the right index at the same height
      if (
        puzzle[topIndex][leftIndex] !== puzzle[topIndex][rightIndex] ||
        puzzle[bottomIndex][leftIndex] !== puzzle[bottomIndex][rightIndex]
      ) {
        return false;
      }

      topIndex++;
      bottomIndex--;
    }

    // Move them out one and check the next cols match
    leftIndex--;
    rightIndex++;
  }

  // If we've made it to here without an early return of false, they all match!
  return true;
}

/**
 * Horizontal is easier - each row is a string, so we can directly check the whole row
 */
function expandAndCheckAllLinesMatchHorizontally(topIndex, bottomIndex, puzzle) {
  while (topIndex >= 0 && bottomIndex < puzzle.length - 1) {
    if (puzzle[topIndex] !== puzzle[bottomIndex]) {
      return false;
    }

    // Move then out one and check the next rows match
    topIndex--;
    bottomIndex++;
  }

  // If we've made it to here without an early return of false, they all match!
  return true;
}
