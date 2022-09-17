class Thing:
    def __init__(self, name, price):
        self.name = name
        self.price = price


class Cart:
    def __init__(self):
        self.goods = []

    def add(self, gd):
        self.goods.append(gd)

    def remove(self, idx):
        self.goods.pop(idx)

    def get_list(self):
        return ['{name}: {price}'.format(name=self.goods[i].name, price=self.goods[i].price) for i in range(len(self.goods))]


class Table(Thing):
    ...


class TV(Thing):
    ...


class Notebook(Thing):
    ...


class Cup(Thing):
    ...


cart = Cart()
cart.add(TV("TV1", "50"))
cart.add(TV("TV2", "100"))
cart.add(Table("Table", "123"))
cart.add(Notebook("Notebook1", "123"))
cart.add(Notebook("Notebook2", "12345"))
cart.add(Cup("Cup", "214"))

