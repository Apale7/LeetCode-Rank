package dal

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
)

func TestGetAcceptedEarlist(t *testing.T) {
	ac, err := GetAcceptedEarlist(context.TODO(), Username("369hh369"))
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	b, _ := json.Marshal(ac)
	fmt.Printf("ac: %v\n", string(b))
}
