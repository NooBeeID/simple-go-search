package db

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConnectionGorm_Success(t *testing.T) {
	db := NewGormDB().GetConnection()
	require.NotNil(t, db.GormDB)
}
