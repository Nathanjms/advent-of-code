import fs from "fs";

const inputPath = "./day18/example-input";

const DIRECTIONS = {
  R: [0, 1],
  L: [0, -1],
  U: [-1, 0],
  D: [1, 0],
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
  let boundaryPoints = 0;
  for (let i = 0; i < digPlan.length; i++) {
    // move the number if steps in the direction, adding this to the path array
    let { dir, steps } = digPlan[i];
    currCoordinate[0] += DIRECTIONS[dir][0] * Number(steps);
    currCoordinate[1] += DIRECTIONS[dir][1] * Number(steps);
    boundaryPoints += Number(steps);
    path.push([...currCoordinate]);
  }

  /* Now let's use pick's theorem and the shoelace formula, wish I'd learned it for day 10 now */
  // A = 1/2 * (SUM from 0 < i <= n of x_i(y_i+1 - y_i-1) )
  // Unsure if we need to adapt because we start from 0 and go down for y, but let's see!

  let sum = 0;
  for (let i = 1; i < path.length; i++) {
    sum += path[i][1] * (path[(i + 1) % path.length][0] - path[i - 1][0]);
  }
  let area = 0.5 * sum;

  /**
   * Pick's theorem says that A = i + b / 2 - 1, where A is the area of the polygon, i is the number of internal integer points, and b is
   * the number of boundary integer points.
   * Rearrange this gives: i = A - 0.5b + 1
   * Then adding on the boundary points b gives i = A + 0.5b + 1
   */
  let numPoints = area + 0.5 * boundaryPoints + 1;

  console.log({ day: 18, part: 1, value: numPoints });
}

export function partTwo(input = null) {
  var input = input || inputPath;
  input = fs.readFileSync(input, "utf8");
  var inputArray = input.trim().split("\n");

  const COLOR_MAP = {
    0: "R",
    1: "D",
    2: "L",
    3: "U",
  };

  let digPlan = inputArray.map((v) => {
    var [_, _, color] = v.split(" ");
    color = color.slice(2, -1);
    const steps = parseInt(color.slice(0, 5), 16);
    const dir = COLOR_MAP[color.slice(-1)];
    return { dir, steps, color };
  });

  let path = [];
  let currCoordinate = [0, 0];
  // a hash indicates a path position:
  path.push([...currCoordinate]);
  let boundaryPoints = 0;
  for (let i = 0; i < digPlan.length; i++) {
    // move the number if steps in the direction, adding this to the path array
    let { dir, steps } = digPlan[i];
    currCoordinate[0] += DIRECTIONS[dir][0] * Number(steps);
    currCoordinate[1] += DIRECTIONS[dir][1] * Number(steps);
    boundaryPoints += Number(steps);
    path.push([...currCoordinate]);
  }

  /* Now let's use pick's theorem and the shoelace formula, wish I'd learned it for day 10 now */
  // A = 1/2 * (SUM from 0 < i <= n of x_i(y_i+1 - y_i-1) )
  // Unsure if we need to adapt because we start from 0 and go down for y, but let's see!

  let sum = 0;
  for (let i = 1; i < path.length; i++) {
    sum += path[i][1] * (path[(i + 1) % path.length][0] - path[i - 1][0]);
  }
  let area = 0.5 * sum;

  /**
   * Pick's theorem says that A = i + b / 2 - 1, where A is the area of the polygon, i is the number of internal integer points, and b is
   * the number of boundary integer points.
   * Rearrange this gives: i = A - 0.5b + 1
   * Then adding on the boundary points b gives i = A + 0.5b + 1
   */
  let numPoints = area + 0.5 * boundaryPoints + 1;

  console.log({ day: 18, part: 2, value: numPoints });
}
