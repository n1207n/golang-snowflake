package golang_snowflake_id_generator

import "sync"

const (
	epoch          = 1288834974657
	workerIDBits   = 10
	sequenceBits   = 12
	workerIDShift  = sequenceBits
	timestampShift = sequenceBits + workerIDBits
	sequenceMask   = -1 ^ (-1 << sequenceBits)
)

type SnowflakeWorker struct {
	mutex      sync.Mutex
	timestamp  int64
	workerId   int64
	sequenceId int64
}

func NewSnowflakeWorker(workerId int64) *SnowflakeWorker {
	return &SnowflakeWorker{workerId: workerId, timestamp: 0, sequenceId: 0}
}
