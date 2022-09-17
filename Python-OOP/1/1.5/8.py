import sys


class ListObject:
    def __init__(self, data):
        self.next_obj = None
        self.data = data

    def link(self, obj):
        self.next_obj = obj


lst_in = list(map(str.strip, sys.stdin.readlines()))
lst = [ListObject(lst_in[0])]
for i in range(1, len(lst_in)):
    lst.append(ListObject(lst_in[i]))
    lst[i-1].link(lst[i])

head_obj = lst[0]



