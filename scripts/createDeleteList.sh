source .env
curl -H 'Content-Type: application/json' \
  -d '{ 
        "1": { 
            "task_name": "Bake a cake",
            "start_time": "2025-08-17T22:00:00+05:30",
            "end_time": "2025-08-17T22:30:00+05:30",
            "completion": false
            },
        "2": {
            "task_name": "Skate at the park",
            "start_time": "2025-08-17T08:30:00+05:30",
            "end_time": "2025-08-17T09:30:00+05:30",
            "completion": false
            },
        "3": {
            "task_name": "Ship application",
            "start_time": "2025-08-17T19:00:00+05:30",
            "end_time": "2025-08-17T20:30:00+05:30",
            "completion": true
            }
    }' \
  -b c.txt \
  -X POST \
  "http://localhost:8080/users/$USER_ID/2025-08-17"
