package models

import (
	"context"
	"portfolio-server/initializers"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Project struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	Title       string             `bson:"title" json:"title"`
	Description string             `bson:"description" json:"description"`
	TextBody    string             `bson:"textBody" json:"textBody,omitempty"`
	TechStack   []string           `bson:"techStack" json:"techStack"`
	GitHubLink  string             `bson:"gitHubLink" json:"gitHubLink"`
	CreatedAt   time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt   time.Time          `bson:"updatedAt" json:"updatedAt"`
}

func GetShortProjects(offset int64, limit int64) ([]Project, error) {
	// database setup
	coll := initializers.DB.Database(initializers.DBName).Collection("projects")

	// create the stages
	unsetStage := bson.D{{"$unset", bson.A{"textBody"}}}
	sortStage := bson.D{{"$sort", bson.D{{"_id", 1}}}}
	skipStage := bson.D{{"$skip", offset}}
	limitStage := bson.D{{"$limit", limit}}

	// pass the pipeline to the Aggregate() method
	cursor, err := coll.Aggregate(context.TODO(), mongo.Pipeline{unsetStage, sortStage, skipStage, limitStage})
	if err != nil {
		return nil, err
	}

	// decode the results
	var results []Project
	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}

	return results, nil
}

func GetProjectByID(id string) (Project, error) {
	// database setup
	coll := initializers.DB.Database(initializers.DBName).Collection("projects")

	// transform id
	ObjID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return Project{}, err
	}

	response := coll.FindOne(context.TODO(), bson.D{{"_id", ObjID}})

	var result Project
	err = response.Decode(&result)
	if err != nil {
		return Project{}, err
	}

	return result, nil
}

func CreateProject(title string, description string, text string, githublink string, techstack []string) Project {

	timenow := time.Now()

	return Project{
		ID:          primitive.NewObjectIDFromTimestamp(timenow),
		Title:       title,
		Description: description,
		TextBody:    text,
		TechStack:   techstack,
		GitHubLink:  githublink,
		CreatedAt:   timenow,
		UpdatedAt:   timenow,
	}
}

func (p Project) Upload() error {
	// database setup
	coll := initializers.DB.Database(initializers.DBName).Collection("projects")

	_, err := coll.InsertOne(context.TODO(), &p)
	if err != nil {
		return err
	}

	return nil
}
