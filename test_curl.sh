#curl -X POST --url http://localhost:8081 -d '{"method":"Store.Reserved","params":[{"store_id": 1, "items_ids": [4, 5, 6], "quantity": 0}],"id":"myID"}'
#curl -X POST --url http://localhost:8081 -d '{"method":"Store.Unreserved","params":[{"store_id": 1, "items_ids": [1, 2, 3], "quantity": 1}],"id":"myID"}'
curl -X POST --url http://localhost:8081 -d '{"method":"Store.Remainder","params":[{"store_id": 1}],"id":"myID"}'
