package main

import (
	"os"

	"github.com/jinzhu/gorm"
	"github.com/vhfuly/codepix-go/application/grpc"
	"github.com/vhfuly/codepix-go/infrastructure/db"
)

var database *gorm.DB

func main() {
	database = db.ConnectDB(os.Getenv("env"))
	grpc.StartGrpcServer(database, 50051)
}
