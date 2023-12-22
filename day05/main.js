import fs from "fs";
import Almanac from "./Almanac.js";

const inputPath = "./day05/example-input";

export function partOne() {
  var input = fs.readFileSync(inputPath, "utf8");

  const almanac = new Almanac(input);

  let valFromSeed = almanac.getValueOfCategory("seed", 79, "location");

  let minValue = Infinity;
  const startCategory = "seed";
  const endCategory = "location";
  for (let seedId of almanac.getSeedNumbers()) {
    let valFromSeed = almanac.getValueOfCategory(startCategory, seedId, endCategory);
    if (valFromSeed < minValue) {
      minValue = valFromSeed;
    }
  }

  console.log({ day: 5, part: 1, value: minValue });
}

export function partTwo() {
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

  /**
   * This is too expensive looping through the millions of seeds.
   * Maybe we could go backwards from location, starting from the lowest value, and check if that seed is present in any of the ranges?
   * Because every number can be a map, I think we need to go from 0 and up, and can't use the locations map.
   * 1. Start from zero as the 'final output' of location
   * 2. Trace this back with the inverse map and check if the seed exists in any range
   * 3. The first match is the seed we want!
   */

  let seedNumber;
  /* Will run multiple terminal sessions and update the range for each */
  let locationNumber = 20_000_000; // Stopped it here so will start it up here again, whoops
  let end = 50_000_000;
  // let locationNumber = 0;
  while (!seedNumber && locationNumber <= end) {
    console.log({ seedNumber, locationNumber });
    let possibleSeedNumber = almanac.inverseGetValueOfCategory("seed", locationNumber, "location");
    if (isNumberInSeedRanges(seedRanges, possibleSeedNumber)) {
      seedNumber = possibleSeedNumber;
    }
    locationNumber++;
  }

  console.log({ day: 5, part: 2, value: locationNumber - 1 });

  /**
   * Determines if the given number is in the range of any of the seed ranges
   * @return {bool}
   */
  function isNumberInSeedRanges(seedRanges, possibleSeedNumber) {
    for (let { start, range } of seedRanges) {
      // Is the possibleSeedNumber in the range?
      if (possibleSeedNumber >= start && possibleSeedNumber < start + range) {
        return true;
      }
    }
    return false;
  }
}
