curl -H 'Content-Type: application/json' \
    -d '{ 
        "1": { 
            "task_name": "Complete login feature",
            "start_time": "2024-07-17T22:00:00+05:30",
            "end_time": "2024-07-17T22:30:00+05:30",
            "completion": false
            },
        "2": {
            "task_name": "Do DSA",
            "start_time": "2024-07-17T08:30:00+05:30",
            "end_time": "2024-07-17T09:30:00+05:30",
            "completion": false
            },
        "3": {
            "task_name": "Eat Chicken",
            "start_time": "2024-07-17T19:00:00+05:30",
            "end_time": "2024-07-17T20:30:00+05:30",
            "completion": true
            }
    }' \
    -b c.txt \
    -X POST \
    http://localhost:8080/users/1234/234

