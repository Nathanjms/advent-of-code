import fs from "fs";

const inputPath = "./day14/example-input";

export function partOne(input = null) {
  var input = input || inputPath;
  input = fs.readFileSync(input, "utf8");
  var dish = input
    .trim()
    .split("\n")
    .map((line) => line.split(""));

  // Store the columns outside of any loops. Cant think of a nicer JS way to 'flip' i and j.
  const dishColumns = [];
  for (let i = 0; i < dish[0].length; i++) {
    dishColumns.push(dish.map((line) => line[i]));
  }

  // We can simply go through, and for each 0, work out where it ends up, and add that.
  let totalLoad = 0;
  for (let i = 0; i < dish.length; i++) {
    for (let j = 0; j < dish[i].length; j++) {
      if (dish[i][j] === "O") {
        // console.log({ i, j, newLoc: getNewRowIndex(i, j) });
        totalLoad += dish.length - getNewRowIndex(i, j);
      }
    }
  }

  console.log({ day: 14, part: 1, value: totalLoad });

  // Instead of 'tilting the dish north' and redrawing the dish, we can calculate, for each 0, it's new position by doing:
  function getNewRowIndex(rowIndex, colIndex) {
    if (rowIndex === 0) return 0; // Can't get any better
    const column = dishColumns[colIndex].slice(0, rowIndex + 1); // Get the column from the top down to the point we are looking at

    // Check if there is a hash (blocker)
    const hashIndex = column.lastIndexOf("#");

    // No hash - just offset by the number of O's (do not need to add one because this rock is included)
    if (hashIndex === -1) {
      return column.filter((val) => val === "O").length - 1;
    }
    // Is a hash - offset by the number of O's below it
    return hashIndex + column.slice(hashIndex).filter((val) => val === "O").length;
  }
}

export function partTwo(input = null) {
  var input = input || inputPath;
  input = fs.readFileSync(input, "utf8");
  var inputArray = input
    .trim()
    .split("\n")
    .map((line) => line.split(""));

  console.log({ day: 14, part: 2, value: "todo" });
}
