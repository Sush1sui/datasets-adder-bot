package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func (c *MongoClient) DeleteUserAccountByEmail(email string) error {
	_, err := c.Client.DeleteOne(context.Background(), bson.M{"email": email})
	if err != nil {
		return err
	}

	return nil
}