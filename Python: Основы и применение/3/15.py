import json

data_json = input()
data_again = json.loads(data_json)
dict1 = {}
c = 1
for i in range(0, len(data_again)):
    for k in range(0, len(data_again)):
        if data_again[i]["name"] in data_again[k]["parents"]:
            data_again[k]["parents"] = data_again[k]["parents"] + data_again[i]["parents"]
for i in range(0, len(data_again)):
    for k in range(0, len(data_again)):
        if data_again[i]["name"] in data_again[k]["parents"]:
            c += 1
    dict1.update({data_again[i]["name"]: c})
    c = 1

list_keys = list(dict1.keys())
list_keys.sort()
for i in list_keys:
    print(i, ':', dict1[i])
