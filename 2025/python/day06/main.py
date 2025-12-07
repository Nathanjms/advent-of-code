import sys
from pathlib import Path

SCRIPT_DIR = Path(__file__).parent
DEFAULT_INPUT = SCRIPT_DIR / "example-input"


def main():
    path = Path(sys.argv[1]) if len(sys.argv) > 1 else DEFAULT_INPUT
    lines = [line.strip() for line in path.read_text().splitlines() if line.strip()]

    entries = [line.split() for line in lines]
    # Last line is the input:
    operators = entries[-1]
    entries = entries[:-1]

    print({"Part": 1, "Value": part_one(entries, operators)})
    print({"Part": 2, "Value": part_two(entries, operators)})


def part_one(entries: list[list[str]], operators: list[str]) -> int:
    sum_of_operations = 0

    for idx, operator in enumerate(operators):
        for jdx, row in enumerate(entries):
            row_value = int(row[idx])
            if jdx == 0:
                result = int(row_value)
                continue
            if operator == "+":
                result += int(row_value)
                continue
            if operator == "*":
                result *= int(row_value)
                continue
        sum_of_operations += result
    return sum_of_operations


def part_two(entries: list[list[str]], operators: list[str]) -> int:
    return 0


if __name__ == "__main__":
    main()
