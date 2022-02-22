package User

import (
	"context"
	"net/http"
	"time"

	"github.com/hasimy-as/go-fiber-api/bin/modules/User/models"
	"github.com/hasimy-as/go-fiber-api/utils/res"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func GetUsers(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var users []models.User
	defer cancel()

	results, err := userCollection.Find(ctx, bson.M{})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(res.ResponseError(err.Error(), http.StatusInternalServerError))
	}

	defer results.Close(ctx)
	for results.Next(ctx) {
		var singleUser models.User
		if err = results.Decode(&singleUser); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(res.ResponseError(err.Error(), http.StatusInternalServerError))
		}

		users = append(users, singleUser)
	}

	return c.Status(http.StatusOK).JSON(res.Response(users, "Users successfully fetched."))
}
