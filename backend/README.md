# Backend

Generating database stuff

```
# Populating database 
docker exec -i postgres psql -U postgres -d pasta < database.sql

# Generating code related to database
$HOME/go/bin/jet  -dsn=postgresql://postgres:password@localhost:5432/postgres?sslmode=disable -path=./.gen
```
