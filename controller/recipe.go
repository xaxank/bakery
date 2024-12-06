package controller

import (
	"context"
	"log"
	"fmt"
	"net/http"
	"encoding/json"
	
	"bakery/models"
	"bakery/adapters"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
)

var ctx = context.TODO()

func SetupRecipeHandler(handler *http.ServeMux) {
	handler.HandleFunc("/recipes", findRecipe)
	handler.HandleFunc("/recipes/create", CreateRecipe)
	// handler.HandleFunc("/recipes/update", UpdateRecipe)
	// handler.HandleFunc("/recipes/delete", DeleteRecipe)
	handler.HandleFunc("/recipes-all", GetAllRecipes)
	
}

func CreateRecipe(w http.ResponseWriter, r *http.Request) {

	recipe := models.Recipe{
		ID: primitive.NewObjectID(),
		Name: "Chocolate Cake",	
		Ingredients: []models.Ingredient{
			{Name: "Flour", Type: "Base"},
			{Name: "Sugar", Type: "Sweetener"},
			{Name: "Cocoa", Type: "Flavoring"},
		},
	}

	collection := adapters.GetClient()

	_, err := collection.InsertOne(ctx, recipe)
	if err != nil {
		fmt.Println("cs is now", err)
		log.Fatal(err)
	}

	j, _ := json.Marshal(recipe)
    w.Write(j)	
}

func findRecipe (w http.ResponseWriter, r *http.Request) {
	collection := adapters.GetClient()
	var recipe models.Recipe

	err := collection.FindOne(ctx, bson.M{"name": "Chocolate Cake"}).Decode(&recipe)
	if err != nil {
		log.Fatal(err)
	}

	j, _ := json.Marshal(recipe)
    w.Write(j)	
}

func GetAllRecipes(w http.ResponseWriter, r *http.Request) {
	collection := adapters.GetClient()
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	defer cursor.Close(ctx)

	var recipes []models.Recipe
	if err = cursor.All(ctx, &recipes); err != nil {
		log.Fatal(err)
	}

	j, _ := json.Marshal(recipes)
    w.Write(j)
	// fmt.Fprintf(w, "Found multiple documents (array of pointers): %+v\n", recipes)

}