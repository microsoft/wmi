// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package host

import (
	"testing"

	wmihost "github.com/microsoft/wmi/pkg/base/host"
	_ "github.com/microsoft/wmi/pkg/base/session"
)

func TestGetComputerSystem(t *testing.T) {
	whost := wmihost.NewWmiLocalHost()
	computersysInfo, err := GetComputerSystem(whost)
	if err != nil {
		t.Fatalf("[%+v]", err)
	}

	t.Logf("Computer System Info [%+v]\n", computersysInfo)
}
