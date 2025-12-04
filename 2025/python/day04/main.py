import sys
from pathlib import Path

SCRIPT_DIR = Path(__file__).parent
DEFAULT_INPUT = SCRIPT_DIR / "example-input"


def main():
    path = Path(sys.argv[1]) if len(sys.argv) > 1 else DEFAULT_INPUT
    lines = [line.strip() for line in path.read_text().splitlines() if line.strip()]

    print({"Part": 1, "Value": part_one(lines)})
    print({"Part": 2, "Value": part_two(lines)})


def check_surroundings(i: int, j: int, grid: list[str]) -> bool:
    roll_count = 0
    for di in [-1, 0, 1]:
        for dj in [-1, 0, 1]:
            new_i = i + di
            new_j = j + dj
            if di == 0 and dj == 0:
                continue  # skip it's own spot!
            if not (0 <= new_i < len(grid[0])) or not (0 <= new_j < len(grid)):
                continue  # skip if out of bounds
            if grid[new_i][new_j] == "@":
                roll_count += 1

    return roll_count < 4


def part_one(lines: list[str]) -> int:
    rolls = 0

    # we'll traverse the grid and for each @, check if there are less than 4 other @ in the grid around them:
    for i, line in enumerate(lines):
        for j in range(len(line)):
            if lines[i][j] == "@" and check_surroundings(i, j, lines):
                rolls += 1

    return rolls


def part_two(lines: list[str]) -> int:
    removed_rolls_count = 0

    # We keep repeating until there's no more to add, where we break
    while True:
        can_be_removed_coordinates: dict = {}
        for i, line in enumerate(lines):
            for j in range(len(line)):
                if lines[i][j] == "@" and check_surroundings(i, j, lines):
                    can_be_removed_coordinates[str(i) + "," + str(j)] = True
        if len(can_be_removed_coordinates) > 0:
            removed_rolls_count += len(can_be_removed_coordinates)

            # Now we rebuild the lines 'grid'.
            for i, line in enumerate(lines):
                line_tmp = []
                for j in range(len(line)):
                    if can_be_removed_coordinates.get(str(i) + "," + str(j)):
                        line_tmp.append(".")
                    else:
                        line_tmp.append(lines[i][j])
                lines[i] = "".join(line_tmp)
        else:
            break

    return removed_rolls_count


if __name__ == "__main__":
    main()
