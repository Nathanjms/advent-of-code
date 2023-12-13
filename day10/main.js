import fs from "fs";
import { floodFill, getEnclosedCoords } from "./helpers.js";

const inputPath = "./day10/example-input";

const LEFT = 0;
const UP = 1;
const RIGHT = 2;
const DOWN = 3;

/**
 * The directions connected by the symbol, from the symbol's perspective (ie. looking 'out' from the symbol)
 */
const AVAILABLE_OUT_DIRECTIONS = {
  "|": [UP, DOWN],
  "-": [LEFT, RIGHT],
  L: [UP, RIGHT],
  J: [UP, LEFT],
  7: [DOWN, LEFT],
  F: [DOWN, RIGHT],
  ".": [],
};

export function partOne(input = null) {
  input = input || inputPath;
  var input = fs.readFileSync(input, "utf8");
  var inputArray = input.trim().split("\n");

  let startCoordinate = findStartCoordinate(inputArray);
  let symbol = determineStartCoordinateSymbol(inputArray, startCoordinate);

  let steps = 0;
  // Define coordinates/symbol for each path
  let directionA = getOppositeDirection(AVAILABLE_OUT_DIRECTIONS[symbol][0]);
  let directionB = getOppositeDirection(AVAILABLE_OUT_DIRECTIONS[symbol][1]);
  let symbolA = symbol;
  let symbolB = symbol;
  let coordinateA = [...startCoordinate];
  let coordinateB = [...startCoordinate];

  while (true) {
    // A:
    let stepResultA = takeStep(directionA, symbolA, ...coordinateA);
    directionA = stepResultA.direction;
    coordinateA = stepResultA.newCoordinate;
    symbolA = inputArray[coordinateA[0]][coordinateA[1]];
    // B:
    let stepResultB = takeStep(directionB, symbolB, ...coordinateB);
    directionB = stepResultB.direction;
    coordinateB = stepResultB.newCoordinate;
    symbolB = inputArray[coordinateB[0]][coordinateB[1]];
    steps++;
    if (coordinateA.join(",") == coordinateB.join(",")) {
      break;
    }
  }

  console.log({ day: 10, part: 1, value: steps });
}

