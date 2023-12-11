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

  // TESTING - coordinate start is row 2 col 0
  let startCoordinate = [2, 0];
  let coordinate = startCoordinate;
  let symbol = "F";
  let direction = AVAILABLE_OUT_DIRECTIONS[symbol][0];
  while (symbol !== "S") {
    let stepOutput = takeStep(direction, symbol, ...coordinate);
    direction = stepOutput.direction;
    coordinate = stepOutput.newCoordinate;
    symbol = inputArray[coordinate[0]][coordinate[1]];
    console.log({ direction, coordinate, symbol });
  }

  console.log({ day: 10, part: 1, value: "todo" });
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
