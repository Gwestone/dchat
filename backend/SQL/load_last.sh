cd dump
cat last.sql | docker exec -i postgres psql -U postgres