export function partTwo(input = null) {
  input = input || inputPath + "3";
  var input = fs.readFileSync(input, "utf8");
  var inputArray = input.trim().split("\n");
  console.log({ inputArray });

  let startCoordinate = findStartCoordinate(inputArray);
  let symbol = determineStartCoordinateSymbol(inputArray, startCoordinate);

  // Loop Coordinates now need to be stored, then we will go through every coordinate, and if it:
  // 1. Is not in the loop coordinates
  // 2. Has loop coordinates every side* (THis doesnt work)
  // Then we say it is surrounded
  const loopCoordinates = [];
  // Define coordinates/symbol for each path
  let directionA = getOppositeDirection(AVAILABLE_OUT_DIRECTIONS[symbol][0]);
  let directionB = getOppositeDirection(AVAILABLE_OUT_DIRECTIONS[symbol][1]);
  let symbolA = symbol;
  let symbolB = symbol;
  let coordinateA = [...startCoordinate];
  let coordinateB = [...startCoordinate];
  loopCoordinates.push(startCoordinate.join(","));

  while (true) {
    // A:
    let stepResultA = takeStep(directionA, symbolA, ...coordinateA);
    directionA = stepResultA.direction;
    coordinateA = stepResultA.newCoordinate;
    symbolA = inputArray[coordinateA[0]][coordinateA[1]];
    loopCoordinates.push(coordinateA.join(","));
    // B:
    let stepResultB = takeStep(directionB, symbolB, ...coordinateB);
    directionB = stepResultB.direction;
    coordinateB = stepResultB.newCoordinate;
    symbolB = inputArray[coordinateB[0]][coordinateB[1]];
    if (coordinateA.join(",") == coordinateB.join(",")) {
      break;
    }
    loopCoordinates.push(coordinateB.join(",")); // Push b here, so we don't duplicate the last one!
  }

  // Apply a flood fill to the grid to work out all coordinates that can escape. Then any remaining not in floodfill and in loop are the enclosed ones.

  // To handle being able to split through pipes, I'll add a special character '#' inbetween each pipe in the x and y directions. Then can handle very similar to before?
  const newInputArray = padInputsWithHashesAndRemap(inputArray);

  let test = floodFill(loopCoordinates, newInputArray, 0, 0);
  let test2 = getEnclosedCoords(newInputArray, test);

  console.log({ test });
  console.log({ test2 });

  let availableCoords = [];
  let coordinatesThatCanEscape = getAvailableCoords(0, 0, newInputArray);

  const enclosedCoords = [];

  for (let i = 0; i < newInputArray.length; i++) {
    for (let j = 0; j < newInputArray[0].length; j++) {
      if (newInputArray[i][j] === "#") {
        continue;
      }
      if (
        !loopCoordinates.includes(`${i / 2},${j / 2}`) &&
        !coordinatesThatCanEscape.includes(`${i},${j}`) &&
        i !== 0 &&
        j !== 0 &&
        i !== newInputArray.length &&
        j !== newInputArray[0].length
      ) {
        // Next we convert back to regular coordinate map. This is simply divide the  part by 2:
        enclosedCoords.push(`${i / 2},${j / 2}`);
      }
    }
  }

  console.log({ enclosedCoords });

  console.log({ day: 10, part: 2, value: enclosedCoords.length });

  function getAvailableCoords(i, j, inputArray, visitedCoords = []) {
    if (loopCoordinates.includes(i / 2 + "," + j / 2) || visitedCoords.includes(`${i / 2},${j / 2}`)) {
      // Stop if we've reached somewhere we've visited or have reached the loop
      return;
    }

    // Add if this coordinate is not in the loop and is not on the edge
    availableCoords.push(`${i},${j}`);

    visitedCoords.push(`${i},${j}`);

    // UP:
    if (
      i > 0 &&
      !visitedCoords.includes(`${i - 1},${j}`) &&
      (inputArray[i - 1][j] === "." || inputArray[i - 1][j] === "#")
    ) {
      availableCoords = getAvailableCoords(i - 1, j, inputArray, visitedCoords);
    }
    // Down:
    if (
      i < inputArray.length - 1 &&
      !visitedCoords.includes(`${i + 1},${j}`) &&
      (inputArray[i + 1][j] === "." || inputArray[i + 1][j] === "#")
    ) {
      availableCoords = getAvailableCoords(i + 1, j, inputArray, visitedCoords);
    }
    // LEFT:
    if (
      j > 0 &&
      !visitedCoords.includes(`${i},${j - 1}`) &&
      (inputArray[i][j - 1] === "." || inputArray[i][j - 1] === "#")
    ) {
      availableCoords = getAvailableCoords(i, j - 1, inputArray, visitedCoords);
    }
    // RIGHT:
    if (
      j < inputArray[0].length &&
      !visitedCoords.includes(`${i},${j + 1}`) &&
      (inputArray[i][j + 1] === "." || inputArray[i][j + 1] === "#")
    ) {
      availableCoords = getAvailableCoords(i, j + 1, inputArray, visitedCoords);
    }

    return availableCoords;
  }

  function padInputsWithHashesAndRemap(inputArray) {
    let newInputArray = [];
    // Also want to add a new line in-between each
    for (let i = 0; i < inputArray.length; i++) {
      newInputArray.push(inputArray[i].replace(/(.{1})/g, "$1#"));
      newInputArray.push(inputArray[i].replace(/(.{1})/g, "##"));
    }

    // Now map each line into array of characters instead of a string
    newInputArray = newInputArray.map((line) => line.split(""));

    // Maintain any pipes in the new input hash lines
    for (let i = 1; i < newInputArray.length - 1; i = i + 2) {
      for (let j = 0; j < newInputArray[i].length; j++) {
        if (newInputArray[i - 1][j] === "|" || newInputArray[i + 1][j] === "|") {
          newInputArray[i][j] = "|";
        }
      }
    }

    for (let i = 0; i < newInputArray.length - 1; i = i + 2) {
      for (let j = 1; j < newInputArray[i].length - 1; j++) {
        if (newInputArray[i][j] === "#" && (newInputArray[i][j - 1] === "-" || newInputArray[i][j + 1] === "-")) {
          newInputArray[i][j] = "-";
        }
      }
    }
    return newInputArray;
  }
}

