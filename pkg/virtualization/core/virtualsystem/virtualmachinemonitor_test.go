// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package virtualsystem

import (
	"fmt"
	"testing"

	_ "github.com/microsoft/wmi/pkg/base/session"
)

type testContext struct {
}

func onCallback(ctx interface{}, data string) {
	fmt.Println("Modified :" + data)
}

func TestCreateVirtualMachineMonitor(t *testing.T) {
	ctx := &testContext{}
	vmMonior := CreateVirtualMachineMonitor(ctx, onCallback, "ID")
	defer vmMonior.Close()
	vmMonior.AddEntity("test")
	vmMonior.AddEntity("test2")
	for {
	}
}
