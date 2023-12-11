import fs from "fs";

// const inputPath = "./day7/example-input";
const inputPath = "./input";

const HAND_TYPES = {
  HIGH_CARD: 1,
  ONE_PAIR: 2,
  TWO_PAIR: 3,
  THREE_OF_A_KIND: 4,
  FULL_HOUSE: 5,
  FOUR_OF_A_KIND: 6,
  FIVE_OF_A_KIND: 7,
};

const CARD_VALUE_MAP = {
  2: 2,
  3: 3,
  4: 4,
  5: 5,
  6: 6,
  7: 7,
  8: 8,
  9: 9,
  T: 10,
  J: 11,
  Q: 12,
  K: 13,
  A: 14,
};

export function partOne() {
  var input = fs.readFileSync(inputPath, "utf8");
  var inputArray = input.trim().split("\n");

  const hands = inputArray.map((line) => {
    const hand = line
      .substring(0, 5)
      .trim()
      .split("")
    //   .sort((a, b) => CARD_VALUE_MAP[b] - CARD_VALUE_MAP[a]);
    return {
      hand,
      bid: Number(line.substring(6).trim()),
      type: getTypeOfHand(hand),
    };
  });

  let totalWinnings = 0;

  // Now we need to order hands by rank. First we will try and do it by type, and if there are any with the same type, we will compare the cards
  const sortedHands = hands.sort((a, b) => {
    if (a.type === b.type) {
      // We need to compare the cards
      for (let i = 0; i < a.hand.length; i++) {
        if (a.hand[i] === b.hand[i]) {
          continue;
        }
        return CARD_VALUE_MAP[b.hand[i]] - CARD_VALUE_MAP[a.hand[i]];
      }
    }
    return b.type - a.type;
  });

  // Total winnings is the rank of the hand * the bid
  for (let i = 0; i < sortedHands.length; i++) {
    let rank = sortedHands.length - i;
    sortedHands[i].rank = rank;
    totalWinnings += sortedHands[i].bid * (sortedHands.length - i);
  }

  console.log({ day: 7, part: 1, value: totalWinnings });
}

export function partTwo() {
  var input = fs.readFileSync(inputPath, "utf8");
  var inputArray = input.trim().split("\n");

  console.log({ day: 7, part: 2, value: "todo" });
}

/**
 *
 * @param {string[]} cardHand
 */
function getTypeOfHand(cardHand) {
  const cardCounts = {};
  cardHand.forEach((card) => {
    if (cardCounts[card]) {
      cardCounts[card]++;
    } else {
      cardCounts[card] = 1;
    }
  });

  // Now we handle each case

  // 1. Five of a kind. If there is only one card, then it must be five of a kind
  if (Object.keys(cardCounts).length === 1) {
    return HAND_TYPES.FIVE_OF_A_KIND;
  }

  // 2. Four of a kind. If there is one card with 4, then it must be four of a kind
  if (Object.values(cardCounts).includes(4)) {
    return HAND_TYPES.FOUR_OF_A_KIND;
  }

  // 3. Full house. If there is one card with 3 and one with 2, then it must be a full house
  if (Object.values(cardCounts).includes(3) && Object.values(cardCounts).includes(2)) {
    return HAND_TYPES.FULL_HOUSE;
  }

  // 4. Three of a kind. If there is one card with 3, then it must be three of a kind
  if (Object.values(cardCounts).includes(3)) {
    return HAND_TYPES.THREE_OF_A_KIND;
  }

  // 5. Two pair. If there are two cards with 2, then it must be two pair
  if (Object.values(cardCounts).filter((count) => count === 2).length === 2) {
    return HAND_TYPES.TWO_PAIR;
  }

  // 6. One pair. If there is one card with 2, then it must be one pair
  if (Object.values(cardCounts).filter((count) => count === 2).length === 1) {
    return HAND_TYPES.ONE_PAIR;
  }

  return HAND_TYPES.HIGH_CARD;
}
