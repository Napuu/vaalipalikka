# clear old data
curl "http://localhost:8281/drop"

# candidates
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"name":"Ehdokas 1","id":"1", "description": "Moi olen ehdokas 1"}' \
  "http://localhost:8281/api?action=candidate&a=add"
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"name":"Ehdokas 2","id":"2", "description": "Moi olen ehdokas 2"}' \
  "http://localhost:8281/api?action=candidate&a=add"
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"name":"Ehdokas 3","id":"3", "description": "Moi olen ehdokas 3"}' \
  "http://localhost:8281/api?action=candidate&a=add"
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"name":"Ehdokas 4","id":"4", "description": "Moi olen ehdokas 4"}' \
  "http://localhost:8281/api?action=candidate&a=add"
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"name":"Ehdokas 5","id":"5", "description": "Moi olen ehdokas 5"}' \
  "http://localhost:8281/api?action=candidate&a=add"


# votings
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"name":"Äänestys 1","id":"1", "description": "Äänestys nro. 1 id 1", "votespertoken": 1}' \
  "http://localhost:8281/api?action=voting&a=add"
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"name":"Äänestys 2","id":"2", "description": "Äänestys nro. 2 id 2", "votespertoken": 1}' \
  "http://localhost:8281/api?action=voting&a=add"
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"name":"Äänestys 3","id":"3", "description": "Äänestys nro. 3 id 3", "votespertoken": 1}' \
  "http://localhost:8281/api?action=voting&a=add"

# voting | candidates
# 1      | 1, 2, 3, 4
# 2      | 2, 3, 4
# 3      | 1, 2, 3, 4, 5
# candidates to voting

# voting 1
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"votingid":"1","candidateid":"1"}' \
  "http://localhost:8281/api?action=availability&a=add"
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"votingid":"1","candidateid":"2"}' \
  "http://localhost:8281/api?action=availability&a=add"
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"votingid":"1","candidateid":"3"}' \
  "http://localhost:8281/api?action=availability&a=add"
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"votingid":"1","candidateid":"4"}' \
  "http://localhost:8281/api?action=availability&a=add"

# voting 2
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"votingid":"2","candidateid":"2"}' \
  "http://localhost:8281/api?action=availability&a=add"
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"votingid":"2","candidateid":"3"}' \
  "http://localhost:8281/api?action=availability&a=add"
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"votingid":"2","candidateid":"4"}' \
  "http://localhost:8281/api?action=availability&a=add"

# voting 3
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"votingid":"3","candidateid":"1"}' \
  "http://localhost:8281/api?action=availability&a=add"
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"votingid":"3","candidateid":"2"}' \
  "http://localhost:8281/api?action=availability&a=add"
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"votingid":"3","candidateid":"3"}' \
  "http://localhost:8281/api?action=availability&a=add"
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"votingid":"3","candidateid":"4"}' \
  "http://localhost:8281/api?action=availability&a=add"
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"votingid":"3","candidateid":"5"}' \
  "http://localhost:8281/api?action=availability&a=add"

curl --request POST "http://localhost:8281/api?action=token&a=generatetokens"
clear
curl --request POST "http://localhost:8281/api?action=token&a=showtokens" -s | awk '{split($0, a, "\""); print a[4]}'

