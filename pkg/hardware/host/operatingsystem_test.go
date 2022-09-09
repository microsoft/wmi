// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package host

import (
	"testing"

	wmihost "github.com/microsoft/wmi/pkg/base/host"
	_ "github.com/microsoft/wmi/pkg/base/session"
)

func TestGetOperatingSystemInfo(t *testing.T) {
	whost := wmihost.NewWmiLocalHost()
	osInfo, err := GetOperatingSystem(whost)
	if err != nil {
		t.Fatalf("[%+v]", err)
	}

	t.Logf("OperatingSystem Info [%+v]\n", osInfo)
}
