import unittest
import json
import sys
sys.path[0] = sys.path[0][:-8]
from scripts.server.flask_server import app

test_user_id = None


class TestApi(unittest.TestCase):
    def test_1_adding_new_user(self):
        sent = {"name": "test_user"}
        with app.test_client() as client:
            url = r'/api/v1/adding_new_user'
            result = client.post(url, json=sent)
        test_user_id = result.json['user_id']
        self.assertEqual(result.json['info'], 'User added successfully.')
    
    def test_2_get_user(self):
        with app.test_client() as client:
            url = r'/api/v1/get_user/1'
            result = client.get(url)
            self.assertEqual(result.json, {'name': 'test_user'})
    
    def test_3_get_user(self):
        with app.test_client() as client:
            url = r'/api/v1/get_user/2'
            result = client.get(url)
            self.assertEqual(result.json, {'Status': False, 'info': 'User with this id does not exist.'})
    
    def test_4_get_user(self):
        with app.test_client() as client:
            url = r'/api/v1/get_user/1'
            result = client.post(url)
            self.assertEqual(result.status_code, 405)
    
    def test_5_get_user(self):
        with app.test_client() as client:
            url = r'/api/v1/get_userxxxxxxx/1'
            result = client.post(url)
            self.assertEqual(result.status_code, 404)
    
    def test_6_get_user(self):
        with app.test_client() as client:
            url = r'/api/v1/user/list'
            result = client.get(url)
            self.assertEqual(result.json['status'], True)
    
    def test_7_update_user(self):
        sent = {"name": "test_user_mod"}
        with app.test_client() as client:
            url = r'/api/v1/update/1'
            result = client.put(url, json=sent)
            self.assertEqual(result.json, {'Status': True, 'info': 'User updated successfully.'})
    
    def test_8_delete_user(self):
        with app.test_client() as client:
            url = r'/api/v1/delete/1'
            result = client.delete(url)
            self.assertEqual(result.json, {'Status': True, 'info': 'User deleted successfully.'})

    def test_9_delete_user(self):
        with app.test_client() as client:
            url = r'/api/v1/delete/2'
            result = client.delete(url)
            self.assertEqual(result.json, {'Status': False, 'info': 'User with this id does not exist.'})


def test_server_run():
    unittest.main()

test_server_run()
