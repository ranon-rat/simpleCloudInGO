FROM golang:alpine
COPY . ./cloud
WORKDIR /go/

WORKDIR /go/cloud/src/public/script
RUN apk add npm
RUN npm install typescript -g
RUN tsc *.ts; rm -rf *.ts
RUN npm uninstall -g typescript
RUN apk del npm
WORKDIR /go/cloud/src/
RUN go build main.go

CMD ["./main"]