docker stop vaalit_db
docker rm vaalit_db
docker run -p 0.0.0.0:4531:5432 -d --name vaalit_db postgres -c 'shared_buffers=256MB' -c 'max_connections=250' # ":D"
docker cp create_db.sql vaalit_db:/create_db.sql
echo "waiting for a while..."
sleep 5
docker exec -it -u postgres vaalit_db /bin/bash -c "psql < /create_db.sql"

