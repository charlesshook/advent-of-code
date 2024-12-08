def partTwo(input):
    inputString = input.readline()
    floor = 0
    count = 0

    for char in inputString:
        if char == '(':
            floor += 1
            count += 1
        elif char == ')':
            floor -= 1
            count += 1

        if floor == -1:
            break

    return count