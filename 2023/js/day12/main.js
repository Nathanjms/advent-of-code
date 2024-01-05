import fs from "fs";

const inputPath = "./day12/example-input";

export function partOne(input = null) {
  var input = input || inputPath;
  input = fs.readFileSync(input, "utf8");
  var inputArray = input.trim().split("\n");

  let total = 0;
  for (let i = 0; i < inputArray.length; i++) {
    let [springs, conditionRecords] = inputArray[i].split(" ");
    conditionRecords = conditionRecords.split(",").map(Number);
    total += countPossibilities(springs, conditionRecords);
  }

  console.log({ day: 12, part: 1, value: total });
}

export function partTwo(input = null) {
  var input = input || inputPath;
  input = fs.readFileSync(input, "utf8");
  var inputArray = input.trim().split("\n");

  let total = 0;
  for (const line of inputArray) {
    let [springs, conditionRecords] = line.split(" ");
    conditionRecords = conditionRecords.split(",").map(Number);

    // Enlarge both by 5:
    springs = (springs + "?").repeat(5).slice(0, -1);
    conditionRecords = Array.from({ length: 5 }, () => [...conditionRecords]).flat();

    total += countPossibilities(springs, conditionRecords);
  }

  console.log({ day: 12, part: 2, value: total });
}

// The hard bit..
/**
 *
 * @param {string} springs
 * @param {string[]} conditionRecords
 * @returns
 */
let cache = {};
function countPossibilities(springs, conditionRecords) {
  if (springs === "") {
    // If there are no more springs, this is possible if there there are no more condition records, else is not
    return conditionRecords.length === 0 ? 1 : 0;
  }

  if (conditionRecords.length === 0) {
    // If there are no more condition records, this is possible ONLY if there are no more '#'
    return springs.includes("#") ? 0 : 1;
  }

  let result = 0;

  let key = springs + conditionRecords;
  if (key in cache) {
    return cache[key];
  }

  /* 2 cases - treat ? as . or as # */

  // Case 1: Treat as .
  if ([".", "?"].includes(springs[0])) {
    // If the first element is a dot or a ?, assume it is a '.'. We then chop it off, and recursively call this again with 1 fewer element.
    // We have not used any of the conditions in this case.
    result += countPossibilities(springs.slice(1), conditionRecords);
  }

  // Case 2: Treat ? as '#'
  if (["#", "?"].includes(springs[0])) {
    if (
      conditionRecords[0] <= springs.length && // The number of springs in a row is no more than the rest of the springs
      springs.slice(0, conditionRecords[0]).indexOf(".") === -1 && // All springs up to the size of the record are NOT '.'
      (springs.length === conditionRecords[0] || springs[conditionRecords[0]] !== "#") // No springs left after OR The next spring MUST NOT be broken '#'
    ) {
      // If we are in here, start a block. So get rid of the first conditionRecord[0]+1 items, and remove the first element from condition records
      // We add one to account for either the . or ? after the last point. In this case, there MUST be a gap and so we KNOW it can't be a .
      result += countPossibilities(springs.slice(conditionRecords[0] + 1), conditionRecords.slice(1));
    } else {
      result += 0;
    }
  }

  cache[key] = result;
  return result;
}
