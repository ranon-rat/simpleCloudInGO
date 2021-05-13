FROM golang:alpine
COPY . ./cloud
WORKDIR /go/

WORKDIR /go/cloud/src/public/script
# update some stuff
RUN apk update && apk upgrade && apk add build-base
# install the dependencies
RUN apk add npm&& apk add sqlite3
# compile the typescript
RUN npm install typescript -g
RUN tsc *.ts; rm -rf *.ts
# delete some stuff
RUN npm uninstall -g typescript
RUN apk del npm
WORKDIR /go/cloud/src/
# exec the query
RUN cat database/init.sql | sqlite3 database/database.db
#compile the main 
RUN go build main.go

CMD ["./main"]