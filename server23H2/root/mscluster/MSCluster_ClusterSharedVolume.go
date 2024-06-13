// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.MSCluster
//////////////////////////////////////////////
package mscluster

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSCluster_ClusterSharedVolume struct
type MSCluster_ClusterSharedVolume struct {
	*MSCluster_LogicalElement

	//
	BackupState uint32

	//
	FaultState uint32

	//
	VolumeName string

	//
	VolumeOffset uint64
}

func NewMSCluster_ClusterSharedVolumeEx1(instance *cim.WmiInstance) (newInstance *MSCluster_ClusterSharedVolume, err error) {
	tmp, err := NewMSCluster_LogicalElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSCluster_ClusterSharedVolume{
		MSCluster_LogicalElement: tmp,
	}
	return
}

func NewMSCluster_ClusterSharedVolumeEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSCluster_ClusterSharedVolume, err error) {
	tmp, err := NewMSCluster_LogicalElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSCluster_ClusterSharedVolume{
		MSCluster_LogicalElement: tmp,
	}
	return
}

// SetBackupState sets the value of BackupState for the instance
func (instance *MSCluster_ClusterSharedVolume) SetPropertyBackupState(value uint32) (err error) {
	return instance.SetProperty("BackupState", (value))
}

// GetBackupState gets the value of BackupState for the instance
func (instance *MSCluster_ClusterSharedVolume) GetPropertyBackupState() (value uint32, err error) {
	retValue, err := instance.GetProperty("BackupState")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetFaultState sets the value of FaultState for the instance
func (instance *MSCluster_ClusterSharedVolume) SetPropertyFaultState(value uint32) (err error) {
	return instance.SetProperty("FaultState", (value))
}

// GetFaultState gets the value of FaultState for the instance
func (instance *MSCluster_ClusterSharedVolume) GetPropertyFaultState() (value uint32, err error) {
	retValue, err := instance.GetProperty("FaultState")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetVolumeName sets the value of VolumeName for the instance
func (instance *MSCluster_ClusterSharedVolume) SetPropertyVolumeName(value string) (err error) {
	return instance.SetProperty("VolumeName", (value))
}

// GetVolumeName gets the value of VolumeName for the instance
func (instance *MSCluster_ClusterSharedVolume) GetPropertyVolumeName() (value string, err error) {
	retValue, err := instance.GetProperty("VolumeName")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetVolumeOffset sets the value of VolumeOffset for the instance
func (instance *MSCluster_ClusterSharedVolume) SetPropertyVolumeOffset(value uint64) (err error) {
	return instance.SetProperty("VolumeOffset", (value))
}

// GetVolumeOffset gets the value of VolumeOffset for the instance
func (instance *MSCluster_ClusterSharedVolume) GetPropertyVolumeOffset() (value uint64, err error) {
	retValue, err := instance.GetProperty("VolumeOffset")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

//
func (instance *MSCluster_ClusterSharedVolume) TurnOnMaintenance() (err error) {
	_, err = instance.InvokeMethodWithReturn("TurnOnMaintenance")
	if err != nil {
		return
	}
	return

}

//
func (instance *MSCluster_ClusterSharedVolume) TurnOffMaintenance() (err error) {
	_, err = instance.InvokeMethodWithReturn("TurnOffMaintenance")
	if err != nil {
		return
	}
	return

}

//
func (instance *MSCluster_ClusterSharedVolume) TurnOnRedirectedAccess() (err error) {
	_, err = instance.InvokeMethodWithReturn("TurnOnRedirectedAccess")
	if err != nil {
		return
	}
	return

}

//
func (instance *MSCluster_ClusterSharedVolume) TurnOffRedirectedAccess() (err error) {
	_, err = instance.InvokeMethodWithReturn("TurnOffRedirectedAccess")
	if err != nil {
		return
	}
	return

}

//

// <param name="HostName" type="string "></param>
func (instance *MSCluster_ClusterSharedVolume) MoveToNewHost( /* IN */ HostName string) (err error) {
	_, err = instance.InvokeMethodWithReturn("MoveToNewHost", HostName)
	if err != nil {
		return
	}
	return

}
