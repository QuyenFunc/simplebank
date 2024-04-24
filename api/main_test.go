package api

import (
	db "github.com/Quyen-2211/simplebank/db/sqlc"
	"github.com/Quyen-2211/simplebank/db/util"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
	"time"
)

func newTestServer(t *testing.T, store db.Store) *Server {
	config := util.Config{
		TokenSymmetricKey:   util.RandomString(32),
		AccessTokenDuration: time.Minute,
	}

	server, err := NewServer(config, store)
	require.NoError(t, err)

	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode) //su dung de du git vao Mode test, cho nhat kys trong gon hon
	os.Exit(m.Run())
}
