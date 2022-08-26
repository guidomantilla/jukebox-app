package cachemanager

import (
	"fmt"
)

func generateKey(cacheName string, key any) string {
	return cacheName + "-" + fmt.Sprintf("%v", key)
}
