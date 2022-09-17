k = int(input())
h = 0
spaces = {}
inputs = []


class Space():
    def __init__(self, name, parent=None):
        self.name = name
        self.parent = parent
        self.vars = []

    def add(self, var):
        self.vars.append(var)


a = Space('global')
spaces.setdefault('global', a)


def get1(name, var):
    a = spaces.get(name)
    if var in a.vars:
        return name
    else:
        if (a.parent is None) and (var in a.vars):
            return 'global'
        elif a.parent is None:
            return 'None'
        else:
            return get1(a.parent, var)


for i in range(0, k):
    com, name, var = input().split()
    if com == 'add':
        a = spaces.get(name)
        a.add(var)
    if com == 'create':
        a = Space(name, var)
        spaces.setdefault(name, a)
    if com == 'get':
        c = get1(name, var)
        inputs.append(c)
        h = h + 1
for i in range(0, h):
    print(inputs[i])
