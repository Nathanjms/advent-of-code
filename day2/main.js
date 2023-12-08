import fs from "fs";
import GameLine from "./GameLine.js";

const inputPath = "./day2/input";

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
    const maxPossiblePerColor = gameLine.getMaxPossiblePerColor();
    let isValidGame = true;
    for (const color in maxPossiblePerColor) {
      if (maxPossiblePerColor[color] > maximumsByColor[color]) {
        isValidGame = false;
        break;
      }
    }

    if (isValidGame) {
      sumOfValidIds += Number(gameId);
    }
  });

  console.log("sumOfValidIds :", sumOfValidIds);
}
