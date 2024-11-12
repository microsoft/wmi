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
	volumeName, err := cn.GetPropertyName()
	if err != nil {
		t.Fatal("Failed " + err.Error())
		return
	}

	t.Logf("Volume Name: %s\n", volumeName)
}

func TestGetClusterSharedVolumeByName(t *testing.T) {
	cn, err := GetClusterSharedVolumebyName(whost, "C:/ClusterStorage/Volume1")
	if err != nil {
		t.Fatal("Failed " + err.Error())
		return
	}
	defer cn.Close()
	volumeName, err := cn.GetPropertyName()
	if err != nil {
		t.Fatal("Failed " + err.Error())
		return
	}

	t.Logf("Volume Name: %s\n", volumeName)
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

		status, err := volume.GetPropertyStatus()
		if err != nil {
			t.Fatal("Failed " + err.Error())
			return
		}
		t.Logf("Volume Status %v \n", status)

		statusokay, err := volume.IsStatusOK()
		if err != nil {
			t.Fatal("Failed " + err.Error())
			return
		}
		t.Logf("Volume StatusOK %v \n", statusokay)
		statusfault, err := volume.IsFaultStateOK()
		if err != nil {
			t.Fatal("Failed " + err.Error())
			return
		}
		t.Logf("Volume FaultStatusOK %v \n", statusfault)

		ownerGroup, err := volume.OwnerGroup()
		if err != nil {
			t.Fatal("Failed " + err.Error())
			return
		}
		t.Logf("Volume OwnerGroup %v \n", ownerGroup)
	}
}
