import sys
from pathlib import Path
from typing import TypedDict, List

SCRIPT_DIR = Path(__file__).parent
DEFAULT_INPUT = SCRIPT_DIR / "example-input"


def main():
    path = Path(sys.argv[1]) if len(sys.argv) > 1 else DEFAULT_INPUT
    lines = [line for line in path.read_text().splitlines() if line.strip()]

    entries_part_one = [line.split() for line in lines]
    # Last line is the input:
    operators_part_one = entries_part_one[-1]
    entries_part_one = entries_part_one[:-1]

    operators_part_two = lines[-1]
    entries_part_two = lines[:-1]

    print({"Part": 1, "Value": part_one(entries_part_one, operators_part_one)})
    print({"Part": 2, "Value": part_two(entries_part_two, operators_part_two)})


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


class Node(TypedDict):
    startIndex: int
    endIndex: int
    operator: str


def part_two(entries: list[str], operators: str) -> int:
    sum_of_operations = 0
    # Build up the list of operators with their start and end
    operators_by_index: List[Node] = []
    current_index = -1  # Fisrt index is always an operator, so start this at -1

    for i in range(0, len(operators), 1):
        if operators[i] in ("*", "+"):
            operators_by_index.append(
                {"startIndex": int(i), "endIndex": -1, "operator": operators[i]}
            )
            current_index += 1
        elif len(operators) - 1 == i:
            operators_by_index[current_index]["endIndex"] = int(i)
        else:
            # Update the end
            operators_by_index[current_index]["endIndex"] = int(i)

    # Now we can loop through each one!
    for operator_entry in operators_by_index:
        result = 0
        for idx in range(
            operator_entry["endIndex"], operator_entry["startIndex"] - 1, -1
        ):
            # We have the index! Now we go down the list to build the number:
            string_value = ""
            for entry in entries:
                if entry[idx] != " ":
                    string_value += entry[idx]
            value = int(string_value) if string_value != "" else 0

            if result == 0:
                result = value
            elif operator_entry["operator"] == "+":
                result += value
            elif operator_entry["operator"] == "*":
                result *= value
            print("Value:", value)
        print(
            "Result: ",
            result,
            "operator:",
            operator_entry["operator"],
        )

        sum_of_operations += result

    return sum_of_operations


if __name__ == "__main__":
    main()
