export default class Almanac {
  CATEGORIES_BY_ORDER = ["seed", "soil", "fertilizer", "water", "light", "temperature", "humidity", "location"];
  valuesPerMap;
  constructor(contents) {
    this.contents = contents.trim();
    this.contentsByLine = this.contents.split("\n");
    this.setValuesForAllMaps();
  }

  getSeedNumbers() {
    return [...this.contentsByLine[0].matchAll(/\d+/g)].map((match) => Number(match[0]));
  }

  computeValuesInMap(map) {
    let rowIndex;
    // Find the row Index that has the map as its title:
    for (let i = 0; i < this.contents.length; i++) {
      if (this.contentsByLine[i]?.match(map)) {
        // We've found the start!
        rowIndex = i;
        break;
      }
    }

    // Now go through each line until we get to a blank one (or reach the end)
    const mapValues = [];
    rowIndex++; // Add one to start (we are currently on the 'title' index)
    while (true) {
      if (rowIndex > this.contentsByLine.length || !this.contentsByLine[rowIndex]) {
        break;
      }

      mapValues.push(this.parseMapLine(this.contentsByLine[rowIndex]));
      rowIndex++;
    }

    return mapValues;
  }

  buildMapTitle(category) {
    this.CATEGORIES_BY_ORDER.indexOf(category);
  }

  parseMapLine(mapLine) {
    // Example Input: '50 98 2'
    let [output, input, length] = mapLine.split(" ").map((n) => Number(n));
    return { input, output, length };
  }

  /**
   * EG. seed, 10, soil will look at the map in order to determine the value.
   * @param {string} startCategory
   * @param {int} starValue
   * @param {string} endCategory
   */
  getValueOfCategory(startCategory, startValue, endCategory) {
    const startIndex = this.CATEGORIES_BY_ORDER.indexOf(startCategory);
    const endIndex = this.CATEGORIES_BY_ORDER.indexOf(endCategory);
    let value = startValue;
    for (let i = startIndex; i < endIndex; i++) {
      // Find out which it is by it's input, then use that output
      for (const map of this.valuesPerMap[i]) {
        // If the value is within the range of the map, then get the actual new number out
        if (value >= map.input && value < map.input + map.length) {
          // The value becomes the output value + the difference between the
          value = map.output + (value - map.input);
          break;
        }
      }
    }

    return value;
  }

  setValuesForAllMaps() {
    const startIndex = 0;
    const endIndex = this.CATEGORIES_BY_ORDER.length - 1;
    let valuesForEachMap = [];
    for (let i = startIndex; i < endIndex; i++) {
      let mapTitle = `${this.CATEGORIES_BY_ORDER[i]}-to-${this.CATEGORIES_BY_ORDER[i + 1]} map:`;

      let mapValues = this.computeValuesInMap(mapTitle);
      valuesForEachMap.push([...Object.values(mapValues)]);
    }

    this.valuesPerMap = valuesForEachMap;
  }

  /**
   * EG. (soil, 10, seed) will look at the map in REVERSE order to determine the value.
   * @param {string} startCategory
   * @param {int} endValue
   * @param {string} endCategory
   */
  inverseGetValueOfCategory(startCategory, endValue, endCategory) {
    const startIndex = this.CATEGORIES_BY_ORDER.indexOf(startCategory);
    const endIndex = this.CATEGORIES_BY_ORDER.indexOf(endCategory);
    let value = endValue;
    for (let i = endIndex; i > startIndex; i--) {
      // Find out which it is by it's output, then use it's input
      for (const map of this.valuesPerMap[i - 1]) {
        // If the value is within the range of the map, then get the actual new number out
        if (value >= map.output && value < map.output + map.length) {
          // The value becomes the input value + the difference
          value = map.input + (value - map.output);
          break;
        }
      }
    }

    return value;
  }
}
