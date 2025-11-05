// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package service

import (
	"log"
	"os"
	"sync"
	"time"

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
	virtualDiskId string,
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
	retryCount := 0

	for {
		result, err1 := method.Execute(inparams, outparams)
		if err1 != nil {
			err = err1
			return
		}

		if !(result.ReturnValue == 4096 || result.ReturnValue == 0) {
			err = errors.Wrapf(errors.Failed, "GetVirtualHardDiskSettingData method failed with [%d]", result.ReturnValue)
			return
		}
		val, ok := result.OutMethodParams["SettingData"]
		if ok && val.Value != nil {
			size, blockSize, lSectorSize, pSectorSize, format, virtualDiskId, err1 = disk.GetVirtualHardDiskSettingDataFromXml(ims.GetWmiHost(), val.Value.(string))
			if err1 != nil {
				err = err1
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
		job, err1 := instance.GetWmiJob(ims.GetWmiHost(), string(constant.Virtualization), val.Value.(string))
		if err1 != nil {
			if job != nil {
				job.Close()
			}
			err = err1
			return
		}

		err1 = job.WaitForJobCompletion(result.ReturnValue, -1)
		if err1 == nil || !errors.IsVhdSystemInUse(err1) {
			if job != nil {
				job.Close()
			}
			err = err1
			return
		}

		// Retry for system in use
		if retryCount < constant.WmiMethodMaxRetries {
			if job != nil {
				job.Close()
			}
			retryCount++
			backoffDuration := time.Duration(retryCount) * constant.WmiMethodRetryDelay
			log.Printf("[WMI] Method [%s] failed with error system is in use. Retrying (%d/%d) after %v...", method.Name, retryCount, constant.WmiMethodMaxRetries, backoffDuration)
			time.Sleep(backoffDuration)
			continue
		}

		// Retries exhausted -> return an error
		if job != nil {
			job.Close()
		}
		err = errors.Wrapf(errors.Failed, "GetVirtualHardDiskSettingData: retries exhausted (%d) due to system in use", retryCount)
		return
	}

	return
}
