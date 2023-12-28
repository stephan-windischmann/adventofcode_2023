import z3

def load_input(filename):
    hailstones = []
    with open(filename) as f:
        for l in f.readlines():
            point = ((int(x) for x in l.split("@")[0].split(",")))
            velocity = ((int(x) for x in l.split("@")[1].split(",")))
            hailstones.append((tuple(point), tuple(velocity)))

    return hailstones


def solve_part2(hailstones):
    x = z3.Int("x")
    y = z3.Int("y")
    z = z3.Int("z")
    dx = z3.Int("dx")
    dy = z3.Int("dy")
    dz = z3.Int("dz")

    s = z3.Solver()

    for i, stone in enumerate(hailstones):
        t = z3.Int(f"i{i}")
        s.add(x + dx * t == stone[0][0] + stone[1][0] * t)
        s.add(y + dy * t == stone[0][1] + stone[1][1] * t)
        s.add(z + dz * t == stone[0][2] + stone[1][2] * t)

    s.check()
    m = s.model()

    return m[x].as_long() + m[y].as_long() + m[z].as_long()


def main():
    hailstones = load_input("../../../input/day24/input")

    r2 = solve_part2(hailstones)

    print(f"Day 24 part 2 solution: {r2}")


if __name__ == "__main__":
    main()
