source .env
curl -H 'Content-Type: application/json' \
  -d '{ 
        "1": { 
            "task_name": "Complete add list feature",
            "start_time": "2025-07-17T14:00:00+05:30",
            "end_time": "2025-07-17T14:15:00+05:30",
            "completion": false
            },
        "2": {
            "task_name": "Grocery shopping",
            "start_time": "2025-07-17T18:30:00+05:30",
            "end_time": "2025-07-17T19:00:00+05:30",
            "completion": false
            }
    }' \
  -b c.txt \
  -X POST \
  http://localhost:8080/users/$USER_ID/2025-07-17/$LIST_ID
echo "http://localhost:8080/users/$USER_ID/2025-07-17/$LIST_ID"
