curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"name":"Äänestys 1","id":"1", "description": "Äänestys nro. 1 id 1", "votespertoken": 1, "open": 1}' \
  "http://localhost:8281/api?action=voting&a=add"
