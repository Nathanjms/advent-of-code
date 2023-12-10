import fs from "fs";
import Almanac from "./Almanac.js";

const inputPath = "./day5/example-input";

export function partOne() {
  var input = fs.readFileSync(inputPath, "utf8");

  const almanac = new Almanac(input);

  console.log({ "valueOfCategory: ": almanac.getValueOfCategory("seed", 79, "location") });

  let minValue = Infinity;

  const startCategory = "seed";
  const endCategory = "location";
  for (let seedId of almanac.getSeedNumbers()) {
    let valFromSeed = almanac.getValueOfCategory(startCategory, seedId, endCategory);
    if (valFromSeed < minValue) {
      minValue = valFromSeed;
    }
  }

  console.log("Min Value :", minValue);
}

export function partTwo() {
  var input = fs.readFileSync(inputPath, "utf8");

  console.log("sum :", 0);
}

function getSeedNumbers(lineZero) {
  return [...lineZero.matchAll(/\d+/g)].map((match) => Number(match[0]));
}
