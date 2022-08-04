// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package host

import (
	wmihost "github.com/microsoft/wmi/pkg/base/host"
	_ "github.com/microsoft/wmi/pkg/base/session"
	"testing"
)

func TestGetPnpEntityCollection(t *testing.T) {
	whost := wmihost.NewWmiLocalHost()
	pnpEntity, err := GetPnpEntityCollection(whost)
	if err != nil {
		t.Fatalf("[%+v]", err)
	}

	t.Logf("PnP Entity Collection [%+v]\n", pnpEntity)
}

func TestGetPnpEntityCollectionByName(t *testing.T) {
	whost := wmihost.NewWmiLocalHost()
	pnpEntity, err := GetPnpEntityCollectionByName(whost, "System Interrupt Controller")
	if err != nil {
		t.Fatalf("[%+v]", err)
	}

	t.Logf("PnP Entity Collection [%+v]\n", pnpEntity)
}
