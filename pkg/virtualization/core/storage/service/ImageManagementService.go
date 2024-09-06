// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package service

import (
	"os"
	"sync"

	"github.com/microsoft/wmi/pkg/base/host"
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/constant"
	"github.com/microsoft/wmi/pkg/errors"
	"github.com/microsoft/wmi/pkg/virtualization/core/storage/disk"
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
	v2 "github.com/microsoft/wmi/server2019/root/virtualization/v2"
)

var (
	serviceStoreMap map[string]*ImageManagementService
	mux             sync.Mutex
)

func init() {
	serviceStoreMap = map[string]*ImageManagementService{}
}

type ImageManagementService struct {
	*v2.Msvm_ImageManagementService
}

// GetVirtualEthernetSwitchManagementService gets the VMMS Switch Management Service
func GetImageManagementService(whost *host.WmiHost) (mgmt *ImageManagementService, err error) {
	if val, ok := serviceStoreMap[whost.HostName]; ok {
		mgmt = val
		return
	}

	mgmt, err = getService(whost)
	if err != nil {
		return
	}

	mux.Lock()
	defer mux.Unlock()
	if val, ok := serviceStoreMap[whost.HostName]; ok {
		mgmt.Close()
		mgmt = val
		return
	}

	serviceStoreMap[whost.HostName] = mgmt
	return
}

func getService(whost *host.WmiHost) (mgmt *ImageManagementService, err error) {
	creds := whost.GetCredential()
	query := query.NewWmiQuery("Msvm_ImageManagementService")
	// TODO: Regenerate wrappers that would take WmiHost directly
	imswmi, err := v2.NewMsvm_ImageManagementServiceEx6(whost.HostName, string(constant.Virtualization), creds.UserName, creds.Password, creds.Domain, query)
	if err != nil {
		return
	}

	mgmt = &ImageManagementService{imswmi}
	return
}

// ResizeDisk
func (ims *ImageManagementService) ResizeDisk(path string, size uint64) (err error) {
	method, err := ims.GetWmiMethod("ResizeVirtualHardDisk")
	if err != nil {
		return
	}
	defer method.Close()

	inparams := wmi.WmiMethodParamCollection{}
	inparams = append(inparams, wmi.NewWmiMethodParam("Path", path))
	inparams = append(inparams, wmi.NewWmiMethodParam("MaxInternalSize", size))

	outparams := wmi.WmiMethodParamCollection{wmi.NewWmiMethodParam("Job", nil)}

	result, err := method.Execute(inparams, outparams)
	if err != nil {
		return
	}

	if result.ReturnValue == 0 {
		return
	}

	if result.ReturnValue != 4096 {
		err = errors.Wrapf(errors.Failed, "Method failed with [%d]", result.ReturnValue)
		return
	}

	val, ok := result.OutMethodParams["Job"]
	if !ok || val.Value == nil {
		err = errors.Wrapf(errors.NotFound, "Job")
		return
	}
	job, err := instance.GetWmiJob(ims.GetWmiHost(), string(constant.Virtualization), val.Value.(string))
	if err != nil {
		return
	}
	defer job.Close()
	return job.WaitForJobCompletion(result.ReturnValue, -1)
}

// CreateDiskEx
func (ims *ImageManagementService) CreateDiskEx(path string,
	logicalSectorSize, physicalSectorSize, blockSize uint32,
	diskSize uint64, dynamic bool, diskFileFormat disk.VirtualHardDiskFormat) (err error) {

	setting, err := disk.GetVirtualHardDiskSettingData(ims.GetWmiHost(), path, logicalSectorSize,
		physicalSectorSize, blockSize, diskSize, dynamic, diskFileFormat)
	if err != nil {
		return
	}
	defer setting.Close()

	defer func() {
		if err == nil {
			return
		}
		// Try to Cleanup.
		// FIXME : This might fail if the VHD is still getting generated and leak
		os.Remove(path)
	}()

	err = ims.CreateDisk(setting)
	return
}

// CreateDisk
func (ims *ImageManagementService) CreateDisk(settings *disk.VirtualHardDiskSettingData) (err error) {
	method, err := ims.GetWmiMethod("CreateVirtualHardDisk")
	if err != nil {
		return
	}
	defer method.Close()
	embeddedInstance, err := settings.EmbeddedXMLInstance()
	if err != nil {
		return err
	}
	inparams := wmi.WmiMethodParamCollection{wmi.NewWmiMethodParam("VirtualDiskSettingData", embeddedInstance)}
	outparams := wmi.WmiMethodParamCollection{wmi.NewWmiMethodParam("Job", nil)}

	result, err := method.Execute(inparams, outparams)
	if err != nil {
		return
	}

	if result.ReturnValue == 0 {
		return
	}

	if result.ReturnValue != 4096 {
		err = errors.Wrapf(errors.Failed, "Method failed with [%d]", result.ReturnValue)
		return
	}

	val, ok := result.OutMethodParams["Job"]
	if !ok || val.Value == nil {
		err = errors.Wrapf(errors.NotFound, "Job")
		return
	}
	job, err := instance.GetWmiJob(ims.GetWmiHost(), string(constant.Virtualization), val.Value.(string))
	if err != nil {
		return
	}
	defer job.Close()
	return job.WaitForJobCompletion(result.ReturnValue, -1)
}

func (ims *ImageManagementService) GetVirtualHardDiskConfig(path string) (size uint64,
	blockSize uint32,
	lSectorSize uint32,
	pSectorSize uint32,
	format uint16,
	err error) {
	method, err := ims.GetWmiMethod("GetVirtualHardDiskSettingData")
	if err != nil {
		return
	}
	defer method.Close()

	inparams := wmi.WmiMethodParamCollection{}
	inparams = append(inparams, wmi.NewWmiMethodParam("Path", path))

	outparams := wmi.WmiMethodParamCollection{}
	outparams = append(outparams, wmi.NewWmiMethodParam("SettingData", nil))
	outparams = append(outparams, wmi.NewWmiMethodParam("Job", nil))

	result, err := method.Execute(inparams, outparams)
	if err != nil {
		return
	}

	if !(result.ReturnValue == 4096 || result.ReturnValue == 0) {
		err = errors.Wrapf(errors.Failed, "GetVirtualHardDiskSettingData method failed with [%d]", result.ReturnValue)
		return
	}
	val, ok := result.OutMethodParams["SettingData"]
	if ok && val.Value != nil {
		size, blockSize, lSectorSize, pSectorSize, format, err = disk.GetVirtualHardDiskSettingDataFromXml(ims.GetWmiHost(), val.Value.(string))
		if err != nil {
			return
		}
	}

	if result.ReturnValue == 0 {
		return
	}

	val, ok = result.OutMethodParams["Job"]
	if !ok || val.Value == nil {
		err = errors.Wrapf(errors.NotFound, "Job")
		return
	}
	job, err := instance.GetWmiJob(ims.GetWmiHost(), string(constant.Virtualization), val.Value.(string))
	if err != nil {
		return
	}
	defer job.Close()
	err = job.WaitForJobCompletion(result.ReturnValue, -1)
	return
}
