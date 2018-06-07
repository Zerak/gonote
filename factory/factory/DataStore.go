package factory

import (
	"errors"
	"fmt"
	"log"
	"sync"
)

var UserNotFoundError = errors.New("User not found")

type DataStore interface {
	Name() string
	ID() int
	//FindUserNameById(id int64) (string, error)
}

type DataStoreFactory func(conf int) (DataStore, error)

func NewPostgreSQLDataStore(conf int) (DataStore, error) {
	return &PostgreSQLDataStore{
		DSN: "",
		//DB:  nil,
	}, nil
}

func NewMemoryDataStore(conf int) (DataStore, error) {
	return &MemoryDataStore{
		Users: map[int64]string{
			1: "mnbbrown",
			0: "root",
		},
		RWMutex: sync.RWMutex{},
	}, nil
}

var datastoreFactories = make(map[int]DataStoreFactory)

func Register(name int, factory DataStoreFactory) {
	if factory == nil {
		log.Panicf("Datastore factory %s does not exist.", name)
	}
	_, registered := datastoreFactories[name]
	if registered {
		fmt.Errorf("Datastore factory %s already registered. Ignoring.", name)
	}
	datastoreFactories[name] = factory
}

func init() {
	//Register("postgres", NewPostgreSQLDataStore)
	//Register("memory", NewMemoryDataStore)
	Register(100, NewPostgreSQLDataStore)
	Register(9999, NewMemoryDataStore)
}

func GetDataStore(conf int) (DataStore, error) {
	// Query configuration for datastore defaulting to "memory".
	engineName := conf

	engineFactory, ok := datastoreFactories[engineName]
	if !ok {
		// Factory has not been registered.
		// Make a list of all available datastore factories for logging.
		//availableDatastores := make([]int, len(datastoreFactories))
		//for k, _ := range datastoreFactories {
		//	availableDatastores = append(availableDatastores, k)
		//}
		//return nil, errors.New(fmt.Sprintf("Invalid Datastore name. Must be one of: %s", strings.Join(availableDatastores, ", ")))
		return nil, fmt.Errorf("err")
	}

	// Run the factory with the configuration.
	return engineFactory(conf)
}
