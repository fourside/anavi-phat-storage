package main

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

func addToFirestore(collectionName string, jsonContent map[string]interface{}) {
	ctx := context.Background()
	opt := option.WithCredentialsFile("./serviceAccountKey.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalln(err)
	}
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	_, _, err = client.Collection(collectionName).Add(ctx, jsonContent)
	if err != nil {
		log.Fatalf("failed to add to store. %v", err)
	}
	defer client.Close()
}
