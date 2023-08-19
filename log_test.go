package logging

import (
	"testing"
)

/**
   @author yhy
   @since 2023/8/19
   @desc //TODO
**/

func TestLog(t *testing.T) {
	New(true, "", "Jie", false)
	Logger.Debug("fsafs")
	Logger.Infoln("===")
}
