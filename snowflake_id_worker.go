package golang_snowflake_id_generator

import (
	"sync"
	"time"
)

const (
	epoch          = 1288834974657
	workerIDBits   = 10
	sequenceBits   = 12
	workerIDShift  = sequenceBits
	timestampShift = sequenceBits + workerIDBits
	sequenceMask   = -1 ^ (-1 << sequenceBits)
)

// SnowflakeWorker - A logical instance that generates a Snowflake ID
type SnowflakeWorker struct {
	mutex      sync.Mutex
	timestamp  int64
	workerId   int64
	sequenceId int64
}

func NewSnowflakeWorker(workerId int64) *SnowflakeWorker {
	return &SnowflakeWorker{workerId: workerId, timestamp: 0, sequenceId: 0}
}

// nextID - Returns a new unique ID given SnowflakeWorker instance
func (s *SnowflakeWorker) nextID() int64 {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	timestamp := time.Now().UnixNano() / int64(time.Millisecond)

	if s.timestamp == timestamp {
		// 1 millisecond has not passed yet, use next sequenceID
		s.sequenceId = (s.sequenceId + 1) & sequenceMask

		// We used all 4096 sequences, recompute the timestamp
		if s.sequenceId == 0 {
			for timestamp <= s.timestamp {
				timestamp = time.Now().UnixNano() / int64(time.Millisecond)
			}
		}
	} else {
		// Reset sequenceID to 0
		s.sequenceId = 0
	}

	s.timestamp = timestamp

	// Construct Snowflake ID
	timestampBits := (timestamp - epoch) << timestampShift
	workerBits := s.workerId << workerIDShift

	newID := timestampBits | workerBits | s.sequenceId
	return newID
}
