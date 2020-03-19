// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_VirtualHardDiskState struct
type Msvm_VirtualHardDiskState struct {
	*cim.WmiInstance

	//
	Alignment uint32

	//
	FileSize uint64

	//
	FragmentationPercentage uint32

	//
	InUse bool

	//
	MinInternalSize uint64

	//
	PhysicalSectorSize uint32

	//
	Timestamp string
}

func NewMsvm_VirtualHardDiskStateEx1(instance *cim.WmiInstance) (newInstance *Msvm_VirtualHardDiskState, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &Msvm_VirtualHardDiskState{
		WmiInstance: tmp,
	}
	return
}

func NewMsvm_VirtualHardDiskStateEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_VirtualHardDiskState, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_VirtualHardDiskState{
		WmiInstance: tmp,
	}
	return
}

// SetAlignment sets the value of Alignment for the instance
func (instance *Msvm_VirtualHardDiskState) SetPropertyAlignment(value uint32) (err error) {
	return instance.SetProperty("Alignment", value)
}

// GetAlignment gets the value of Alignment for the instance
func (instance *Msvm_VirtualHardDiskState) GetPropertyAlignment() (value uint32, err error) {
	retValue, err := instance.GetProperty("Alignment")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFileSize sets the value of FileSize for the instance
func (instance *Msvm_VirtualHardDiskState) SetPropertyFileSize(value uint64) (err error) {
	return instance.SetProperty("FileSize", value)
}

// GetFileSize gets the value of FileSize for the instance
func (instance *Msvm_VirtualHardDiskState) GetPropertyFileSize() (value uint64, err error) {
	retValue, err := instance.GetProperty("FileSize")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFragmentationPercentage sets the value of FragmentationPercentage for the instance
func (instance *Msvm_VirtualHardDiskState) SetPropertyFragmentationPercentage(value uint32) (err error) {
	return instance.SetProperty("FragmentationPercentage", value)
}

// GetFragmentationPercentage gets the value of FragmentationPercentage for the instance
func (instance *Msvm_VirtualHardDiskState) GetPropertyFragmentationPercentage() (value uint32, err error) {
	retValue, err := instance.GetProperty("FragmentationPercentage")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInUse sets the value of InUse for the instance
func (instance *Msvm_VirtualHardDiskState) SetPropertyInUse(value bool) (err error) {
	return instance.SetProperty("InUse", value)
}

// GetInUse gets the value of InUse for the instance
func (instance *Msvm_VirtualHardDiskState) GetPropertyInUse() (value bool, err error) {
	retValue, err := instance.GetProperty("InUse")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMinInternalSize sets the value of MinInternalSize for the instance
func (instance *Msvm_VirtualHardDiskState) SetPropertyMinInternalSize(value uint64) (err error) {
	return instance.SetProperty("MinInternalSize", value)
}

// GetMinInternalSize gets the value of MinInternalSize for the instance
func (instance *Msvm_VirtualHardDiskState) GetPropertyMinInternalSize() (value uint64, err error) {
	retValue, err := instance.GetProperty("MinInternalSize")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPhysicalSectorSize sets the value of PhysicalSectorSize for the instance
func (instance *Msvm_VirtualHardDiskState) SetPropertyPhysicalSectorSize(value uint32) (err error) {
	return instance.SetProperty("PhysicalSectorSize", value)
}

// GetPhysicalSectorSize gets the value of PhysicalSectorSize for the instance
func (instance *Msvm_VirtualHardDiskState) GetPropertyPhysicalSectorSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("PhysicalSectorSize")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTimestamp sets the value of Timestamp for the instance
func (instance *Msvm_VirtualHardDiskState) SetPropertyTimestamp(value string) (err error) {
	return instance.SetProperty("Timestamp", value)
}

// GetTimestamp gets the value of Timestamp for the instance
func (instance *Msvm_VirtualHardDiskState) GetPropertyTimestamp() (value string, err error) {
	retValue, err := instance.GetProperty("Timestamp")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
