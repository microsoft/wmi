// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package virtualnetworkadapter

import (
	"fmt"
	_ "github.com/microsoft/wmi/pkg/base/session"
	"testing"
)

type testContext struct {
}

func onCallback(ctx interface{}, data string) {
	fmt.Println("Modified :" + data)
}

func TestCreateVirtualNetworkAdapterMonitor(t *testing.T) {
	ctx := &testContext{}
	vnicMonior := CreateVirtualNetworkAdapterMonitor(ctx, onCallback)
	defer vnicMonior.Close()

	vnicMonior.AddEntityToMonitorForGuestNetworkChanges("Network Adapter Guest",
		"DB2591A8-01CD-41A9-9586-85714961324F")
	vnicMonior.AddEntityToMonitorAdapterChanges("Network Adapter",
		"DB2591A8-01CD-41A9-9586-85714961324F")
	vnicMonior.AddEntityToMonitorPortChanges("Network Adapter Port",
		"DB2591A8-01CD-41A9-9586-85714961324F")
	for {
	}
}
