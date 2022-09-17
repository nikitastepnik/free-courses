# здесь объявите класс TriangleChecker
import numbers


class TriangleChecker:
    def __init__(self, a, b, c):
        self.a = a
        self.b = b
        self.c = c

    def is_triangle(self):
        for elem in [self.a, self.b, self.c]:
            if not type(elem) in [int, float] or elem <= 0:
                return 1
        if self.a > self.b + self.c or self.b > self.a + self.c or self.c > self.a + self.b:
            return 2
        return 3


a, b, c = map(int, input().split())  # эту строчку не менять
# здесь создайте экземпляр tr класса TriangleChecker и вызовите метод is_triangle() с выводом информации на экран
tr = TriangleChecker(a, b, c)
print(tr.is_triangle())
