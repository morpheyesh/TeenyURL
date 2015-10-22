package main

import (
	"crypto/rand"
	"encoding/json"
	//"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/garyburd/redigo/redis"
	"regexp"
	"time"
)

const (
	ALPHANUM = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
)

var pool *redis.Pool

type Url struct {
	Id        string
	LongUrl   string
	CreatedAt time.Time
}

func redisConnection() {
	pool = &redis.Pool{
		MaxIdle:     5,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", "localhost:6379")
			if err != nil {
				return nil, err
			}
			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

func GetKey(url string) (*Url, error) {
	redisConnection()

	conn := pool.Get()
	defer conn.Close()

	session := &Url{LongUrl: url, CreatedAt: time.Now()}
	//check whether url isnt empty
	if len(url) == 0 {
		log.Error("Url is empty. Provide a URL")
	}

	if httpcheck, _ := regexp.MatchString("^https?", url); !httpcheck {
		url = "http://" + url
	} else if matches, _ := regexp.MatchString("[.]+", url); !matches {
		log.Error("Please provide a valid URL")
		return nil, nil
	}

	bytes := make([]byte, 6)
	for {
		rand.Read(bytes)
		for i, b := range bytes {
			bytes[i] = ALPHANUM[b%byte(len(ALPHANUM))]
		}
		id := string(bytes)

		//1. new url -> create an id-lurl map
		//2. same url -> fetch and add it up for this session
		if urlexists, _ := redis.Bool(conn.Do("EXISTS", id)); !urlexists {
			session.Id = id
			break
		}

	}
	session.save()
	return session, nil
}

func (s *Url) save() error {
	redisConnection()

	conn := pool.Get()
	defer conn.Close()

	urlData, err := json.Marshal(s)
	if err != nil {
		return err
	}

	reply, err := conn.Do("SET", s.Id, urlData)
	if err == nil && reply != "OK" {
		log.Error("Invalid response from Redis")
	}

	if err != nil {
		return err
	}

	return nil

}

func GetlongUrl(id string) (*Url, error) {
	//TODO: Move this sonofgun out to init
	redisConnection()

	conn := pool.Get()
	defer conn.Close()

	r, err := conn.Do("GET", id)
	if err != nil {
		log.Error("redis connection error")
	}

	d, _ := redis.Bytes(r, err)
	urlz := Url{}
	err = json.Unmarshal(d, &urlz)

	return &urlz, nil

}
