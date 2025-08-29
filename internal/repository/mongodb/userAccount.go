package mongodb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func (c *MongoClient) DeleteUserAccountByEmail(email string) (int, error) {
	result, err := c.Client.DeleteOne(context.Background(), bson.M{"email": email})
	if err != nil {
		return 0, err
	}

	// log if there was a deleted account
	if result.DeletedCount > 0 {
		fmt.Println("Deleted user account with email:", email)
	}

	return int(result.DeletedCount), nil
}