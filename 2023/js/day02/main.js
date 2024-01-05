import fs from "fs";
import GameLine from "./GameLine.js";

const inputPath = "./day02/example-input";

export function partOne() {
  var input = fs.readFileSync(inputPath, "utf8");
  var inputArray = input.trim().split("\n");

  const maximumsByColor = {
    red: 12,
    green: 13,
    blue: 14,
  };

  let sumOfValidIds = 0;

  inputArray.forEach((line) => {
    const gameLine = new GameLine(line);
    const gameId = gameLine.getGameId();
    const minPossiblePerColor = gameLine.getMinPossiblePerColor();
    let isValidGame = true;
    for (const color in minPossiblePerColor) {
      if (minPossiblePerColor[color] > maximumsByColor[color]) {
        isValidGame = false;
        break;
      }
    }

    if (isValidGame) {
      sumOfValidIds += Number(gameId);
    }
  });

  console.log({ day: 2, part: 1, value: sumOfValidIds });
}

export function partTwo() {
  var input = fs.readFileSync(inputPath, "utf8");
  var inputArray = input.trim().split("\n");

  let sum = 0;

  inputArray.forEach((line) => {
    const gameLine = new GameLine(line);
    const gameId = gameLine.getGameId();

    const minPossiblePerColor = gameLine.getMinPossiblePerColor();
    // The power of a set of cubes is equal to the numbers of red, green, and blue cubes multiplied together.
    const powerOfCubes = minPossiblePerColor.red * minPossiblePerColor.green * minPossiblePerColor.blue;
    sum += powerOfCubes;
  });

  console.log({ day: 2, part: 2, value: sum });
}
