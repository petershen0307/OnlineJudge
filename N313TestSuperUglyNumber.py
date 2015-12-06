import unittest
import N313SuperUglyNumber


class TestCase(unittest.TestCase):
    def test_function(self):
        self.assertEqual(N313SuperUglyNumber.nthSuperUglyNumber(1, []), 1)
        # [1, 2, 4, 7, 8, 13, 14, 16, 19, 26, 28, 32]
        self.assertEqual(N313SuperUglyNumber.nthSuperUglyNumber(12, [2, 7, 13, 19]), 32)
        # [1, 2, 3, 4, 5, 6, 8, 9, 10]
        self.assertEqual(N313SuperUglyNumber.nthSuperUglyNumber(9, [2, 3, 5]), 10)
        self.assertEqual(N313SuperUglyNumber.nthSuperUglyNumber(650,
                                                                [3, 11, 13, 19, 29, 31, 41, 43, 47, 59, 61, 71, 79, 83,
                                                                 89, 97, 103, 107, 137, 139, 149, 163, 173, 179, 191,
                                                                 197, 211, 223, 229, 239]), 10919)


if __name__ == '__main__':
    unittest.main()
