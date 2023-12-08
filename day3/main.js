import fs from "fs";

const inputPath = "./day3/input";

export function partOne() {
  var input = fs.readFileSync(inputPath, "utf8");
  var inputArray = input.trim().split("\n");

  const numbersToUse = [];

  for (let rowIndex = 0; rowIndex < inputArray.length; rowIndex++) {
    let line = inputArray[rowIndex];
    // Go through and find a number
    let matches = line.matchAll(/\d+/g);

    [...matches].forEach((match) => {
      let hasSymbol = false;
      let digitIndex = 0;
      // Split number to array of digits:
      let digits = match[0].split("");
      // For each digit of the number, check the surrounding coordinates for a symbol, until we have either found one, or run out of digits.
      while (!hasSymbol && digitIndex < digits.length) {
        // Get the coordinates surrounding the digit:
        const coordinates = getCoordinatesAroundPoint(
          match.index + digitIndex,
          rowIndex,
          line.length,
          inputArray.length
        );
        // Check if any of these coordinates have a symbol:
        for (const coordinate of coordinates) {
          if (inputArray[coordinate[0]][coordinate[1]].match(/[^\s\d\.\w]/)) {
            hasSymbol = true;
            break;
          }
        }
        digitIndex++;
      }
      // If we have found a symbol, add the number to the list:
      if (hasSymbol) {
        numbersToUse.push(Number(match[0]));
        return;
      }
    });
  }

  // We could sum as we go, but this helps with debugging, so reduce the array to a sum:
  console.log(
    "Sum: ",
    numbersToUse.reduce((a, b) => a + b, 0)
  );
}

export function partTwo() {
  var input = fs.readFileSync(inputPath, "utf8");
  var inputArray = input.trim().split("\n");

  const gearRatios = [];

  // We now need to find any * characters, and check if they are surrounded by numbers. If they are surrounded by MORE THAN ONE, store these numbers
  for (let rowIndex = 0; rowIndex < inputArray.length; rowIndex++) {
    let line = inputArray[rowIndex];
    // Get the *'s on this line:
    let matches = line.matchAll(/\*/g);

    const numbersByFirstCoordinate = {};
    [...matches].forEach((match) => {
      // Get the coordinates surrounding the *:
      const coordinates = getCoordinatesAroundPoint(match.index, rowIndex, line.length, inputArray.length);
      // Check if any of these coordinates have a number:
      for (const coordinate of coordinates) {
        if (inputArray[coordinate[0]][coordinate[1]].match(/\d/)) {
          // If we have a digit, we need to get the full number from it by going left and right until there is not a digit:
          let leftDigitIndex = coordinate[0] - 1;
          let number = inputArray[coordinate[0]][coordinate[1]];
          // Go left:
          while (leftDigitIndex >= 0 && inputArray[leftDigitIndex][coordinate[1]].match(/\d/)) {
            number = inputArray[leftDigitIndex][coordinate[1]].toString() + number.toString();
            leftDigitIndex--;
          }
          // Go right:
          let rightDigitIndex = coordinate[0] + 1;
          while (rightDigitIndex < line.length && inputArray[rightDigitIndex][coordinate[1]].match(/\d/)) {
            number = number.toString() + inputArray[rightDigitIndex][coordinate[1]].toString();
            rightDigitIndex++;
          }

          // If this number is not already in the list, add it:
          if (!numbersByFirstCoordinate.hasOwnProperty((leftDigitIndex + 1).toString() + "," + coordinate[1])) {
            numbersByFirstCoordinate[(leftDigitIndex + 1).toString() + "," + coordinate[1]] = Number(number);
          }
        }
      }
    });

    // if there are exactly two numbers, add their product to the list:
    if (Object.keys(numbersByFirstCoordinate).length === 2) {
      gearRatios.push(
        numbersByFirstCoordinate[Object.keys(numbersByFirstCoordinate)[0]] *
          numbersByFirstCoordinate[Object.keys(numbersByFirstCoordinate)[1]]
      );
    }
  }

  console.log({ gearRatios });

  console.log(
    "Sum: ",
    gearRatios.reduce((a, b) => a + b, 0)
  );
}

function getCoordinatesAroundPoint(i, j, maxI, maxJ) {
  let coordinates = [];
  // 8 Possible coordinates around a point, loop these and check if they are valid:
  for (let x = i - 1; x <= i + 1; x++) {
    for (let y = j - 1; y <= j + 1; y++) {
      coordinates.push([y, x]);
    }
  }

  // Filter out any coordinates that are out of bounds or the coordinate itself:
  coordinates = coordinates
    .filter((coordinate) => coordinate[0] >= 0 && coordinate[0] < maxJ && coordinate[1] >= 0 && coordinate[1] < maxI)
    .filter((coordinate) => !(coordinate[0] === j && coordinate[1] === i));

  return coordinates;
}
