import sys
import json
import random
import string


class WriteError(Exception):
    """Database write error."""


class User:
    def __init__(self, u_id):
        self.u_id = u_id
        self.name = WrapperDB().get_name(self.u_id)
        self.password = WrapperDB().get_password(self.u_id)

    @classmethod
    def create(cls, name, password):
        try:
            print('Создание пользователя')
            status = WrapperDB().create_user(name, password)
            if status:
                return {"Status": True, "info": "User added successfully."}
            return {"Status": False, "info": "A user with the same name already exists."}
        except WriteError:
            return {"Status": False, "info": "user not added"}
        


class WrapperDB:
    def __init__(self):
        self.data = self.get_json()
    
    def get_json(self):
        with open(sys.path[0] + "/DataBase.json") as read_file:
            data = json.load(read_file)
            return data
    
    def commit(self):
        """Saving data to the database."""
        try:
            with open(sys.path[0] + "/DataBase.json", "w") as write_file:
                json.dump(self.data, write_file)
            return True
        except:
            raise WriteError

    def is_user_name(self, user_name):
        try:
            for info in list(self.data['users'].values()):
                if info['name'] == user_name:
                    return True
            return False
        except KeyError:
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
    
    def create_user(self, name, password):
        if not WrapperDB().is_user_name(name):
            user_id = self.get_new_id()
            self.data['users'][user_id] = {"name": name, "password": password}
            self.commit()
            return True
        return False

    def get_all_users(self):
        try:
            data = {"users":[], "status": False}
            user_list = list()
            for user_id in self.data['users']:
                self.data['users'][user_id]['id'] = user_id
                data['users'].append(self.data['users'][user_id])
            data['status'] = True
            return data
        except KeyError:
            return data
        
