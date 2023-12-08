import fs from "fs";

const inputPath = "./day3/example-input";

export function partOne() {
  var input = fs.readFileSync(inputPath, "utf8");
  var inputArray = input.trim().split("\n");

  // Each line has numbers or symbols. We find each symbol, and then find ALL the numbers next to it (including diagonals).
  // So each line is the same length, so we can use [i][j] syntax to get the symbol at a given position.

  // First we will loop through to find all of the i,j coordinates of the numbers we need.
  // Then we will extract these. This isn' easy because they may span multiple i coordinates.
  // Because of this, maybe we should find the LHS-most number for each, then unique them, then extract them?

  const coordinatesOfNumbers = [];

  // The coordinates surrounding any symbol.
  let coordinatesToCheck = [];

  for (let rowIndex = 0; rowIndex < inputArray.length; rowIndex++) {
    let line = inputArray[rowIndex];
    for (let j = 0; j < line.length; j++) {
      // If is any of *,#,+,$, then we need to find the numbers around it:
      if (["*", "#", "+", "$"].includes(line[j])) console.log("match");
      coordinatesToCheck = [
        ...coordinatesToCheck,
        ...getCoordinatesAroundPoint(rowIndex, j, inputArray.length, line.length),
      ];
    }
  }

  console.log(coordinatesToCheck);

  function getCoordinatesAroundPoint(i, j, maxI, maxJ) {
    const coordinates = [];
    // 8 Possible coordinates around a point, loop these and check if they are valid:
    for (let x = i - 1; x <= i + 1; x++) {
      for (let y = j - 1; y <= j + 1; y++) {
        // If the coordinate is valid, add it to the list:
        if (x >= 0 && x < maxI && y >= 0 && y < maxJ) {
          coordinates.push([x, y]);
        }
      }
    }
    return coordinates;
  }
}
