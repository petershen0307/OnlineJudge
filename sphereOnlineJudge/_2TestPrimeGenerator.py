import unittest
from _2PrimeGenerator import prime_generator
from _2PrimeGenerator import is_prime_number


class TestPrimeGenerator(unittest.TestCase):
    def test_single_number(self):
        self.assertEqual(is_prime_number(2), True)
        self.assertEqual(is_prime_number(3), True)
        self.assertEqual(is_prime_number(97), True)
        self.assertEqual(is_prime_number(1), False)
        self.assertEqual(is_prime_number(4), False)
        self.assertEqual(is_prime_number(15), False)
        self.assertEqual(is_prime_number(16), False)
        self.assertEqual(is_prime_number(100000000), False)

    def test_prime_number(self):
        self.assertEqual(prime_generator(3, 5), [3, 5])
        self.assertEqual(prime_generator(1, 10), [2, 3, 5, 7])
        self.assertEqual(prime_generator(11, 20), [11, 13, 17, 19])
        self.assertEqual(prime_generator(21, 30), [23, 29])

if __name__ == '__main__':
    unittest.main()
