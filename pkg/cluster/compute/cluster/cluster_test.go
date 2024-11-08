// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package cluster

import (
	"github.com/microsoft/wmi/pkg/base/host"
	_ "github.com/microsoft/wmi/pkg/base/session"
	"testing"
)

var (
	whost *host.WmiHost
)

func init() {
	whost = host.NewWmiLocalHost()
}

func TestGetCluster(t *testing.T) {
	cn, err := GetCluster(whost)
	if err != nil {
		t.Fatal("Failed " + err.Error())
		return
	}
	defer cn.Close()
}

func TestGetCluster_IsHealthy(t *testing.T) {
	cn, err := GetCluster(whost)
	if err != nil {
		t.Fatal("Failed " + err.Error())
		return
	}
	defer cn.Close()

	cn.IsHealthy()
}
