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
  var dish = input
    .trim()
    .split("\n")
    .map((line) => line.split(""));

  // Store the columns outside of any loops. Cant think of a nicer JS way to 'flip' i and j.
  const dishColumns = [];
  for (let i = 0; i < dish[0].length; i++) {
    dishColumns.push(dish.map((line) => line[i]));
  }

  /**
   * Pt 2 wants 1,000,000,000 (billion) iterations of 'cycles' of the dish. A cycle is defined as:
   * Tilt the dish north, west, south, east.
   * We probably wont want to do this 1,000,000,000 times, so we need to find a pattern.
   * These are done in cycles, and so we can simply find the number of times until we are back at the start.
   * Then the modulo of the number of cycles we want, and the number of cycles until we are back at the start, will give us the number of cycles we need to do!
   */

  // Store the original dish, being careful not to mutate it.
  const originalDish = cloneArray(dish);
  let cache = {};
  let cycle = 0;
  while (cycle < 1_000_000_000) {
    dish = handleCycle(dish);
    cycle++;
    if (dish.join("") === originalDish.join("")) {
      break;
    }
  }
  console.log({ cycle });

  // We can simply go through, and for each 0, work out where it ends up, and add that.
  let totalLoad = 0;
  for (let i = 0; i < dish.length; i++) {
    for (let j = 0; j < dish[i].length; j++) {
      if (dish[i][j] === "O") {
        totalLoad += dish.length - getNewRowIndex(i, j);
      }
    }
  }

  console.log({ day: 14, part: 2, value: totalLoad });

  function handleCycle(dish) {
    const key = dish.join("");
    if (key in cache) {
      return cache[key];
    }

    // Now the hard bit! We did pt 1 without redrawing the dish, so we need to now do that I think :(
    display(dish);
    console.log("Tilting dish north");
    console.log("");
    dish = tiltDishNorth(dish);
    display(dish);
    console.log("Tilting dish west");
    console.log("");
    dish = tiltDishWest(dish);
    display(dish);
    console.log("Tilting dish south");
    console.log("");
    dish = tiltDishSouth(dish);
    display(dish);
    console.log("Tilting dish east");
    console.log("");
    dish = tiltDishEast(dish);
    display(dish);
    cache[key] = dish;
    return dish;
  }

  function tiltDishNorth(dish) {
    // Go through each one, and get where it ends up. Replace the old spot with a '.' and the new one with an 'O'.
    let newDish = cloneArray(dish);

    for (let i = 0; i < dish.length; i++) {
      for (let j = 0; j < dish[i].length; j++) {
        if (dish[i][j] === "O") {
          const newRowIndex = getNewNorthRowIndex(i, j, dish);
          newDish[i][j] = ".";
          newDish[newRowIndex][j] = "O";
        }
      }
    }

    return newDish;
  }

  function tiltDishWest(dish) {
    // Go through each one, and get where it ends up. Replace the old spot with a '.' and the new one with an 'O'.
    let newDish = cloneArray(dish);

    for (let i = 0; i < dish.length; i++) {
      for (let j = 0; j < dish[i].length; j++) {
        if (dish[i][j] === "O") {
          const newColIndex = getNewWestColIndex(i, j, dish);
          if (newColIndex !== j) {
            newDish[i][j] = ".";
            newDish[i][newColIndex] = "O";
          }
        }
      }
    }

    return newDish;
  }

  function tiltDishSouth(dish) {
    // Go through each one, and get where it ends up. Replace the old spot with a '.' and the new one with an 'O'.
    let newDish = cloneArray(dish);

    for (let i = 0; i < dish.length; i++) {
      for (let j = 0; j < dish[i].length; j++) {
        if (dish[i][j] === "O") {
          const newRowIndex = getNewSouthRowIndex(i, j, dish);
          newDish[i][j] = ".";
          newDish[newRowIndex][j] = "O";
        }
      }
    }

    return newDish;
  }

  function tiltDishEast(dish) {
    // Go through each one, and get where it ends up. Replace the old spot with a '.' and the new one with an 'O'.
    let newDish = cloneArray(dish);

    for (let i = 0; i < dish.length; i++) {
      for (let j = 0; j < dish[i].length; j++) {
        if (dish[i][j] === "O") {
          const newColIndex = getNewEastColIndex(i, j, dish);
          newDish[i][j] = ".";
          newDish[i][newColIndex] = "O";
          display(newDish);
        }
      }
    }

    return newDish;
  }

  function getNewNorthRowIndex(rowIndex, colIndex, dish) {
    if (rowIndex === 0) return 0; // Can't get any better
    const column = dish.map((line) => line[colIndex]).slice(0, rowIndex + 1); // Get the column from the top down to the point we are looking at

    // Check if there is a hash (blocker)
    const hashIndex = column.lastIndexOf("#");

    // No hash - just offset by the number of O's (do not need to add one because this rock is included)
    if (hashIndex === -1) {
      return column.filter((val) => val === "O").length - 1;
    }
    // Is a hash - offset by the number of O's below it
    return hashIndex + column.slice(hashIndex).filter((val) => val === "O").length;
  }

  function getNewWestColIndex(rowIndex, colIndex, dish) {
    if (colIndex === 0) return 0; // Can't get any better
    const row = dish[rowIndex].slice(0, colIndex + 1); // Get the row from the left to the point we are looking at

    // Check if there is a hash (blocker)
    const hashIndex = row.lastIndexOf("#");

    // No hash - just offset by the number of O's (do not need to add one because this rock is included)
    if (hashIndex === -1) {
      return row.filter((val) => val === "O").length - 1;
    }
    // Is a hash - offset by the number of O's below it
    return hashIndex + row.slice(hashIndex).filter((val) => val === "O").length;
  }

  /**
   * South - means that the rowIndex is INCREASING
   * @returns
   */
  function getNewSouthRowIndex(rowIndex, colIndex, dish) {
    if (rowIndex === dish.length - 1) return dish.length - 1; // Can't get any better, already on the edge
    const column = dish.map((line) => line[colIndex]).slice(rowIndex); // Get from point to the end

    // Check if there is a hash (blocker)
    const hashIndex = column.indexOf("#");

    // No hash - just offset by the number of O's AHEAD of it
    if (hashIndex === -1) {
      return dish.length - column.filter((val) => val === "O").length;
    }
    // Start at the rowIndex and we can add as many gaps up to the hashIndex
    return rowIndex + (hashIndex - column.slice(0, hashIndex).filter((val) => val === "O").length);
  }

  function getNewEastColIndex(rowIndex, colIndex, dish) {
    if (colIndex === dish[rowIndex].length - 1) return dish[rowIndex].length - 1; // Can't get any better
    const row = dish[rowIndex].slice(colIndex);

    // Check if there is a hash (blocker)
    const hashIndex = row.indexOf("#");

    // No hash - go to the end then offset by the number of 0's (note row includes itself)
    if (hashIndex === -1) {
      return dish[rowIndex].length - row.filter((val) => val === "O").length;
    }
    // Is a hash - move it left up to hash for as many places as there are available (ie. no 0s)
    return colIndex + ((hashIndex - colIndex) - row.slice(colIndex, hashIndex).filter((val) => val === "O").length);
  }
}

function display(dish) {
  console.log(dish.map((line) => line.join("")).join("\n"));
  console.log();
}

// This is grim but I can't stop the array being mutated without it :c
function cloneArray(array) {
  return JSON.parse(JSON.stringify(array));
}
