import fs from "fs";

const inputPath = "./day11/example-input";

export function partOne(input = null) {
  var input = input || inputPath;
  input = fs.readFileSync(input, "utf8");
  var inputArray = input
    .trim()
    .split("\n")
    .map((line) => line.split(""));

  let { galaxyCoordinates, rowsWithoutGalaxies, colsWithoutGalaxies } = getOriginalGalaxyCoordinates(inputArray);

  const scale = 2;

  // // Formula for number of unique pairs is (n * (n-1))/2. where n is number of points
  // let numPoints = (galaxyCount * (galaxyCount - 1)) / 2;
  let galaxiesToCompare = [...galaxyCoordinates];
  let totalSteps = 0;
  for (let i = 0; i < galaxyCoordinates.length - 1; i++) {
    // Iterate through the remaining galaxies to compare
    for (let j = 1; j < galaxiesToCompare.length; j++) {
      // Number of steps is te sum of dy and dx
      let diff =
        Math.abs(galaxyCoordinates[i][0] - galaxiesToCompare[j][0]) +
        (scale - 1) *
          getNumberOfExpansionsBetween(rowsWithoutGalaxies, galaxyCoordinates[i][0], galaxiesToCompare[j][0]) +
        Math.abs(galaxyCoordinates[i][1] - galaxiesToCompare[j][1]) +
        (scale - 1) *
          getNumberOfExpansionsBetween(colsWithoutGalaxies, galaxyCoordinates[i][1], galaxiesToCompare[j][1]);

      totalSteps += diff;
    }
    galaxiesToCompare.shift(); // remove the current one
  }

  console.log({ day: 11, part: 1, value: totalSteps });
}

export function partTwo(input = null) {
  var input = input || inputPath;
  input = fs.readFileSync(input, "utf8");
  var inputArray = input
    .trim()
    .split("\n")
    .map((line) => line.split(""));

  /**
   * For this one, we need to not just expand the map as it's enormous. Instead we will:
   * 1. Get original galaxy coordinates and rows/cold without a map
   * 2. When computing differences:
   *     a. For rows, do |dx| + n*(scale-1), where n is number of rows BETWEEN without galaxies and scale is the gap between
   *     b. For columns, do |dy| + n*(scale-1), where n is number of rows BETWEEN without galaxies and scale is the gap between
   */

  let { galaxyCoordinates, rowsWithoutGalaxies, colsWithoutGalaxies } = getOriginalGalaxyCoordinates(inputArray);

  const scale = 1000000;

  // // Formula for number of unique pairs is (n * (n-1))/2. where n is number of points
  // let numPoints = (galaxyCount * (galaxyCount - 1)) / 2;
  let galaxiesToCompare = [...galaxyCoordinates];
  let totalSteps = 0;
  for (let i = 0; i < galaxyCoordinates.length - 1; i++) {
    // Iterate through the remaining galaxies to compare
    for (let j = 1; j < galaxiesToCompare.length; j++) {
      // Number of steps is te sum of dy and dx
      let diff =
        Math.abs(galaxyCoordinates[i][0] - galaxiesToCompare[j][0]) +
        (scale - 1) *
          getNumberOfExpansionsBetween(rowsWithoutGalaxies, galaxyCoordinates[i][0], galaxiesToCompare[j][0]) +
        Math.abs(galaxyCoordinates[i][1] - galaxiesToCompare[j][1]) +
        (scale - 1) *
          getNumberOfExpansionsBetween(colsWithoutGalaxies, galaxyCoordinates[i][1], galaxiesToCompare[j][1]);

      totalSteps += diff;
    }
    galaxiesToCompare.shift(); // remove the current one
  }

  console.log({ day: 11, part: 2, value: totalSteps });
}

function getOriginalGalaxyCoordinates(inputArray) {
  // Start with a list of all row/col indexes, and remove if we detect a galaxy
  let rowsWithoutGalaxies = new Set([...Array(inputArray.length).keys()]);
  let colsWithoutGalaxies = new Set([...Array(inputArray[0].length).keys()]);
  const galaxyCoordinates = [];
  for (let i = 0; i < inputArray.length; i++) {
    for (let j = 0; j < inputArray[0].length; j++) {
      if (inputArray[i][j] === "#") {
        galaxyCoordinates.push([i, j]);
        // Convert to arrays for ease of indexing
        rowsWithoutGalaxies.delete(i);
        colsWithoutGalaxies.delete(j);
      }
    }
  }

  rowsWithoutGalaxies = [...rowsWithoutGalaxies];
  colsWithoutGalaxies = [...colsWithoutGalaxies];

  return { galaxyCoordinates, rowsWithoutGalaxies, colsWithoutGalaxies };
}

function getNumberOfExpansionsBetween(rowOrColWithoutGalaxies, positionOne, positionTwo) {
  // Need to handle 3 cases: positions equal, one > two and two > one
  if (positionOne === positionTwo) {
    return 0;
  }
  if (positionOne > positionTwo) {
    return rowOrColWithoutGalaxies.filter((rowIndex) => rowIndex > positionTwo && rowIndex < positionOne).length;
  }
  if (positionOne < positionTwo) {
    return rowOrColWithoutGalaxies.filter((rowIndex) => rowIndex > positionOne && rowIndex < positionTwo).length;
  }
}