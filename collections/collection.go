package collections

import (
	"github.com/shyamjith94/go-gin/configuration"
	"go.mongodb.org/mongo-driver/mongo"
)

var UserCollection *mongo.Collection = configuration.GetCollection(configuration.DbClient, "users")
var ProductCollection *mongo.Collection = configuration.GetCollection(configuration.DbClient, "Product")
