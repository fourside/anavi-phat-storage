package main

import (
	"bufio"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
);

func main() {

  var collection string
  flag.StringVar(&collection, "collection", "", "a collection name")
  flag.Parse()
  if collection == "" {
    log.Fatal("pass --collection {collection}")
  }
  fmt.Println(collection)

  fmt.Println(os.Args[1:])
  stdin := bufio.NewScanner(os.Stdin)
  stdin.Scan()
  text := stdin.Text()
  fmt.Println(text)
  var result map[string]interface{}
  json.Unmarshal([]byte(text), &result)
  result["date"] = time.Now().Format(time.RFC3339)
  fmt.Println(result)

  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file", err)
  }
  // user := os.Getenv("MONGO_USER")
  // password := os.Getenv("MONGO_PASSWORD")
  database := os.Getenv("MONGO_DATABASE")
  // mongoUri := fmt.Sprintf("mongodb://%s:%s@localhost:27017", user, password)
  mongoUri := "mongodb://localhost:27017"

  client, err := mongo.NewClient(options.Client().ApplyURI(mongoUri))
  if err != nil {
    log.Fatal("mongo.NewClient: ", err)
  }

  ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
  defer cancel()
  err = client.Connect(ctx)
  if err != nil {
    log.Fatal("client.Connect: ", err)
  }

  dbCollection := client.Database(database).Collection(collection)

  res, err := dbCollection.InsertOne(ctx, result)
  if err != nil {
    log.Fatal("collection.InsertOne: ", err)
  }
  fmt.Println("hooray!", res)
  client.Disconnect(ctx)
  };
