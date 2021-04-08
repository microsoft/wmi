// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package service

import (
	"os"

	"github.com/microsoft/wmi/pkg/base/host"
	_ "github.com/microsoft/wmi/pkg/base/session"
	"testing"

	"github.com/microsoft/wmi/pkg/virtualization/core/storage/disk"
)

var (
	whost   *host.WmiHost
	timeout uint16 = 30
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

func TestCreateVirtualHardDisk(t *testing.T) {
	ims, err := GetImageManagementService(whost)
	if err != nil {
		t.Fatal("Failed " + err.Error())
	}
	path := "c:\\test\\tmp.vhdx"
	setting, err := disk.GetVirtualHardDiskSettingData(whost, path, 512, 512, 0, 1024*1024*10, true)
	if err != nil {
		t.Fatal("Failed " + err.Error())
	}
	defer setting.Close()
	err = ims.CreateDisk(setting, timeout)
	if err != nil {
		t.Fatal("Create Failed " + err.Error())
	}
	defer os.RemoveAll(path)

	err = ims.ResizeDisk(path, 1024*1024*1024*100, timeout)
	if err != nil {
		t.Fatal("Resize Failed " + err.Error())
	}
}
