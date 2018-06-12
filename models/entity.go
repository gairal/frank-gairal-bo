package models

import (
	"context"
	"log"

	"google.golang.org/appengine/datastore"
	// "cloud.google.com/go/datastore"
)

// Entity - abstract entity
type Entity struct {
	// DbClient *datastore.Client
}

// InitDbClient - Init Data Store Client
// func InitDbClient() *Entity {
// 	ctx := context.Background()
// 	projectID := "com-gairal-frank-api"

// 	client, err := datastore.NewClient(ctx, projectID)
// 	if err != nil {
// 		log.Fatalf("Failed to create client: %v", err)
// 		panic(err)
// 	}

// 	e := &Entity{DbClient: client}

// 	return e
// }

// GetAll - Return all of type entities
func (e *Entity) getAll(ctx context.Context, q *datastore.Query, entities interface{}, orderCol string, descending bool) {
	// ctx := context.Background()

	if orderCol != "" {
		var order = ""
		if descending {
			order = "-"
		}

		q = q.Order(order + orderCol)
	}

	if _, err := q.GetAll(ctx, entities); err != nil {
		log.Fatalf("Failed to get entities: %v", err)
		panic(err)
	}
}

// Get - Return an entity by its key
func (e *Entity) get(ctx context.Context, kind string, k int64, entity interface{}) {
	key := datastore.NewKey(ctx, kind, "", k, nil)
	// key := datastore.IDKey(kind, k, nil)

	if err := datastore.Get(ctx, key, entity); err != nil {
		log.Fatalf("Failed to get entity: %v", err)
		panic(err)
	}
}
