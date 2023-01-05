package repository

import "log"

type repo struct {
	Caching *cache
	Auth    *auth
}

// NewRepository creates a new repository
func NewRepository() *repo {
	//I need to build a in memory caching layer which takes over incase of the redis caching layer failing.
	cache, err := newCache()
	if err != nil {
		log.Println(err)
	}
	auth, err := newAuth()
	if err != nil {
		log.Fatalln(err.Error())
	}

	return &repo{
		Caching: cache,
		Auth:    auth,
	}
}
