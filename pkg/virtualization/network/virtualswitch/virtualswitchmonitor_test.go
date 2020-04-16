// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package virtualswitch

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

func TestCreateVirtualSwitchMonitor(t *testing.T) {
	//return
	ctx := &testContext{}
	vnicMonior := CreateVirtualSwitchMonitor(ctx, onCallback)
	defer vnicMonior.Close()
	vnicMonior.AddEntity("test")
	for {
	}
}
