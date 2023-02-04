package models

import (
	"context"
	"portfolio-server/initializers"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type BlogArticle struct {
	ID primitive.ObjectID `bson:"_id" json:"id"`

	Title       string `bson:"title" json:"title"`
	Description string `bson:"description" json:"description"`
	TextBody    string `bson:"textBody" json:"textBody,omitempty"`

	Views int64    `bson:"views" json:"views"`
	Tags  []string `bson:"tags" json:"tags"`

	CreatedAt time.Time `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt" json:"updatedAt"`
}

func GetShortPosts(offset int64, limit int64) ([]BlogArticle, error) {
	// database setup
	coll := initializers.DB.Database(initializers.DBName).Collection("posts")

	// create the stages
	unsetStage := bson.D{{"$unset", bson.A{"textBody"}}}
	sortStage := bson.D{{"$sort", bson.D{{"_id", -1}}}}
	skipStage := bson.D{{"$skip", offset}}
	limitStage := bson.D{{"$limit", limit}}

	// pass the pipeline to the Aggregate() method
	cursor, err := coll.Aggregate(context.TODO(), mongo.Pipeline{unsetStage, sortStage, skipStage, limitStage})
	if err != nil {
		return nil, err
	}

	// decode the results
	var results []BlogArticle
	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}

	return results, nil
}

func GetPostByID(id string) (BlogArticle, error) {
	// database setup
	coll := initializers.DB.Database(initializers.DBName).Collection("posts")

	// transform id
	ObjID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return BlogArticle{}, err
	}

	response := coll.FindOne(context.TODO(), bson.D{{"_id", ObjID}})

	var result BlogArticle
	err = response.Decode(&result)
	if err != nil {
		return BlogArticle{}, err
	}

	filter := bson.D{{"_id", ObjID}}
	update := bson.D{{"$inc", bson.D{{"views", 1}}}}
	opts := options.Update().SetUpsert(false)
	coll.UpdateOne(context.TODO(), filter, update, opts)

	return result, nil
}

func CreatePost(title string, description string, text string, tags []string) BlogArticle {

	timenow := time.Now()

	return BlogArticle{
		ID:          primitive.NewObjectIDFromTimestamp(timenow),
		Title:       title,
		Description: description,
		TextBody:    text,
		Tags:        tags,
		Views:       0,
		CreatedAt:   timenow,
		UpdatedAt:   timenow,
	}
}

func (p BlogArticle) Upload() error {
	// database setup
	coll := initializers.DB.Database(initializers.DBName).Collection("posts")

	_, err := coll.InsertOne(context.TODO(), &p)
	if err != nil {
		return err
	}

	return nil
}
