package wlog

import "github.com/rs/zerolog"

// NewNopLogger returns a Logger where all operations are no-op.
func NewNopLogger() Logger {
	return BasicLogger{zlog: zerolog.Nop()}
}
