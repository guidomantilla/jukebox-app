package main

import (
	"log"

	"github.com/bradfitz/gomemcache/memcache"
)

func main() {

	mc := memcache.New("localhost:11211")

	err := mc.Set(&memcache.Item{Key: "foo", Value: []byte("my value")})
	if err != nil {
		log.Fatal(err)
		return
	}

	it, err := mc.Get("foo")
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Println(it.Key, string(it.Value), it.Flags, it.Expiration)
}
