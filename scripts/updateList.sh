source .env
curl -H 'Content-Type: application/json' \
	-d '{ 
        "d4e7be0c-a07e-4c90-8b6e-bd598956026e": { 
            "task_name": "Complete signup feature",
            "start_time": "2025-07-17T22:00:00+05:30",
            "end_time": "2025-07-17T22:30:00+05:30",
            "completion": true 
            },
        "4000c0d1-86ba-42a3-a88c-f2ae4dd0e6b8": {
            "task_name": "Do DSA",
            "start_time": "2025-07-17T08:30:00+05:30",
            "end_time": "2025-07-17T09:30:00+05:30",
            "completion": true 
            },
        "4fb2d2b3-d549-424a-8f8c-1d738ce4d9e4": {
            "task_name": "Eat Chicken",
            "start_time": "2025-07-17T19:00:00+05:30",
            "end_time": "2025-07-17T20:30:00+05:30",
            "completion": false 
            }
    }' \
	-b c.txt \
	-X PUT \
	http://localhost:8080/users/$USER_ID/2025-07-17/$LIST_ID
