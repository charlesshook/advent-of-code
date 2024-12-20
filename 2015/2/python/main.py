import argparse
import part_one
import part_two


def main():
    parser = argparse.ArgumentParser()
    parser.add_argument("-part", type=int, required=True)
    args = parser.parse_args()
    print("Running challenge part {part}.".format(part=args.part))

    input = open("../input.txt", "r")

    if args.part == 1:
        answer = part_one.partOne(input)
        print("The solution to part 1 is: {answer}.".format(answer=answer))
    elif args.part == 2:
        answer = part_two.partTwo(input)
        print("The solution to part 2 is: {answer}.".format(answer=answer))
    else:
        print("Part {part} is not a valid option.".format(part=args.part))

    input.close()


if __name__ == "__main__":
    main()
