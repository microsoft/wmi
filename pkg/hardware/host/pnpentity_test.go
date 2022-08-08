// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package host

import (
	wmihost "github.com/microsoft/wmi/pkg/base/host"
	_ "github.com/microsoft/wmi/pkg/base/session"
	"testing"
)

func TestGetPnpEntityByName(t *testing.T) {
	whost := wmihost.NewWmiLocalHost()
	pnpEntity, err := GetPnpEntityByName(whost, "System Interrupt Controller")
	if err != nil {
		t.Fatalf("[%+v]", err)
	}
	defer pnpEntity.Close()

	t.Logf("PnP Entity [%+v]\n", pnpEntity)
}
