import sys
from pathlib import Path
from itertools import combinations
from operator import itemgetter
from functools import lru_cache


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
    coordinate_set = frozenset(coordinates)  # Make a set for faster lookup later
    # Make a list of all green tiles, start with just the borders then add the inside:
    green_tile_edges = set()
    for idx, coordinate in enumerate(coordinates):
        # Determine the direction:
        dir: tuple[int, int] = (
            coordinate[0] - coordinates[idx - 1][0],
            coordinate[1] - coordinates[idx - 1][1],
        )

        # Build the green tiles set. Probably a smarter way of doing this but it'll work
        if dir[0] > 0:
            # We're going right
            y_coord = coordinate[1]
            for x in range(coordinates[idx - 1][0] + 1, coordinate[0]):
                green_tile_edges.add((x, y_coord))
        elif dir[0] < 0:
            # We're going left
            y_coord = coordinate[1]
            for x in range(coordinates[idx - 1][0] - 1, coordinate[0], -1):
                green_tile_edges.add((x, y_coord))
        elif dir[1] > 0:
            # We're going down
            x_coord = coordinate[0]
            for y in range(coordinates[idx - 1][1] + 1, coordinate[1]):
                green_tile_edges.add((x_coord, y))
        else:
            # We're going up
            x_coord = coordinate[0]
            for y in range(coordinates[idx - 1][1] - 1, coordinate[1], -1):
                green_tile_edges.add((x_coord, y))

    max_x = max(coordinate_set, key=itemgetter(0))[0]
    max_y = max(coordinate_set, key=itemgetter(1))[1]

    green_tile_edges = frozenset(green_tile_edges)

    # We use pt 1 to give us in desc order the best area sizes, then we can abort once we get the first that works:
    areas_by_combo = []
    for coord_pair in combinations(coordinates, 2):
        area = (abs(coord_pair[1][0] - coord_pair[0][0]) + 1) * (
            abs(coord_pair[1][1] - coord_pair[0][1]) + 1
        )

        areas_by_combo.append((area, coord_pair))

    areas_by_combo.sort(reverse=True, key=lambda area_combo: area_combo[0])

    total = len(areas_by_combo)

    for idx, area_coord_pair in enumerate(areas_by_combo):
        print(f"{idx}/{total}")
        coord_pair = area_coord_pair[1]
        # We now need to go through each tile in this area and check that ALL of them are inside (or on boundary) of the grid of red/green
        if is_inside_grid(coordinate_set, green_tile_edges, coord_pair, max_x, max_y):
            return area_coord_pair[0]

    return 0


# Look at the Point-in-polygon (2023 day 10) algorithm!
@lru_cache(None)  # Cache results of this function
def is_inside_grid(
    coordinate_set: frozenset,  # Changed to frozenset for caching
    green_tile_edges: frozenset,  # Changed to frozenset for caching
    coord_pair: tuple[tuple, tuple],
    max_x: int,
    max_y: int,
):
    # Note: Boundaries are tricky since if we are on an edge, we ignore it

    edges_points = rectangle_edges_full(coord_pair)

    for x, y in edges_points:
        if (x, y) in coordinate_set or (x, y) in green_tile_edges:
            # It's one of the edges and so is 100% in the space
            continue

        # If we're here, then we need to check the number of passthroughs. If odd, then it's INSIDE
        if (
            get_number_passthroughs(
                x, y, coordinate_set, green_tile_edges, max_x, max_y
            )
            % 2
            == 1
        ):
            return False  # This entire coord_pair is invalid, so early return!

    return True

    # Check every single element
    for x in range(
        min(coord_pair[0][0], coord_pair[1][0]),
        max(coord_pair[0][0], coord_pair[1][0]) + 1,
    ):
        for y in range(
            min(coord_pair[0][1], coord_pair[1][1]),
            max(coord_pair[0][1], coord_pair[1][1]) + 1,
        ):
            if (x, y) in coordinate_set or (x, y) in green_tile_edges:
                # It's one of the edges and so is 100% in the space
                continue

            # If we're here, then we need to check the number of passthroughs. If odd, then it's INSIDE
            if (
                get_number_passthroughs(
                    x, y, coordinate_set, green_tile_edges, max_x, max_y
                )
                % 2
                == 1
            ):
                return False  # This entire coord_pair is invalid, so early return!

    return True


@lru_cache(None)
def get_number_passthroughs(x, y, coordinate_set, green_tile_edges, max_x, max_y):
    passthrough_count = 0
    # we could be super smart and choose whichever is closest the the edge of the grid to apply point-in-polygon direction
    # for now let's do this just in the x-plane always to the right:

    # Closer to right wall so make line go from x to max
    passed_corner_type = None  # Can be L J 7 or F - see 2023 day 10
    for i in range(x, max_x + 1, 1):
        if (i, y) not in coordinate_set and (i, y) not in green_tile_edges:
            continue

        if (i, y) in coordinate_set and passed_corner_type == None:
            # May not be at a boundary, depending on which type of corner:
            passed_corner_type = deduce_corner_type(i, y, green_tile_edges)
        elif (i, y) in coordinate_set and passed_corner_type != None:
            new_passed_corner_type = deduce_corner_type(i, y, green_tile_edges)
            passed_corner_type = None  # Reset
            if (passed_corner_type == "L" and new_passed_corner_type == "7") or (
                passed_corner_type == "F" and new_passed_corner_type == "J"
            ):
                passthrough_count += 1
        elif (
            (i, y) in green_tile_edges
            and (i + 1, y) not in coordinate_set
            and (i + 1, y) not in green_tile_edges
        ):
            passthrough_count += 1

    return passthrough_count


# Boldly assume no coords are right next to each other
@lru_cache(None)
def deduce_corner_type(i, y, green_tile_edges):
    if (i + 1, y) in green_tile_edges and (i, y + 1) in green_tile_edges:
        return "F"
    elif (i + 1, y) in green_tile_edges and (i, y - 1) in green_tile_edges:
        return "L"
    elif (i - 1, y) in green_tile_edges and (i, y + 1) in green_tile_edges:
        return "7"
    elif (i - 1, y) in green_tile_edges and (i, y - 1) in green_tile_edges:
        return "J"


def rectangle_edges_full(coord_pair: tuple[tuple, tuple]):
    # Extract the two opposite corners
    (x1, y1), (x2, y2) = coord_pair

    # Get all points along the edges of the rectangle
    points = []

    # Horizontal edges (constant y-values)
    for x in range(min(x1, x2), max(x1, x2) + 1):  # For x between min and max of x1, x2
        points.append((x, y1))  # Bottom edge
        points.append((x, y2))  # Top edge

    # Vertical edges (constant x-values)
    for y in range(min(y1, y2), max(y1, y2) + 1):  # For y between min and max of y1, y2
        points.append((x1, y))  # Left edge
        points.append((x2, y))  # Right edge

    # Remove duplicates in case of overlap at corners
    return list(set(points))


# Test the example input now has green tiles etc!
def test_print_grid(coordinate_set, green_tile_edges):
    for y in range(max(coordinate_set, key=itemgetter(1))[1] + 2):
        line = []
        for x in range(max(coordinate_set, key=itemgetter(0))[0] + 2):
            if (x, y) in coordinate_set:
                line.append("#")
            elif (x, y) in green_tile_edges:
                line.append("X")
            else:
                line.append(".")
        print("".join(line))


if __name__ == "__main__":
    main()
