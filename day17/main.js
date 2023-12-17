import fs from "fs";

const inputPath = "./day17/example-input";

const DIRECTIONS = {
  LEFT: "left",
  RIGHT: "right",
  UP: "up",
  DOWN: "down",
};

export function partOne(input = null) {
  var input = input || inputPath;
  input = fs.readFileSync(input, "utf8");
  var grid = input
    .trim()
    .split("\n")
    .map((line) => line.split(""));

  const queue = [{ coordinate: [0, 0], heatLoss: 0, currDirection: null, numSteps: 0 }];

  const endCoordinate = [grid.length - 1, grid[grid.length - 1].length - 1];

  // Seen coordinates with a combo of i,j,newI,newJ,n
  let seen = new Set();

  const maxRun = 3;
  let minHeatLoss = Infinity;

  while (queue.length) {
    // Get the one with the lowest heat loss
    let { coordinate, heatLoss, currDirection, numSteps } = queue.sort((a, b) => a.heatLoss - b.heatLoss).shift();
    // We've found it!
    if (coordinate.join(",") === endCoordinate.join(",")) {
      minHeatLoss = heatLoss;
      break;
    }

    if (seen.has(cacheKey({ coordinate, currDirection, numSteps }))) {
      // It's been removed from the queue, so we just skip it!
      continue;
    }

    seen.add(cacheKey({ coordinate, currDirection, numSteps }));

    /* Now to get the next possible coordinates! */
    // Case 1: We are stepping in a direction, and have not gone up to three yet!
    if (numSteps < maxRun && currDirection) {
      let [newI, newJ] = step(currDirection, coordinate[0], coordinate[1]);

      // Add if inside the grid
      if (newI >= 0 && newI < grid.length && newJ >= 0 && newJ < grid[0].length) {
        queue.push({
          coordinate: [newI, newJ],
          heatLoss: Number(heatLoss) + Number(gridVal(newI, newJ)),
          currDirection,
          numSteps: numSteps + 1,
        });
      }
    }

    // Case2: We can change direction too
    for (const direction of Object.values(DIRECTIONS)) {
      // We can't go opposite if we are moving
      if (currDirection && direction === getOppositeDirection(currDirection)) {
        continue;
      }
      // We handle stepping forwards above, so skip that too
      if (currDirection === direction) {
        continue;
      }

      let [newI, newJ] = step(direction, coordinate[0], coordinate[1]);

      if (newI >= 0 && newI < grid.length && newJ >= 0 && newJ < grid[0].length) {
        queue.push({
          coordinate: [newI, newJ],
          heatLoss: Number(heatLoss) + Number(gridVal(newI, newJ)),
          currDirection: direction,
          numSteps: 1, // First step in this new direction
        });
      }
    }
  }

  console.log({ day: 17, part: 1, value: minHeatLoss });

  function gridVal(i, j) {
    return grid[i][j];
  }
}

export function partTwo(input = null) {
  var input = input || inputPath;
  input = fs.readFileSync(input, "utf8");
  var grid = input
    .trim()
    .split("\n")
    .map((line) => line.split(""));

  const queue = [{ coordinate: [0, 0], heatLoss: 0, currDirection: null, numSteps: 0 }];

  const endCoordinate = [grid.length - 1, grid[grid.length - 1].length - 1];

  // Seen coordinates with a combo of i,j,newI,newJ,n
  let seen = new Set();

  const maxRun = 10;
  const minRun = 4;
  let minHeatLoss = Infinity;

  while (queue.length) {
    // Get the one with the lowest heat loss
    let { coordinate, heatLoss, currDirection, numSteps } = queue.shift();
    console.log({ heatLoss });
    // We've found it!
    if (coordinate.join(",") === endCoordinate.join(",") && numSteps >= 4) {
      minHeatLoss = heatLoss;
      break;
    }

    if (seen.has(cacheKey({ coordinate, currDirection, numSteps }))) {
      // It's been removed from the queue, so we just skip it!
      continue;
    }

    seen.add(cacheKey({ coordinate, currDirection, numSteps }));

    /* Now to get the next possible coordinates! */
    // Case 1: We are stepping in a direction, and need to be less than 10 not gone up to three yet!
    if (numSteps < maxRun && currDirection) {
      let [newI, newJ] = step(currDirection, coordinate[0], coordinate[1]);

      // Add if inside the grid
      if (newI >= 0 && newI < grid.length && newJ >= 0 && newJ < grid[0].length) {
        insertIntoPosition({
          coordinate: [newI, newJ],
          heatLoss: Number(heatLoss) + Number(gridVal(newI, newJ)),
          currDirection,
          numSteps: numSteps + 1,
        });
      }
    }

    // Case2: We can change direction too, but only after 4 steps or if we're not moving
    if (numSteps >= minRun || !currDirection) {
      for (const direction of Object.values(DIRECTIONS)) {
        // We can't go opposite if we are moving
        if (currDirection && direction === getOppositeDirection(currDirection)) {
          continue;
        }
        // We handle stepping forwards above, so skip that too
        if (currDirection === direction) {
          continue;
        }

        let [newI, newJ] = step(direction, coordinate[0], coordinate[1]);

        if (newI >= 0 && newI < grid.length && newJ >= 0 && newJ < grid[0].length) {
          insertIntoPosition({
            coordinate: [newI, newJ],
            heatLoss: Number(heatLoss) + Number(gridVal(newI, newJ)),
            currDirection: direction,
            numSteps: 1, // First step in this new direction
          });
        }
      }
    }
  }

  console.log({ day: 17, part: 2, value: minHeatLoss });

  function gridVal(i, j) {
    return grid[i][j];
  }

  function insertIntoPosition(newVal) {
    let heatLoss = newVal.heatLoss;
    if (!queue.length || queue[queue.length - 1].heatLoss < heatLoss) {
      queue.push(newVal);
      return;
    }
    if (heatLoss < queue[0].heatLoss) {
      queue.unshift(newVal);
      return;
    }
    for (let i = 1; i < queue.length; i++) {
      if (queue[i].heatLoss > heatLoss) {
        queue.splice(i - 1, 0, newVal);
        return;
      }
    }

    queue.push(newVal);
    return;
  }
}

function cacheKey(vals) {
  let key = "";
  for (const val of Object.values(vals)) {
    key += val + "-";
  }
  return key;
}
function getOppositeDirection(direction) {
  if (direction === DIRECTIONS.LEFT) return DIRECTIONS.RIGHT;
  if (direction === DIRECTIONS.RIGHT) return DIRECTIONS.LEFT;
  if (direction === DIRECTIONS.UP) return DIRECTIONS.DOWN;
  if (direction === DIRECTIONS.DOWN) return DIRECTIONS.UP;
}

function step(direction, i, j) {
  if (direction === DIRECTIONS.RIGHT) {
    return [i, j + 1];
  }
  if (direction === DIRECTIONS.LEFT) {
    return [i, j - 1];
  }
  if (direction === DIRECTIONS.DOWN) {
    return [i + 1, j];
  }
  if (direction === DIRECTIONS.UP) {
    return [i - 1, j];
  }
}
