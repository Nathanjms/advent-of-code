import fs from "fs";
import { init as z3Init } from "z3-solver";

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

  // const MIN = 200000000000000;
  // const MAX = 400000000000000;

  const MIN = 7;
  const MAX = 21;

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

  console.log({ day: 24, part: 1, value: seenPoints });
}

export async function partTwo(input = null) {
  var input = input || inputPath;
  input = fs.readFileSync(input, "utf8");
  var hailstones = parseInput(input);
  const { Context, em } = await z3Init();
  const { Real, Solver } = Context("main");

  // We want to now write the equations to then solve. the rock's x,y,x and velocity in x y z will be denoted below:
  const x = Real.const("x");
  const y = Real.const("y");
  const z = Real.const("z");

  const v_x = Real.const("v_x");
  const v_y = Real.const("v_y");
  const v_z = Real.const("v_z");

  const solver = new Solver();

  for (let i = 0; i < hailstones.length; i++) {
    const h = hailstones[i];
    const tVal = Real.const("t_" + i);

    solver.add(tVal.ge(0));

    // 3 equations for each rock:
    solver.add(tVal.mul(v_x).add(x).eq(tVal.mul(h.vel[0]).add(h.pos[0])));
    solver.add(tVal.mul(v_y).add(y).eq(tVal.mul(h.vel[1]).add(h.pos[1])));
    solver.add(tVal.mul(v_z).add(z).eq(tVal.mul(h.vel[2]).add(h.pos[2])));
  }

  const satisfied = await solver.check();

  if (!satisfied) {
    throw Error("Not satisfied the equations :c");
  }

  const model = solver.model();

  let result = 0;
  [x, y, z].forEach((val) => {
    result += Number(model.eval(val));
  });

  em.PThread.terminateAllThreads(); // For some reason, node process doesn't stop unless I put this in? src: https://github.com/Z3Prover/z3/issues/7070#issuecomment-1871017371

  console.log({ day: 24, part: 2, value: result });
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
