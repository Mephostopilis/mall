package biz

import (
	"edu/comet/internal/conf"
	"edu/pkg/bytes"
	"edu/pkg/time"
)

// RoundOptions round options.
type RoundOptions struct {
	Timer        int
	TimerSize    int
	Reader       int
	ReadBuf      int
	ReadBufSize  int
	Writer       int
	WriteBuf     int
	WriteBufSize int
}

// Round userd for connection round-robin get a reader/writer/timer for split big lock.
type Round struct {
	readers []bytes.Pool
	writers []bytes.Pool
	timers  []time.Timer
	options RoundOptions
}

// NewRound new a round struct.
func NewRound(c *conf.Server) (r *Round) {
	var i int
	r = &Round{
		options: RoundOptions{
			Reader:       int(c.Tcp.Reader),
			ReadBuf:      int(c.Tcp.ReadBuf),
			ReadBufSize:  int(c.Tcp.ReadBufSize),
			Writer:       int(c.Tcp.Writer),
			WriteBuf:     int(c.Tcp.WriteBuf),
			WriteBufSize: int(c.Tcp.WriteBufSize),
			Timer:        int(c.Protocol.Timer),
			TimerSize:    int(c.Protocol.TimerSize),
		}}
	// reader
	r.readers = make([]bytes.Pool, r.options.Reader)
	for i = 0; i < r.options.Reader; i++ {
		r.readers[i].Init(r.options.ReadBuf, r.options.ReadBufSize)
	}
	// writer
	r.writers = make([]bytes.Pool, r.options.Writer)
	for i = 0; i < r.options.Writer; i++ {
		r.writers[i].Init(r.options.WriteBuf, r.options.WriteBufSize)
	}
	// timer
	r.timers = make([]time.Timer, r.options.Timer)
	for i = 0; i < r.options.Timer; i++ {
		r.timers[i].Init(r.options.TimerSize)
	}
	return
}

// Timer get a timer.
func (r *Round) Timer(rn int) *time.Timer {
	return &(r.timers[rn%r.options.Timer])
}

// Reader get a reader memory buffer.
func (r *Round) Reader(rn int) *bytes.Pool {
	return &(r.readers[rn%r.options.Reader])
}

// Writer get a writer memory buffer pool.
func (r *Round) Writer(rn int) *bytes.Pool {
	return &(r.writers[rn%r.options.Writer])
}
