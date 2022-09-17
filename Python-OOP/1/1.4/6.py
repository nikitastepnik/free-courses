class Translator:
    dictionary = {}

    def add(self, eng, rus):
        if self.dictionary.get(eng):
            if rus not in self.dictionary[eng]:
                self.dictionary[eng].append(rus)
        else:
            self.dictionary[eng] = [rus]

    def remove(self, eng):
        del self.dictionary[eng]

    def translate(self, eng):
        return self.dictionary[eng]


tr = Translator()
tr.add("tree", "дерево")
tr.add("car", "машина")
tr.add("leaf", "лист")
tr.add("river", "река")
tr.add("go", "идти")
tr.add("go", "ехать")
tr.add("go", "ходить")
tr.add("milk", "молоко")

tr.remove("car")
print(*tr.translate("go"))
