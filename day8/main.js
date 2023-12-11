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

export function partTwo(input = null) {
  input = input || inputPath + "3";
  var input = fs.readFileSync(input, "utf8");
  var inputArray = input.trim().split("\n");

  const { steps, coordinateMap } = parseInput(inputArray);

  let currentCoordinates = Object.keys(coordinateMap).filter((coordinate) => coordinate[2] === "A"); // All that end with A. We know they are 3 digits long

  /**
   * 'Brute force' will be very slow here, so we will use the fact that these are cyclic values and each only has one valid end value, and then look to find the lowest common multiple
   * To do this, we will:
   * 1. Get the number of steps needed to reach the coordinate ending in Z FOR EACH coordinate
   * 2. Compute the lowest common multiple of all of these different values
   *
   * Note that these being cyclic was not mentioned in the question, but appears to be the case.
   */

  const repetitionsUntilReachZ = {};
  currentCoordinates.forEach((coordinate) => {
    let currentCoordinate = coordinate;
    let numSteps = 0;
    let instructionIndex = 0;
    while (true) {
      if (currentCoordinate[2] === "Z") {
        break;
      }
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
    repetitionsUntilReachZ[coordinate] = numSteps;
  });

  const lcmValue = lcmOfArray(Object.values(repetitionsUntilReachZ));

  console.log({ day: 8, part: 2, value: lcmValue });

  function gcd(num1, num2) {
    //if num2 is 0 return num1;
    if (num2 === 0) {
      return num1;
    }

    //call the same function recursively
    return gcd(num2, num1 % num2);
  }

  function lcm(num1, num2) {
    return (num1 * num2) / gcd(num1, num2);
  }

  function lcmOfArray(arr) {
    let result = arr[0];

    for (let i = 1; i < arr.length; i++) {
      console.log({ result });
      result = lcm(result, arr[i]);
    }

    return result;
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

function originalPartTwo(inputPath = null) {
  var input = fs.readFileSync(inputPath, "utf8");
  var inputArray = input.trim().split("\n");

  const { steps, coordinateMap } = parseInput(inputArray);

  let currentCoordinates = Object.keys(coordinateMap).filter((coordinate) => coordinate[2] === "A"); // All that end with A. We know they are 3 digits long
  let instructionIndex = 0;
  let numSteps = 0;
  while (!allEndInZ(currentCoordinates)) {
    console.log({ currentCoordinates });
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
