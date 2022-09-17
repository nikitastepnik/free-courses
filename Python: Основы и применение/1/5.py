class MoneyBox:
    def __init__(self, capacity):
        self.capacity = capacity
        self.n = 0

    def can_add(self, v):
        if self.n + v <= self.capacity:
            return True
        else:
            return False

    def add(self, v):
        if (a.can_add):
            self.n = self.n + v


a = MoneyBox(100)
a.add(99)
