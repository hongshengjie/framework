package blotdb

import (
	"log"
	"os"
	"time"

	"go.etcd.io/bbolt"
)

var database string
var fileMode os.FileMode = 0600 // owner can read and write
var db *bbolt.DB

// DB DB
func DB() *bbolt.DB {
	return db
}

// InitBolt inits bolt database. Create the file if not exist.
// By default it opens the file in 0600 mode, with a 10 seconds timeout period
func InitBolt(path string, buckets []string) error {
	log.Println("Trying to open databse")
	database = path
	var err error
	// open the target file, file mode fileMode, and a 10 seconds timeout period
	db, err = bbolt.Open(database, fileMode, &bbolt.Options{Timeout: 10 * time.Second})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Trying to create buckets")
	err = db.Update(func(tx *bbolt.Tx) error {
		for _, value := range buckets {
			_, err := tx.CreateBucketIfNotExists([]byte(value))
			if err != nil {
				return err
			}
		}

		return nil
	})

	log.Println("Trying to return error")
	return err
}

// Close bolt db
func Close() {
	db.Close()
}

// Get value from bucket by key
func Get(bucket []byte, key []byte) []byte {
	var value []byte

	err := db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket(bucket)
		v := b.Get(key)
		if v != nil {
			value = append(value, b.Get(key)...)
		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	return value
}

// AutoIncr AutoIncr
func AutoIncr(bucket []byte) (uint64, error) {
	var id uint64
	err := db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket(bucket)
		var err error
		id, err = b.NextSequence()
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return 0, err
	}
	return id, nil
}

// Put a key/value pair into target bucket
func Put(bucket []byte, key []byte, value []byte) error {
	err := db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket(bucket)
		err := b.Put(key, value)
		return err
	})

	return err
}

// Delete a key from target bucket
func Delete(bucket []byte, key []byte) error {
	err := db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket(bucket)
		err := b.Delete(key)
		return err
	})

	return err
}

// GetAllKeys get all keys from the target bucket
func GetAllKeys(bucket []byte) [][]byte {
	var keys [][]byte

	db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket(bucket)
		b.ForEach(func(k, v []byte) error {
			dst := make([]byte, len(k))
			copy(dst, k)
			keys = append(keys, dst)
			return nil
		})
		return nil
	})

	return keys
}

// BoltPair is a struct to store key/value pair data
type BoltPair struct {
	Key   []byte
	Value []byte
}

// GetAllKeyValues get all key/value pairs from a bucket in BoltPair struct format
func GetAllKeyValues(bucket []byte) []BoltPair {
	var pairs []BoltPair

	db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket(bucket)
		b.ForEach(func(k, v []byte) error {
			dstk := make([]byte, len(k))
			dstv := make([]byte, len(v))
			copy(dstk, k)
			copy(dstv, v)

			pair := BoltPair{dstk, dstv}
			pairs = append(pairs, pair)
			return nil
		})

		return nil
	})

	return pairs
}

// LimitKeyValuesAfterKey LimitKeyValuesAfterKey
func LimitKeyValuesAfterKey(bucket, key []byte, limit int64) []BoltPair {
	var pairs []BoltPair
	db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket(bucket)
		c := b.Cursor()
		var k, v []byte

		if key == nil {
			k, v = c.First()
		} else {
			c.Seek(key)
			k, v = c.Next()

		}
		for i := int64(0); k != nil && i < limit; k, v = c.Next() {
			dstk := make([]byte, len(k))
			dstv := make([]byte, len(v))
			copy(dstk, k)
			copy(dstv, v)
			pair := BoltPair{dstk, dstv}
			pairs = append(pairs, pair)
			i++
		}
		return nil
	})

	return pairs
}

// LimitKeyValuesBeforeKey LimitKeyValuesBeforeKey
func LimitKeyValuesBeforeKey(bucket, key []byte, limit int64) []BoltPair {
	var pairs []BoltPair
	db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket(bucket)
		c := b.Cursor()
		var k, v []byte
		if key == nil {
			k, v = c.Last()
		} else {
			c.Seek(key)
			k, v = c.Prev()

		}
		for i := int64(0); k != nil && i < limit; k, v = c.Prev() {
			dstk := make([]byte, len(k))
			dstv := make([]byte, len(v))
			copy(dstk, k)
			copy(dstv, v)
			pair := BoltPair{dstk, dstv}
			pairs = append(pairs, pair)
			i++
		}
		return nil
	})

	return pairs
}
