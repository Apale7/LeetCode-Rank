package proxy

import "testing"

func TestGetProxy(t *testing.T) {
	for i := 0; i < 10; i++ {
		got, err := GetProxy()
		if err != nil {
			t.Errorf("GetProxy error: %v", err)
			continue
		}
		t.Logf("got: %s", got)
	}
}
