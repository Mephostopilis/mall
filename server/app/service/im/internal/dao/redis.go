package dao

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"edu/service/im/internal/model"

	"github.com/go-redis/redis/v8"
	"github.com/zhenjl/cityhash"
)

const (
	_prefixMidServer    = "mid_%d" // mid -> key:server
	_prefixKeyServer    = "key_%s" // key -> server
	_prefixServerOnline = "ol_%s"  // server -> online
)

func keyMidServer(mid int64) string {
	return fmt.Sprintf(_prefixMidServer, mid)
}

func keyKeyServer(key string) string {
	return fmt.Sprintf(_prefixKeyServer, key)
}

func keyServerOnline(key string) string {
	return fmt.Sprintf(_prefixServerOnline, key)
}

// pingRedis check redis connection.
func (d *dao) pingRedis(c context.Context) (err error) {
	err = d.redis.Ping(c).Err()
	return
}

// AddMapping add a mapping.
// Mapping:
//	mid -> key_server
//	key -> server
func (d *dao) AddMapping(c context.Context, mid int64, key, server string) (err error) {
	p := d.redis.Pipeline()
	var n = 2
	if mid > 0 {
		p.HSet(c, keyMidServer(mid), key, server)
		p.Expire(c, keyMidServer(mid), d.redisExpire)
		n += 2
	}
	p.Set(c, keyKeyServer(key), server, time.Duration(d.redisExpire))
	p.Expire(c, keyKeyServer(key), time.Duration(d.redisExpire))

	rplys, err := p.Exec(c)
	if err != nil {
		d.log.Errorf("conn.Flush() error(%v)", err)
		return
	}
	for i := 0; i < n; i++ {
		if err = rplys[i].Err(); err != nil {
			d.log.Errorf("conn.Receive() error(%v)", err)
			return
		}
	}
	return
}

// ExpireMapping expire a mapping.
func (d *dao) ExpireMapping(c context.Context, mid int64, key string) (has bool, err error) {

	p := d.redis.Pipeline()
	var n = 1
	if mid > 0 {
		p.Expire(c, keyMidServer(mid), d.redisExpire)
		n++
	}
	p.Expire(c, keyKeyServer(key), d.redisExpire)
	replys, err := p.Exec(c)
	if err != nil {
		return
	}
	for i := 0; i < n; i++ {
		if err = replys[i].Err(); err != nil {
			d.log.Errorf("conn.Receive() error(%v)", err)
			return
		}
	}
	return
}

// DelMapping del a mapping.
func (d *dao) DelMapping(c context.Context, mid int64, key, server string) (has bool, err error) {
	p := d.redis.Pipeline()
	n := 1
	if mid > 0 {
		p.HDel(c, keyMidServer(mid), key)
		n++
	}
	p.Del(c, keyKeyServer(key))
	repls, err := p.Exec(c)
	if err != nil {
		return
	}
	for i := 0; i < n; i++ {
		if err = repls[i].Err(); err != nil {
			d.log.Errorf("conn.Receive() error(%v)", err)
			return
		}
	}
	return
}

// ServersByKeys get a server by key.
func (d *dao) ServersByKeys(c context.Context, keys []string) (res []string, err error) {
	var args []string
	for _, key := range keys {
		args = append(args, keyKeyServer(key))
	}
	sli := d.redis.MGet(c, args...)
	if sli.Err() != nil {
		return
	}
	res = make([]string, 0)
	for i := 0; i < len(sli.Val()); i++ {
		res = append(res, sli.Val()[i].(string))
	}
	return
}

// KeysByMids get a key server by mid.
func (d *dao) KeysByMids(c context.Context, mids []int64) (ress map[string]string, olMids []int64, err error) {
	ress = make(map[string]string)
	cmds := make(map[int64]*redis.StringStringMapCmd, 0)
	p := d.redis.Pipeline()
	for _, mid := range mids {
		cmd := p.HGetAll(c, keyMidServer(mid))
		cmds[mid] = cmd
	}
	_, err = p.Exec(c)
	if err != nil {
		d.log.Errorf("conn.Receive() error(%v)", err)
		return
	}
	for idx := 0; idx < len(mids); idx++ {
		var (
			res map[string]string
		)
		cmd := cmds[mids[idx]]
		res, err = cmd.Result()
		if err != nil {
			return
		}
		if len(res) > 0 {
			olMids = append(olMids, mids[idx])
		}
		for k, v := range res {
			ress[k] = v
		}
	}
	return
}

// AddServerOnline add a server online.
func (d *dao) AddServerOnline(c context.Context, server string, online *model.Online) (err error) {
	roomsMap := map[uint32]map[string]int32{}
	for room, count := range online.RoomCount {
		rMap := roomsMap[cityhash.CityHash32([]byte(room), uint32(len(room)))%64]
		if rMap == nil {
			rMap = make(map[string]int32)
			roomsMap[cityhash.CityHash32([]byte(room), uint32(len(room)))%64] = rMap
		}
		rMap[room] = count
	}
	key := keyServerOnline(server)
	for hashKey, value := range roomsMap {
		err = d.addServerOnline(c, key, strconv.FormatInt(int64(hashKey), 10), &model.Online{RoomCount: value, Server: online.Server, Updated: online.Updated})
		if err != nil {
			return
		}
	}
	return
}

func (d *dao) addServerOnline(c context.Context, key string, hashKey string, online *model.Online) (err error) {
	b, _ := json.Marshal(online)

	p := d.redis.Pipeline()
	p.HSet(c, key, hashKey, b)
	p.Expire(c, key, d.redisExpire)
	cmds, err := p.Exec(c)
	if err != nil {
		return
	}
	for i := 0; i < 2; i++ {
		cmd := cmds[i]
		if err = cmd.Err(); err != nil {
			d.log.Errorf("conn.Receive() error(%v)", err)
			return
		}
	}
	return
}

// ServerOnline get a server online.
func (d *dao) ServerOnline(c context.Context, server string) (online *model.Online, err error) {
	online = &model.Online{RoomCount: map[string]int32{}}
	key := keyServerOnline(server)
	for i := 0; i < 64; i++ {
		ol, err := d.serverOnline(c, key, strconv.FormatInt(int64(i), 10))
		if err == nil && ol != nil {
			online.Server = ol.Server
			if ol.Updated > online.Updated {
				online.Updated = ol.Updated
			}
			for room, count := range ol.RoomCount {
				online.RoomCount[room] = count
			}
		}
	}
	return
}

func (d *dao) serverOnline(c context.Context, key string, hashKey string) (online *model.Online, err error) {
	cmd := d.redis.HGet(c, key, hashKey)
	if cmd.Err() != nil {
		if cmd.Err() != redis.Nil {
			d.log.Errorf("conn.Do(HGET %s %s) error(%v)", key, hashKey, err)
		}
		return
	}
	b, err := cmd.Bytes()
	online = new(model.Online)
	if err = json.Unmarshal(b, online); err != nil {
		d.log.Errorf("serverOnline json.Unmarshal(%s) error(%v)", b, err)
		return
	}
	return
}

// DelServerOnline del a server online.
func (d *dao) DelServerOnline(c context.Context, server string) (err error) {
	key := keyServerOnline(server)
	cmd := d.redis.Del(c, key)
	if err = cmd.Err(); err != nil {
		d.log.Errorf("conn.Do(DEL %s) error(%v)", key, err)
	}
	return
}
