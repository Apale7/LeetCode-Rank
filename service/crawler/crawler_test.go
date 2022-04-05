package crawler

import (
	"fmt"
	"sync"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestGetUserAcInfo(t *testing.T) {
	GetUserPublicProfile("33qwg5mua3")
}

func TestGetUserQuestionProgress(t *testing.T) {
	wg := sync.WaitGroup{}
	logrus.SetLevel(logrus.DebugLevel)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			ac := GetUserQuestionProgress("apale")
			fmt.Printf("%+v\n", ac)
			wg.Done()
		}()
	}
	wg.Wait()
}
