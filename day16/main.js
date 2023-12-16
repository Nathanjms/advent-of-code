import fs from "fs";
import { stringify } from "querystring";

const inputPath = "./day16/example-input";

const DIRECTIONS = {
  LEFT: "left",
  RIGHT: "right",
  UP: "up",
  DOWN: "down",
};

export function partOne(input = null) {
  var input = input || inputPath;
  input = fs.readFileSync(input, "utf8");
  var inputArray = input.trim().split("\n");

  let traversedTiles = new Set();

  let lightBeamPositions = [
    { key: 0, coordinates: [0, 0], startCoordinates: [0, 0], direction: DIRECTIONS.RIGHT, isFinished: false },
  ]; // Start with single beam at 0,0

  // Note: We need to handle 0,0 to determine which direction the beam should be inside the loop. For ease, I've manually done this:
  lightBeamPositions[0].direction = DIRECTIONS.DOWN;

  traversedTiles.add("0,0");
  // Store whether we have been to this spot with these steps before, and break if so
  let cache = [];

  do {
    // Take a step to the right
    takeStep(lightBeamPositions);
  } while (lightBeamPositions.some((lightBeam) => !lightBeam.isFinished));

  console.log({ day: 16, part: 1, value: traversedTiles.size });

  function takeStep() {
    for (const lightBeam of lightBeamPositions.filter((beam) => !beam.isFinished)) {
      let [newI, newJ] = step(lightBeam.direction, ...lightBeam.coordinates);
      let key = newI + "," + newJ + "," + lightBeam.direction;

      // If the beam has exited or is back where it started (not sure this 2nd one is right), then we can say its finished.
      if (
        newI < 0 ||
        newI > inputArray.length - 1 ||
        newJ < 0 ||
        newJ > inputArray[1].length - 1 ||
        cache.includes(key)
        // (newI === lightBeam.startCoordinates[0] && newJ === lightBeam.startCoordinates[1])
      ) {
        lightBeam.isFinished = true;
        break;
      }
      cache.push(key);

      lightBeam.coordinates = [newI, newJ];
      traversedTiles.add(newI + "," + newJ);

      // Now the big ol' if statements
      if (inputArray[newI][newJ] === ".") {
        // Easiest case - nothing happens so keep the direction the same.
        break;
      }

      if (inputArray[newI][newJ] === "-") {
        // Two options here: (1) we approach from up or down and (2) we approach from left and right
        if (lightBeam.direction === DIRECTIONS.LEFT || lightBeam.direction === DIRECTIONS.RIGHT) {
          // Pass through as if nothing has happened!
          break;
        } else {
          // THis one we need to split the beam! First we'll change the direction of the Original, then clone and att the opposite direction
          lightBeam.direction = DIRECTIONS.LEFT;
          lightBeamPositions.push(cloneBeamAndChangeDirection(lightBeam));
          break;
        }
      }
      if (inputArray[newI][newJ] === "|") {
        // Two options here, similar to above but swap left/right with up/down etc
        if (lightBeam.direction === DIRECTIONS.UP || lightBeam.direction === DIRECTIONS.DOWN) {
          // Pass through as if nothing has happened!
          break;
        } else {
          // THis one we need to split the beam! First we'll change the direction of the Original, then clone and att the opposite direction
          lightBeam.direction = DIRECTIONS.UP;
          lightBeamPositions.push(cloneBeamAndChangeDirection(lightBeam));
          break;
        }
      }
      if (inputArray[newI][newJ] === "/") {
        lightBeam.direction = handleForwardSlash(lightBeam.direction);
        break;
      }
      if (inputArray[newI][newJ] === "\\") {
        // This will be the opposite as if it was a forward slash (i think!)
        lightBeam.direction = getOppositeDirection(handleForwardSlash(lightBeam.direction));
        break;
      }
    }
  }

  function handleForwardSlash(direction) {
    if (direction === DIRECTIONS.DOWN) return DIRECTIONS.LEFT;
    if (direction === DIRECTIONS.LEFT) return DIRECTIONS.DOWN;
    if (direction === DIRECTIONS.UP) return DIRECTIONS.RIGHT;
    if (direction === DIRECTIONS.RIGHT) return DIRECTIONS.UP;
  }
}

export function partTwo(input = null) {
  var input = input || inputPath;
  input = fs.readFileSync(input, "utf8");
  var inputArray = input
    .trim()
    .split("\n")
    .map((line) => line.split(""));

  console.log({ day: 16, part: 2, value: "todo" });
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

function cloneBeamAndChangeDirection(beam) {
  const newBeam = JSON.parse(JSON.stringify(beam));
  newBeam.direction = getOppositeDirection(beam.direction);
  return newBeam;
}
