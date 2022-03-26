#!/bin/bash

echo "List of all options:"
echo "1. Create schema file"
echo "2. Pull docker image for postgres"
echo "3. Create database"
echo "4. Drop database"
echo "5. migrate up"
echo "6. migrate down"
echo "7. sqlc init"
echo "8. sqlc generate"
echo "9. run unit tests"

# read the option
read -p "*  Enter your option: " option

# check if the option is 1
if [ $option -eq 1 ]; then
    # read the schema file name
    read -p "*  Enter the schema file name: " schema_file_name
    # show the schema file name
    echo "*  Your schema file name is: $schema_file_name"
    migrate create -ext sql -dir db/migration -seq $schema_file_name
    echo "*  Create schema file successfully"
fi

if [ $option -eq 2 ]; then
    # read the schema file name
    read -p "name: " image_name
    read -p "password: " password
    docker run --name $image_name -e POSTGRES_PASSWORD=$password -v postgres_data:/var/lib/postgresql/data -p 5432:5432 -d postgres:14.2-alpine3.15
fi

if [ $option -eq 3 ]; then
    # read the schema file name
    read -p "*  Enter the image name: " image_name
    read -p "*  Enter the database name: " database_name
    docker exec -it $image_name psql -U postgres -c "CREATE DATABASE $database_name;"
fi

if [ $option -eq 4 ]; then
    # read the schema file name
    read -p "*  Enter the image name: " image_name
    read -p "*  Enter the database name: " database_name
    docker exec -it $image_name psql -U postgres -c "DROP DATABASE $database_name;"
fi

if [ $option -eq 5 ]; then
    # read the schema file name
    read -p "*  Enter the database name: " database_name
    read -p "*  Enter the password: " password
    migrate -path db/migration --database "postgres://postgres:$password@localhost:5432/$database_name?sslmode=disable" --verbose up
fi

if [ $option -eq 6 ]; then
    # read the schema file name
    read -p "*  Enter the database name: " database_name
    read -p "*  Enter the password: " password
    migrate -path db/migration --database "postgres://postgres:$password@localhost:5432/$database_name?sslmode=disable" --verbose down
fi

if [ $option -eq 7 ]; then
    # read the schema file name
    sqlc init
fi

if [ $option -eq 8 ]; then
    # read the schema file name
    sqlc generate
fi

if [ $option -eq 9 ]; then
    # read the schema file name
    go test -v -cover ./...
fi