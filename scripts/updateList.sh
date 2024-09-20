source .env
curl -H 'Content-Type: application/json' \
  -d '{ 
        "1555ca73-c471-4ca7-adc6-c242b1c70907": { 
            "task_name": "Complete signup feature",
            "start_time": "2025-07-17T22:00:00+05:30",
            "end_time": "2025-07-17T22:30:00+05:30",
            "completion": true 
            },
        "cfcbda5c-f16d-4d97-b0c7-9f9324e95a55": {
            "task_name": "Do DSA",
            "start_time": "2025-07-17T08:30:00+05:30",
            "end_time": "2025-07-17T09:30:00+05:30",
            "completion": true 
            },
        "dacbf5ff-6a8b-4047-bc35-3ed9d1d7f124": {
            "task_name": "Eat Chicken",
            "start_time": "2025-07-17T19:00:00+05:30",
            "end_time": "2025-07-17T20:30:00+05:30",
            "completion": false 
            }
    }' \
  -b c.txt \
  -X PUT \
  http://localhost:8080/users/$USER_ID/2025-07-17/$LIST_ID
