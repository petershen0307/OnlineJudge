import unittest
import time
import N313SuperUglyNumber


class TestCase(unittest.TestCase):
    def test_function(self):
        methods = [N313SuperUglyNumber.method3]  # N313SuperUglyNumber.method1,
        for method in methods:
            started = time.time()
            self.assertEqual(method(1, []), 1)
            print('1 spend time:', (time.time() - started))

            # [1, 2, 4, 7, 8, 13, 14, 16, 19, 26, 28, 32]
            started = time.time()
            self.assertEqual(method(12, [2, 7, 13, 19]), 32)
            print('12 spend time:', (time.time() - started))

            # [1, 2, 3, 4, 5, 6, 8, 9, 10]
            started = time.time()
            self.assertEqual(method(9, [2, 3, 5]), 10)
            print('9 spend time:', (time.time() - started))

            started = time.time()
            self.assertEqual(method(650, [3, 11, 13, 19, 29, 31, 41, 43, 47, 59, 61, 71, 79, 83,
                                          89, 97, 103, 107, 137, 139, 149, 163, 173, 179, 191,
                                          197, 211, 223, 229, 239]), 10919)
            print('650 spend time:', (time.time() - started))

            started = time.time()
            self.assertEqual(method(200000,
                                    [2, 3, 5, 13, 19, 29, 31, 41, 43, 53, 59, 73, 83, 89, 97, 103, 107, 109, 127, 137,
                                     139, 149, 163, 173, 179, 193, 197, 199, 211, 223, 227, 229, 239, 241, 251, 257,
                                     263, 269, 271, 281, 317, 331, 337, 347, 353, 359, 367, 373, 379, 389, 397, 409,
                                     419, 421, 433, 449, 457, 461, 463, 479, 487, 509, 521, 523, 541, 547, 563, 569,
                                     577, 593, 599, 601, 613, 619, 631, 641, 659, 673, 683, 701, 709, 719, 733, 739,
                                     743, 757, 761, 769, 773, 809, 811, 829, 857, 859, 881, 919, 947, 953, 967, 971]),
                             3610419)
            print('200000 spend time(s):', (time.time() - started))


if __name__ == '__main__':
    unittest.main()
