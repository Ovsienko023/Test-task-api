from flask import Flask, request
from scripts.logic.utilities import WrapperDB
from scripts.logic.utilities import User
app = Flask(__name__)


@app.route('/api/v1/adding_new_user', methods=['POST'])
def adding_new_user():
    """
    {"name": "Bob", password: "123"}
    """
    data = request.json
    name = data['name']
    password = data['password']
    status = User.create(name, password)
    return status


@app.route('/api/v1/user/list', methods=['GET'])
def user_list():
    """Get all user."""
    user_list = WrapperDB().get_all_users()
    return user_list

@app.route('/api/v1/update/<user_id>', methods=['PUT'])
def update_user(user_id):
    return "{}"
