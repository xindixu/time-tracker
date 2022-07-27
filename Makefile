DB_NAME = "todo-time-tracker.db"

setup-db:
	touch ${DB_NAME}

drop-db:
	rm ${DB_NAME}

reinit-db: drop-db setup-db