def find_all_paths(graph, start, end, path=None):
    if path is None:
        path = []
    path = path + [start]

    if start == end:
        return [path]

    if start not in graph.keys():
        return []

    paths = []

    for node in graph[start]:

        if node not in path:

            newpaths = find_all_paths(graph, node, end, path)

            for newpath in newpaths:
                paths.append(newpath)

    return paths


k = int(input())
dict1 = {}
for i in range(0, k):
    a, *b = input().replace(":", " ").split()
    if a in dict1.keys():
        c = dict1.get(a)
        dict1[a] = c + b
    else:
        dict1[a] = b

k = int(input())
for i in range(0, k):
    a, b = input().split()
    if find_all_paths(dict1, b, a):
        print("Yes")
    else:
        print("No")
