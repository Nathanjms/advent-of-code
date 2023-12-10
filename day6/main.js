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

  console.log({ time, recordDistance, possibilitiesPerGame });

  console.log({ day: 6, part: 1, value: possibilitiesPerGame.reduce((a, b) => a * b, 1) });
}

export function partTwo() {
  var input = fs.readFileSync(inputPath, "utf8");
}
