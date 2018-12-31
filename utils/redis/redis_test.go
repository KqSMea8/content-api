package redis_test

import (
	"code.byted.org/learning_fe/pathfinder-api/utils/redis"
	"testing"
)

// go test -run="NoopsGetContentListWithFilter"
func TestInitRedis(t *testing.T) {
	err := redis.Init(1)
	if err != nil {
		t.Fatal(err)
	}
	// t.Log(resp)
	// realResp := resp.RealResponse().(*noops.GetContentListWithFilterResponse)
	// t.Log(realResp.ContentId)
}
