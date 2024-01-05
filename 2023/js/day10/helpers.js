let count = 0;

export function floodFill(loopPieceCoordinates, inputArray, i, j) {
  let validSlots = [];
  // Run the fill function starting at the position given...
  fill(loopPieceCoordinates, inputArray, i, j, validSlots);
  return validSlots;
}

function fill(loopPieceCoordinates, inputArray, i, j, validSlots = []) {
  count++;
  // If i is less than 0 or greater equals to the length of inputArray...
  // Or, If j is less than 0 or greater equals to the length of inputArray[0]...
  if (i < 0 || i >= inputArray.length || j < 0 || j >= inputArray[0].length) return;
  // If i,j is a loop piece, we can stop. Note that we have doubled the matrix so half them here!
  if (loopPieceCoordinates.includes(i * 2 + "," + j * 2)) return;
  // If is already valid, skip
  if (validSlots.includes(i + "," + j)) return;
  // Add to valid slots
  validSlots.push(i + "," + j);
  // Make four recursive calls to the function with (i-1, j), (i+1, j), (i, j-1) and (i, j+1)...
  fill(loopPieceCoordinates, inputArray, i - 1, j, validSlots);
  fill(loopPieceCoordinates, inputArray, i + 1, j, validSlots);
  fill(loopPieceCoordinates, inputArray, i, j - 1, validSlots);
  fill(loopPieceCoordinates, inputArray, i, j + 1, validSlots);
}

export function getEnclosedCoords(newInputArray, coordinatesThatCanEscape) {
  let enclosedCoords = [];
  for (let i = 0; i < newInputArray.length; i++) {
    for (let j = 0; j < newInputArray[0].length; j++) {
      if (newInputArray[i][j] === "#") {
        continue;
      }
      if (
        !coordinatesThatCanEscape.includes(`${i},${j}`) &&
        i !== 0 &&
        j !== 0 &&
        i !== newInputArray.length &&
        j !== newInputArray[0].length
      ) {
        // Next we convert back to regular coordinate map. This is simply divide the  part by 2:
        enclosedCoords.push(`${i / 2},${j / 2}`);
      }
    }
  }
  return enclosedCoords;
}
