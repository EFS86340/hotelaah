package hotelaah

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

type RedisStandardOP interface {
	Commit() error
}

// TODO
type HotelaahRedisOp struct {
}

type Redisor struct {
	addr   string
	passwd string
	db     int

	rdb *redis.Client
}

// func pinyinScore(s string) int {
//
// }

func NewRedisor(a, p string, d int) *Redisor {
	return &Redisor{
		addr:   a,
		passwd: p,
		db:     d,
	}
}

func (r *Redisor) Init() error {
	r.rdb = redis.NewClient(&redis.Options{
		Addr:     r.addr,
		Password: r.passwd,
		DB:       r.db,
	})
	_, err := r.rdb.Ping(ctx).Result()
	return err
}

// set all score to 0
func (r *Redisor) SetCity(c, p string) {
	// r.rdb.Set(ctx, c, p, 0).Err()
	r.rdb.ZAdd(ctx, p, &redis.Z{Score: 0, Member: c}).Err()
}

func (r *Redisor) SetStringPair(sp StringPair) {
	r.SetCity(sp.First, sp.Second)
}

func (r *Redisor) GetCity(c string) (string, error) {
	return r.rdb.Get(ctx, c).Result()
}

// here DumbDo can receive any types of inteface
func (r *Redisor) DumbDo(op interface{}) {

}
