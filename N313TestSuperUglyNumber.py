import unittest
import N313SuperUglyNumber

class TestCase(unittest.TestCase):
    def test_n_is_1(self):
        self.assertEqual(N313SuperUglyNumber.nthSuperUglyNumber(1, []), 1)
