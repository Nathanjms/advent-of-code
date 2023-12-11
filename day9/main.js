import fs from "fs";

const inputPath = "./day9/example-input";

export function partOne(input = null) {
  input = input || inputPath;
  var input = fs.readFileSync(input, "utf8");
  var inputArray = input.trim().split("\n");

  let sum = 0;

  for (const line of inputArray) {
    let differences = [line.split(" ").map(Number)];
    do {
      differences.push(buildNextLine(differences));
    } while (lastRowIsNotAllZeros(differences));

    // Now we work from the 2nd-from-bottom sequence, and add on an extra element, which is the current last element + the value in the row 'below':
    for (let i = differences.length - 2; i >= 0; i--) {
      const thisLine = differences[i];
      const lastLine = differences[i + 1];
      thisLine.push(thisLine[thisLine.length - 1] + lastLine[lastLine.length - 1]);
    }

    // Add the last element in the first line of differences (the original) to the sum:
    sum += differences[0][differences[0].length - 1];
  }

  console.log({ day: 9, part: 1, value: sum });
}

export function partTwo(input = null) {
  input = input || inputPath;
  var input = fs.readFileSync(input, "utf8");
  var inputArray = input.trim().split("\n");

  let sum = 0;

  for (const line of inputArray) {
    let differences = [line.split(" ").map(Number)];
    do {
      differences.push(buildNextLine(differences));
    } while (lastRowIsNotAllZeros(differences));

    // Now we work from the 2nd-from-bottom sequence, and add on an extra element AT THE START, which is the current first element - the value in the row 'below''s first element:
    for (let i = differences.length - 2; i >= 0; i--) {
      const thisLine = differences[i];
      const lastLine = differences[i + 1];
      thisLine.unshift(thisLine[0] - lastLine[0]);
    }

    // Add the first element in the first line of differences (the original) to the sum:
    sum += differences[0][0];
  }

  console.log({ day: 9, part: 2, value: sum });
}

/**
 * Builds the next line of differences, by taking the difference between each element, then finally removing the last element
 */
function buildNextLine(differences) {
  return differences[differences.length - 1].map((_, i, arr) => arr[i + 1] - arr[i]).slice(0, -1);
}

function lastRowIsNotAllZeros(differences) {
  return differences[differences.length - 1].some((item) => item !== 0);
}
