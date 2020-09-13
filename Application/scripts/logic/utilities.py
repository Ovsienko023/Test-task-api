import sys
import json
import random
import string


class WriteError(Exception):
    """Database write error."""


class User:
    def __init__(self, user_id):
        self.user_id = user_id
        self.name = WrapperDB().get_name(self.user_id)

    @classmethod
    def create(cls, name):
        try:
            print('Создание пользователя')
            status, user_id = WrapperDB().create_user(name)
            if status:
                return {"Status": True,
                        "info": "User added successfully.",
                        "user_id": user_id}
        except WriteError:
            return {"Status": False, "info": "user not added"}
    
    def update(self, data):
        try:
            for update_data in data:
                if update_data == 'name':
                    self.name = data[update_data]
            return (True, update)
        except:
            return False
        
    def __str__(self):
        return f"user_id: {self.user_id}, username: {self.name}"

    def __repr__(self):
        return f"user_id: {self.user_id}, username: {self.name}"


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

    def get_new_id(self):
        """Generates a unique id for a new user."""
        numbers = string.digits
        unique_id = ''
        for _ in range(7):
            unique_id += random.choice(numbers)

        if self.data.get(unique_id):
            self.get_new_id()
        return unique_id

    def get_name(self, u_id):
        """Returns username by id."""
        return self.data['users'][u_id]['name']
    
    def create_user(self, name):
        user_id = self.get_new_id()
        self.data['users'][user_id] = {"name": name}
        self.commit()
        return (True, user_id)

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

    def update_user(self, obj_user):
        """Updates the user in the database."""
        print(obj_user)
        user_id = obj_user.user_id
        username = obj_user.name
        self.data['users'][user_id]['name'] = username
        try:
            if self.commit():
                return {"Status": True, "info": "User updated successfully."}
        except WriteError:
            return {"Status": True, "info": "User not updated."}

    def delete_user(self, obj_user):
        """Deleting a user from the database."""
        user_id = obj_user.user_id
        del self.data['users'][user_id]
        try:
            if self.commit():
                return {"Status": True, "info": "User deleted successfully."}
        except WriteError:
            return {"Status": True, "info": "User not deleted."}

    def get_user(self, obj_user):
        """Returns information about the user."""
        user_id = obj_user.user_id
        try:
            user_info =  self.data['users'][user_id]
            return user_info
        except KeyError:
            return {"Status": False, "info": f"User with id:{user_id} cannot be retrieved."}
