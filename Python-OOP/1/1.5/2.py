import random


def rand_params():
    return random.random(), random.random(), random.random(), random.random()


class Line:
    def __init__(self, a, b, c, d):
        self.sp = (a, b)
        self.ep = (c, d)


class Rect:
    def __init__(self, a, b, c, d):
        self.sp = (a, b)
        self.ep = (c, d)


class Ellipse:
    def __init__(self, a, b, c, d):
        self.sp = (a, b)
        self.ep = (c, d)


elements = [random.choice([Line, Rect, Ellipse])(*rand_params()) for i in range(217)]

for i in range(len(elements)):
    if isinstance(elements[i], Line):
        elements[i].__init__(0, 0, 0, 0)





