curl -X POST --url http://localhost:8081 -d '{"method":"Store.Reserved","params":[{"store_id": 3, "items_for_reserved": [{"id": 1, "quantity": 3}, {"id": 2, "quantity": 2}, {"id": 3, "quantity": 5}]}],"id":"myID"}'
#curl -X POST --url http://localhost:8081 -d '{"method":"Store.Unreserved","params":[{"store_id": 1, "items_for_unreserved": [{"id": 1, "quantity": 3}, {"id": 2, "quantity": 2}, {"id": 3, "quantity": 5}]}],"id":"myID"}'
curl -X POST --url http://localhost:8081 -d '{"method":"Store.Remainder","params":[{"store_id": 3}],"id":"myID"}'
