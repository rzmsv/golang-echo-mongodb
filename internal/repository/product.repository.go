package repository

import (
	"context"
	"fmt"
	"github.com/rzmsv/golang-project/config"
	mongo "github.com/rzmsv/golang-project/internal/db"
	mongo_model "github.com/rzmsv/golang-project/internal/model/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	mongoDriver "go.mongodb.org/mongo-driver/mongo"
	"time"
)

type ProductRepository struct {
	configs *config.AppConfigs
}

func NewProductRepository(configs *config.AppConfigs) *ProductRepository {
	return &ProductRepository{
		configs: configs,
	}
}

/* ----------------------------------- GET ---------------------------------- */
func (R *ProductRepository) ProductList() ([]mongo_model.Products, error) {
	var productList []mongo_model.Products
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collection := mongo.GetCollection("myApp", "products")
	list, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	err = list.All(ctx, &productList)
	if err != nil {
		return nil, err
	}
	return productList, nil
}
func (R *ProductRepository) ProductById(id string) (mongo_model.Products, error) {
	var product mongo_model.Products
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := mongo.GetCollection("myApp", "products")
	pId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return mongo_model.Products{}, err
	}
	result := collection.FindOne(ctx, bson.M{"_id": pId})
	if result.Err() == mongoDriver.ErrNoDocuments {
		return mongo_model.Products{}, nil
	}
	err = result.Decode(&product)
	if err != nil {
		return mongo_model.Products{}, fmt.Errorf("error decoding result: %w", err)
	}
	return product, nil
}

/* ---------------------------------- POST ---------------------------------- */
func (R *ProductRepository) CreateNewProduct(data *mongo_model.Products) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := mongo.GetCollection("myApp", "products")
	result, err := collection.InsertOne(ctx, data)
	if err != nil {
		return "", err
	}
	objectID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", fmt.Errorf("failed to convert InsertedID to ObjectID")
	}
	return objectID.Hex(), nil
}

/* ---------------------------------- PATCH --------------------------------- */
func (R *ProductRepository) UpdateProduct(id string, data mongo_model.Products) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := mongo.GetCollection("myApp", "products")
	pID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return "", err
	}

	result, err := collection.UpdateOne(ctx, bson.M{"_id": pID}, bson.M{"$set": data})
	if err != nil {
		return "", err
	} else if result.MatchedCount > 0 && result.ModifiedCount == 0 {
		return "", fmt.Errorf("Can not updated but id exist!")
	} else if result.MatchedCount == 0 {
		return "not_found!", nil
	}
	return "updated", nil
}

/* --------------------------------- DELETE --------------------------------- */
func (R *ProductRepository) DeleteProduct(id string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := mongo.GetCollection("myApp", "products")
	pID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return "", err
	}
	result, err := collection.DeleteOne(ctx, bson.M{"_id": pID})
	if err != nil {
		return "", err
	}
	if result.DeletedCount == 0 {
		return "not_found!", nil
	}
	return "deleted!", nil
}
