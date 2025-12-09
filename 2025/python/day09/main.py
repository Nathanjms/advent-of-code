import sys
from pathlib import Path
from itertools import combinations

SCRIPT_DIR = Path(__file__).parent
DEFAULT_INPUT = SCRIPT_DIR / "example-input"


def main():
    path = Path(sys.argv[1]) if len(sys.argv) > 1 else DEFAULT_INPUT
    lines = [line.strip() for line in path.read_text().splitlines() if line.strip()]

    coordinates = [tuple(map(int, line.split(","))) for line in lines]

    print({"Part": 1, "Value": part_one(coordinates)})
    print({"Part": 2, "Value": part_two(coordinates)})


def part_one(coordinates: list[tuple]) -> int:
    max_area = 0
    for coord_pair in combinations(coordinates, 2):
        area = (abs(coord_pair[1][0] - coord_pair[0][0]) + 1) * (
            abs(coord_pair[1][1] - coord_pair[0][1]) + 1
        )

        if area > max_area:
            max_area = area
    return max_area


def part_two(coordinates: list[tuple]) -> int:
    return 0


if __name__ == "__main__":
    main()
