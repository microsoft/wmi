// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package host

import (
	wmihost "github.com/microsoft/wmi/pkg/base/host"
	_ "github.com/microsoft/wmi/pkg/base/session"
	"testing"
)

func TestGetProcessor(t *testing.T) {
	whost := wmihost.NewWmiLocalHost()
	proc, err := GetTotalProcessor(whost)
	if err != nil {
		t.Fatalf("[%+v]", err)
	}

	t.Logf("Total Processor [%+v]\n", proc)
}
