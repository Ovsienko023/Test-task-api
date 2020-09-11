build :
	python3 Application/testing/create_db.py

test :
	python3 Application/testing/test_server.py

run :
	python3 Application/main.py