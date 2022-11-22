package api

import (
	"github.com/brucewangzhihua/sentinel-golang/core/base"
	"github.com/brucewangzhihua/sentinel-golang/logging"
)

// TraceError records the provided error to the given SentinelEntry.
func TraceError(entry *base.SentinelEntry, err error) {
	defer func() {
		if e := recover(); e != nil {
			logging.Panicf("Failed to TraceError, panic error: %+v", e)
			return
		}
	}()
	if entry == nil || err == nil {
		return
	}

	entry.SetError(err)
}
