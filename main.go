package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

type Quotes struct {
	Id      	uint64    		`json: "id"`
	Content 	string 		`json: "content"`
	Author  	string 		`json: "author"`
	Created_at 	string 	`json: "created_at"`
	Like 		uint64 		`json: "like"`
}

var quotes []Quotes


//postgresql://postgres:[YOUR-PASSWORD]@db.qxfmtqrnjdfiptrfpzsq.supabase.co:5432/postgres

func getAllQuotes(c *gin.Context) {
	//select all quotes [id, content, author, created_at, like]
	//connection to the postgres database on supabase
	
	var conn *pgx.Conn

	conn, err := pgx.Connect(context.Background(), "postgresql://postgres.qxfmtqrnjdfiptrfpzsq:jevaisalecole@aws-0-eu-north-1.pooler.supabase.com:6543/postgres")
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}
	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(), `select "id", "content", "author", "created_at","like" from quote;`, pgx.QueryExecModeSimpleProtocol)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(context.Background())

	i := len(rows.FieldDescriptions())

	fmt.Println(i)

	for rows.Next() {
		var q Quotes
		err := rows.Scan(&q.Id, &q.Content, &q.Author, &q.Created_at, &q.Like)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(q.Id, "'"+q.Content+"'", q.Author, q.Created_at, q.Like)
		//quote := quote{Id: id, Content: content, Author: author, Created_at: created_at, Like: like}
		quotes = append(quotes, q)
	}
	c.IndentedJSON(200, quotes)
}

func main() {

	//initialize a router
	router := gin.Default()

	//handling all quotes
	router.GET("/quotes", getAllQuotes)

	//serve on port 8080
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	if err := router.Run(":" + port); err != nil {
		log.Panicf("error: %s", err)
	}
}
