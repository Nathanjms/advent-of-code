import fs from "fs";

const inputPath = "./day23/example-input";

const DIRECTIONS = {
  LEFT: "left",
  RIGHT: "right",
  UP: "up",
  DOWN: "down",
};

const SLOPE_TO_DIRECTION = {
  "^": DIRECTIONS.UP,
  ">": DIRECTIONS.RIGHT,
  v: DIRECTIONS.DOWN,
  "<": DIRECTIONS.LEFT,
};

export function partOne(input = null) {
  var input = input || inputPath;
  input = fs.readFileSync(input, "utf8");
  var trailMap = input
    .trim()
    .split("\n")
    .map((line) => line.split(""));

  let start = [0, trailMap[0].indexOf(".")];
  let end = [trailMap.length - 1, trailMap[trailMap.length - 1].indexOf(".")];

  /* Use Edge Contraction to simplify the input first, before using brute force to find the longest path */

  let points = [[...start], [...end]]; // STart with the 'points of interest' as just the start and end
  for (let r = 0; r < trailMap.length; r++) {
    for (let c = 0; c < trailMap[r].length; c++) {
      const point = trailMap[r][c];
      if (point === "#") {
        continue;
      }
      let neighbours = 0;
      for (const dir of Object.values(DIRECTIONS)) {
        // If we are currently on a 'Peak', restrict
        if (["^", ">", "v", "<"].includes(trailMap[r][c]) && dir !== SLOPE_TO_DIRECTION[trailMap[r][c]]) {
          continue;
        }
        let newCoord = step(dir, r, c);
        if (
          newCoord[0] >= 0 &&
          newCoord[0] <= trailMap.length - 1 &&
          newCoord[1] >= 0 &&
          newCoord[1] < trailMap[0].length - 1 &&
          trailMap[newCoord[0]][newCoord[1]] !== "#"
        ) {
          neighbours++;
        }
      }
      // If we have at least 2 places to go (+ 1 because we've stepped from a valid point), then this is an 'edge' point
      if (neighbours >= 3) {
        points.push([r, c]);
      }
    }
  }

  let graph = points.reduce((agg, a) => {
    agg[a.join(",")] = [];
    return agg;
  }, {});

  // Now we have a list of points that are where the 'splits' occur. Now we can build a graph of a key f the start, and a set of vals
  // that each contain the number of steps if go that route. Use depths first search for this for each point

  for (const pt of points) {
    let queue = [{ pt, numSteps: 0 }];
    let seen = new Set([pt.join(",")]);

    while (queue.length) {
      const el = queue.pop();
      const numSteps = el.numSteps;
      const [newR, newC] = [...el.pt];
      const key = newR + "," + newC;

      if (numSteps !== 0 && points.some((v) => v[0] == newR && v[1] == newC)) {
        graph[pt.join(",")].push({ next: [newR, newC], numSteps });
        continue;
      }

      for (const dir of Object.values(DIRECTIONS)) {
        // If we are currently on a 'Peak', restrict
        if (["^", ">", "v", "<"].includes(trailMap[newR][newC]) && dir !== SLOPE_TO_DIRECTION[trailMap[newR][newC]]) {
          continue;
        }
        let newCoord = step(dir, newR, newC);
        if (
          !seen.has(newCoord[0] + "," + newCoord[1]) &&
          newCoord[0] >= 0 &&
          newCoord[0] <= trailMap.length - 1 &&
          newCoord[1] >= 0 &&
          newCoord[1] < trailMap[0].length - 1 &&
          trailMap[newCoord[0]][newCoord[1]] !== "#"
        ) {
          queue.push({ pt: [...newCoord], numSteps: numSteps + 1 });
          seen.add(key);
        }
      }
    }
  }

  // Now we simply start at the start and do a DFS for each paths we can take, outputting the maximum one!
  const result = depthFirstSearch(...start);

  console.log({ day: 23, part: 1, value: result });

  function depthFirstSearch(i, j) {
    if (i == end[0] && j == end[1]) {
      return 0; // At the end so no steps to take
    }

    let steps = -Infinity;

    for (let { next, numSteps } of graph[i + "," + j]) {
      steps = Math.max(steps, depthFirstSearch(...next) + numSteps);
    }

    return steps;
  }
}

