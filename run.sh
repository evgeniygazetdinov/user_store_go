docker-compose -f docker-compose.yml up -d db &&
go run main.go handlers.go models.go