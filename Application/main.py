import os
import json
import tempfile
from scripts.server.flask_server import app
from scripts.logic.utilities import storage_path

# storage_path = os.path.join(tempfile.gettempdir(), 'storage.data')


def main():
    app.run(*conf_server())


def conf_server():
    """Returns tuple(host, server) from the file: config.json"""
    path = os.getcwd() + "/Application/config.json"
    with open(path) as config:
        json_str = config.read()
        json_str = json.loads(json_str)

    host = json_str['server']['host']
    port = json_str['server']['port']
    return host, port


def is_file():
    """Return True, if file available else: return False."""
    return os.path.exists(storage_path)


def cread_file_data():
    if not is_file():
        with open(storage_path, 'w') as writer:
            writer.write('{}')
    return is_file()


if __name__ == "__main__":
    cread_file_data()
    main()
