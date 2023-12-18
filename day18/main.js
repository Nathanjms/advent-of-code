import fs from "fs";

const inputPath = "./day18/example-input";

const DIRECTIONS = {
  R: [0, 1],
  L: [0, -1],
  U: [-1, 0],
  D: [1, 0],
};

const AVAILABLE_OUT_DIRECTIONS = {
  "|": [DIRECTIONS.U, DIRECTIONS.D],
  "-": [DIRECTIONS.L, DIRECTIONS.R],
  L: [DIRECTIONS.U, DIRECTIONS.R],
  J: [DIRECTIONS.U, DIRECTIONS.L],
  7: [DIRECTIONS.D, DIRECTIONS.L],
  F: [DIRECTIONS.D, DIRECTIONS.R],
  ".": [],
};

export function partOne(input = null) {
  var input = input || inputPath;
  input = fs.readFileSync(input, "utf8");
  var inputArray = input.trim().split("\n");

  let digPlan = inputArray.map((v) => {
    let [dir, steps, color] = v.split(" ");
    color = color.slice(2, -1);
    return { dir, steps, color };
  });
  let path = [];
  let currCoordinate = [0, 0];
  // a hash indicates a path position:
  path.push([...currCoordinate]);
  for (let i = 0; i < digPlan.length; i++) {
    // move the number if steps in the direction, adding this to the path array
    let { dir, steps } = digPlan[i];
    for (let n = 1; n <= steps; n++) {
      currCoordinate[0] += DIRECTIONS[dir][0];
      currCoordinate[1] += DIRECTIONS[dir][1];
      // if we are back at the start, break:
      if (currCoordinate.toString() === "0,0") {
        break;
      }
      path.push([...currCoordinate]);
    }
  }

  // now we have the path, build out the grid by filling gaps with '.'.
  // first find the max size of the grid:
  let maxI = Math.max(...path.map(([i, _]) => i));
  let minI = Math.min(...path.map(([i, _]) => i));
  let maxJ = Math.max(...path.map(([_, j]) => j));
  let minJ = Math.min(...path.map(([_, j]) => j));

  maxI = maxI - minI;
  maxJ = maxJ - minJ;
  // now build it. I'm going to recreate the pipes etc from day 10 to then reuse that logic for getting the number of interior points.
  let grid = Array.from({ length: maxI + 1 }).map(() => Array.from({ length: maxJ + 1 }).fill("."));
  for (let i = 0; i < path.length; i++) {
    let prevElement = i === 0 ? path[path.length - 1] : path[i - 1];
    let nextElement = i === path.length - 1 ? path[0] : path[i + 1];
    const element = path[i];

    // dy is opposite because rows decrease!
    let prevElementDiff = [prevElement[0] - element[0], prevElement[1] - element[1]];
    const nextElementDiff = [nextElement[0] - element[0], nextElement[1] - element[1]];

    // Can either be |, -, L, F, J, 7

    // Find the available direction with both the diffs direction
    let type = null;
    for (const direction in AVAILABLE_OUT_DIRECTIONS) {
      const element = AVAILABLE_OUT_DIRECTIONS[direction];
      // If both are there then it's the right type of symbol
      if (
        element.findIndex((el) => el.toString() === prevElementDiff.toString()) !== -1 &&
        element.findIndex((el) => el.toString() === nextElementDiff.toString()) !== -1
      ) {
        type = direction;
        break; // found it!
      }
    }
    if (!type) {
      throw Error("no dir :0");
    }
    grid[element[0] - minI][element[1] - minJ] = type;
  }

  // Now we can loop through each element and check if it is trapped.
  // Note we skip the boundaries as they can't be trapped by definition of being on the edge
  // At most they will be coordinate
  const trappedElements = [];
  for (let i = 1; i < grid.length - 1; i++) {
    for (let j = 1; j < grid[0].length - 1; j++) {
      console.log({ i, j });
      if (path.findIndex((el) => el.toString() === i + "," + j) !== -1) {
        // If this is a boundary, it can't be a trapped element.
        continue;
      }
      if (getNumberOfPassThroughs(i, j, grid[i], path) % 2 === 1) {
        trappedElements.push(i + "," + j);
      }
    }
  }

  function getNumberOfPassThroughs(i, j, line, path) {
    // Go from this point to the end by incrementing j
    let sum = 0;
    let lineIndex = j;

    while (lineIndex < line.length) {
      /**
       * If it is a boundary point, then it can be considered as crossing if it is a | (easy), or if it is F---J or L---7, where '-' is any
       * natural number N. Because we are checking boundary coordinates only, we know there MUST be a J if there is an F, and a 7 if an L.
       * Therefore, below we only need to check three: |, J, L
       */
      if (path.findIndex((el) => el.toString() === i + "," + lineIndex) && ["|", "J", "L"].includes(line[lineIndex])) {
        sum++;
      }
      lineIndex++;
    }
    return sum;
  }
  // count combo for answer
  const answer = trappedElements.length + path.length;

  console.log({ day: 18, part: 1, value: answer });
}

export function partTwo(input = null) {
  var input = input || inputPath;
  input = fs.readFileSync(input, "utf8");
  var inputArray = input.trim().split("\n");

  console.log({ day: 18, part: 2, value: "todo" });
}

function display(grid) {
  console.log(grid.map((val) => val.join("")).join("\n"));
  console.log();
}
