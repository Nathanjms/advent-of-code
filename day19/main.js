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
  const { workflowsByKey } = parseInput(input);
  const ranges = {
    x: [1, 4000],
    m: [1, 4000],
    a: [1, 4000],
    s: [1, 4000],
  };

  const result = getPossibilities(ranges, "in");

  console.log({ day: 19, part: 2, value: result });

  function getPossibilities(ranges, workflowKey) {
    if (workflowKey === "R") {
      // Ended up at a reject, so ignore all of the ranges in this pathway
      return 0;
    }

    if (workflowKey === "A") {
      // Made it to an "Accept", so all of the remaining are valid possibilities!
      return Object.values(ranges).reduce((product, range) => {
        return product * (range[1] - range[0] + 1);
      }, 1);
    }

    /* Now we get to the head-scratching bit with recursion! */
    let result = 0;

    // For this workflow, we have a set of rules. We need to split the ranges up to allow for any of these.
    const workflow = workflowsByKey[workflowKey];
    // LOOP EACH RULE IN THE WORKFLOW. The > and < are so similar that maybe could refactor to a nicer solution, not sure if it's worth the effort though.
    for (let i = 0; i < workflow.length; i++) {
      if (!workflow[i].operator) {
        // LAST RULE IN THE FLOW - Default rule - all remaining ranges here go through to the next workflowKey:
        return result + getPossibilities(cloneArray(ranges), workflow[i].output);
      } else if (workflow[i].operator === ">") {
        // Range splits for the ones larger and smaller
        let value = workflow[i].value;
        let rangeKey = workflow[i].input;
        // Split into two - those that satisfy and move on to the next direction, and those that do not

        // Value larger than or equal to top of range - do nothing
        if (value >= ranges[rangeKey][1]) {
          // None of it!
        } else if (value < ranges[rangeKey][0]) {
          // All of it - so can stop here
          return result + getPossibilities(cloneArray(ranges), workflow[i].output);
        } else {
          // Some of it - chop up ranges
          let newAltRange = cloneArray(ranges);
          // New one will go from eg 500-4000 to 3001-4000 if value is 3000 and og will continue as 500-3000
          newAltRange[rangeKey][0] = value + 1;
          result += getPossibilities(newAltRange, workflow[i].output);
          // Ones left behind use the value as upper bound
          ranges[rangeKey][1] = value; // Update this for the next loop of this workflow
        }
      } else if (workflow[i].operator === "<") {
        // Range splits for the ones larger and smaller
        let value = workflow[i].value;
        let rangeKey = workflow[i].input;

        // Value less than or equal to bottom or range - do nothing
        if (value <= ranges[rangeKey][0]) {
          // None of it goes this way
        } else if (value > ranges[rangeKey][1]) {
          // All of it, no need to continue inside this
          return result + getPossibilities(cloneArray(ranges), workflow[i].output);
        } else {
          // Some of it - chop up
          let newAltRange = cloneArray(ranges);
          // New one will go from eg 500-4000 to 500-2999 if value is 3000 and og will continue as 3000-4000
          newAltRange[rangeKey][1] = value - 1;
          result += getPossibilities(newAltRange, workflow[i].output);
          ranges[rangeKey][0] = value;
        }
      } else {
        throw Error("Operator Not supported?");
      }
    }

    return result;
  }
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
          return { input, output, operator, value: Number(value) };
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

function cloneArray(arr) {
  // Nice and inefficient :) structuredClone might be what I need but needs node v18+
  return JSON.parse(JSON.stringify(arr));
}
