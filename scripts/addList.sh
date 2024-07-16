
curl -H 'Content-Type: application/json' \
    -d '{ 
        "1": { 
            "task_name": "Complete add list feature",
            "start_time": "2024-07-17T14:00:00+05:30",
            "end_time": "2024-07-17T14:15:00+05:30",
            "completion": false
            },
        "2": {
            "task_name": "Grocery shopping",
            "start_time": "2024-07-17T18:30:00+05:30",
            "end_time": "2024-07-17T19:00:00+05:30",
            "completion": false
            }
    }' \
    -b c.txt \
    -X POST \
    http://localhost:8080/users/4984ba83-d59f-445e-9434-1464ca7b0673/2024-07-17/0d62adcf-d066-4fd2-ba9b-291d765dfbb0
