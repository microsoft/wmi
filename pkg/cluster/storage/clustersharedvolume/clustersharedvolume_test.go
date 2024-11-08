// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package clustersharedvolume

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

func TestGetClusterSharedVolume(t *testing.T) {
	cn, err := GetClusterSharedVolume(whost, "C:\\\\ClusterStorage\\\\Volume1")
	if err != nil {
		t.Fatal("Failed " + err.Error())
		return
	}
	defer cn.Close()
}

func TestGetClusterSharedVolumes(t *testing.T) {
	nc, err := GetClusterSharedVolumes(whost)
	if err != nil {
		t.Fatal("Failed " + err.Error())
		return
	}
	defer nc.Close()
	t.Logf("Volumes returned %d\n", len(nc))

	for _, volume := range nc {
		volumeName, err := volume.GetPropertyName()
		if err != nil {
			t.Fatal("Failed " + err.Error())
			return
		}
		t.Logf("Volume Name: %s\n", volumeName)
	}
}
