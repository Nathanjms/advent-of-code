import fs from "fs";

const inputPath = "./day22/example-input";

export function partOne(input = null) {
  var input = input || inputPath;
  input = fs.readFileSync(input, "utf8");
  var sandBlocks = parseInput(input);

  // First we need to 'drop' all of the sandBlocks into their base positions, going down in Z

  // It's worth noting that we've typically dealt in i,j (and k) - but these sandBlocks are in x,y & z.

  // We will start with the lowest Z index block and 'drop' it as far as it can. So let's start by ordering the blocks by the lowest first
  sandBlocks = sandBlocks.sort((a, b) => {
    // The first coordinate for each is the lower one, so grab the first coordinate, and then compare the 3rd element (z-axis)
    return a[0][2] - b[0][2];
  });

  /* Drop */
  for (let i = 0; i < sandBlocks.length; i++) {
    const sandBlock = sandBlocks[i];
    // The floor is at 0, so the best z value is 1
    let bestZValue = 1;
    if (sandBlock[0][2] === bestZValue) {
      // Any already at the floor can chill
      continue;
    }

    // Otherwise, we can compare against all the blocks, and keep checking until we find a blocker. If no blocker, then we're at max
    // We can use the fact that the blocks are ordered, so that we only need to compare overlaps with any previous ones (lower down blocks)
    for (let j = i - 1; j >= 0; j--) {
      // We can drop past it if the x & y values do not overlap, so we only need to check in the x-y plane
      let compareBlock = sandBlocks[j];
      if (overlapsInXY(compareBlock, sandBlock)) {
        bestZValue = Math.max(compareBlock[1][2] + 1, bestZValue);
      }
    }

    // If we can drop, go to that level
    if (bestZValue) {
      sandBlocks[i] = [
        [sandBlock[0][0], sandBlock[0][1], bestZValue],
        [sandBlock[1][0], sandBlock[1][1], sandBlock[1][2] - (sandBlock[0][2] - bestZValue)],
      ];
    }
  }

  // Sort them again in-case they have now got out of z order:
  sandBlocks = sandBlocks.sort((a, b) => {
    // The first coordinate for each is the lower one, so grab the first coordinate, and then compare the 3rd element (z-axis)
    return a[0][2] - b[0][2];
  });

  console.log({ day: 22, part: 1, value: getBlastPossibilities() });

  function getBlastPossibilities() {
    let blastPossibilities = 0;

    for (let i = 0; i < sandBlocks.length; i++) {
      const lowerSandBlock = sandBlocks[i];
      // Are there any this one is holding up?
      const sandBlocksBeingHeldUp = sandBlocks
        .filter((b) => b[0][2] === lowerSandBlock[1][2] + 1)
        .filter((b) => overlapsInXY(b, lowerSandBlock));

      // Simple case: it's holding up none and can be blasted!
      if (sandBlocksBeingHeldUp.length === 0) {
        blastPossibilities++;
        continue;
      }

      // Otherwise, for ALL of the blocks being held, is there another that can do the job?
      const blocksAtSameHeight = sandBlocks
        .filter((_, idx) => i !== idx)
        .filter((b) => b[1][2] === lowerSandBlock[1][2]);
      // First if there are no others at this height, its a definite no.
      if (!blocksAtSameHeight.length) {
        continue;
      }

      let canBlast = true;
      for (const blockBeingHeldUp of sandBlocksBeingHeldUp) {
        // If NONE of the blocks at the same height overlap, then the original is crucial and can't be blasted
        if (!blocksAtSameHeight.some((b) => overlapsInXY(blockBeingHeldUp, b))) {
          canBlast = false;
          break;
        }
      }

      // If we've made it to here, we can blast (I think?)
      if (canBlast) {
        blastPossibilities++;
      }
    }
    return blastPossibilities;
  }
}

