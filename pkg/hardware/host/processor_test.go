// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package host

import (
	"testing"

	wmihost "github.com/microsoft/wmi/pkg/base/host"
	_ "github.com/microsoft/wmi/pkg/base/session"
)

func TestGetProcessor(t *testing.T) {
	whost := wmihost.NewWmiLocalHost()
	proc, err := GetTotalProcessor(whost)
	if err != nil {
		t.Fatalf("[%+v]", err)
	}

	t.Logf("Total Processor [%+v]\n", proc)
}

func TestGetProcessorInfo(t *testing.T) {
	whost := wmihost.NewWmiLocalHost()
	process, err := GetProcessor(whost)
	if err != nil {
		t.Fatalf("[%+v]", err)
	}

	t.Logf("ProcessorInfo [%+v]\n", process)
}

func TestGetComputerSystem(t *testing.T) {
	whost := wmihost.NewWmiLocalHost()
	process, err := GetComputerSystem(whost)
	if err != nil {
		t.Fatalf("[%+v]", err)
	}

	t.Logf("ProcessorInfo [%+v]\n", process)
}
