import sys
from pathlib import Path

SCRIPT_DIR = Path(__file__).parent
DEFAULT_INPUT = SCRIPT_DIR / "example-input"


def main():
    path = Path(sys.argv[1]) if len(sys.argv) > 1 else DEFAULT_INPUT
    lines = [line.strip() for line in path.read_text().splitlines() if line.strip()]

    print({"Day": 2, "Part": 1, "Value": part_one(lines)})
    print({"Day": 2, "Part": 2, "Value": part_two(lines)})

def largest_subsequence(s: str, k: int) -> str:
    result = []
    start = 0
    remaining = k

    while remaining > 0:
        # What are the index that we're allowed to choose from?
        window_end = len(s) - remaining + 1

        # Select best possible digit for this position based on the window
        best_char = max(s[start:window_end])
        result.append(best_char)

        # Move start for the next letter past the chosen digit for the next step. index will always use the first occurrence, which works for us
        start = s.index(best_char, start) + 1

        remaining -= 1

    return "".join(result)


def part_one(lines: list[str]) -> int:
    sum_of_maxes = 0
    for line in lines:
        sum_of_maxes += int(largest_subsequence(line, 2))
        
    return sum_of_maxes

def part_two(lines: list[str]) -> int:
    sum_of_maxes = 0
    for line in lines:
        sum_of_maxes += int(largest_subsequence(line, 12))
    return sum_of_maxes



if __name__ == "__main__":
    main()
