package di

import (
	"EM/internal/domain"
	"EM/internal/handlers"
	"EM/internal/handlers/middleware"
	"EM/internal/repository"
	"EM/internal/usecase"
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"net/http"
	"os"

	_ "github.com/jackc/pgx/v5/pgxpool"
)

type Container struct {
	router http.Handler
	pool   *pgxpool.Pool

	personRepository  *repository.PersonRepository
	createPerson      *usecase.CreatePersonUseCase
	postPersonHandler *handlers.POSTPersonHandler
}

func NewContainer(ctx context.Context) *Container {
	pool, err := CreateConnection(ctx)
	if err != nil {
		fmt.Println("error: %w", err)
	}
	return &Container{
		pool: pool,
	}
}

func (c *Container) POSTPersonHandler() *handlers.POSTPersonHandler {
	if c.postPersonHandler == nil {
		c.postPersonHandler = handlers.NewPOSTPersonHandler(c.CreatePerson())
	}
	return c.postPersonHandler
}

func (c *Container) CreatePerson() *usecase.CreatePersonUseCase {
	if c.createPerson == nil {
		c.createPerson = usecase.NewCreatePersonUseCase(c.PersonRepository())
	}
	return c.createPerson
}

func (c *Container) PersonRepository() domain.PersonRepository {
	if c.personRepository == nil {
		c.personRepository = repository.NewPersonRepository(c.pool)
	}
	return c.personRepository
}

func (c *Container) HTTPRouter() http.Handler {
	if c.router != nil {
		return c.router
	}
	router := mux.NewRouter()
	router.Use(middleware.Recover)

	router.Handle("/person", c.POSTPersonHandler()).Methods(http.MethodPost)
	c.router = router
	return c.router

}

func CreateConnection(ctx context.Context) (*pgxpool.Pool, error) {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("error: %w", err)
		fmt.Println(".env file not found")
	}
	dns := os.Getenv("DATABASE_URL")
	pool, err := pgxpool.New(ctx, dns)
	if err != nil {
		return nil, err
	}
	err = pool.Ping(ctx)
	if err != nil {
		return nil, err
	}
	return pool, err
}

func (c *Container) Close() {
	c.pool.Close()
}
