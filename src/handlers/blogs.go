package handlers

import (
	"BLOGS_GOLANG/src/models"
	"BLOGS_GOLANG/src/utils"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func AddABlog(c echo.Context) error {
	var blogInput models.Blog
	errInput := json.NewDecoder(c.Request().Body).Decode(&blogInput)
	resp := models.Response{}
	if errInput != nil {
		resp = models.Response{Message: errInput.Error(), Status: "500"}
	} else {
		result, err := utils.BlogsCollection.InsertOne(utils.Ctx, blogInput)
		if err != nil {
			resp = models.Response{Message: err.Error(), Status: "500"}
		} else {
			id, ok := result.InsertedID.(primitive.ObjectID)
			if !ok {
				resp = models.Response{Message: "inserted document doesn't have a primitive id.", Status: "500"}
			} else {
				blogInput.ID = id
				resp = models.Response{Message: "Record created successfully", Data: blogInput, Status: "200"}
			}
		}
	}
	return c.JSON(http.StatusOK, resp)
}

func GetABlog(c echo.Context) error {
	id := c.Param("id")
	objID, errId := primitive.ObjectIDFromHex(id)
	resp := models.Response{}
	if errId != nil {
		resp = models.Response{Message: errId.Error(), Status: "400"}
	}
	fmt.Println(objID)

	filter := bson.D{{Key: "_id", Value: objID}}
	var blogData models.Blog
	err := utils.BlogsCollection.FindOne(utils.Ctx, filter).Decode(&blogData)
	if err != nil {
		resp = models.Response{Message: err.Error(), Status: "204"}
	} else {
		resp = models.Response{Message: "Data retried successfully", Status: "200", Data: blogData}
	}
	return c.JSON(http.StatusOK, resp)
}
func GetAllBlog(c echo.Context) error {
	resp := models.Response{}
	cursor, err := utils.BlogsCollection.Find(utils.Ctx, bson.M{})
	if err != nil {
		resp = models.Response{Message: err.Error(), Status: "500"}
	}
	blogsArr := make([]interface{}, 0)
	defer cursor.Close(utils.Ctx)
	for cursor.Next(utils.Ctx) {
		var blogData models.Blog
		if err = cursor.Decode(&blogData); err != nil {
			fmt.Println(err)
		}
		fmt.Println(blogData,"@@@@@@@@@@@@@@@@@@@@@@@@@@@@")
		blogsArr = append(blogsArr, blogData)
	}

	resp = models.Response{Message: "Data retried successfully", Status: "200", Data: blogsArr}
	return c.JSON(http.StatusOK, resp)
}

func EditABlog(c echo.Context) error {
	id := c.Param("id")
	var blog models.Blog
	errInput := json.NewDecoder(c.Request().Body).Decode(&blog)
	jsonBodyInput := make(map[string]interface{})
	json.NewDecoder(c.Request().Body).Decode(&jsonBodyInput)
	resp := models.Response{}
	if errInput != nil {
		resp = models.Response{Message: errInput.Error(), Status: "400"}
	}
	jsonBody := make(map[string]interface{})

	_, detailsExist := jsonBodyInput["details"]
	if detailsExist {
		jsonBody["details"] = blog.Details
	}
	_, titleExist := jsonBodyInput["title"]
	if titleExist {
		jsonBody["title"] = blog.Title
	}
	_, subtitleExist := jsonBodyInput["subtitle"]
	if subtitleExist {
		jsonBody["subtitle"] = blog.Subtitle
	}
	_, minAgeExist := jsonBodyInput["minAge"]
	if minAgeExist {
		jsonBody["minAge"] = blog.MinAge
	}
	_, maxAgeExist := jsonBodyInput["maxAge"]
	if maxAgeExist {
		jsonBody["maxAge"] = blog.MaxAge
	}
	_, interestsExist := jsonBodyInput["interests"]
	if interestsExist {
		jsonBody["interests"] = blog.Interests
	}
	documentId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{"_id", documentId}}

	result, err := utils.BlogsCollection.UpdateOne(utils.Ctx, filter, jsonBody)
	if err != nil {
		resp = models.Response{Message: err.Error(), Status: "500"}
	} else {
		fmt.Println(result, "@@@@@@@@@@@@@@@@@@@@@@@@", jsonBody)
		resp = models.Response{Message: "Data retried successfully", Status: "200", Data: result}
	}
	return c.JSON(http.StatusOK, resp)
}
