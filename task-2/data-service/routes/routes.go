package routes

import (
	"cisdi-technical-assessment/REST/data-service/controller"
	"cisdi-technical-assessment/REST/data-service/model/dto"
	"cisdi-technical-assessment/REST/data-service/repository"
	"cisdi-technical-assessment/REST/data-service/service"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupRouter(cfg *dto.Config) *gin.Engine {
	r := gin.Default()

	// Connect to the database
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.Name,
		cfg.Database.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("failed to connect to database: %w", err))
	}

	// Setup repositories
	bookRepo := repository.NewBookRepository(db)
	authorRepo := repository.NewAuthorRepository(db)
	publisherRepo := repository.NewPublisherRepository(db)

	// Setup services
	bookService := service.NewBookService(bookRepo, authorRepo, publisherRepo)
	authorService := service.NewAuthorService(authorRepo)
	publisherService := service.NewPublisherService(publisherRepo)

	// Setup controllers
	bookController := controller.NewBookController(bookService)
	authorController := controller.NewAuthorController(authorService)
	publisherController := controller.NewPublisherController(publisherService)

	// Define routes
	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			data := v1.Group("/data")
			{
				author := data.Group("/author")
				{
					author.GET("/all", authorController.GetAllAuthors)
					author.GET("/:id", authorController.GetAuthorByID)
					author.POST("/create", authorController.CreateAuthor)
					author.POST("/update", authorController.UpdateAuthor)
					author.POST("/delete", authorController.DeleteAuthor)
				}

				publisher := data.Group("/publisher")
				{
					publisher.GET("/all", publisherController.GetAllPublishers)
					publisher.GET("/:id", publisherController.GetAllPublishers)
					publisher.POST("/create", publisherController.CreatePublisher)
					publisher.POST("/update", publisherController.UpdatePublisher)
					publisher.POST("/delete", publisherController.DeletePublisher)
				}

				book := data.Group("/book")
				{
					book.GET("/all", bookController.GetAllBooks)
					book.GET("/:id", bookController.GetBookByID)
					book.POST("/create", bookController.CreateBook)
				}
			}
		}
	}

	return r
}
