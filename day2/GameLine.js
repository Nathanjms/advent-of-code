export default class GameLine {
  constructor(line) {
    this.line = line;
  }

  getMaxPossiblePerColor() {
    const colors = ["blue", "red", "green"];
    const colorCounts = colors.map((color) => {
      return this.getCountForColor(color);
    });

    return {
      blue: colorCounts[0],
      red: colorCounts[1],
      green: colorCounts[2],
    };
  }

  getGameId() {
    const regex = /Game\s(\d+):/;
    const matches = this.line.match(regex);
    return matches[1];
  }

  getCountForColor(color) {
    // Example regex: /(\d)\sred/
    const regex = new RegExp(`(\\d+)\\s${color}`, "g");
    const matches = this.line.matchAll(regex);
    if (!matches) {
      return 0;
    }
    // The max at one go is the highest we know we have per colour
    return Math.max(...[...matches].map((match) => Number(match[1])));
  }
}
