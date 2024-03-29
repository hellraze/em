package di

import (
	"EM/internal/domain"
	"EM/internal/handlers"
	"EM/internal/handlers/middleware"
	"EM/internal/pkg/logging"
	"EM/internal/repository"
	"EM/internal/usecase"
	"context"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
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

	deletePerson        *usecase.DeletePersonUseCase
	deletePersonHandler *handlers.DeletePersonHandler

	putPerson        *usecase.PutPersonUseCase
	putPersonHandler *handlers.PutPersonHandler

	readPerson       *usecase.ReadPersonUseCase
	getPersonHandler *handlers.GETPeopleHandler
}

func NewContainer(ctx context.Context) *Container {
	log := logging.NewLog()
	log.Init()

	pool, err := CreateConnection(ctx)
	if err != nil {
		log.Log.WithFields(logrus.Fields{
			"container.go": "NewContainer() 46",
		}).Errorf("Ошибка %w", err)
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

func (c *Container) DeletePerson() *usecase.DeletePersonUseCase {
	if c.deletePerson == nil {
		c.deletePerson = usecase.NewDeletePersonUseCase(c.PersonRepository())
	}
	return c.deletePerson
}

func (c *Container) DeletePersonHandler() *handlers.DeletePersonHandler {
	if c.deletePersonHandler == nil {
		c.deletePersonHandler = handlers.NewDeletePersonHandler(c.DeletePerson())
	}
	return c.deletePersonHandler
}

func (c *Container) PutPerson() *usecase.PutPersonUseCase {
	if c.putPerson == nil {
		c.putPerson = usecase.NewPutPersonUseCase(c.PersonRepository())
	}
	return c.putPerson
}

func (c *Container) PutPersonHandler() *handlers.PutPersonHandler {
	if c.putPersonHandler == nil {
		c.putPersonHandler = handlers.NewPutPersonHandler(c.PutPerson())
	}
	return c.putPersonHandler
}

func (c *Container) ReadPerson() *usecase.ReadPersonUseCase {
	if c.readPerson == nil {
		c.readPerson = usecase.NewReadPersonUseCase(c.PersonRepository())
	}
	return c.readPerson
}

func (c *Container) GETPersonHandler() *handlers.GETPeopleHandler {
	if c.getPersonHandler == nil {
		c.getPersonHandler = handlers.NewGETPeopleHandler(c.ReadPerson())
	}
	return c.getPersonHandler
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

	router.Handle("/api/people", c.POSTPersonHandler()).Methods(http.MethodPost)
	router.Handle("/api/people", c.DeletePersonHandler()).Methods(http.MethodDelete)
	router.Handle("/api/people", c.PutPersonHandler()).Methods(http.MethodPut)
	router.Handle("/api/people", c.GETPersonHandler()).Methods(http.MethodGet)
	c.router = router
	return c.router

}

func CreateConnection(ctx context.Context) (*pgxpool.Pool, error) {
	log := logging.NewLog()
	log.Init()

	err := godotenv.Load(".env")
	if err != nil {
		log.Log.WithFields(logrus.Fields{
			"container.go": "CreateConnection() 136",
		}).Errorf("Ошибка %w", err)
	}
	dns := os.Getenv("DATABASE")
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
