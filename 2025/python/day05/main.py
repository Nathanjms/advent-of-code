import sys
from pathlib import Path

SCRIPT_DIR = Path(__file__).parent
DEFAULT_INPUT = SCRIPT_DIR / "example-input"


def main():
    path = Path(sys.argv[1]) if len(sys.argv) > 1 else DEFAULT_INPUT
    block_one, block_two = path.read_text().strip().split("\n\n")

    coordinates = [
        tuple(map(int, line.split("-")))
        for line in block_one.splitlines()
        if line.strip()
    ]
    values = [int(x) for x in block_two.splitlines()]

    print({"Part": 1, "Value": part_one(coordinates, values)})
    print({"Part": 2, "Value": part_two(coordinates)})


def part_one(coordinates: list[tuple], values: list[int]) -> int:
    # We'll start by ordering the coordinates. I think what we may end up doing is that 'boundary shift' thing, but start simpler.
    coordinates.sort(key=lambda coord: coord[0])

    # Try a for loop and see how it performs
    num_fresh = 0
    for val in values:
        is_fresh = False
        if val < coordinates[0][0]:
            continue  # We've ordered them by start, so can immediately skip as not fresh if below the first coord

        for coord in coordinates:
            if coord[0] <= val <= coord[1]:
                is_fresh = True
                break
        if is_fresh:
            num_fresh += 1

    return num_fresh


def part_two(coordinates: list[tuple]) -> int:
    # this now is about combining ranges, how is this done hmm
    coordinates.sort(key=lambda coord: coord[0])

    # Instead of (start,end), we build a list of (start, range), so 3->5 is now (3,2) since it goes from 3 for 2 positions
    start_and_range_list: list[tuple] = []

    # Since we've sorted the start values, let's try a for loop where we store the size ansd not the end coord?
    for coord in coordinates:
        # First case differs, we'll handle it explicitly for ease
        if len(start_and_range_list) == 0:
            start_and_range_list.append((coord[0], coord[1] - coord[0]))
            continue

        idx = len(start_and_range_list) - 1

        # Next, we check if the next one can be combined, and add to it to the latest one in the list if so, otherwise we start a new one
        if coord[0] <= start_and_range_list[idx][0] + start_and_range_list[idx][1] + 1:
            # Can be combined - we take whichever ends up with the largest as the new range
            if (
                start_and_range_list[idx][0] + start_and_range_list[idx][1] + 1
                <= coord[1]
            ):
                # Starts from the same start, but now finishes at end of of the new coord, so use the new range it covers
                start_and_range_list[idx] = (
                    start_and_range_list[idx][0],
                    coord[1] - start_and_range_list[idx][0],
                )
        else:
            # Start a new range instead!
            start_and_range_list.append((coord[0], coord[1] - coord[0]))

    # The total is now the sum of all the ranges, but because they are inclusive, we need to add 1 for each list item:
    sum_of_ranges = len(start_and_range_list) + sum(
        range_val for _, range_val in start_and_range_list
    )

    return sum_of_ranges


if __name__ == "__main__":
    main()
