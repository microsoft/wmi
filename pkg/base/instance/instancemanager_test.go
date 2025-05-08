// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package instance

import (
	"strconv"
	"sync"
	"testing"
)

func TestGetWmiInstanceManagerConcurrently(t *testing.T) {
	n := 2
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(name string) {
			defer wg.Done()
			GetWmiInstanceManager(name, "", "", "", "")
		}(strconv.Itoa(n))
	}
	wg.Wait()
}
