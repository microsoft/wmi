// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package service

import (
	"testing"

	"github.com/microsoft/wmi/pkg/base/host"
	_ "github.com/microsoft/wmi/pkg/base/session"
)

var (
	whost *host.WmiHost
)

func init() {
	whost = host.NewWmiLocalHost()
}

func TestGetClusterService(t *testing.T) {
	_, err := GetLocalClusterService(whost)
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
}

func TestGetClusterServices(t *testing.T) {
	_, err := GetClusterServices(whost)
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
}
