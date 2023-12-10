import fs from "fs";
import Almanac from "./Almanac.js";

const inputPath = "./day5/input";

export function partOne() {
  var input = fs.readFileSync(inputPath, "utf8");

  const almanac = new Almanac(input);

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

  var input = fs.readFileSync(inputPath, "utf8");

  const almanac = new Almanac(input);
  const seedNumbers = almanac.getSeedNumbers();

  const seedRanges = [];
  for (let i = 0; i < seedNumbers.length; i += 2) {
    seedRanges.push({
      start: seedNumbers[i],
      range: seedNumbers[i + 1],
    });
  }

  console.log({ seedRanges });

  /**
   * This is too expensive loopng through the millions of seeds.
   * Maybe we could go backwards from location, starting from the lowest value, and check if that seed is present in any of the ranges?
   * 1. Order Locations by lowest to highest of their outputs
   * 2. Loop these until we find a match by reversing the map
   * 3. The first match is the seed we want!
   */

  let locationsSmallestToLargest = almanac.getLocationsBySize();

  // console.log("Min Value :", minValue);
}
