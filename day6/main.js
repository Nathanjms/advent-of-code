import fs from "fs";

const inputPath = "./day6/input";

export function partOne() {
  var input = fs.readFileSync(inputPath, "utf8");
  var inputArray = input.trim().split("\n");

  const time = [...inputArray[0].substring(5).trim().matchAll(/\d+/g)].map((match) => Number(match[0]));
  const recordDistance = [...inputArray[1].substring(9).trim().matchAll(/\d+/g)].map((match) => Number(match[0]));
  const possibilitiesPerGame = time.map(() => 0);

  // For each game, we want to go through the possibilities
  for (let i = 0; i < time.length; i++) {
    // Loop each option, skipping the first and last cases as they will never win
    for (let timeHolding = 1; timeHolding < time[i]; timeHolding++) {
      // Distance is the time remaining * speed. The speed is the time spent holding
      const timeRemaining = time[i] - timeHolding;
      const speed = timeHolding;
      if (timeRemaining * speed > recordDistance[i]) {
        possibilitiesPerGame[i]++;
      }
    }
  }
  console.log({ day: 6, part: 1, value: possibilitiesPerGame.reduce((a, b) => a * b, 1) });
}

export function partTwo() {
  var input = fs.readFileSync(inputPath, "utf8");
  var inputArray = input.trim().split("\n");

  const time = Number(inputArray[0].substring(5).replaceAll(" ", ""));
  const recordDistance = Number(inputArray[1].substring(9).replaceAll(" ", ""));

  /**
   * Now the numbers are much bigger, we can optimise this a bit. Let's do the following:
   * Go from both left and right, and stop for each one once we have beaten the record!
   */

  let leftValue, rightValue;
  let leftIndex = 1;
  let rightIndex = time;

  while (!(leftValue && rightValue) && leftIndex < rightIndex) {
    // if we haven't found the match for the left side yet, lets check then add one to the index
    if (!leftValue) {
      let timeRemaining = time - leftIndex;
      if (timeRemaining * leftIndex > recordDistance) {
        leftValue = leftIndex;
      } else {
        leftIndex++;
      }
    }
    // Similar for the right side
    if (!rightValue) {
      let timeRemaining = time - rightIndex;
      if (timeRemaining * rightIndex > recordDistance) {
        rightValue = rightIndex;
      } else {
        rightIndex--;
      }
    }
  }
  console.log({ day: 6, part: 2, value: rightIndex - leftIndex + 1 });
}
