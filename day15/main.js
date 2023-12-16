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

  console.log({ sums });
  console.log({ day: 15, part: 1, value: sums.reduce((a, b) => a + b, 0) });

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
}

export function partTwo(input = null) {
  var input = input || inputPath;
  input = fs.readFileSync(input, "utf8");
  var inputArray = input
    .trim()
    .split("\n")
    .map((line) => line.split(""));

  console.log({ day: 15, part: 2, value: "todo" });
}
