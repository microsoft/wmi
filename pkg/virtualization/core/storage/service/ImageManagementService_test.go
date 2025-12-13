// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package service

import (
	"os"
	"strings"
	"testing"

	"github.com/microsoft/wmi/pkg/base/host"
	_ "github.com/microsoft/wmi/pkg/base/session"

	"github.com/microsoft/wmi/pkg/virtualization/core/storage/disk"
)

var (
	whost *host.WmiHost
)

func init() {
	whost = host.NewWmiLocalHost()
}

func TestGetImageManagementService(t *testing.T) {
	_, err := GetImageManagementService(whost)
	if err != nil {
		t.Fatal("Failed " + err.Error())
	}
}

func TestCreateDynamicVirtualHardDisk(t *testing.T) {
	ims, err := GetImageManagementService(whost)
	if err != nil {
		t.Fatal("Failed " + err.Error())
	}
	path := "c:\\test\\tmp.vhdx"
	setting, err := disk.GetVirtualHardDiskSettingData(whost, path, 512, 512, 0, 1024*1024*10, true, disk.VirtualHardDiskFormat_2)
	if err != nil {
		t.Fatal("Failed " + err.Error())
	}
	defer setting.Close()
	err = ims.CreateDisk(setting)
	if err != nil {
		t.Fatal("Create Failed " + err.Error())
	}
	defer os.RemoveAll(path)

	err = ims.ResizeDisk(path, 1024*1024*1024*100)
	if err != nil {
		t.Fatal("Resize Failed " + err.Error())
	}
}

func TestCreateStaticVirtualHardDisk(t *testing.T) {
	ims, err := GetImageManagementService(whost)
	if err != nil {
		t.Fatal("Failed " + err.Error())
	}
	path := "c:\\test\\tmp.vhdx"
	setting, err := disk.GetVirtualHardDiskSettingData(whost, path, 512, 512, 0, 1024*1024*10, false, disk.VirtualHardDiskFormat_2)
	if err != nil {
		t.Fatal("Failed " + err.Error())
	}
	defer setting.Close()
	err = ims.CreateDisk(setting)
	if err != nil {
		t.Fatal("Create Failed " + err.Error())
	}
	defer os.RemoveAll(path)

	err = ims.ResizeDisk(path, 1024*1024*100)
	if err != nil {
		t.Fatal("Resize Failed " + err.Error())
	}
}

func TestGetVirtualHardDiskConfig(t *testing.T) {
	ims, err := GetImageManagementService(whost)
	if err != nil {
		t.Fatal("Failed " + err.Error())
	}
	path := "c:\\test\\tmp.vhdx"
	var disksize uint64 = 1024 * 1024 * 10
	var lsectorSize uint32 = 512
	var psectorSize uint32 = 512
	setting, err := disk.GetVirtualHardDiskSettingData(whost, path, lsectorSize, psectorSize, 0, disksize, true, disk.VirtualHardDiskFormat_2)
	if err != nil {
		t.Fatal("Failed " + err.Error())
	}
	defer setting.Close()
	err = ims.CreateDisk(setting)
	if err != nil {
		t.Fatal("Create Failed " + err.Error())
	}
	defer os.RemoveAll(path)

	readSize, _, readLsectorSize, readPsectorSize, _, readVirtualDiskId, dynamic, err := ims.GetVirtualHardDiskConfig(path)
	if err != nil {
		t.Fatal("Get vhd configuration failed " + err.Error())
	}

	if readSize != disksize {
		t.Fatal("Get vhd configuration size mismatch")
	}

	if readLsectorSize != lsectorSize {
		t.Fatal("Get vhd configuration logical sector mismatch")
	}

	if readPsectorSize != psectorSize {
		t.Fatal("Get vhd configuration physical sector mismatch")
	}

	if readVirtualDiskId == "" {
		t.Fatal("Get vhd configuration virtual disk id is empty")
	}

	if !dynamic {
		t.Fatal("Get vhd configuration dynamic flag mismatch")
	}

	t.Logf("Virtual Disk Id: %s", readVirtualDiskId)
	hostId := strings.ReplaceAll(readVirtualDiskId, "-", "")
	if len(hostId) != 32 {
		t.Fatal("Get vhd configuration virtual disk id length is not equal to 32")
	}
}
