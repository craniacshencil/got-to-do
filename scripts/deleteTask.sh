source .env
curl \
	-b c.txt \
	-X DELETE "http://localhost:8080/users/$USER_ID/2025-08-17/task/$TASK_ID_2"
