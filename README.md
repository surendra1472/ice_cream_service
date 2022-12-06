
## Prerequisites
[GoLang](https://golang.org/dl/) <br>

## Development setup
### Project setup
Set the go path in ~/.bash_profile
```
vi ~/.bash_profile
```

```
export GOPATH="<desired path>"
eg: export GOPATH="/Users/me/go"
```
Create the following folder structure from the go path
```
cd <go path>
eg:cd /Users/me/go
```

### Build and Run Locally
Go to the git repo and run dep to fetch vendor libraries
```
dep ensure
```
Build project
```
go build
```
Run project
```
./ic-service
```

### Kafka setup 

    
```
Install kafka_2.11-1.0.0
```
Run Zookeeper
```
kafka_2.11-1.0.0/bin/zookeeper-server-start.sh kafka_2.11-1.0.0/config/zookeeper.properties
```
Run kafka server
```
kafka_2.11-1.0.0/bin/kafka-server-start.sh kafka_2.11-1.0.0/config/server.properties
```

#### Run Postgres 
Please refer `migration.up.sql.go` to view the database structure for icecream

### Developer Notes

## Swagger

The project is configured with Swagger to leverage the API documentation and also helpful to sync up with icecream-indexer-service from ic-service through ic-indexer-worker.

### JWT Authentication

This service is authenticated by JWT token. Please use the below login when making request to the system.

```
 username: zalora
 password: zalora
```

### Producer

Here, All actions are transactional. The complete transaction includes create/update the postgres database and sending the business object to Kafka as producer model.

### POST, PUT, DELETE - Authenticated

This service solely responsible for all the `POST`, `PUT`. `DELETE` requests. 


### TestCases:

Test cases are written and it has 100% of test coverage and code coverage. It includes both *unit test cases*.

## Environment based Config files

The project uses config files based on the following environments

    - dev
    - staging
    - uat
    
## Configuration

The configurations need to run the project 

    - server Port Number
    - postgres configuration
    - kafka configuration
        
