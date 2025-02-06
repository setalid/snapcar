package memory

import (
	"context"
	"errors"
	"sync"
)

var (
	ErrAlreadyExist = errors.New("already exists")
	ErrNotFound     = errors.New("not found")
)

type DB struct {
	lock sync.Mutex
	m    map[string]string
}

func New() *DB {
	return &DB{
		m: make(map[string]string),
	}
}

func (db *DB) Get(_ context.Context, key string) (string, error) {
	db.lock.Lock()
	defer db.lock.Unlock()

	v, exists := db.m[key]
	if !exists {
		return "", ErrNotFound
	}

	return v, nil
}

func (db *DB) Set(_ context.Context, key string, v string, force bool) error {
	db.lock.Lock()
	defer db.lock.Unlock()

	if force {
		db.m[key] = v
		return nil
	}

	if _, exists := db.m[key]; exists {
		return ErrAlreadyExist
	}

	db.m[key] = v

	return nil
}
