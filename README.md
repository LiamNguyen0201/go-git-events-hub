# go-git-events-hub
A center hub that polling events from multiple Git repositories then forwarding them to multiple HTTP endpoints

## Dependencies

```shell
## Environment variables
go get github.com/joho/godotenv

## Google authentication
go get golang.org/x/oauth2
go get golang.org/x/oauth2/google

## JWT
go get github.com/golang-jwt/jwt/v5

## Logging
go get github.com/sirupsen/logrus
go get github.com/bshuster-repo/logrus-logstash-hook
go get github.com/yukitsune/lokirus
go get github.com/grafana/loki-client-go/loki

## MongoDB
go get go.mongodb.org/mongo-driver/mongo
go get go.mongodb.org/mongo-driver/mongo/options

## SQL database
go get gorm.io/gorm
go get gorm.io/driver/sqlite

## Web
go get github.com/gin-gonic/gin

## Validator
go get github.com/go-playground/validator/v10
```

## Hot reload

```shell
## Install air via go install
go install github.com/air-verse/air@latest

## Generate configuration file
air init

## Run the Gin Server with air
air
```

## Reference
- [List a projects visible events](https://docs.gitlab.com/ee/api/events.html#list-a-projects-visible-events)
- [?](https://gilangprambudi.medium.com/streamlining-log-management-in-go-with-grafana-loki-integration-8b124f2e4121)
