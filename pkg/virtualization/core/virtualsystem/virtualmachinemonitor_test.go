// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package virtualsystem

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

func TestCreateVirtualMachineMonitor(t *testing.T) {
	return
	ctx := &testContext{}
	vnicMonior := CreateVirtualMachineMonitor(ctx, onCallback, "ID")
	defer vnicMonior.Close()
	vnicMonior.AddEntity("test")
	vnicMonior.AddEntity("test2")
	for {
	}
}
