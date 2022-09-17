objects = [i for i in range(100)]
uniq = []
for obj in objects:
    if obj in uniq:
        pass
    else:
        uniq.append(obj)
print(len(uniq))
