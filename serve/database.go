package serve

import "time"

type Database struct {
	Name    string    `json:"name"`
	Created bool      `json:"created"`
	Updated time.Time `json:"updated"`
}

type Databases []Database
