package models

import (
	"context"
	"portfolio-server/initializers"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ContactSubmission struct {
	ID primitive.ObjectID `bson:"_id" json:"id"`

	FullName string `bson:"fullName" json:"fullName"`
	Email    string `bson:"email" json:"email"`
	Subject  string `bson:"subject" json:"subject"`
	Message  string `bson:"message" json:"message"`

	CreatedAt time.Time `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt" json:"updatedAt"`
}

func CreateContactSubmission(fullname string, email string, subject string, message string) ContactSubmission {

	timenow := time.Now()

	return ContactSubmission{
		ID: primitive.NewObjectIDFromTimestamp(timenow),

		FullName: fullname,
		Email:    email,
		Subject:  subject,
		Message:  message,

		CreatedAt: timenow,
		UpdatedAt: timenow,
	}
}

func (cs ContactSubmission) Upload() error {
	// database setup
	coll := initializers.DB.Database(initializers.DBName).Collection("contacts")

	_, err := coll.InsertOne(context.TODO(), &cs)
	if err != nil {
		return err
	}

	return nil
}
