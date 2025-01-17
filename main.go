package main

import (
    "context"
    "database/sql"
    _"github.com/lib/pq"
    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    "github.com/go-redis/redis/v8"
    "log"
    "strconv"
    "fmt"
    "os"
    "go-simple-crud/controllers"
    "go-simple-crud/repositories"
    "go-simple-crud/services"
)

func main() {
	r := gin.Default()
    
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    db_type := os.Getenv("DB_TYPE")

    postgres_host := os.Getenv("POSTGRES_DB_HOST")
    postgres_port := os.Getenv("POSTGRES_DB_PORT")
    postgres_username := os.Getenv("POSTGRES_DB_USERNAME")
    postgres_password := os.Getenv("POSTGRES_DB_PASSWORD")
    postgres_db := os.Getenv("POSTGRES_DB_NAME")

    redis_host := os.Getenv("REDIS_HOST")
    redis_port := os.Getenv("REDIS_PORT")
    redis_password := os.Getenv("REDIS_PASSWORD")
    redis_db := os.Getenv("REDIS_DB")

    var albumRepo services.AlbumRepository

    switch db_type {
        case "postgres":
            connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", postgres_host, postgres_port, postgres_username, postgres_password, postgres_db)
            db, err := sql.Open("postgres", connStr)
            if err != nil {
                log.Fatal(err)
            }
            defer db.Close()
            createTableSQL := `
                CREATE TABLE IF NOT EXISTS albums (
                    id SERIAL PRIMARY KEY,
                    title VARCHAR(100),
                    artist VARCHAR(100),
                    price NUMERIC
                );`
        
            _, err = db.Exec(createTableSQL)
            if err != nil {
                log.Fatalf("Failed to create table: %v", err)
            }
        
            fmt.Println("Albums table created or already exists.")
            albumRepo = &repositories.PostgresAlbumRepository{DB: db}

        case "redis":
            address := fmt.Sprintf("%s:%s", redis_host, redis_port)
            num, err := strconv.Atoi(redis_db)
            if err != nil {
                fmt.Println("Unsupported Redis DB Type:", err)
                return
            }
            redisClient := redis.NewClient(&redis.Options{
                Addr: address,
                Password: redis_password,
                DB: num,
            })
            ctx := context.Background()
            albumRepo = &repositories.RedisAlbumRepository{Client: redisClient, Ctx: ctx}

        default:
            log.Fatal("Unsupported DB_TYPE. Use 'postgres' or 'redis'.")
    }

    albumService := services.NewAlbumService(albumRepo)

    r.GET("/albums", controllers.GetAlbums(albumService))
    r.POST("/albums", controllers.CreateAlbum(albumService))
	r.PUT("/albums/:id", controllers.UpdateAlbum(albumService))
	r.DELETE("/albums/:id", controllers.DeleteAlbum(albumService))
	r.Run("localhost:8080")
}