# vaalipalikka

Voting machine for student guilds [(Rakennusinsinöörikilta)](https://rakennusinsinoorikilta.fi) election meetings.

Secure to use in in scenarios where people can be trusted to not do bad things. Eg. ddos is fairly effective against this as proper logging in procedure is missing.

## setup

Completely dockerized package with reverse-proxy included.
```
docker-compose up -d # to get stack up

# it's easiest to add new tokens directly through psql
docker-compose exec postgres psql -U vaalit -c "insert into mastertoken(value) values('123');" # insert admin token
docker-compose exec postgres psql -U vaalit -c "insert into token(value, valid) values('4321', 0);" # insert regular token

service is up at localhost:3000

```