function findStartCoordinate(inputArray) {
  for (let i = 0; i < inputArray.length; i++) {
    let startIndex = inputArray[i].indexOf("S");
    if (startIndex > -1) {
      return [i, startIndex]; // return (i,j) - ie. the (rowIndex, colIndex)
    }
  }
}

function determineStartCoordinateSymbol(inputArray, startCoordinate) {
  const surroundingValidDirections = [];
  // Look UP if we are not on the top row
  if (startCoordinate[0] > 0) {
    let testCoord = [startCoordinate[0] - 1, startCoordinate[1]];
    let letter = inputArray[testCoord[0]][testCoord[1]];
    // Does the letter have an output of DOWN
    if (AVAILABLE_OUT_DIRECTIONS[letter].includes(DOWN)) {
      surroundingValidDirections.push(UP);
    }
  }
  // Look DOWN if we are not on the bottom row
  if (startCoordinate[0] < inputArray.length) {
    let testCoord = [startCoordinate[0] + 1, startCoordinate[1]];
    let letter = inputArray[testCoord[0]][testCoord[1]];
    // Does the letter have an output of UP
    if (AVAILABLE_OUT_DIRECTIONS[letter].includes(UP)) {
      surroundingValidDirections.push(DOWN);
    }
  }

  // Look RIGHT if we are not on the right-most col  (note: all rows are the same length)
  if (startCoordinate[1] < inputArray[0].length) {
    let testCoord = [startCoordinate[0], startCoordinate[1] + 1];
    let letter = inputArray[testCoord[0]][testCoord[1]];
    // Does the letter have an output of LEFT
    if (AVAILABLE_OUT_DIRECTIONS[letter].includes(LEFT)) {
      surroundingValidDirections.push(RIGHT);
    }
  }

  // Look LEFT is we are not on the first col
  if (startCoordinate[1] > 0) {
    let testCoord = [startCoordinate[0], startCoordinate[1] - 1];
    let letter = inputArray[testCoord[0]][testCoord[1]];
    // Does the letter have an output of RIGHT
    if (AVAILABLE_OUT_DIRECTIONS[letter].includes(RIGHT)) {
      surroundingValidDirections.push(LEFT);
    }
  }

  for (const letter in AVAILABLE_OUT_DIRECTIONS) {
    // If every value in AVAILABLE_OUT_DIRECTIONS[letter] is in surroundingValidDirections, return that letter
    // Note: If >1 match, we maybe need to go through and handle errors?
    if (
      AVAILABLE_OUT_DIRECTIONS[letter].every((val) => {
        return surroundingValidDirections.indexOf(val) !== -1;
      })
    ) {
      return letter;
    }
  }
}

function getOppositeDirection(direction) {
  if (direction === LEFT) return RIGHT;
  if (direction === RIGHT) return LEFT;
  if (direction === UP) return DOWN;
  if (direction === DOWN) return UP;
}

/**
 * NOTE: i is the row, and j is the column - DO NOT MISTAKE FOR (x,y)!
 */
function takeStep(fromDirection, symbol, i, j) {
  // We have come from x, so get the inverse as the way it came into the symbol. The other remaining direction if the way we move.
  let directionToRemove = getOppositeDirection(fromDirection);
  let directionToMove = AVAILABLE_OUT_DIRECTIONS[symbol].find((direction) => direction != directionToRemove);

  if (directionToMove === LEFT) {
    return {
      direction: LEFT,
      newCoordinate: [i, j - 1],
    };
  }
  if (directionToMove === RIGHT) {
    return {
      direction: RIGHT,
      newCoordinate: [i, j + 1],
    };
  }
  if (directionToMove === UP) {
    return {
      direction: UP,
      newCoordinate: [i - 1, j],
    };
  }
  if (directionToMove === DOWN) {
    return {
      direction: DOWN,
      newCoordinate: [i + 1, j],
    };
  }
  throw Error("Invalid Direction!");
}
