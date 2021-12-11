package redis

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	redisClient "github.com/gomodule/redigo/redis"
	"github.com/tech-thinker/linkly/base62"
	"github.com/tech-thinker/linkly/models"
	"github.com/tech-thinker/linkly/storage"
)

type redis struct{ pool *redisClient.Pool }

// New  creates a new redis storage
func New(host, port, password string) (storage.Service, error) {
	pool := &redisClient.Pool{
		MaxIdle:     10,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redisClient.Conn, error) {
			return redisClient.Dial("tcp", fmt.Sprintf("%s:%s", host, port))
		},
	}

	return &redis{pool}, nil
}

// isUsed checks if the id is used
func (r *redis) isUsed(id uint64) bool {
	conn := r.pool.Get()
	defer conn.Close()

	exists, err := redisClient.Bool(conn.Do("EXISTS", "Shortener:"+strconv.FormatUint(id, 10)))
	if err != nil {
		return false
	}
	return exists
}

// Save saves a new link
func (r *redis) Save(url string, expires time.Time) (string, error) {
	conn := r.pool.Get()
	defer conn.Close()

	var id uint64

	for used := true; used; used = r.isUsed(id) {
		id = rand.Uint64()
	}

	shortLink := models.URL{id, url, expires.Format("2006-01-02 15:04:05.728046 +0300 EEST"), 0}

	_, err := conn.Do("HMSET", redisClient.Args{"Shortener:" + strconv.FormatUint(id, 10)}.AddFlat(shortLink)...)
	if err != nil {
		return "", err
	}

	_, err = conn.Do("EXPIREAT", "Shortener:"+strconv.FormatUint(id, 10), expires.Unix())
	if err != nil {
		return "", err
	}

	return base62.Encode(id), nil
}

// Load visits count
func (r *redis) Load(code string) (string, error) {
	conn := r.pool.Get()
	defer conn.Close()

	decodedId, err := base62.Decode(code)
	if err != nil {
		return "", err
	}

	urlString, err := redisClient.String(conn.Do("HGET", "Shortener:"+strconv.FormatUint(decodedId, 10), "url"))
	if err != nil {
		return "", err
	} else if len(urlString) == 0 {
		return "", errors.New("no link found")
	}

	_, err = conn.Do("HINCRBY", "Shortener:"+strconv.FormatUint(decodedId, 10), "visits", 1)

	return urlString, nil
}

// isAvailable checks if the id is available
func (r *redis) isAvailable(id uint64) bool {
	conn := r.pool.Get()
	defer conn.Close()

	exists, err := redisClient.Bool(conn.Do("EXISTS", "Shortener:"+strconv.FormatUint(id, 10)))
	if err != nil {
		return false
	}
	return !exists
}

// LoadInfo returns the link info
func (r *redis) LoadInfo(code string) (*models.URL, error) {
	conn := r.pool.Get()
	defer conn.Close()

	decodedId, err := base62.Decode(code)
	if err != nil {
		return nil, err
	}

	values, err := redisClient.Values(conn.Do("HGETALL", "Shortener:"+strconv.FormatUint(decodedId, 10)))
	if err != nil {
		return nil, err
	} else if len(values) == 0 {
		return nil, errors.New("no link found")
	}
	var shortLink models.URL
	err = redisClient.ScanStruct(values, &shortLink)
	if err != nil {
		return nil, err
	}

	return &shortLink, nil
}

// Close closes the redis connection
func (r *redis) Close() error {
	return r.pool.Close()
}
