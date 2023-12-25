## Cricket-Api

## Features

* using `chi` for the http router, which is lightweight and Designed for composable APIs
* using `sqlc` to generate type-safe code from SQL
* using `go-migrate` for database migrations
* using `air` to live reload Go apps for dev

## How to run it

### Prerequisite

There are a lot of config that can be set, please check `./config`.

the most necessary one could be 

```bash
export DATABASE_URL="root:password@/cricket_db?parseTime=true"
```


> [!WARNING]
> the user:password is only used for development

### Start the services

```bash
docker compose up
```

it will start a mysql server and do migrations, but at this point, the tables are still empty

you can run the following command to seed to table

```bash
go run cmd/load_csv_to_db/main.go
```


## Api document
