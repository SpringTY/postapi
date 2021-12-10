package rpc

import (
	"encoding/json"
	"testing"
)

func TestGetStaticData(t *testing.T) {
	ds, err := GetStaticDealData()
	print(err == nil)
	s, _ := json.Marshal(ds)
	print(string(s))
}
