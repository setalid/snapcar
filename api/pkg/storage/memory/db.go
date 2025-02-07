package memory

import (
	"context"
	"encoding/json"
	"errors"
	"strings"
	"sync"
)

var (
	ErrAlreadyExist = errors.New("already exists")
	ErrNotFound     = errors.New("not found")
)

type DB struct {
	lock   sync.Mutex
	tables map[string]map[string][]byte
}

func New() *DB {
	return &DB{
		tables: make(map[string]map[string][]byte),
	}
}

func (db *DB) createTableIfNotExists(table string) {
	if _, exists := db.tables[table]; !exists {
		db.tables[table] = make(map[string][]byte)
	}
}

func (db *DB) Get(_ context.Context, table, key string, v any) error {
	db.lock.Lock()
	defer db.lock.Unlock()

	if _, exists := db.tables[table]; !exists {
		return ErrNotFound
	}

	b, exists := db.tables[table][key]
	if !exists {
		return ErrNotFound
	}

	return json.Unmarshal(b, v)
}

func (db *DB) Set(_ context.Context, table, key string, v any, force bool) error {
	db.lock.Lock()
	defer db.lock.Unlock()

	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	db.createTableIfNotExists(table)

	if force {
		db.tables[table][key] = b
		return nil
	}

	if _, exists := db.tables[table][key]; exists {
		return ErrAlreadyExist
	}

	db.tables[table][key] = b

	return nil
}

func (db *DB) All(_ context.Context, table string, v any) error {
	db.lock.Lock()
	defer db.lock.Unlock()

	if _, exists := db.tables[table]; !exists {
		return ErrNotFound
	}

	var builder strings.Builder
	builder.WriteString("[")

	first := true
	for _, jsonBytes := range db.tables[table] {
		if !first {
			builder.WriteString(",")
		}
		first = false
		builder.Write(jsonBytes)
	}
	builder.WriteString("]")

	return json.Unmarshal([]byte(builder.String()), v)
}
