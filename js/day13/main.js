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
    sum += isVertical ? count : 100 * count;
  }

  console.log({ day: 13, part: 1, value: sum });
}

export function partTwo(input = null) {
  var input = input || inputPath;
  input = fs.readFileSync(input, "utf8");

  let puzzles = input
    .trim()
    .split("\n\n")
    .map((line) => line.split("\n"));

  let sum = 0;
  for (let puzzle of puzzles) {
    const { isVertical, count } = findVerticalOrHorizontalReflectionLineWithSmudge(puzzle);
    sum += isVertical ? count : 100 * count;
  }

  console.log({ day: 13, part: 2, value: sum });
}

function findVerticalOrHorizontalReflectionLine(puzzle) {
  let reflectionPoint = 0.5; // In-between 2 points.

  let validVerticalReflectionPoints = [];
  let validHorizontalReflectionPoints = [];

  while (reflectionPoint > 0 && (reflectionPoint < puzzle[0].length - 1 || reflectionPoint < puzzle.length - 1)) {
    const lowIndex = Math.floor(reflectionPoint); // left for vertical, top for horizontal
    const highIndex = Math.ceil(reflectionPoint); // right for vertical, bottom for vertical
    // Horizontally...
    if (reflectionPoint < puzzle[0].length - 1) {
      if (expandAndCheckAllLinesMatchVertically(lowIndex, highIndex, puzzle)) {
        // We've found the line! So we know it's (a) vertical and (b) at reflectionPoint. This means there are highIndex columns to the left of it
        validVerticalReflectionPoints.push(highIndex);
      }
    }

    // Vertically...
    if (reflectionPoint < puzzle.length - 1) {
      if (expandAndCheckAllLinesMatchHorizontally(lowIndex, highIndex, puzzle)) {
        // We've found the line! So we know it's (a) horizontal and (b) at reflectionPoint. This means there are highIndex columns above of it
        validHorizontalReflectionPoints.push(highIndex);
      }
    }
    reflectionPoint++;
  }

  if (validVerticalReflectionPoints.length) {
    return { isVertical: true, count: validVerticalReflectionPoints[validVerticalReflectionPoints.length - 1] };
  }

  if (validHorizontalReflectionPoints.length) {
    return { isVertical: false, count: validHorizontalReflectionPoints[validHorizontalReflectionPoints.length - 1] };
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
  while (topIndex >= 0 && bottomIndex < puzzle.length) {
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

function findVerticalOrHorizontalReflectionLineWithSmudge(puzzle) {
  let reflectionPoint = 0.5; // In-between 2 points.

  let validVerticalReflectionPoints = [];
  let validHorizontalReflectionPoints = [];

  while (reflectionPoint > 0 && (reflectionPoint < puzzle[0].length - 1 || reflectionPoint < puzzle.length - 1)) {
    const lowIndex = Math.floor(reflectionPoint); // left for vertical, top for horizontal
    const highIndex = Math.ceil(reflectionPoint); // right for vertical, bottom for vertical
    // Horizontally...
    if (reflectionPoint < puzzle[0].length - 1) {
      if (expandAndCheckAllLinesMatchVerticallyWithSmudge(lowIndex, highIndex, puzzle)) {
        // We've found the line! So we know it's (a) vertical and (b) at reflectionPoint. This means there are highIndex columns to the left of it
        validVerticalReflectionPoints.push(highIndex);
      }
    }

    // Vertically...
    if (reflectionPoint < puzzle.length - 1) {
      if (expandAndCheckAllLinesMatchHorizontallyWithSmudge(lowIndex, highIndex, puzzle)) {
        // We've found the line! So we know it's (a) horizontal and (b) at reflectionPoint. This means there are highIndex columns above of it
        validHorizontalReflectionPoints.push(highIndex);
      }
    }
    reflectionPoint++;
  }

  if (validVerticalReflectionPoints.length) {
    return { isVertical: true, count: validVerticalReflectionPoints[validVerticalReflectionPoints.length - 1] };
  }

  if (validHorizontalReflectionPoints.length) {
    return { isVertical: false, count: validHorizontalReflectionPoints[validHorizontalReflectionPoints.length - 1] };
  }

  throw Error("No Reflection found :c");
}

/**
 * For a given 'line' (left and right index next to it), check all points on it, and then 'expand' left and right and look again.
 * To deal with the smudge, we are after the one(s) that has ONE incorrect value but ALL OTHERS match.
 */
function expandAndCheckAllLinesMatchVerticallyWithSmudge(leftIndex, rightIndex, puzzle) {
  let incorrectMatchCount = 0;
  while (leftIndex >= 0 && rightIndex < puzzle[0].length) {
    let topIndex = 0;
    let bottomIndex = puzzle.length - 1;
    // We now need to handle one '.' being a '#' or vice-versa. We will check for when lines match ALL BUT ONE instead of ALL
    while (topIndex <= bottomIndex) {
      if (puzzle[topIndex][leftIndex] !== puzzle[topIndex][rightIndex]) {
        incorrectMatchCount++;
      }

      // Only do 2nd comparison when the indexes are not equal to avoid duplicating results
      if (
        topIndex !== bottomIndex &&
        leftIndex !== rightIndex &&
        puzzle[bottomIndex][leftIndex] !== puzzle[bottomIndex][rightIndex]
      ) {
        incorrectMatchCount++;
      }

      if (incorrectMatchCount > 1) {
        return false;
      }

      topIndex++;
      bottomIndex--;
    }

    // Move them out one and check the next cols match
    leftIndex--;
    rightIndex++;
  }

  return incorrectMatchCount === 1;
}

/**
 * Horizontal is no longer easier :c
 */
function expandAndCheckAllLinesMatchHorizontallyWithSmudge(topIndex, bottomIndex, puzzle) {
  let incorrectMatchCount = 0;
  while (topIndex >= 0 && bottomIndex < puzzle.length) {
    let leftIndex = 0;
    let rightIndex = puzzle[0].length - 1;
    while (leftIndex <= rightIndex) {
      if (puzzle[topIndex][leftIndex] !== puzzle[bottomIndex][leftIndex]) {
        incorrectMatchCount++;
      }

      // Only do 2nd comparison when the indexes are not equal to avoid duplicating results
      if (
        leftIndex !== rightIndex &&
        topIndex !== bottomIndex &&
        puzzle[topIndex][rightIndex] !== puzzle[bottomIndex][rightIndex]
      ) {
        incorrectMatchCount++;
      }

      if (incorrectMatchCount > 1) {
        return false;
      }

      // Increment both
      leftIndex++;
      rightIndex--;
    }

    // Move then out one and check the next rows match
    topIndex--;
    bottomIndex++;
  }

  // If we've made it to here without an early return of false, they all match!
  return incorrectMatchCount === 1;
}
