package paulog_test

import (
	"github.com/paust-team/paulog"
	"testing"
)

func TestLog(t *testing.T) {
	defer paulog.ClearLogLevels()
	logger := paulog.GetLogger("pau.log.a")
	logger.Debugf("debugf")
	logger.Infof("infof")
}

func TestPackagePrefix(t *testing.T) {
	defer paulog.ClearLogLevels()
	paulog.SetLevel("pau", paulog.ERROR)
	logger := paulog.GetLogger("pau.log.a")
	logger.Debugf("debugf")
	logger.Infof("infof")
	logger.Errorf("errorf")

	paulog.SetLevel("pau.log", paulog.INFO)
	logger.Debugf("debugf")
	logger.Infof("infof")
	logger.Errorf("errorf")
}
