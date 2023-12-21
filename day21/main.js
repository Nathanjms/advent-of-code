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
  while (steps < 200_000) {
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
  var inputArray = input
    .trim()
    .split("\n")
    .map((line) => line.split(""));

  console.log({ day: 21, part: 2, value: "todo" });
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

function getOppositeDirection(direction) {
  if (direction === DIRECTIONS.LEFT) return DIRECTIONS.RIGHT;
  if (direction === DIRECTIONS.RIGHT) return DIRECTIONS.LEFT;
  if (direction === DIRECTIONS.UP) return DIRECTIONS.DOWN;
  if (direction === DIRECTIONS.DOWN) return DIRECTIONS.UP;
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
