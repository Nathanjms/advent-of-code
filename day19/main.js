import fs from "fs";

const inputPath = "./day19/example-input";

export function partOne(input = null) {
  var input = input || inputPath;
  input = fs.readFileSync(input, "utf8");

  const { workflowsByKey, parts } = parseInput(input);

  let acceptedSum = 0;

  for (const part of parts) {
    let position = "in"; // Always starts at 'in'
    while (!["A", "R"].includes(position)) {
      position = goThroughWorkflow(workflowsByKey[position], part);
    }

    if (position === "A") {
      acceptedSum += Object.values(part).reduce((a, b) => a + b, 0);
    }
  }

  console.log({ day: 19, part: 1, value: acceptedSum });
}

export function partTwo(input = null) {
  var input = input || inputPath;
  input = fs.readFileSync(input, "utf8");
  var inputArray = input.trim().split("\n");

  console.log({ day: 19, part: 2, value: "todo" });
}

function parseInput(input) {
  var [workflows, parts] = input.split("\n\n").map((chunk) => chunk.split("\n"));

  const workflowsByKey = workflows.reduce((accum, workflow) => {
    var [key, rest] = workflow.split("{");
    accum[key] = rest
      .slice(0, -1)
      .split(",")
      .map((val) => {
        if (val.indexOf(":") === -1) {
          return { output: val };
        } else {
          let input = val.slice(0, 1);
          let operator = val.slice(1, 2);
          let [value, output] = val.slice(2).split(":");
          return { input, output, operator, value };
        }
      });
    return accum;
  }, {});

  parts = parts.map((part) =>
    part
      .slice(1, -1)
      .split(",")
      .reduce((accum, prt) => {
        const [k, v] = prt.split("=");
        accum[k] = parseInt(v);
        return accum;
      }, {})
  );

  return { workflowsByKey, parts };
}

function goThroughWorkflow(workflow, part) {
  for (let i = 0; i < workflow.length; i++) {
    if (!workflow[i].operator) {
      return workflow[i].output;
    }
    if (workflow[i].operator === ">") {
      if (part[workflow[i].input] > workflow[i].value) {
        return workflow[i].output;
      }
    } else if (workflow[i].operator === "<") {
      if (part[workflow[i].input] < workflow[i].value) {
        return workflow[i].output;
      }
    } else {
      throw Error("Operator Not supported?");
    }
  }
}
