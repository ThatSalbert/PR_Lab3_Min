package items

import (
	"sync"
	"time"
)

var lock = &sync.Mutex{}

type Table struct {
	Message string `json:"message"`
	Time    int64  `json:"time"`
}

type Datastore map[int]Table

var Instance *Datastore

func GetInstance() *Datastore {
	if Instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if Instance == nil {
			Instance = &Datastore{}
		}
	}
	return Instance
}

func (ds Datastore) Create(UserId int, Message string) *Table {
	if ds.Check(UserId) {
		return nil
	}
	ds[UserId] = Table{
		Message: Message,
		Time:    time.Now().UnixNano(),
	}
	var returnTable = ds[UserId]
	return &returnTable
}

func (ds Datastore) Read(UserId int) *Table {
	if !ds.Check(UserId) {
		return nil
	}
	var returnTable = ds[UserId]
	return &returnTable
}

func (ds Datastore) Update(UserId int, Message string) *Table {
	if !ds.Check(UserId) {
		return nil
	}
	ds[UserId] = Table{
		Message: Message,
		Time:    time.Now().UnixNano(),
	}
	var returnTable = ds[UserId]
	return &returnTable
}

func (ds Datastore) Delete(UserId int) {
	if ds.Check(UserId) {
		delete(ds, UserId)
	}
}

func (ds Datastore) Check(UserId int) (ok bool) {
	_, ok = ds[UserId]
	return ok
}