export function partTwo(input = null) {
  var input = input || inputPath;
  input = fs.readFileSync(input, "utf8");
  var sandBlocks = parseInput(input);
  /* Drop the sand blocks the same as last time! */

  // We will start with the lowest Z index block and 'drop' it as far as it can. So let's start by ordering the blocks by the lowest first
  sandBlocks = sandBlocks.sort((a, b) => {
    // The first coordinate for each is the lower one, so grab the first coordinate, and then compare the 3rd element (z-axis)
    return a[0][2] - b[0][2];
  });

  /* Drop */
  for (let i = 0; i < sandBlocks.length; i++) {
    const sandBlock = sandBlocks[i];
    // The floor is at 0, so the best z value is 1
    let bestZValue = 1;
    if (sandBlock[0][2] === bestZValue) {
      // Any already at the floor can chill
      continue;
    }

    // Otherwise, we can compare against all the blocks, and keep checking until we find a blocker. If no blocker, then we're at max
    // We can use the fact that the blocks are ordered, so that we only need to compare overlaps with any previous ones (lower down blocks)
    for (let j = i - 1; j >= 0; j--) {
      // We can drop past it if the x & y values do not overlap, so we only need to check in the x-y plane
      let compareBlock = sandBlocks[j];
      if (overlapsInXY(compareBlock, sandBlock)) {
        bestZValue = Math.max(compareBlock[1][2] + 1, bestZValue);
      }
    }

    // If we can drop, go to that level
    if (bestZValue) {
      sandBlocks[i] = [
        [sandBlock[0][0], sandBlock[0][1], bestZValue],
        [sandBlock[1][0], sandBlock[1][1], sandBlock[1][2] - (sandBlock[0][2] - bestZValue)],
      ];
    }
  }

  // Sort them again in-case they have now got out of z order:
  sandBlocks = sandBlocks.sort((a, b) => {
    return a[0][2] - b[0][2];
  });

  /**
   * This time, instead of checking to see which would cause others to fall and avoid them, we want to actually blast each one, and find the one
   * with the longest chain reaction of falling sand
   */

  let chainByBlock = sandBlocks.map((v) => 0);

  // Add a key to each block as it's 3rd element (2nd index):
  sandBlocks = sandBlocks.map((v, idx) => [...v, idx]);

  for (let i = 0; i < sandBlocks.length; i++) {
    let fallingBlocks = new Set();
    const queue = [sandBlocks[i]];
    while (queue.length) {
      let block = queue.shift(1);
      let brickKey = block[2];
      fallingBlocks.add(block[2]);

      // Are there any this one is holding up?
      const sandBlocksBeingHeldUp = sandBlocks
        .filter((b) => b[0][2] === block[1][2] + 1)
        .filter((b) => overlapsInXY(b, block));

      const blocksAtSameHeight = sandBlocks
        .filter((v) => v[2] !== brickKey)
        .filter((b) => b[1][2] === block[1][2])
        .filter((v) => !fallingBlocks.has(v[2]));

      if (!sandBlocksBeingHeldUp.length) {
        // Nothing to add
      } else if (!blocksAtSameHeight.length) {
        // ALL added to queue:
        sandBlocksBeingHeldUp.forEach((block) => queue.push(block));
      } else {
        // (Possibly) SOME to add
        // Get all of the blocks being held up, and check if they are still supported by any non-falling blocks
        for (const blockBeingHeldUp of sandBlocksBeingHeldUp) {
          if (!blocksAtSameHeight.some((b) => overlapsInXY(blockBeingHeldUp, b))) {
            queue.push(blockBeingHeldUp);
          }
        }
      }
    }

    // Add these to the queue, subtracting 1 for the original block that has not fallen, but has been blasted 🔫
    chainByBlock[i] = fallingBlocks.size - 1;
  }

  console.log({ day: 22, part: 2, value: chainByBlock.reduce((a, b) => a + b, 0) });
}

function parseInput(input) {
  return input
    .trim()
    .split("\n")
    .map((line) => line.split("~").map((coord) => coord.split(",").map((v) => Number(v))));
}

function overlapsInXY(blockOne, blockTwo) {
  // Intersect if the lower value of each range is less than  (or equal to) the higher value of the other range. Can be simplified with min/max
  return (
    Math.max(blockOne[0][0], blockTwo[0][0]) <= Math.min(blockOne[1][0], blockTwo[1][0]) &&
    Math.max(blockOne[0][1], blockTwo[0][1]) <= Math.min(blockOne[1][1], blockTwo[1][1])
  );
}
