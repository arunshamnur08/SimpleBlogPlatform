package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	constant "simple-blogging-platform/constants"
	"simple-blogging-platform/mongo_client"
	"simple-blogging-platform/utils"
	"time"
)

type Blog struct {
	Ctx          context.Context
	log          log.Logger
	Mongo_client *mongo.Client
}

var blogs []utils.Blog

func (s *Blog) Createblog(c *gin.Context) {
	var b utils.BlogDto
	if err := c.ShouldBind(&b); err == nil {
		blog := utils.Blog{
			PostID:    fmt.Sprintf("%s-%d-%s", "Blog", time.Now().Year(), b.Title),
			Title:     b.Title,
			Content:   b.Content,
			CreatedAt: time.Now(),
		}

		err = mongo_client.Insert(s.Ctx, s.Mongo_client, blog)
		if err != nil {
			fmt.Printf("unable to insert to mongodb %s", err.Error())
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"messgae": fmt.Sprintf("unable to create Blog err is %s", err.Error())})
			return
		}
		c.IndentedJSON(http.StatusCreated, gin.H{"message": "blog has been created"})

	} else {
		log.Println("error in binding request body , err is %s", err)
	}
}

func (s *Blog) GetAllBlogs(c *gin.Context) {
	filter := bson.D{}
	cursor := mongo_client.QueryAll(s.Mongo_client, "sample_training", "posts", filter)
	var results []utils.Blog
	if err := cursor.All(context.TODO(), &results); err != nil {
		log.Println("got error while reading from cursor ", err)
	}
	for _, result := range results {
		res, _ := json.Marshal(result)
		log.Println("Matched Results ", string(res))
	}
	c.IndentedJSON(http.StatusOK, results)
}
func (s *Blog) GetBlogById(c *gin.Context) {
	blogId := c.Param(constant.BLOGID)
	filter := bson.D{{Key: "postid", Value: blogId}}
	cursor := mongo_client.QueryAll(s.Mongo_client, "sample_training", "posts", filter)
	var results []utils.Blog
	if err := cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	for _, result := range results {
		res, _ := json.Marshal(result)
		log.Println("Matched Results", string(res))
	}
	c.IndentedJSON(http.StatusOK, results)
}

func (s *Blog) UpdateBlogById(c *gin.Context) {
	blogId := c.Param(constant.BLOGID)
	var b utils.BlogDtoForUpdate
	if err := c.ShouldBind(&b); err == nil {
		filter := bson.D{{"postid", blogId}}
		update := bson.D{{"$set", bson.D{{"title", b.Title}, {"content", b.Content}}}}
		result, err := mongo_client.Update(s.Mongo_client, "sample_training", "posts", filter, update)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, fmt.Sprintf("unable to update Blog with blogid %s err is %s", blogId, err.Error()))
			return
		}
		log.Println("Updated Results Count", result.MatchedCount)
	} else {
		log.Println("error in binding request body , err is %s", err)
		return
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Blog Is Updated Successfully"})
}

func (s *Blog) DeleteBlogBYID(c *gin.Context) {
	blogId := c.Param(constant.BLOGID)
	filter := bson.D{{"postid", blogId}}
	result, err := mongo_client.Delete(s.Mongo_client, "sample_training", "posts", filter)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, fmt.Sprintf("unable to delete Blog with blogid %s err is %s", blogId, err.Error()))
		return
	}
	if result.DeletedCount == 0 {
		log.Println("Number of documents deleted:", result.DeletedCount)
		c.IndentedJSON(http.StatusNotFound, fmt.Sprintf("Blog with blogid %s not found", blogId))
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Number of documents deleted: %d\n", result.DeletedCount)})
}
