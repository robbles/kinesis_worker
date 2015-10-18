package kinesis_worker

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStreamWorkerInitialize(t *testing.T) {
	stream := StreamWorker{}
	err := stream.initialize()

	assert.Nil(t, err, "Initialization of stream should succeed")
	assert.Equal(t, DefaultRegion, stream.Region)
	assert.Equal(t, "", stream.StreamName)
	assert.Equal(t, ShardIteratorTypeLatest, stream.IteratorType)
	assert.Nil(t, stream.StartingSequenceNumber, "StartingSequenceNumber should be a nil pointer if not explicitly set")
	assert.Equal(t, DefaultSleepTime, stream.SleepTime)
	assert.Equal(t, int64(DefaultBatchSize), stream.BatchSize)
	assert.NotNil(t, stream.Client)
	assert.NotNil(t, stream.Output)
	assert.Nil(t, stream.workers)
}
