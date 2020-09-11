from flask import Flask
from scripts.logic.utilities import WrapperDB
from scripts.logic.utilities import User
app = Flask(__name__)


@app.route('/api/v1/user/list', methods=['GET'])
def user_list():
    """ GET all user """
    a = WrapperDB()
    # a.data['users']['2'] = {"name": "Kop", "password": "1213"}
    # print(a.data)
    # print(a.data['users']['1'])
    # a.commit()
    # print(a.get_new_id())
    
    return ""