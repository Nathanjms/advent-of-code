import fs from "fs";

const inputPath = "./day15/example-input";

export function partOne(input = null) {
  var input = input || inputPath;
  input = fs.readFileSync(input, "utf8");
  var instructions = input.trim().split(",");

  const sums = [];

  for (const instruction of instructions) {
    sums.push(applyHashAlgorithm(instruction));
  }

  console.log({ day: 15, part: 1, value: sums.reduce((a, b) => a + b, 0) });
}

export function partTwo(input = null) {
  var input = input || inputPath;
  input = fs.readFileSync(input, "utf8");
  var instructions = input
    .trim()
    .split(",")
    .map((instruction) => {
      let [_, label, operator, value] = instruction.match(/(\w+)([=\-])(\d*)/);
      return {
        label,
        operator,
        value,
      };
    });

  let boxes = Array.from({ length: 256 }, () => []);

  for (const instruction of instructions) {
    let boxId = applyHashAlgorithm(instruction.label);
    if (instruction.operator === "-") {
      boxes[boxId] = boxes[boxId].filter((val) => val.slice(0, instruction.label.length) !== instruction.label);
    } else {
      // If equals, we either replace the number of the existing, else add to the end
      let existingIndex = boxes[boxId].findIndex((val) => val.slice(0, instruction.label.length) === instruction.label);
      if (existingIndex === -1) {
        boxes[boxId].push(instruction.label + "," + instruction.value + "," + boxId);
      } else {
        boxes[boxId][existingIndex] = instruction.label + "," + instruction.value + "," + boxId;
      }
    }
  }

  // Compute focusing power:
  let focusingPower = 0;
  // Counts per box to keep track of index:
  let countPerBox = Array.from({ length: 256 }, () => 1);
  // Flatten and just go through them:
  boxes.flat().forEach((val) => {
    let [_, focalLength, boxId] = val.split(",");
    focusingPower += (1 + Number(boxId)) * countPerBox[Number(boxId)] * Number(focalLength);
    countPerBox[boxId]++;
  });

  console.log({ day: 15, part: 2, value: focusingPower });
}

function applyHashAlgorithm(string, currentVal = 0) {
  /**
   * Determine the ASCII code for the current character of the string.
   * Increase the current value by the ASCII code you just determined.
   * Set the current value to itself multiplied by 17.
   * Set the current value to the remainder of dividing itself by 256.
   */
  let asciiCode = string.charCodeAt(0);
  currentVal += asciiCode;
  currentVal = (currentVal * 17) % 256;

  if (string.length > 1) {
    currentVal = applyHashAlgorithm(string.slice(1), currentVal);
  }

  return currentVal;
}
