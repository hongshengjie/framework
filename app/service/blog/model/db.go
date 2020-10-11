package model

import (
	"encoding/binary"
	"framework/app/api/blog"
	"framework/conf"
	"framework/database/blotdb"
)

// Init Init
func Init() {
	var config struct {
		Title string
		Db    struct {
			Dbpath  string
			Buckets []string
		}
		Nav []*blog.Nav
	}
	err := conf.Get("config.toml", &config)
	if err != nil {
		panic(err)
	}
	nav = config.Nav

	err = blotdb.InitBolt(config.Db.Dbpath, config.Db.Buckets)
	if err != nil {
		panic(err)
	}
}

// Close close the db
func Close() {
	blotdb.Close()
}

func itob(v uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, v)
	return b
}