export function partTwo(input = null) {
  var input = input || inputPath;
  input = fs.readFileSync(input, "utf8");
  var trailMap = input
    .trim()
    .split("\n")
    .map((line) => line.split(""));

  let start = [0, trailMap[0].indexOf(".")];
  let end = [trailMap.length - 1, trailMap[trailMap.length - 1].indexOf(".")];

  /* Use Edge Contraction to simplify the input first, before using brute force to find the longest path */

  let points = [[...start], [...end]]; // Start with the 'points of interest' as just the start and end
  for (let r = 0; r < trailMap.length; r++) {
    for (let c = 0; c < trailMap[r].length; c++) {
      const point = trailMap[r][c];
      if (point === "#") {
        continue;
      }
      let neighbours = 0;
      for (const dir of Object.values(DIRECTIONS)) {
        let newCoord = step(dir, r, c);
        if (
          newCoord[0] >= 0 &&
          newCoord[0] <= trailMap.length - 1 &&
          newCoord[1] >= 0 &&
          newCoord[1] < trailMap[0].length - 1 &&
          trailMap[newCoord[0]][newCoord[1]] !== "#"
        ) {
          neighbours++;
        }
      }
      // If we have at least 2 places to go (+ 1 because we've stepped from a valid point), then this is an 'edge' point
      if (neighbours >= 3) {
        points.push([r, c]);
      }
    }
  }

  let graph = points.reduce((agg, a) => {
    agg[a.join(",")] = [];
    return agg;
  }, {});

  // Now we have a list of points that are where the 'splits' occur. Now we can build a graph of a key f the start, and a set of vals
  // that each contain the number of steps if go that route. Use depths first search for this for each point

  for (const pt of points) {
    let queue = [{ pt, numSteps: 0 }];
    let seen = new Set([pt.join(",")]);

    while (queue.length) {
      const el = queue.pop();
      const numSteps = el.numSteps;
      const [newR, newC] = [...el.pt];
      const key = newR + "," + newC;

      if (numSteps !== 0 && points.some((v) => v[0] == newR && v[1] == newC)) {
        graph[pt.join(",")].push({ next: [newR, newC], numSteps });
        continue;
      }

      for (const dir of Object.values(DIRECTIONS)) {
        let newCoord = step(dir, newR, newC);
        if (
          !seen.has(newCoord[0] + "," + newCoord[1]) &&
          newCoord[0] >= 0 &&
          newCoord[0] <= trailMap.length - 1 &&
          newCoord[1] >= 0 &&
          newCoord[1] < trailMap[0].length - 1 &&
          trailMap[newCoord[0]][newCoord[1]] !== "#"
        ) {
          queue.push({ pt: [...newCoord], numSteps: numSteps + 1 });
          seen.add(key);
        }
      }
    }
  }

  // Now we simply start at the start and do a DFS for each paths we can take, outputting the maximum one!
  let seen = new Set([start.join(",")]);
  const result = depthFirstSearch(...start, seen);

  console.log({ day: 23, part: 2, value: result });

  function depthFirstSearch(i, j, seen) {
    if (i == end[0] && j == end[1]) {
      return 0; // At the end so no steps to take
    }

    let steps = -Infinity;

    for (let { next, numSteps } of graph[i + "," + j]) {
      if (!seen.has(next[0] + "," + next[1])) {
        // Now add to set, but do not manipulate the original variable
        let newSet = new Set(Array.from(seen));
        newSet.add(next.join(","));
        steps = Math.max(steps, depthFirstSearch(...next, newSet) + numSteps);
      }
    }

    return steps;
  }
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
