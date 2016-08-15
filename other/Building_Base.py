class Building:
    def __init__(self, south, west, width_we, width_ns, height=10):
        self.south = south
        self.west = west
        self.width_we = width_we
        self.width_ns = width_ns
        self.height = height

    @property
    def __repr__(self):
        return 'Building({south}, {west}, {width_we}, {width_ns}, {height})'.format(south=self.south, west=self.west,
                                                                                    width_we=self.width_we,
                                                                                    width_ns=self.width_ns,
                                                                                    height=self.height)

    def __str__(self):
        return self.__repr__

    def corners(self):
        return {'north-west': [self.south + self.width_ns, self.west],
                'north-east': [self.south + self.width_ns, self.west + self.width_we],
                'south-west': [self.south, self.west],
                'south-east': [self.south, self.west + self.width_we]}

    def area(self):
        return self.width_we * self.width_ns

    def volume(self):
        return self.width_we * self.width_ns * self.height

if __name__ == '__main__':
    print(Building(10, 10, 1, 2, 2))
