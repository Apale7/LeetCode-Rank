package crawler

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestGetUserAcInfo(t *testing.T) {
	GetUserPublicProfile("33qwg5mua3")
}

func TestGetUserQuestionProgress(t *testing.T) {
	logrus.SetLevel(logrus.DebugLevel)
	GetUserQuestionProgress("apale")
}
