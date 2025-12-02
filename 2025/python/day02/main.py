import sys
from pathlib import Path

SCRIPT_DIR = Path(__file__).parent
DEFAULT_INPUT = SCRIPT_DIR / "example-input"


def main():
    path = Path(sys.argv[1]) if len(sys.argv) > 1 else DEFAULT_INPUT
    lines = [line.strip() for line in path.read_text().split(",") if line.strip()]
    #  Split by the - character inside each lines record to get a start and end. Also make them ints:
    ranges = [line.split("-") for line in lines]

    print({"Day": 2, "Part": 1, "Value": part_one(ranges)})
    print({"Day": 2, "Part": 2, "Value": part_two(ranges)})


def digits_repeat_twice(value: int) -> bool:
    if value < 10:
        return False
    # Get the midpoint and compare if the value on each side matches:
    str_value = str(value)
    midpoint = len(str_value) // 2
    return str_value[:midpoint] == str_value[midpoint:]


def part_one(ranges: list[list[str]]) -> int:
    sum_of_repeating = 0

    for range_var in ranges:
        #  Start from the 0th index and go up to the 1st index:
        for i in range(int(range_var[0]), int(range_var[1]) + 1):
            # Do all the numbers match for this number?
            if digits_repeat_twice(i):
                sum_of_repeating += i

    return sum_of_repeating


def part_two(ranges: list[list[str]]) -> int:
    return 0


if __name__ == "__main__":
    main()
