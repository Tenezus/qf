package main

import (
	"log"
	"os"
	"time"
	"math/rand"
	"github.com/gin-gonic/gin"
)

type quote struct{
	Id      int    `json: "id"`
	Author  string `json: "author"`
	content string `json: "content"`
}

var quotes = []quote{
	{Id: 1, Author: "Anabella", content: "life is too short, i can't waste my time for them"},
	{Id: 2, Author: "Antoine de saint exuperi", content: "No stress, stress them"},
	{Id: 3, Author: "BITHO Essowassam Fortune", content: "Plus rien ne m'etonnera, a part s'il changent"},
	{Id: 4, Author: "Camille de Kara", content: "If they hurt you, let hurt them"},
	{Id: 5, Author: "Phantom", content: "coder dans l'ombre', illuminer le monde"},
}

func getAllQuotes(c *gin.Context){
	c.IndentedJSON(200, quotes)
}

func getRandomQuote(c *gin.Context){
	rand.Seed(time.Now().Unix())
	c.IndentedJSON(200, quotes[rand.Intn(len(quotes))])
}

func main(){

	router := gin.Default()
	router.GET("/quotes", getAllQuotes)
	router.GET("/randomquote", getRandomQuote)
	port := os.Getenv("PORT")

	if(port == ""){
		port = "8080"
	}

	if err:= router.Run(":"+port); err != nil{
		log.Panicf("error: %s", err)
	}
}