from flask import Flask, request
from scripts.logic.utilities import WrapperDB
from scripts.logic.utilities import User
app = Flask(__name__)


@app.route('/api/v1/adding_new_user', methods=['POST'])
def adding_new_user():
    """
    Adding a new user.
    """
    data = request.json
    try:
        name = data['name']
        status = User.create(name)
        return status
    except KeyError:
        return {"Status": False, "info": "Invalid data."}
   


@app.route('/api/v1/get_user/<user_id>', methods=['GET'])
def get_user(user_id):
    """Give user by id."""
    try: 
        user = User(user_id)
        status = WrapperDB().get_user(user)
        return status
    except KeyError:
        return {'Status': False, 'info': 'User with this id does not exist.'}



@app.route('/api/v1/user/list', methods=['GET'])
def user_list():
    """Get all user."""
    user_list = WrapperDB().get_all_users()
    return user_list


@app.route('/api/v1/update/<user_id>', methods=['PUT'])
def update_user(user_id):
    """User update according to the specified parameters."""
    try:
        data = request.json
        user = User(user_id)
        user.update(data)
        status = WrapperDB().update_user(user)
        return status
    except KeyError:
        return {'Status': False, 'info': 'User with this id does not exist.'}


@app.route('/api/v1/delete/<user_id>', methods=['DELETE'])
def delete_user(user_id):
    """Removing a user by his id."""
    try: 
        user = User(user_id)
        status = WrapperDB().delete_user(user)
        return status
    except KeyError:
        return {'Status': False, 'info': 'User with this id does not exist.'}
