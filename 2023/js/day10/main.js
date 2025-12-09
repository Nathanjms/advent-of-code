import fs from "fs";

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
  input = input || inputPath + "4";
  var input = fs.readFileSync(input, "utf8");
  var inputArray = input
    .trim()
    .split("\n")
    .map((line) => line.split(""));

  /**
   * lets use the Point In Polygon, Ray Casting Algorithm to solve this.
   * https://en.wikipedia.org/wiki/Point_in_polygon
   * For each point, count in one direction to the edge, the number of times the boundary is passed.
   * If even, the point is OUTSIDE, if odd, INSIDE.
   * We will go to the right.
   * The slight difficulty comes from being able to go in-between pipes.
   * This means for F---7 and L---J, this can be gone around and so does not count as a border
   * We count as 1 boundary for the opposite, ie. F----J or L-----7
   */

  let startCoordinate = findStartCoordinate(inputArray);
  let symbol = determineStartCoordinateSymbol(inputArray, startCoordinate);

  // First lets get the boundary coordinates
  const boundaryCoordinates = getLoopCoordinates(symbol, startCoordinate);

  // next let's swap 'S' in the input for it's symbol:
  inputArray[startCoordinate[0]][startCoordinate[1]] = symbol;

  // Now we can loop through each element and check if it is trapped.
  // Note we skip the boundaries as they can't be trapped by definition of being on the edge
  // At most they will be coordinate
  const trappedElements = [];
  for (let i = 1; i < inputArray.length - 1; i++) {
    for (let j = 1; j < inputArray[0].length - 1; j++) {
      if (boundaryCoordinates.includes(i + "," + j)) {
        // If this is a boundary, it can't be a trapped element.
        continue;
      }
      if (getNumberOfPassThroughs(i, j, inputArray[i], boundaryCoordinates) % 2 === 1) {
        trappedElements.push(i + "," + j);
      }
    }
  }

  console.log({ day: 10, part: 2, value: trappedElements.length });

  function getNumberOfPassThroughs(i, j, line, boundaryCoordinates) {
    // Go from this point to the end by incrementing j
    let sum = 0;
    let lineIndex = j;

    while (lineIndex < line.length) {
      /**
       * If it is a boundary point, then it can be considered as crossing if it is a | (easy), or if it is F---J or L---7, where '-' is any
       * natural number N. Because we are checking boundary coordinates only, we know there MUST be a J if there is an F, and a 7 if an L.
       * Therefore, below we only need to check three: |, J, L
       */
      if (boundaryCoordinates.includes(i + "," + lineIndex) && ["|", "J", "L"].includes(line[lineIndex])) {
        sum++;
      }
      lineIndex++;
    }
    return sum;
  }

  function getLoopCoordinates(symbol, startCoordinate) {
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

    return loopCoordinates;
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
