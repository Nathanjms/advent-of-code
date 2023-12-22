import fs from "fs";

const inputPath = "./day21/example-input";

const DIRECTIONS = {
  LEFT: "left",
  RIGHT: "right",
  UP: "up",
  DOWN: "down",
};

export function partOne(input = null) {
  var input = input || inputPath;
  input = fs.readFileSync(input, "utf8");
  var garden = input
    .trim()
    .split("\n")
    .map((line) => line.split(""));

  // First get start position
  let startCoordinate = getStartPosition(garden);
  let nextCoordinates = new Set([startCoordinate.toString()]);

  let steps = 0;
  while (steps < 64) {
    steps++;
    // For each coordinate, build the set of 'next coordinates'
    let currentCoordinates = structuredClone(nextCoordinates);
    // Clear this ready to fill in
    nextCoordinates = new Set();
    for (const coord of currentCoordinates) {
      // Move and add to the next queue
      for (const dir of Object.values(DIRECTIONS)) {
        let nextCoord = step(dir, ...coordFromString(coord));

        if (
          nextCoord[0] >= 0 &&
          nextCoord[0] < garden.length &&
          nextCoord[1] >= 0 &&
          nextCoord[1] < garden[0].length &&
          garden[nextCoord[0]][nextCoord[1]] !== "#"
        ) {
          nextCoordinates.add(nextCoord.toString());
        }
      }
    }
  }

  console.log({ day: 21, part: 1, value: nextCoordinates.size });
}

export function partTwo(input = null) {
  var input = input || inputPath;
  input = fs.readFileSync(input, "utf8");
  var garden = input
    .trim()
    .split("\n")
    .map((line) => line.split(""));

  // First get start position
  let startCoordinate = getStartPosition(garden);
  let nextCoordinates = new Set([startCoordinate.toString()]);

  /**
   * We now have an infinitely repeating garden. Observing that there are no obstacles in any 'direct' up/down/left/right, and that the
   *  maze is square, we can look for a pattern every time the equivalent 'start' points are reached, which is every maze.length steps!
   */
  let steps = 0;
  let valueAfterSteps = {};
  let stepGoal = 26501365;
  while (steps < stepGoal) {
    steps++;
    // For each coordinate, build the set of 'next coordinates'
    let currentCoordinates = structuredClone(nextCoordinates);
    // Clear this ready to fill in
    nextCoordinates = new Set();
    for (const coord of currentCoordinates) {
      // Move and add to the next queue
      for (const dir of Object.values(DIRECTIONS)) {
        let nextCoord = step(dir, ...coordFromString(coord));

        // Shift to the 'first' garden for the check - Javascript handles modulo of negatives unexpectedly so this looks a bit odd vs a nice "coord % size", but seems to be necessary
        const boundedI = (((nextCoord[0] + garden.length) % garden.length) + garden.length) % garden.length;
        const boundedJ = (((nextCoord[1] + garden.length) % garden.length) + garden.length) % garden.length;
        if (garden[boundedI][boundedJ] !== "#") {
          nextCoordinates.add(nextCoord.toString());
        }
      }
    }

    // If we're back at another 'edge-to-edge' point, store this
    if (steps % garden.length === stepGoal % garden.length) {
      console.log({ steps, val: nextCoordinates.size });
      valueAfterSteps[steps] = nextCoordinates.size;
    }

    // Only need 3 of these for the quadratic function below - this was upped to > 4 to build for the example at the bottom of the page
    if (Object.keys(valueAfterSteps).length > 4) {
      break;
    }
  }

  console.log({
    day: 21,
    part: 2,
    value: solveQuadratic(Object.values(valueAfterSteps), Math.floor(stepGoal / garden.length)),
  });

  /**
   * From hints online, and by confirming via day 9 showing the difference 2 levels deep is the same (see bottom of page for example), we
   * can simply solve for a quadratic and then plug in the number of gardens traversed (total steps % garden size)
   */
  function solveQuadratic(points, numGardensTraversed) {
    const y0 = parseInt(points[0]);
    const y1 = parseInt(points[1]);
    const y2 = parseInt(points[2]);
    const a = (y2 + y0 - 2 * y1) / 2;
    const b = y1 - y0 - a;
    const c = y0;
    return a * Math.pow(numGardensTraversed, 2) + b * numGardensTraversed + c;
  }
}

function getStartPosition(gardenMap) {
  for (let i = 0; i < gardenMap.length; i++) {
    for (let j = 0; j < gardenMap[i].length; j++) {
      if (gardenMap[i][j] === "S") {
        return [i, j];
      }
    }
  }
}

/* Converts "1,2" into [1,2] */
function coordFromString(str) {
  return str.split(",").map((v) => Number(v));
}

function step(direction, i, j) {
  i = parseInt(i);
  j = parseInt(j);
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

/**
 * Example - we got the following when going up by the number of steps across the board
 * { steps: 0, val: 1 }
 * { steps: 131, val: 15436 }
 * { steps: 262, val: 61343 }
 * { steps: 393, val: 137722 }
 * { steps: 524, val: 244573 }
 * { steps: 655, val: 381896 }
 * 1       15436        61343        137722         244573        381896
 *   15435        45907        76379         106851        137323
 *         30472       30472         30472          30472 <--------- Diffs are all the same 2 down - so it is a quadratic
 */

/**
 * Example - we got the following when going up by the number of steps across the board
 * { steps: 65, val: 3889 }
 * { steps: 196, val: 34504 }
 * { steps: 327, val: 95591 }
 * { steps: 458, val: 187150 }
 * { steps: 589, val: 309181 }
 * 3889       34504        95591      187150        309181
 *      30615        61087       91559       122031
 *            30472       30472         30472 <--------- Diffs are all the same 2 down - so it is a quadratic
 */
