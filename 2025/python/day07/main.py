import sys
from pathlib import Path
from collections import defaultdict


SCRIPT_DIR = Path(__file__).parent
DEFAULT_INPUT = SCRIPT_DIR / "example-input"


def main():
    path = Path(sys.argv[1]) if len(sys.argv) > 1 else DEFAULT_INPUT
    lines = [line.strip() for line in path.read_text().splitlines() if line.strip()]

    grid = [[char for char in line] for line in lines]

    print({"Part": 1, "Value": part_one(grid)})
    print({"Part": 2, "Value": part_two(grid)})


def part_one(grid: list[list[str]]) -> int:
    # Init with the location just in one place
    number_of_splits = 0
    beam_locations: list[int] = [grid[0].index("S")]
    current_row_idx = 0

    while current_row_idx + 1 < len(grid):
        current_row_idx += 1
        current_row = grid[current_row_idx]
        # Check for any splitters on this new row:
        splitters = [idx for idx, char in enumerate(current_row) if char == "^"]

        new_beam_locations = set()
        # Do any beams now spit?
        for beam_idx in beam_locations:
            print(beam_idx)
            if beam_idx in splitters:
                new_locations = [beam_idx - 1, beam_idx + 1]
                for new_loc in new_locations:
                    if (
                        0 <= new_loc < len(grid[0])
                        and new_loc not in new_beam_locations
                    ):
                        new_beam_locations.add(new_loc)
                # Bump splits by 1 as at least one of the above will be valid
                number_of_splits += 1
            else:
                new_beam_locations.add(beam_idx)

        beam_locations = list(new_beam_locations)

    return number_of_splits


def part_two(grid):
    rows = len(grid)
    cols = len(grid[0])

    # Find S
    start_col = grid[0].index("S")

    # states[row][col] = timeline_count at that position
    curr = {start_col: 1}

    for r in range(rows):
        next_row = defaultdict(int)
        row = grid[r]

        for col, count in curr.items():
            if row[col] == "^":  # splitter
                if col > 0:
                    next_row[col - 1] += count
                if col < cols - 1:
                    next_row[col + 1] += count
            else:  # empty
                next_row[col] += count

        curr = next_row

    # final number of timelines
    return sum(curr.values())


if __name__ == "__main__":
    main()
