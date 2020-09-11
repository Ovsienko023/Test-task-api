import sys
import json
import random
import string


class User:
    def __init__(self, u_id):
        self.u_id = u_id
        self.name = WrapperDB().get_name(self.u_id)
        self.password = WrapperDB().get_password(self.u_id)

    @classmethod
    def create(cls, name, password):
        print('Создание пользователя')
        pass



class WrapperDB:
    def __init__(self):
        self.data = self.get_json()
    
    def get_json(self):
        with open(sys.path[0] + "/DataBase.json") as read_file:
            data = json.load(read_file)
            return data
    
    def commit(self):
        with open(sys.path[0] + "/DataBase.json", "w") as write_file:
            json.dump(self.data, write_file)

    def is_user_name(self, user_name):
        for info in self.data.values():
            if info['name'] == user_name:
                return True
        return False

    def get_new_id(self):
        """Generates a unique id for a new user."""
        numbers = string.digits
        unique_id = ''
        for _ in range(5):
            unique_id += random.choice(numbers)

        if self.data.get(unique_id):
            self.get_new_id()
        return unique_id

    def get_name(self, u_id):
        """Returns username by id."""
        return self.data['users'][u_id]['name']
    
    def get_password(self, u_id):
        """Returns the user's password by its id."""
        return self.data['users'][u_id]['password']