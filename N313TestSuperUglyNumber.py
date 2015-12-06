import unittest
import N313SuperUglyNumber


class TestCase(unittest.TestCase):
    def test_function(self):
        self.assertEqual(N313SuperUglyNumber.nthSuperUglyNumber(1, []), 1)
        # [1, 2, 4, 7, 8, 13, 14, 16, 19, 26, 28, 32]
        self.assertEqual(N313SuperUglyNumber.nthSuperUglyNumber(12, [2, 7, 13, 19]), 32)

if __name__ == '__main__':
    unittest.main()
