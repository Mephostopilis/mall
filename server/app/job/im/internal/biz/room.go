package biz

import (
	"errors"
	"time"

	comet "edu/api/comet/grpc"
	"edu/job/im/internal/conf"
	"edu/pkg/bytes"

	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrComet commet error.
	ErrComet = errors.New("comet rpc is not available")
	// ErrCometFull comet chan full.
	ErrCometFull = errors.New("comet proto chan full")
	// ErrRoomFull room chan full.
	ErrRoomFull = errors.New("room proto chan full")

	roomReadyProto = new(comet.Proto)
)

// Room room.
type Room struct {
	c     *conf.Server_Room
	job   *Job
	id    string
	proto chan *comet.Proto
	log   *log.Helper
}

// NewRoom new a room struct, store channel room info.
func NewRoom(job *Job, id string, c *conf.Server_Room) (r *Room) {
	r = &Room{
		c:     c,
		id:    id,
		job:   job,
		proto: make(chan *comet.Proto, c.Batch*2),
	}
	go r.pushproc(int(c.Batch), c.Signal.AsDuration())
	return
}

// Push push msg to the room, if chan full discard it.
func (r *Room) Push(op int32, msg []byte) (err error) {
	var p = &comet.Proto{
		Ver:  1,
		Op:   op,
		Body: msg,
	}
	select {
	case r.proto <- p:
	default:
		err = ErrRoomFull
	}
	return
}

// pushproc merge proto and push msgs in batch.
func (r *Room) pushproc(batch int, sigTime time.Duration) {
	var (
		n    int
		last time.Time
		p    *comet.Proto
		buf  = bytes.NewWriterSize(int(comet.MaxBodySize))
	)
	r.log.Infof("start room:%s goroutine", r.id)
	td := time.AfterFunc(sigTime, func() {
		select {
		case r.proto <- roomReadyProto:
		default:
		}
	})
	defer td.Stop()
	for {
		if p = <-r.proto; p == nil {
			break // exit
		} else if p != roomReadyProto {
			// merge buffer ignore error, always nil
			p.WriteTo(buf)
			if n++; n == 1 {
				last = time.Now()
				td.Reset(sigTime)
				continue
			} else if n < batch {
				if sigTime > time.Since(last) {
					continue
				}
			}
		} else {
			if n == 0 {
				break
			}
		}
		_ = r.job.broadcastRoomRawBytes(r.id, buf.Buffer())
		// TODO use reset buffer
		// after push to room channel, renew a buffer, let old buffer gc
		buf = bytes.NewWriterSize(buf.Size())
		n = 0
		if r.c.Idle.AsDuration() != 0 {
			td.Reset(time.Duration(r.c.Idle.AsDuration()))
		} else {
			td.Reset(time.Minute)
		}
	}
	r.job.delRoom(r.id)
	r.log.Infof("room:%s goroutine exit", r.id)
}

func (j *Job) delRoom(roomID string) {
	j.roomsMutex.Lock()
	delete(j.rooms, roomID)
	j.roomsMutex.Unlock()
}

func (j *Job) getRoom(roomID string) *Room {
	j.roomsMutex.RLock()
	room, ok := j.rooms[roomID]
	j.roomsMutex.RUnlock()
	if !ok {
		j.roomsMutex.Lock()
		if room, ok = j.rooms[roomID]; !ok {
			room = NewRoom(j, roomID, j.c.Room)
			j.rooms[roomID] = room
		}
		j.roomsMutex.Unlock()
		j.log.Infof("new a room:%s active:%d", roomID, len(j.rooms))
	}
	return room
}
