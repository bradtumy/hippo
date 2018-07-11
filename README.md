# hippo - API Gateway

## Config
Update config/config.json.example to match your environment and then save as config.json (do not check in config.json)

## Run
go run main.go

## Docker
### Create Binary
CGO_ENABLED=0 GOOS=linux GARCH=386 go build -a -installsuffix cgo -ldflags '-s' -o server

### Create Docker Image
docker build -t testserver .

### Run Docker Instance
docker run --rm -p 9091:9091 testserver
