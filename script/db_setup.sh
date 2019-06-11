#!/bin/bash
echo Wait for servers to be up
sleep 10

HOSTPARAMS="--host roach-one --insecure"
SQL="/cockroach/cockroach.sh sql $HOSTPARAMS"

# https://www.cockroachlabs.com/docs/stable/create-database.html
$SQL -e "CREATE DATABASE accounts;"
$SQL -e "CREATE DATABASE dashboard;"
$SQL -e "CREATE DATABASE wallet;"

# https://www.cockroachlabs.com/docs/stable/create-user.html
$SQL -e "CREATE USER trustkeeper;"

# https://www.cockroachlabs.com/docs/stable/training/users-and-privileges.html
# https://www.cockroachlabs.com/docs/stable/grant.html
$SQL -e "GRANT ALL ON DATABASE accounts, dashboard, wallet to trustkeeper;"
# $SQL -e "GRANT ALL ON DATABASE dashboard to trustkeeper;"
# $SQL -e "GRANT ALL ON DATABASE wallet to trustkeeper;"

# $SQL -d tarun -e "CREATE TABLE articles(name VARCHAR);"
