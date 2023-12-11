import fs from "fs";

const inputPath = "./day8/example-input";
// const inputPath = "./input";

export function partOne() {
  var input = fs.readFileSync(inputPath + "2", "utf8");
  var inputArray = input.trim().split("\n");

  const { steps, coordinateMap } = parseInput(inputArray);

  let currentCoordinate = "AAA"; // Start at AAA
  let instructionIndex = 0;
  let numSteps = 0;
  while (currentCoordinate !== "ZZZ") {
    const [left, right] = coordinateMap[currentCoordinate];
    if (steps[instructionIndex] === "L") {
      currentCoordinate = left;
    } else if (steps[instructionIndex] === "R") {
      currentCoordinate = right;
    } else {
      throw new Error("Invalid instruction");
    }
    numSteps++;
    // Increment instruction index, but if it goes over the length of the steps, reset it to 0 by using modulo
    instructionIndex = (instructionIndex + 1) % steps.length;
  }

  console.log({ day: 8, part: 1, value: numSteps });
}

export function partTwo() {
  var input = fs.readFileSync(inputPath + "3", "utf8");
  var inputArray = input.trim().split("\n");

  const { steps, coordinateMap } = parseInput(inputArray);

  let currentCoordinates = Object.keys(coordinateMap).filter((coordinate) => coordinate[2] === "A"); // All that end with A. We know they are 3 digits long
  let instructionIndex = 0;
  let numSteps = 0;
  while (!allEndInZ(currentCoordinates)) {
    currentCoordinates = currentCoordinates.map((currentCoordinate) => {
      const [left, right] = coordinateMap[currentCoordinate];
      if (steps[instructionIndex] === "L") {
        return left;
      } else if (steps[instructionIndex] === "R") {
        return right;
      } else {
        throw new Error("Invalid instruction");
      }
    });
    numSteps++;
    // Increment instruction index, but if it goes over the length of the steps, reset it to 0 by using modulo
    instructionIndex = (instructionIndex + 1) % steps.length;
  }

  console.log({ day: 8, part: 2, value: numSteps });

  function allEndInZ(coordinates) {
    return coordinates.filter((coordinate) => coordinate[2] === "Z").length === coordinates.length;
  }
}

function parseInput(inputArray) {
  // Start from line 0, the steps are each line until the blank line.
  let steps = "";
  let startMapLine = 0;
  for (let i = 0; i < inputArray.length; i++) {
    steps += inputArray[i];
    if (inputArray[i] === "") {
      startMapLine = i + 1;
      break;
    }
  }

  // Now go through the next lines and store each map with key of the coordinate, and val of [LEFT, RIGHT] from input of eg RGT = (HDG, QJV)
  const coordinateMap = {};
  for (let i = startMapLine; i < inputArray.length; i++) {
    [...inputArray[i].matchAll(/(\w{3}) = \((\w{3}), (\w{3})\)/g)].forEach((match) => {
      let [_, key, left, right] = match;
      coordinateMap[key] = [left, right];
    });
  }

  return { steps, coordinateMap };
}
