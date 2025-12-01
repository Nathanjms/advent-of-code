import sys


def main():
    defaultContentPath = "example-input"
    #  If user has input an argument, then that becomes the path:
    if len(sys.argv) > 1:
        defaultContentPath = sys.argv[1]

    partOne(defaultContentPath)
    partTwo(defaultContentPath)


def partOne(defaultContentPath):
    lines = []

    with open(defaultContentPath) as f:
        for line in f:
            lines.append(line.strip())

    startingVal = 50
    timesOnZero = 0

    for line in lines:
        modifier = 1 if line.startswith("R") else -1
        startingVal = (startingVal + (modifier * int(line[1:])) + 100) % 100

        if startingVal == 0:
            timesOnZero += 1

    print(
        {
            "Day": 1,
            "Part": 1,
            "Value": timesOnZero,
        }
    )


def partTwo(defaultContentPath):
    lines = []

    with open(defaultContentPath) as f:
        for line in f:
            lines.append(line.strip())

    startingVal = 50
    timesOnZero = 0

    for line in lines:
        # First, strip out all of the full cycles:
        clicks, remainder = divmod(int(line[1:]), 100)

        # ...after adding them to the number of times past zero:
        timesOnZero += clicks

        # Now we're good to deal with the actual spinning
        modifier = 1 if line.startswith("R") else -1

        tempVal = startingVal + (modifier * remainder)

        # Check if we crossed zero in this last bit, but exclude if we started on zero, since we would've counted this already
        if startingVal != 0 and ((tempVal >= 100) or (tempVal <= 0)):
            timesOnZero += 1

        startingVal = (tempVal + 100) % 100

    print({"Day": 1, "Part": 2, "Value": timesOnZero})


if __name__ == "__main__":
    main()

    exit(0)
