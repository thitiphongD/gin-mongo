package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/thitiphongD/gin-mongo/configs"
	"github.com/thitiphongD/gin-mongo/models"
	"github.com/thitiphongD/gin-mongo/responses"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

var mangaCollection *mongo.Collection = configs.GetCollection(configs.DB, "mangas")
var validate = validator.New()

func CreateManga() gin.HandlerFunc {
	return func(c *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

		var manga models.Manga
		defer cancel()

		if err := c.BindJSON(&manga); err != nil {
			c.JSON(http.StatusBadRequest, responses.MangaResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()},
			})
			return
		}

		if validationErr := validate.Struct(&manga); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.MangaResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data:    map[string]interface{}{"data": validationErr.Error()},
			})
			return
		}

		newManga := models.Manga{
			Id:        primitive.NewObjectID(),
			Name:      manga.Name,
			Character: manga.Character,
			Type:      manga.Type,
		}

		result, err := mangaCollection.InsertOne(ctx, newManga)
		if err != nil {
			c.JSON(http.StatusCreated, responses.MangaResponse{
				Status:  http.StatusCreated,
				Message: "success",
				Data:    map[string]interface{}{"data": result},
			})
		}
	}
}

func GetAllManga() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var mangas []models.Manga
		defer cancel()

		results, err := mangaCollection.Find(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.MangaResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()},
			})
			return
		}

		defer results.Close(ctx)
		for results.Next(ctx) {
			var singleManga models.Manga
			if err = results.Decode(&singleManga); err != nil {
				c.JSON(http.StatusInternalServerError, responses.MangaResponse{
					Status:  http.StatusInternalServerError,
					Message: "error",
					Data:    map[string]interface{}{"data": err.Error()},
				})
			}
			mangas = append(mangas, singleManga)
		}
		c.JSON(http.StatusOK, responses.MangaResponse{
			Status:  http.StatusOK,
			Message: "success",
			Data:    map[string]interface{}{"data": mangas},
		})
	}
}

func GetManga() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		mangaID := c.Param("mangaID")
		var manga models.Manga
		defer cancel()

		objID, _ := primitive.ObjectIDFromHex(mangaID)

		err := mangaCollection.FindOne(ctx, bson.M{"id": objID}).Decode(&manga)

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.MangaResponse{
				Status:  http.StatusOK,
				Message: "success",
				Data:    map[string]interface{}{"data": manga},
			})
		}

	}
}

func EditManga() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		mangaID := c.Param("mangaID")
		var manga models.Manga
		defer cancel()
		objID, _ := primitive.ObjectIDFromHex(mangaID)

		if err := c.BindJSON(&manga); err != nil {
			c.JSON(http.StatusBadRequest, responses.MangaResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()},
			})
			return
		}

		if validateErr := validate.Struct(&manga); validateErr != nil {
			c.JSON(http.StatusBadRequest, responses.MangaResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data:    map[string]interface{}{"data": validateErr.Error()},
			})
			return
		}

		update := bson.M{"name": manga.Name, "character": manga.Character, "type": manga.Type}
		result, err := mangaCollection.UpdateOne(ctx, bson.M{"id": objID}, bson.M{"$set": update})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.MangaResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()},
			})
			return
		}

		var updateManga models.Manga
		if result.MatchedCount == 1 {
			err := mangaCollection.FindOne(ctx, bson.M{"id": objID}).Decode(&updateManga)
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.MangaResponse{
					Status:  http.StatusInternalServerError,
					Message: "error",
					Data:    map[string]interface{}{"data": err.Error()},
				})
				return
			}
		}

		c.JSON(http.StatusOK, responses.MangaResponse{
			Status:  http.StatusOK,
			Message: "success",
			Data:    map[string]interface{}{"data": updateManga},
		})
	}
}

func DeleteManga() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		mangaID := c.Param("mangaID")
		defer cancel()

		objID, _ := primitive.ObjectIDFromHex(mangaID)

		result, err := mangaCollection.DeleteOne(ctx, bson.M{"id": objID})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.MangaResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()},
			})
			return
		}

		if result.DeletedCount < 1 {
			c.JSON(http.StatusNotFound, responses.MangaResponse{
				Status:  http.StatusNotFound,
				Message: "error",
				Data:    map[string]interface{}{"data": "Manga with specified ID not found!"},
			})
			return
		}

		c.JSON(http.StatusOK, responses.MangaResponse{
			Status:  http.StatusOK,
			Message: "success",
			Data:    map[string]interface{}{"data": "Manga successfully deleted!"},
		})

	}
}
