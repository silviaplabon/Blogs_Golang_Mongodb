package initiator

import (
	"BLOGS_GOLANG/src/handlers"
	"BLOGS_GOLANG/src/utils"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Initialize() {
	var err error
	err = godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	utils.MongoClient, err = initializeMongodb()
	if err != nil {
		log.Fatalln(err)
	}
	defer DisconnectMongodb()
	e := echo.New()
	e.Use(middleware.CORS())

	api := e.Group("/api")
	api.POST("/blogs/addABlog", handlers.AddABlog)
	api.GET("/blogs/getABlog/:id", handlers.GetABlog)
	api.PUT("/blogs/editABlog/:id", handlers.EditABlog)
	api.GET("/blogs/getAllBlogs", handlers.GetAllBlog)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", os.Getenv("PORT"))))
}

func initializeMongodb() (*mongo.Client, error) {
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("Please add MONGODB_URI at environment variable.")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	} else {
		utils.BlogsCollection = client.Database("Blogs").Collection("blogs_collection")
	}
	defer func() {
		// if err := client.Disconnect(context.TODO()); err != nil {
		// 	panic(err)
		// }
	}()
	return client, nil

}
func DisconnectMongodb() {
	if err := utils.MongoClient.Disconnect(utils.Ctx); err != nil {
		panic(err)
	}
}
