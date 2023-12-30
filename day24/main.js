import fs from "fs";

const inputPath = "./day24/example-input";

export function partOne(input = null) {
  var input = input || inputPath;
  input = fs.readFileSync(input, "utf8");
  var points = parseInput(input);

  // for each line, find the '+c' value of their equation
  for (let i = 0; i < points.length; i++) {
    const point = points[i];
    let m = Number(point.vel[1]) / Number(point.vel[0]);
    let c = Number(point.pos[1]) - m * point.pos[0];
    point["c"] = c;
    point["m"] = m;
  }

  // Store indexes of all compared to not repeat
  let compared = new Set();
  let seenPoints = 0;

  const MIN = 200000000000000;
  const MAX = 400000000000000;

  // const MIN = 7;
  // const MAX = 21;

  for (let i = 0; i < points.length; i++) {
    const el1 = points[i];
    for (let j = 0; j < points.length; j++) {
      const el2 = points[j];
      if (i === j) {
        continue;
      }
      if (compared.has(j)) {
        continue;
      }

      if (el1.m === el2.m) {
        // Parallel
        continue;
      }

      const x = (el2.c - el1.c) / (el1.m - el2.m);
      const y = x * el1.m + el1.c;

      // check if we have gone into the past to get this - if the point is increasing and has gone down, or vice versa (for both x and y)
      const el1Time = (x - el1.pos[0]) / el1.vel[0];
      const el2Time = (x - el2.pos[0]) / el2.vel[0];

      if (el1Time >= 0 && el2Time >= 0 && x >= MIN && x <= MAX && y >= MIN && y <= MAX) {
        seenPoints++;
      }
    }
    compared.add(i);
  }

  console.log(seenPoints);

  console.log({ day: 24, part: 1, value: seenPoints });
}

export function partTwo(input = null) {
  var input = input || inputPath;
  input = fs.readFileSync(input, "utf8");
  var inputArray = input
    .trim()
    .split("\n")
    .map((line) => line.split(""));

  console.log({ day: 24, part: 2, value: "todo" });
}

function parseInput(input) {
  return input
    .trim()
    .split("\n")
    .map((line) => {
      let [pos, vel] = line.split(" @ ").map((coord) =>
        coord
          .replace(/\s/g, "")
          .split(",")
          .map((el) => Number(el))
      );
      return { pos, vel };
    });
}
