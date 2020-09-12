build :
	python3 Application/testing/create_db.py

test :
	python3 Application/testing/test_json.py
	python3 Application/testing/test_server.py
	python3 Application/testing/test_run_json.py

run :
	python3 Application/main.py