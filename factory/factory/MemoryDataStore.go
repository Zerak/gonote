package factory

import "sync"

//The second implementation.
type MemoryDataStore struct {
	sync.RWMutex
	Users map[int64]string
}

func (mds *MemoryDataStore) ID() int {
	return 100
}

func (mds *MemoryDataStore) Name() string {
	return "MemoryDataStore"
}

func (mds *MemoryDataStore) FindUserNameById(id int64) (string, error) {
	mds.RWMutex.RLock()
	defer mds.RWMutex.RUnlock()
	username, ok := mds.Users[id]
	if !ok {
		return "", UserNotFoundError
	}
	return username, nil
}
