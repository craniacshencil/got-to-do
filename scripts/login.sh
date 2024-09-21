curl -H 'Content-Type: application/json' \
  -d '{"username": "Amanda", "password": "Password@1234"}' \
  -c c.txt \
  -X POST \
  http://localhost:8080/login
