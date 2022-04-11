package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const MONGO_COLLECTION = "tmp";

type Product struct {
	name   string
	price  int64
	amount int32
}

var mongoClient *mongo.Client;

func main() {

	var err error

	fmt.Printf("Program started...\n")

	// connect to the db
	mongoClient, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGO_URI")));

	if err != nil {
		panic(err);
	}

	if err = mongoClient.Ping(context.TODO(), readpref.Primary()); err != nil {
        panic(err);
	}

	fmt.Printf("Database initialization seems ok...\n");

	http.HandleFunc("/ping", ping);
	http.HandleFunc("/get", get);
	http.HandleFunc("/list", list);
	http.HandleFunc("/update", update);
	http.HandleFunc("/delete", delete);

	http.ListenAndServe(":8080", nil);

}

func ping(wr http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(wr, "pong")

}

func update(wr http.ResponseWriter, req *http.Request) {

	params := req.URL.Query();

	coll := mongoClient.Database("godb").Collection(MONGO_COLLECTION);

	product := Product{
		name: params.Get("name"),
		price: str2int(params.Get("price")),
		amount: int32(str2int(params.Get("amount"))),
	};

	doc := bson.D{
		{"name", params.Get("name")}, 
		{"price", str2int(params.Get("price"))}, 
		{"amount", int32(str2int(params.Get("amount")))},
	};

	result, err := coll.InsertOne(context.TODO(), doc);
	
	if err != nil {
		fmt.Fprintf(wr, "Error\nvalue: %i", err);
		fmt.Printf("Error while inserting!\n");
	} else {
		fmt.Fprintf(wr, "Inserted\nID: %i", result.InsertedID);
		fmt.Printf("Inserted: \n\tname: %s\n\tprice: %v\n\tamount: %v\n", product.name, product.price, product.amount);
	}

}

func get(wr http.ResponseWriter, req *http.Request) {

	params := req.URL.Query();

	coll := mongoClient.Database("godb").Collection(MONGO_COLLECTION);

	//filter := bson.D{
	//		{"name", params.Get("name")},
	//};

	filter := bson.M{
		"name": params.Get("name"),
	};

	cursor, err := coll.Find(context.TODO(), filter);
	
	if err != nil {
		fmt.Fprintf(wr, "Error\nvalue: %i", err);
		fmt.Printf("Error while geting cursor!\n");
		return;
	}
	
	var results []bson.M;
	if err = cursor.All(context.TODO(), &results); err != nil {
		fmt.Fprintf(wr, "Error\nvalue: %i", err);
		fmt.Printf("Error while reading!\n");
		return;
	}

	var sb strings.Builder;
	for i := 0; i < len(results); i++ {
		var result bson.M = results[i];
		sb.WriteString(
			fmt.Sprintf("Product:\n\tname: %s\n\tprice: %v\n\tamount: %v\n", result["name"], result["price"], result["amount"]),
		);
	}

	var finalString string = sb.String();
	fmt.Fprintf(wr, finalString);
	fmt.Println(finalString);

}

func list(wr http.ResponseWriter, req *http.Request) {

	coll := mongoClient.Database("godb").Collection(MONGO_COLLECTION);

	cursor, err := coll.Find(context.TODO(), bson.M{});
	
	if err != nil {
		fmt.Fprintf(wr, "Error\nvalue: %i", err);
		fmt.Printf("Error while geting cursor!\n");
		return;
	}
	
	var results []bson.M;
	if err = cursor.All(context.TODO(), &results); err != nil {
		fmt.Fprintf(wr, "Error\nvalue: %i", err);
		fmt.Printf("Error while reading!\n");
		return;
	}

	var sb strings.Builder;
	for i := 0; i < len(results); i++ {
		var result bson.M = results[i];
		sb.WriteString(
			fmt.Sprintf("Product:\n\tname: %s\n\tprice: %v\n\tamount: %v\n", result["name"], result["price"], result["amount"]),
		);
	}

	var finalString string = sb.String();
	fmt.Fprintf(wr, finalString);
	fmt.Println(finalString);

}

func delete(wr http.ResponseWriter, req *http.Request) {

	params := req.URL.Query();

	coll := mongoClient.Database("godb").Collection(MONGO_COLLECTION);

	filter := bson.M{
		"name": params.Get("name"),
	};
	//options := options.Delete();

	result, err := coll.DeleteMany(context.TODO(), filter);
	
	if err != nil {
		fmt.Fprintf(wr, "Error\nvalue: %i", err);
		fmt.Printf("Error while deleting!\n");
		return;
	}

	fmt.Fprintf(wr, "Deleted: \ncount: %v", result.DeletedCount);
	fmt.Printf("Deleted: \ncount: %v\n", result.DeletedCount);

}

func str2int(str string) int64 {

	val, err := strconv.ParseInt(str, 10, 64);
	
	if (err != nil) {
		return 0;
	}

	return val;

}
