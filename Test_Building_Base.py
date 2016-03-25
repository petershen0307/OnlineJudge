import unittest
from Building_Base import Building


class TestCase(unittest.TestCase):
    def test_function(self):
        self.assertEqual(Building(1, 2, 2, 2).corners(),
                         {'north-west': [3, 2], 'north-east': [3, 4], 'south-west': [1, 2], 'south-east': [1, 4]})
        self.assertEqual(Building(1, 2.5, 4.2, 1.25).area(), 5.25)
        self.assertEqual(Building(1, 2.5, 4.2, 1.25, 101).volume(), 530.25)
        self.assertEqual(str(Building(0, 0, 10.5, 2.546)), 'Building(0, 0, 10.5, 2.546, 10)')

if __name__ == '__main__':
    unittest.main()
