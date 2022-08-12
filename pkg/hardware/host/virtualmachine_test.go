// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package host

import (
	wmihost "github.com/microsoft/wmi/pkg/base/host"
	_ "github.com/microsoft/wmi/pkg/base/session"
	"testing"
)

func TestGetVirtualMachines(t *testing.T) {
	whost := wmihost.NewWmiLocalHost()
	vms, err := GetVirtualMachines(whost)
	if err != nil {
		t.Fatalf("[%+v]", err)
	}
	defer vms.Close()

	t.Logf("Virtual Machines [%+v]\n", vms)
}
