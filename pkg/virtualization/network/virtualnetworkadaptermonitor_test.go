// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package network

import (
	"testing"
)

type testContext struct {
}

func onCallback(ctx interface{}, data string) {
}

func TestCreateVirtualNetworkAdapterMonitor(t *testing.T) {
	ctx := &testContext{}
	vnicMonior := CreateVirtualNetworkAdapterMonitor(ctx, onCallback, "Name")
	defer vnicMonior.Close()
}
