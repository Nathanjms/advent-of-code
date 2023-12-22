import fs from "fs";

const inputPath = "./day07/example-input";
// const inputPath = "./input";

const HAND_TYPES = {
  HIGH_CARD: 1,
  ONE_PAIR: 2,
  TWO_PAIR: 3,
  THREE_OF_A_KIND: 4,
  FULL_HOUSE: 5,
  FOUR_OF_A_KIND: 6,
  FIVE_OF_A_KIND: 7,
};

export function partOne() {
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

  var input = fs.readFileSync(inputPath, "utf8");
  var inputArray = input.trim().split("\n");

  const hands = inputArray.map((line) => {
    const hand = line.substring(0, 5).trim().split("");
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
}

export function partTwo() {
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
    J: 0, // J is now super weak on it's own
    Q: 12,
    K: 13,
    A: 14,
  };
  var input = fs.readFileSync(inputPath, "utf8");
  var inputArray = input.trim().split("\n");

  const hands = inputArray.map((line) => {
    const hand = line.substring(0, 5).trim().split("");
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

  console.log({ day: 7, part: 2, value: totalWinnings });

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

    // Now we handle each case. The difference here is the 'J' is wildcard and can act like whatever card makes the hand best.

    let handType = HAND_TYPES.HIGH_CARD;

    // 1. Five of a kind. If there is only one card, then it must be five of a kind
    if (Object.keys(cardCounts).length === 1) {
      handType = HAND_TYPES.FIVE_OF_A_KIND;
    } else if (Object.values(cardCounts).includes(4)) {
      // 2. Four of a kind. If there is one card with 4, then it must be four of a kind
      handType = HAND_TYPES.FOUR_OF_A_KIND;
    } else if (Object.values(cardCounts).includes(3) && Object.values(cardCounts).includes(2)) {
      // 3. Full house. If there is one card with 3 and one with 2, then it must be a full house
      handType = HAND_TYPES.FULL_HOUSE;
    } else if (Object.values(cardCounts).includes(3)) {
      // 4. Three of a kind. If there is one card with 3, then it must be three of a kind
      handType = HAND_TYPES.THREE_OF_A_KIND;
    } else if (Object.values(cardCounts).filter((count) => count === 2).length === 2) {
      // 5. Two pair. If there are two cards with 2, then it must be two pair
      handType = HAND_TYPES.TWO_PAIR;
    } else if (Object.values(cardCounts).filter((count) => count === 2).length === 1) {
      // 6. One pair. If there is one card with 2, then it must be one pair
      handType = HAND_TYPES.ONE_PAIR;
    }

    // 7. High card. If there are no other matches, then it must be a high card. This is already set as the default

    // If the hand contains any J's, we need to replace them with the best card and update the hand type
    if (cardHand.includes("J")) {
      // A lot more cases here
      // 1. Five of a kind. We can skip this as no need to convert
      // 2. Four of a kind. We can promote to five of a kind
      if (handType === HAND_TYPES.FOUR_OF_A_KIND) {
        return HAND_TYPES.FIVE_OF_A_KIND;
      }
      // 3. Full house. Two cases here: 3 J's and 2 of the same card, or 2 J's and 3 of the same card. Either way, we can end with 5 of a kind!
      if (handType === HAND_TYPES.FULL_HOUSE) {
        return HAND_TYPES.FIVE_OF_A_KIND;
      }
      // 4. Three of a kind. This becomes a bit more complicated as there are more options. break them down below
      if (handType === HAND_TYPES.THREE_OF_A_KIND) {
        // 3 J's can be converted to 4 of a kind. 1 J can be converted to make four of a kind.
        return HAND_TYPES.FOUR_OF_A_KIND;
      }
      // 5. Two pair. Complicated so broken down below
      if (handType === HAND_TYPES.TWO_PAIR) {
        // 5a. If there are 2 J's, there are 2 of another card. We can convert to a 4 of a kind
        if (cardCounts["J"] === 2) {
          // This means there is also another set of 2, so we can convert to 4 of a kind
          return HAND_TYPES.FOUR_OF_A_KIND;
        }
        // Otherwise, we can convert to make 3 match and 2-match, meaning we can have a full house
        return HAND_TYPES.FULL_HOUSE;
      }
      // 6. One pair. If there are 2 J's or 2 of any other card, then we can convert to a three of a kind
      if (handType === HAND_TYPES.ONE_PAIR) {
        return HAND_TYPES.THREE_OF_A_KIND;
      }
      // 7. High card. We can promote to One Pair
      if (handType === HAND_TYPES.HIGH_CARD) {
        return HAND_TYPES.ONE_PAIR;
      }
    }

    return handType;
  }
}
