package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_"github.com/lib/pq"    //still include as we are not direclty callign it 
)

func main(){
	fmt.Println("hello world")
	godotenv.Load(".env")

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("Port is not found in the environment")
	}

	type apiConfig struct{
		DB *database.Queries
	}

	db_URL := os.Getenv("DB_URL")
	if db_URL != nil{
		log.Fatal("DB_URL is not found in the environment")
	}

	conn,err := sql.open("postgres",db_URL)
	if err != nil {
		log.Fatal("err while connectinig to the database")
	}


	router :=  chi.NewRouter()
	

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: false,
		MaxAge: 300,
	}))

	v1Router := chi.NewRouter() 

	v1Router.Get("/healtz",handlerReadiness)
	v1Router.Get("/error",handleErr)
	router.Mount("/v1",v1Router)


	srv := &http.Server{
		Handler: router,
		Addr: ":"+portString, 
	}


	fmt.Printf("Server is starting on %v ",portString)
	err := srv.ListenAndServe()
	if err != nil{
		log.Fatal("error in starting the server")
	}
}