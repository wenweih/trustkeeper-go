#!/bin/bash
echo Wait for servers to be up
sleep 10

HOSTPARAMS="--host roach-one --insecure"
SQL="/cockroach/cockroach.sh sql $HOSTPARAMS"

# https://www.cockroachlabs.com/docs/stable/create-database.html
$SQL -e "CREATE DATABASE account;"
$SQL -e "CREATE DATABASE dashboard;"
$SQL -e "CREATE DATABASE wallet;"

# https://www.cockroachlabs.com/docs/stable/create-user.html
$SQL -e "CREATE USER trustkeeper;"

# https://www.cockroachlabs.com/docs/stable/training/users-and-privileges.html
# https://www.cockroachlabs.com/docs/stable/grant.html
$SQL -e "GRANT ALL ON DATABASE account, dashboard, wallet to trustkeeper;"
$SQL -d wallet -e 'CREATE TABLE "chains" ("id" serial,"created_at" timestamp with time zone,"updated_at" timestamp with time zone,"deleted_at" timestamp with time zone,"name" text,"coin" text,"bip44id" integer,"status" boolean ,"decimal" bigint, PRIMARY KEY ("id"));'
$SQL -d wallet -e "INSERT  INTO "chains" ("name","coin","bip44id","status", "decimal") VALUES ('Bitcoincore','BTC',0,true,100000000)"
$SQL -d wallet -e "INSERT  INTO "chains" ("name","coin","bip44id","status", "decimal") VALUES ('Ethereum','ETH',60,true,1000000000000000000)"
