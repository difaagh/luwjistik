# Pre-requisite
Go 1.15 or latest

execute sql script in db_backup.sql file
#

# Getting started
```
cd $GOPATH
git clone https://github.com/difaagh/luwjistik.git
go mod vendor
touch .env
```

# env variable
edit .env file and add this with yours
```
MYSQL_HOST=yours_mysql_host
MYSQL_PORT=yours_mysql_port
MYSQL_USERNAME=yours_mysql_username
MYSQL_PASSWORD=yours_mysql_password
MYSQL_DATABASE=yours_mysql_database
REDIS_URI=yours_redis_uri:port
REDIS_PASSWORD=yours_redis_password
REDIS_USERNAME=yours_redis_username
```

# Start
go run main.go

# Test
before test, replace the contents of file .env.test like .env, the key is same :
```
MYSQL_HOST=yours_mysql_host
MYSQL_PORT=yours_mysql_port
MYSQL_USERNAME=yours_mysql_username
MYSQL_PASSWORD=yours_mysql_password
MYSQL_DATABASE=yours_mysql_database
REDIS_URI=yours_redis_uri:port
REDIS_PASSWORD=yours_redis_password
REDIS_USERNAME=yours_redis_username
```
or you can change to the different connection or database for testing purpose

the run :
```
go test ./...
```