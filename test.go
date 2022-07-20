package morpc

import (
	"context"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/fagongzi/goetty/v2"
	"github.com/matrixorigin/matrixone/pkg/logutil"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestCreateServerWithOptions(t *testing.T) {
	testRPCServer(t, func(rs *server) {
		assert.Equal(t, 100, rs.options.batchSendSize)
		assert.Equal(t, 200, rs.options.bufferSize)
	}, WithServerBatchSendSize(100),
		WithServerSessionBufferSize(200))
}
