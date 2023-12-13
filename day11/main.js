import fs from "fs";

const inputPath = "./day11/example-input";

export function partOne(input = null) {
  var input = input || inputPath;
  input = fs.readFileSync(input, "utf8");
  var inputArray = input
    .trim()
    .split("\n")
    .map((line) => line.split(""));

  var { inputArray, galaxyCoordinates } = handleExpansion(inputArray);

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
        Math.abs(galaxyCoordinates[i][1] - galaxiesToCompare[j][1]);
      console.log({ diff, i, j });
      totalSteps += diff;
    }
    galaxiesToCompare.shift(); // remove the current one
  }

  console.log({ totalSteps });

  console.log({ day: 11, part: 1, value: totalSteps });
}

function display(array) {
  console.log(array.map((line) => line.join("")).join("\n"));
}

function handleExpansion(inputArray) {
  // First let's handle the expansion.
  // Start with a list of all row/col indexes, and remove if we detect a galaxy
  let rowsWithoutGalaxies = new Set([...Array(inputArray.length).keys()]);
  let colsWithoutGalaxies = new Set([...Array(inputArray[0].length).keys()]);
  for (let i = 0; i < inputArray.length; i++) {
    for (let j = 0; j < inputArray[0].length; j++) {
      if (inputArray[i][j] === "#") {
        rowsWithoutGalaxies.delete(i);
        colsWithoutGalaxies.delete(j);
      }
    }
  }

  // Convert to arrays for ease of indexing
  rowsWithoutGalaxies = [...rowsWithoutGalaxies];
  colsWithoutGalaxies = [...colsWithoutGalaxies];

  // Now we need to squeeze in these extra rows/cols!
  // Rows are easier so start there:
  let insertionCount = 0;
  let emptyRow = [...Array(inputArray[0].length).keys()].map(() => ".");
  while (insertionCount < rowsWithoutGalaxies.length) {
    // The index is offset by the insertion count
    inputArray.splice(rowsWithoutGalaxies[insertionCount] + insertionCount, 0, [...emptyRow]);

    insertionCount++;
  }

  // Now do cols... bit trickier?
  insertionCount = 0;
  while (insertionCount < colsWithoutGalaxies.length) {
    for (let i = 0; i < inputArray.length; i++) {
      inputArray[i].splice(colsWithoutGalaxies[insertionCount] + insertionCount, 0, ".");
    }
    insertionCount++;
  }

  // Now get the galaxy coordinates. We could do a clever thing of deciding how much to add based on the above, but loop is easier
  let galaxyCoordinates = [];
  for (let i = 0; i < inputArray.length; i++) {
    for (let j = 0; j < inputArray[0].length; j++) {
      if (inputArray[i][j] === "#") {
        galaxyCoordinates.push([i, j]);
      }
    }
  }

  display(inputArray);

  return { inputArray, galaxyCoordinates };
}
