import sys
from pathlib import Path

SCRIPT_DIR = Path(__file__).parent
DEFAULT_INPUT = SCRIPT_DIR / "example-input"


def main():
    path = Path(sys.argv[1]) if len(sys.argv) > 1 else DEFAULT_INPUT
    lines = [line.strip() for line in path.read_text().splitlines() if line.strip()]

    print({"Day": 2, "Part": 1, "Value": part_one(lines)})
    print({"Day": 2, "Part": 2, "Value": part_two(lines)})

def find_largest_concatenation(string_of_numbers: str) -> int:
    # Go through each letter
    current_largest = '00'
    for idx, char in enumerate(string_of_numbers):
        # This is checking the 'ten' digit and so can be skipped if it's smaller than the current 'ten' digit
        if (idx == len(string_of_numbers)-1) or (int(char) < int(current_largest[0])):
            continue
        
        # Otherwise, we then iterate over the remaining values to find the largest
        for j in range(len(string_of_numbers[idx+1:])):
            new_possible_largest = char + string_of_numbers[idx+1+j]
            if int(new_possible_largest) > int(current_largest):
                current_largest = new_possible_largest[0] + new_possible_largest[1]
        
    # print("largest for ", string_of_numbers, " is ", current_largest)
    return int(current_largest)

def part_one(lines: list[str]) -> int:
    sum_of_maxes = 0
    for line in lines:
        sum_of_maxes += find_largest_concatenation(line)
        
    
    return sum_of_maxes

def part_two(lines: list[str]) -> int:
    return 0


if __name__ == "__main__":
    main()
