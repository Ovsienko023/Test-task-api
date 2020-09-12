import json
import sys
sys.path[0] = sys.path[0][:-8]


data = {"users": {"1": {"name": "test_user"}}}
with open(sys.path[0] + "/DataBase.json", "w") as write_file:
    json.dump(data, write_file)
