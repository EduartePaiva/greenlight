# Greenlight project, from let's go further

# using migrations

new migration: migrate create -seq -ext .sql  -dir ./migrations some_table

up command: migrate -path=./migrations -database=$GREENLIGHT_DB_DNS up