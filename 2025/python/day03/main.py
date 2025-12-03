import sys
from pathlib import Path
from functools import lru_cache

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

def find_largest_concatenation_of_length_n(s: str, n: int) -> str:

    # Recursive function, make use of handy lru_cache for performance
    @lru_cache(None)
    def helper(start, n):
        substring = s[start:]
        # Base case: need 1 digit, so return largest digit in remainder
        if n == 1:
            return max(substring)

        # If exact number of chars left, must take all of them in that order
        if len(s) - start == n:
            return substring

        current_largest = "0" * n

        # We can only choose a leading digit from:
        # [start .. len(s)-n]
        limit = len(s) - n + 1

        for idx in range(start, limit):
            char = s[idx]
            candidate = char + helper(idx + 1, n - 1)
            if candidate > current_largest:
                current_largest = candidate

        return current_largest

    return helper(0, n)



def part_two(lines: list[str]) -> int:
    sum_of_maxes = 0
    for line in lines:
        sum_of_maxes += int(find_largest_concatenation_of_length_n(line, 12))
        
    
    return sum_of_maxes


if __name__ == "__main__":
    main()
