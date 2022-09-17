import sys

# программу не менять, только добавить два метода
lst_in = list(map(str.strip, sys.stdin.readlines()))  # считывание списка строк из входного потока


class DataBase:
    lst_data = []
    FIELDS = ('id', 'name', 'old', 'salary')

    def select(self, a, b):
        return self.lst_data[a:b+1]

    def insert(self, data):
        lst = []
        for elem in data:
            lst.append(dict(zip(self.FIELDS, elem.split())))
        self.lst_data.extend(lst)



db = DataBase()
db.insert(lst_in)
