def partTwo(input):
    l = 0
    w = 0
    h = 0
    totalLengthOfRibbon = 0

    for line in input:
        temp = line.rstrip()
        items = temp.split("x")
        sides = [int(items[0]), int(items[1]), int(items[2])]
        sides.sort()
        totalLengthOfRibbon += (
            2 * sides[0] + 2 * sides[1] + sides[0] * sides[1] * sides[2]
        )

    return totalLengthOfRibbon
