// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package service

import (
	"fmt"
	"os"
	"strings"
	"syscall"
	"testing"

	"github.com/microsoft/wmi/pkg/base/host"
	_ "github.com/microsoft/wmi/pkg/base/session"
	"github.com/microsoft/wmi/pkg/constant"
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

	readSize, _, readLsectorSize, readPsectorSize, _, readVirtualDiskId, err := ims.GetVirtualHardDiskConfig(path)
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

	t.Logf("Virtual Disk Id: %s", readVirtualDiskId)
	hostId := strings.ReplaceAll(readVirtualDiskId, "-", "")
	if len(hostId) != 32 {
		t.Fatal("Get vhd configuration virtual disk id length is not equal to 32")
	}
}

func openExclusiveSyscall(path string) (*os.File, error) {
	p, err := syscall.UTF16PtrFromString(path)
	if err != nil {
		return nil, err
	}
	h, err := syscall.CreateFile(p,
		syscall.GENERIC_READ,
		0, // no sharing
		nil,
		syscall.OPEN_EXISTING,
		syscall.FILE_ATTRIBUTE_NORMAL,
		0)
	if err != nil {
		return nil, err
	}
	return os.NewFile(uintptr(h), path), nil
}

func TestGetVirtualHardDiskConfigWithSystemInUse(t *testing.T) {
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

	// Keep a file handle opened to simulate SystemInUse error
	f, err := openExclusiveSyscall(path)
	if err != nil {
		t.Fatal("Failed to open " + err.Error())
	}
	defer f.Close()

	_, _, _, _, _, _, err = ims.GetVirtualHardDiskConfig(path)

	if err == nil {
		t.Fatal("Get vhd configuration didn't fail when file is in use.")
	}

	expected := fmt.Sprintf("GetVirtualHardDiskSettingData: retries exhausted (%d) due to system in use: Failed", constant.WmiMethodMaxRetries)
	if err.Error() != expected {
		t.Fatalf("Get vhd configuration error does not match. [Expected]: %s [Actual]: %s", expected, err.Error())
	}
}
