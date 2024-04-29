// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package gpu

import (
	wmihost "github.com/microsoft/wmi/pkg/base/host"
	_ "github.com/microsoft/wmi/pkg/base/session"
	"testing"
)

func TestGetPartitionableGpuCollection(t *testing.T) {
	whost := wmihost.NewWmiLocalHost()
	gpuCollection, err := GetPartitionableGpuCollection(whost)
	if err != nil {
		t.Fatalf("[%+v]", err)
	}

	t.Logf("Partitionable GPUs [%+v]\n", gpuCollection)
}
