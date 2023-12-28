import day24_part2
from unittest import TestCase


class Test(TestCase):
    def test_solve_part2(self):
        hailstones = day24_part2.load_input("../../../input/day24/input_test")

        r = day24_part2.solve_part2(hailstones)

        self.assertEqual(47, r)
