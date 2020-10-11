package model

import (
	"github.com/hongshengjie/framework/app/api/blog"
	bb "github.com/hongshengjie/framework/database/blotdb"

	"github.com/golang/protobuf/ptypes"
	"google.golang.org/protobuf/proto"
)

// Upsert Upsert
func Upsert(p *blog.Post) error {
	bkt := pstbkt()
	p.Mtime = ptypes.TimestampNow()
	if p.Id == 0 {
		id, _ := bb.AutoIncr(bkt)
		p.Id = int64(id)
		p.Ctime = ptypes.TimestampNow()
	}
	v, err := proto.Marshal(p)
	if err != nil {
		return err
	}
	return bb.Put(bkt, itob(uint64(p.Id)), v)
}

//PostByID PostByID
func PostByID(id uint64) (*blog.Post, error) {
	var p blog.Post
	bkt := pstbkt()
	v := bb.Get(bkt, itob(id))
	if err := proto.Unmarshal(v, &p); err != nil {
		return nil, err
	}
	return &p, nil
}

// PostList PostList
func PostList(prekey int64, limit int64, asc bool) ([]*blog.Post, error) {
	var r []bb.BoltPair
	var key []byte
	if prekey > 0 {
		key = itob(uint64(prekey))
	}
	if asc {
		r = bb.LimitKeyValuesAfterKey(pstbkt(), key, limit)

	} else {
		r = bb.LimitKeyValuesBeforeKey(pstbkt(), key, limit)
	}
	var ret []*blog.Post
	for _, v := range r {
		var p blog.Post
		if err := proto.Unmarshal(v.Value, &p); err != nil {
			continue
		}
		ret = append(ret, &p)

	}
	return ret, nil

}

func pstbkt() []byte {
	var p *blog.Post
	return []byte(p.ProtoReflect().Descriptor().Name())
}
