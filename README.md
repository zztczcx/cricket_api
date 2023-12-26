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

## API Authentication

using 'Authorization: BEARER T' request header, Token can generate by using the following tool.

Secret can be configured through env varaible `JWT_SECRET`, by default its value is `secret`

### Util
See https://github.com/goware/jwtutil for utility to help you generate JWT tokens.

go install github.com/goware/jwtutil

Usage: jwtutil -secret=secret -encode -claims='{"user_id":111}'

Output: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxMTF9._cLJn0xFS0Mdr_4L_8XF8-8tv7bHyOQJXyWaNsSqlEs

### Example

```
curl http://localhost:8080/api/v1/players/most_runs\?careerEndYear\=2010 -H 'Authorization: BEARER eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxMTF9._cLJn0xFS0Mdr_4L_8XF8-8tv7bHyOQJXyWaNsSqlEs'

{"name":"Mohammad Yousuf (Asia/PAK)","runs":9720}
```
