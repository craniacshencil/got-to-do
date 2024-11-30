source .env.local
# Running migrations and signing up
make up
echo ""
make signup
echo ""

# Creating a .env file for running the scripts
touch scripts/.env

# Logging in and adding user_id to .env
make login
user_id=$(psql -U ${DB_USER} -d ${DB_NAME} -h ${DB_HOST} -p ${DB_PORT} -tc 'SELECT ID from users;')
user_id=$(echo "$user_id" | xargs)
echo "user-id: $user_id"
echo "USER_ID=$user_id" >>scripts/.env
echo ""

# Creating the todo-list for all actions except deletion, storing its list_id in .env
make create
list_id=$(psql -U ${DB_USER} -d ${DB_NAME} -h ${DB_HOST} -p ${DB_PORT} -tc "SELECT list_id from todo_lists WHERE date='2025-07-17';")
list_id=$(echo "$list_id" | xargs)
echo "list-id: $list_id"
echo "LIST_ID=$list_id" >>scripts/.env
echo ""

# Creating the todo-list for all delete-list endpoint, storing its list_id in .env
make create-delete
delete_list_id=$(psql -U ${DB_USER} -d ${DB_NAME} -h ${DB_HOST} -p ${DB_PORT} -tc "SELECT list_id from todo_lists WHERE date='2025-08-17';")
delete_list_id=$(echo "$delete_list_id" | xargs)
echo "delete-list-id: $delete_list_id"
echo "DELETE_LIST_ID=$delete_list_id" >>scripts/.env
echo ""

# For delete-task endpoint, storing a task's task_id in .env
task_1=$(psql -U ${DB_USER} -d ${DB_NAME} -h ${DB_HOST} -p ${DB_PORT} -tc "SELECT task_id from tasks WHERE task_name like 'Bake%';")
task_1=$(echo "$task_1" | xargs)
echo "task_1: $task_1"
echo "TASK_ID_1=$task_1" >>scripts/.env
echo ""

# For delete-task endpoint, storing another task's task_id in .env
task_2=$(psql -U ${DB_USER} -d ${DB_NAME} -h ${DB_HOST} -p ${DB_PORT} -tc "SELECT task_id from tasks WHERE task_name like 'Skate%';")
task_2=$(echo "$task_2" | xargs)
echo "task_2: $task_2"
echo "TASK_ID_1=$task_1" >>scripts/.env

# Showing the user the outcome
echo ""
echo "Created .env file looks like this"
echo ""
echo "====================================================================="
echo ""
cat scripts/.env
echo ""
echo "====================================================================="
echo ""
echo "Script executed successfully! You can now run the makefile commands"
