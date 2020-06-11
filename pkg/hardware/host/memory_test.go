// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package host

import (
	wmihost "github.com/microsoft/wmi/pkg/base/host"
	_ "github.com/microsoft/wmi/pkg/base/session"
	"testing"
)

func TestGetTotalMemory(t *testing.T) {
	whost := wmihost.NewWmiLocalHost()
	mem, err := GetTotalPhysicalMemory(whost)
	if err != nil {
		t.Fatalf("[%+v]", err)
	}

	t.Logf("Total Memory [%+v]\n", mem)
}

func TestGetFreeMemory(t *testing.T) {
	whost := wmihost.NewWmiLocalHost()
	mem, err := GetFreePhysicalMemory(whost)
	if err != nil {
		t.Fatalf("[%+v]", err)
	}

	t.Logf("Free Memory [%+v]\n", mem)
}
