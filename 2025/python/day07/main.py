import sys
from pathlib import Path
from dataclasses import dataclass
from collections import deque


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


@dataclass
class State:
    beam_idx: int
    grid_row_idx: int
    path: str


def part_two(grid: list[list[str]]) -> int:
    # Init with the location just in one place
    queue: deque[State] = deque()
    start = grid[0].index("S")
    num_timelines = 0
    visited: set[tuple[int, str]] = set()

    # Init with the second line
    queue.append(State(beam_idx=start, grid_row_idx=0, path=str(start)))

    while queue:
        state = queue.popleft()
        # Are we on the last line?
        if state.grid_row_idx == len(grid):
            num_timelines += 1
            continue

        state_key = (state.beam_idx, state.path)
        if (state_key) in visited:
            continue

        print(state.grid_row_idx)

        visited.add(state_key)

        if grid[state.grid_row_idx][state.beam_idx] != "^":
            # We just keep goin' down, no new timelines
            queue.append(
                State(
                    beam_idx=state.beam_idx,
                    grid_row_idx=state.grid_row_idx + 1,
                    path=state.path + str(state.beam_idx),
                )
            )
            continue
        else:
            new_locations = [state.beam_idx - 1, state.beam_idx + 1]
            for new_loc in new_locations:
                if 0 <= new_loc < len(grid[0]):
                    # Valid state, add to the queue!
                    # num_timelines += 1
                    queue.append(
                        State(
                            beam_idx=new_loc,
                            grid_row_idx=state.grid_row_idx + 1,
                            path=state.path + str(new_loc),
                        )
                    )

    return num_timelines


if __name__ == "__main__":
    main()
