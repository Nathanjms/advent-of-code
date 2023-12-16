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
  const originalDishKey = dish.join("");
  let dishArrangementCycle = {
    originalDishKey: 0,
  };
  let cycle = 0;
  while (cycle < 1_000_000_000) {
    const key = dish.join("");
    if (key in dishArrangementCycle) {
      break;
    }
    dishArrangementCycle[key] = cycle;
    dish = handleCycle(dish, cycle);
    cycle++;
  }

  // After the first x (dishArrangementCycle[dish.join("")]), it repeats every repeatCycleNumber - x times

  // So we now run it 1,000,000,000 % repeatCycleNumber - x times more times!
  let timesToRun =
    (1_000_000_000 - dishArrangementCycle[dish.join("")]) % (cycle - dishArrangementCycle[dish.join("")]);
  for (let i = 0; i < timesToRun; i++) {
    dish = handleCycle(dish, cycle);
  }

  let totalLoad = 0;
  for (let i = 0; i < dish.length; i++) {
    for (let j = 0; j < dish[i].length; j++) {
      if (dish[i][j] === "O") {
        totalLoad += dish.length - i;
      }
    }
  }

  console.log({ day: 14, part: 2, value: totalLoad });

  function handleCycle(dish) {
    // Now the hard bit! We did pt 1 without redrawing the dish, so we need to now do that I think :(
    dish = tiltDishNorth(dish);
    dish = tiltDishWest(dish);
    dish = tiltDishSouth(dish);
    dish = tiltDishEast(dish);
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
    // West can just be an array flip of east, then move the index back!

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

    for (let i = dish.length - 1; i >= 0; i--) {
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
      for (let j = dish[i].length - 1; j >= 0; j--) {
        if (dish[i][j] === "O") {
          const newColIndex = getNewEastColIndex(i, j, dish);
          newDish[i][j] = ".";
          newDish[i][newColIndex] = "O";
        }
      }
    }

    return newDish;
  }

  function getNewNorthRowIndex(rowIndex, colIndex, dish, flip = false) {
    if (rowIndex === 0) return 0; // Can't get any better
    let column = dish.map((line) => line[colIndex]);
    if (flip) {
      column = column.slice().reverse();
    }
    column = column.slice(0, rowIndex + 1); // Get the column from the top down to the point we are looking at

    // Check if there is a hash (blocker)
    const hashIndex = column.lastIndexOf("#");

    // No hash - just offset by the number of O's (do not need to add one because this rock is included)
    if (hashIndex === -1) {
      return column.filter((val) => val === "O").length - 1;
    }
    // Is a hash - offset by the number of O's below it
    return hashIndex + column.slice(hashIndex).filter((val) => val === "O").length;
  }

  function getNewWestColIndex(rowIndex, colIndex, dish, flip = false) {
    if (colIndex === 0) return 0; // Can't get any better
    let row = dish[rowIndex];
    if (flip) {
      row = row.slice().reverse();
    }
    row = row.slice(0, colIndex + 1); // Get the row from the left to the point we are looking at

    // Check if there is a hash (blocker)
    const hashIndex = row.lastIndexOf("#");

    // No hash - just offset by the number of O's (do not need to add one because this rock is included)
    if (hashIndex === -1) {
      return row.filter((val) => val === "O").length - 1;
    }
    // Is a hash - offset by the number of O's to the right of it
    return hashIndex + row.slice(hashIndex).filter((val) => val === "O").length;
  }

  function getNewSouthRowIndex(rowIndex, colIndex, dish) {
    // South can just be an array flip of North, then move the index back!
    return dish.length - 1 - getNewNorthRowIndex(dish.length - 1 - rowIndex, colIndex, dish, true);
  }

  function getNewEastColIndex(rowIndex, colIndex, dish) {
    // East can just be an array flip of West, then move the index back!
    return dish[rowIndex].length - 1 - getNewWestColIndex(rowIndex, dish[rowIndex].length - 1 - colIndex, dish, true);
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
