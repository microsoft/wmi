// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package controller

import (
	"fmt"
	"log"

	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/virtualization/core/storage/drive"
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
	"github.com/microsoft/wmi/server2019/root/virtualization/v2"
)

type SCSIControllerSettings struct {
	*v2.Msvm_ResourceAllocationSettingData
}

// NewSCSIControllerSettings
func NewSCSIControllerSettings(instance *wmi.WmiInstance) (*SCSIControllerSettings, error) {
	wmivm, err := v2.NewMsvm_ResourceAllocationSettingDataEx1(instance)
	if err != nil {
		return nil, err
	}
	return &SCSIControllerSettings{wmivm}, nil
}

func (settings *SCSIControllerSettings) GetFreeLocation() (int32, error) {
	// Get all drives - Is this
	col, err := settings.getResourceAllocationSettingData(v2.ResourceAllocationSettingData_ResourceType_Disk_Drive)
	if err != nil {
		return -1, err
	}
	defer col.Close()
	dvdcol, err := settings.getResourceAllocationSettingData(v2.ResourceAllocationSettingData_ResourceType_DVD_drive)
	if err != nil {
		return -1, err
	}
	col = append(col, dvdcol...)

	freeLocation := 0
	defer func() {
		log.Printf("[WMI] - FreeLocation [%d] - Total Drives [%d]\n", freeLocation, len(col))
	}()
	for _, _ = range col {
		exists, err := checkIfLocationExists(col, freeLocation)
		if err != nil {
			return -1, err
		}
		if exists {
			freeLocation = freeLocation + 1
			continue
		}
		break
	}
	return int32(freeLocation), nil
}

func checkIfLocationExists(col wmi.WmiInstanceCollection, locationNumber int) (bool, error) {
	for _, inst := range col {
		vmdrive, err := drive.NewSyntheticDiskDrive(inst)
		if err != nil {
			return false, err
		}
		loc, err := vmdrive.GetControllerLocation()
		if err != nil {
			return false, err
		}

		if loc == fmt.Sprintf("%d", locationNumber) {
			return true, nil
		}

	}
	return false, nil
}

func (settings *SCSIControllerSettings) getResourceAllocationSettingData(rtype v2.ResourceAllocationSettingData_ResourceType) (col wmi.WmiInstanceCollection, err error) {
	resourceType := fmt.Sprintf("%d", int32(rtype))
	query := query.NewWmiQuery("Cim_ResourceAllocationSettingData", "ResourceType", resourceType)
	col, err = settings.GetAllRelatedWithQuery(query)
	return
}
