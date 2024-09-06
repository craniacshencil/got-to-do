curl -H 'Content-Type: application/json' \
	-d '{
  "username": "Amanda", 
  "first_name": "Amanda", 
  "last_name": "Hickory", 
  "confirm_password": "Rishi@1234",
  "password": "Rishi@1234"
  }' \
	-X POST \
	http://localhost:8080/signup
