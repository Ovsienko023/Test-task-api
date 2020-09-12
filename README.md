# Test-task-api
REST API по работе с сущностью User

## Для начала работы с api необходимо:

1. Установить все зависимости 
`$ pip install -r requirements.txt`
2. В файле config.json записать сервера:
		    {
                "server":{"host": "Хост сервера", 
                       "port": "Порт Сервера"}
            }
3. Выполнить в терминале команды: 
		$ make test
		ok
		$ make run
	После этого будет запущен сервер и api готов к работе.

## Начало работы с api
Примеры работы:

#### Добавление User
            def create_user():
                url = r'http://127.0.0.1:5555/api/v1/adding_new_user'
                data = {"name": "Zoomer"}
                status = requests.post(url, json=data)
                return status.json()
Возвращает id созданного пользователя

#### Получение списка User
            def get_users():
                url = f'http://127.0.0.1:5555/api/v1/user/list'
                status = requests.get(url)
                return status.json()

#### Получение User по Id
            def get_user(user_id):
                url = f'http://127.0.0.1:5555/api/v1/get_user/{user_id}'
                status = requests.get(url)
                return status.json()

#### Редактирование User по Id
            def update_user(user_id):
                url = f'http://127.0.0.1:5555/api/v1/update/{user_id}'
                data = {"name": "Boob"}
                status = requests.put(url, json=data)
                return status.json()

#### Удаление User по Id
            def dalete_user(user_id):
                url = f'http://127.0.0.1:5555/api/v1/delete/{user_id}'
                status = requests.delete(url)
                return status.json()
