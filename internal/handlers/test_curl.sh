curl -X POST --url http://localhost:8081 -d '{"method":"Store.Reserved","params":[{"items_ids": [4, 5, 6]}],"id":"myID"}'
curl -X POST --url http://localhost:8081 -d '{"method":"Store.Unreserved","params":[{"items_ids": [1, 2, 3]}],"id":"myID"}'
curl -X POST --url http://localhost:8081 -d '{"method":"Store.Remainder","params":[{"store_id": 1}],"id":"myID"}'
