def partOne(input):
    l = 0
    w = 0
    h = 0
    totalSquareFeet = 0

    for line in input:
        temp = line.rstrip()
        items = temp.split("x")
        l = int(items[0])
        w = int(items[1])
        h = int(items[2])
        sides = [l * w, w * h, h * l]
        sides.sort()
        totalSquareFeet += 2 * sides[0] + 2 * sides[1] + 2 * sides[2] + sides[0]

    return totalSquareFeet
