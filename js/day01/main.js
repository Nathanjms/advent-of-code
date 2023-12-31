import fs from "fs";

const inputPath = "./day01/example-input";

export function partOne() {
  var input = fs.readFileSync(inputPath + "1", "utf8");
  var inputArray = input.trim().split("\n");
  const sum = inputArray.reduce((sum, line) => {
    return sum + getDigitsFromLine(line);
  }, 0);

  console.log({ day: 1, part: 1, value: sum });

  function getDigitsFromLine(line) {
    console.log(line);
    let leftIndex = 0;
    let rightIndex = line.length - 1;
    let leftDigit = null;
    let rightDigit = null;
    // Go from the left until you find a digit:
    while (leftDigit === null && leftIndex <= line.length) {
      // If is a digit, set leftDigit to that digit
      if (!isNaN(line[leftIndex])) {
        leftDigit = line[leftIndex];
        break;
      }
      leftIndex++;
    }

    while (rightDigit === null && rightIndex >= 0) {
      // If is a digit, set rightDigit to that digit
      if (!isNaN(line[rightIndex])) {
        rightDigit = line[rightIndex];
        break;
      }
      rightIndex--;
    }

    // concat the digits together and return the result
    return Number(leftDigit.toString() + rightDigit.toString());
  }
}

export function partTwo() {
  const wordMap = {
    one: 1,
    two: 2,
    three: 3,
    four: 4,
    five: 5,
    six: 6,
    seven: 7,
    eight: 8,
    nine: 9,
  };

  var input = fs.readFileSync(inputPath + "2", "utf8");
  var inputArray = input.trim().split("\n");

  const sum = inputArray.reduce((sum, line) => {
    return sum + getDigitsFromLine(line);
  }, 0);

  console.log({ day: 1, part: 2, value: sum });

  function getDigitsFromLine(line) {
    return Number(getFirstNumber(line).toString() + getLastNumber(line).toString());
  }

  function getFirstNumber(line) {
    // For each line, go by character until we either hit a digit or a word in the wordMap:
    let characters = "";
    for (let i = 0; i < line.length; i++) {
      if (!isNaN(line[i])) {
        // We've found a digit before finding a word, so return the digit:
        return Number(line[i]);
      }
      characters += line[i];
      let num = getNumberInWordMap(characters);
      if (num) {
        return num;
      }
    }
  }

  function getLastNumber(line) {
    // For each line, go by character until we either hit a digit or a word in the wordMap:
    let characters = "";
    for (let i = line.length - 1; i >= 0; i--) {
      if (!isNaN(line[i])) {
        // We've found a digit before finding a word, so return the digit:
        return Number(line[i]);
      }
      characters = line[i] + characters;
      // If the characters container any of the words in the word map, return the digit for that word:
      let num = getNumberInWordMap(characters);
      if (num) {
        return num;
      }
    }
  }

  function getNumberInWordMap(characters) {
    for (let word in wordMap) {
      if (characters.includes(word)) {
        return wordMap[word];
      }
    }
  }
}
