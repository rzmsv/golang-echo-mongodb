package db

import (
	"context"
	"fmt"

	"github.com/rzmsv/golang-project/config"
	"github.com/rzmsv/golang-project/core"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func InitMongoDB(appConfigs *config.AppConfigs) {
	ctx := context.Background()
	/* ----------------------------- Set MongoDB URI ---------------------------- */
	clientOptions := options.Client().ApplyURI(appConfigs.AppConfig("MONGO_DB_URL"))

	/* ------------------------- Setting auth credential ------------------------ */
	clientOptions.SetAuth(options.Credential{
		Username: appConfigs.AppConfig("MONGO_USERNAME"),
		Password: appConfigs.AppConfig("MONGO_PASSWORD"),
	})
	/* --------------------------- Connect to Mongodb --------------------------- */
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		core.PanicHandler(fmt.Sprintf("❌ Failed to connect to MongoDB: %v", err.Error()))
	}
	client.Database(appConfigs.AppConfig("MONGO_DB_NAME"))
	Client = client
	fmt.Println("✅ MongoDB connected!")
}

func GetCollection(dbName, colName string) *mongo.Collection {
	return Client.Database(dbName).Collection(colName)
}
