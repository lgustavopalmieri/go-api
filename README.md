# go-api

go mod init github.com/lgustavopalmieri/go-api


### gorilla mux

go get -u github.com/gorilla/mux


para pegar o ip do container docker do postgres
### docker exec -it (id do container) sh
### hostname -i
ou
### docker inspect (id do container) | grep IPAdress


gorm https://gorm.io/docs/
### go get -u gorm.io/gorm

### go get gorm.io/driver/postgres


gorilla mux handlers
### go get github.com/gorilla/handlers