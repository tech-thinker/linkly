package redis

import (
	"context"
	"time"

	redisClient "github.com/go-redis/redis/v8"
	"github.com/tech-thinker/linkly/models"
	"github.com/tech-thinker/linkly/storage"
)

type redis struct{ rdb *redisClient.Client }

// New  creates a new redis storage
func New(host, port, password string) (storage.Service, error) {
	rdb := redisClient.NewClient(&redisClient.Options{
		Addr:     host + ":" + port,
		Password: password,
		DB:       0,
	})
	return &redis{rdb}, nil
}

// Save saves a new link
func (r *redis) Save(url string, expires time.Time) (string, error) {
	return "hash", nil
}

// Load return a link
func (r *redis) Load(hash string) (string, error) {
	return "url", nil
}

// LoadAll return all links
func (r *redis) LoadInfo(code string) (*models.URL, error) {
	return &models.URL{}, nil
}

// Close closes the redis connection
func (r *redis) Close() error {
	var ctx = context.Background()
	conn := r.rdb.Conn(ctx)
	defer conn.Close()
	if err := conn.ClientSetName(ctx, "client").Err(); err != nil {
		return err
	}
	return nil
}
