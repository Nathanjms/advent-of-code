// Loop through each line of input.txt and print to console:
var fs = require("fs");

main();

function main(input = "input.txt") {
  var input = fs.readFileSync(input, "utf8");
  var inputArray = input.trim().split("\n");
  const sum = inputArray.reduce((sum, line) => {
    return sum + getDigitsFromLine(line);
  }, 0);

  console.log("sum: ", sum);
}

function getDigitsFromLine(line) {
  let leftIndex = 0;
  let rightIndex = line.length - 1;
  let leftDigit = null;
  let rightDigit = null;
  // Go from the left until you find a digit:
  while (leftDigit === null) {
    // If is a digit, set leftDigit to that digit
    if (!isNaN(line[leftIndex])) {
      leftDigit = line[leftIndex];
      break;
    }
    leftIndex++;
  }

  while (rightDigit === null) {
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
