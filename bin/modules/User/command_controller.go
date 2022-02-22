package User

import (
	"context"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/hasimy-as/go-fiber-api/bin/db"
	"github.com/hasimy-as/go-fiber-api/bin/modules/User/models"
	"github.com/hasimy-as/go-fiber-api/utils/res"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var validate = validator.New()
var userCollection *mongo.Collection = db.GetCollection(db.DB, "users")

func CreateUser(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var user models.User
	defer cancel()

	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(res.ResponseError(err.Error(), http.StatusBadRequest))
	}

	if validationErr := validate.Struct(&user); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(res.ResponseError(validationErr.Error(), http.StatusBadRequest))
	}

	newUser := models.User{
		Id:        primitive.NewObjectID(),
		UserId:    uuid.New(),
		Name:      user.Name,
		Address:   user.Address,
		CreatedAt: time.Now(),
	}

	result, err := userCollection.InsertOne(ctx, newUser)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(res.ResponseError(err.Error(), http.StatusInternalServerError))
	}

	return c.Status(http.StatusOK).JSON(res.Response(result, "User successfully inserted!"))
}
