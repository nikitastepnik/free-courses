from xml.etree import ElementTree

dict1 = {"red": 0, "green": 0, "blue": 0}


def f(root, level=1):
    dict1[root.attrib['color']] += level
    level += 1
    for child in root:
        f(child, level)


root = ElementTree.fromstring(input())
f(root)
print(dict1['red'], dict1['green'], dict1['blue'])
