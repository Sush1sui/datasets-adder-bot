package mongodb

import (
	"context"
	"fmt"

	"github.com/Sush1sui/datasets_adder/internal/models"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (c *MongoClient) GetAllUserAccounts() ([]models.UserAccount, error) {
	var userAccounts []models.UserAccount
	cursor, err := c.Client.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var userAccount models.UserAccount
		if err := cursor.Decode(&userAccount); err != nil {
			return nil, err
		}
		userAccounts = append(userAccounts, userAccount)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return userAccounts, nil
}

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