def partOne(input):
    inputString = input.readline()
    left = 0
    right = 0

    for char in inputString:
        if char == '(':
            left += 1
        elif char == ')':
            right += 1

    return left - right