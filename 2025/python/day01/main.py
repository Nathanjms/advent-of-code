import sys
from pathlib import Path

SCRIPT_DIR = Path(__file__).parent
DEFAULT_INPUT = SCRIPT_DIR / "example-input"


def main():
    path = Path(sys.argv[1]) if len(sys.argv) > 1 else DEFAULT_INPUT
    lines = [line.strip() for line in path.read_text().splitlines() if line.strip()]

    print({"Day": 1, "Part": 1, "Value": part_one(lines)})
    print({"Day": 1, "Part": 2, "Value": part_two(lines)})


def part_one(lines: list[str]) -> int:
    position = 50
    zero_hits = 0

    for line in lines:
        direction_modifier = 1 if line.startswith("R") else -1
        position = (position + (direction_modifier * int(line[1:]))) % 100

        if position == 0:
            zero_hits += 1

    return zero_hits


def part_two(lines: list[str]) -> int:
    position = 50
    zero_hits = 0

    for line in lines:
        # First, strip out all of the full cycles:
        full_rotations, partial_rotation_amount = divmod(int(line[1:]), 100)

        # ...after adding them to the number of times past zero:
        zero_hits += full_rotations

        # Now we're good to deal with the actual spinning
        direction_modifier = 1 if line.startswith("R") else -1

        new_position_with_overflow = position + (
            direction_modifier * partial_rotation_amount
        )

        # Check if we crossed zero in this last bit, but exclude if we started on zero, since we would've counted this already
        if position != 0 and not (0 < new_position_with_overflow < 100):
            zero_hits += 1

        position = new_position_with_overflow % 100

    return zero_hits


if __name__ == "__main__":
    main()
