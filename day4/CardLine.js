export default class CardLine {
  winningNumbers;
  userNumbers;
  id;

  constructor(line) {
    this.line = line;

    const regex = /Card\s+(\d+):(.*)\|(.*)/;
    const match = this.line.match(regex);
    this.id = match[1];
    this.winningNumbers = new Set(
      match[2]
        .trim()
        .split(" ")
        .filter((number) => number !== "")
        .map((number) => Number(number))
    );
    this.userNumbers = new Set(
      match[3]
        .trim()
        .split(" ")
        .filter((number) => number !== "")
        .map((number) => Number(number))
    );
  }

  getWinningNumbers() {
    return this.winningNumbers;
  }

  getUserNumbers() {
    return this.userNumbers;
  }

  getId() {
    return this.id;
  }
}
