// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.MSCluster
//////////////////////////////////////////////
package mscluster

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSCluster_ClusterDisk struct
type MSCluster_ClusterDisk struct {
	*MSCluster_LogicalElement

	//
	GptGuid string

	//
	Id string

	//
	Number uint32

	//
	ScsiBus uint32

	//
	ScsiLun uint32

	//
	ScsiPort uint32

	//
	ScsiTargetId uint32

	//
	Signature uint32

	//
	Size uint64

	//
	StoragePoolId string

	//
	VirtualDiskId string
}

func NewMSCluster_ClusterDiskEx1(instance *cim.WmiInstance) (newInstance *MSCluster_ClusterDisk, err error) {
	tmp, err := NewMSCluster_LogicalElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSCluster_ClusterDisk{
		MSCluster_LogicalElement: tmp,
	}
	return
}

func NewMSCluster_ClusterDiskEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSCluster_ClusterDisk, err error) {
	tmp, err := NewMSCluster_LogicalElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSCluster_ClusterDisk{
		MSCluster_LogicalElement: tmp,
	}
	return
}

// SetGptGuid sets the value of GptGuid for the instance
func (instance *MSCluster_ClusterDisk) SetPropertyGptGuid(value string) (err error) {
	return instance.SetProperty("GptGuid", value)
}

// GetGptGuid gets the value of GptGuid for the instance
func (instance *MSCluster_ClusterDisk) GetPropertyGptGuid() (value string, err error) {
	retValue, err := instance.GetProperty("GptGuid")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetId sets the value of Id for the instance
func (instance *MSCluster_ClusterDisk) SetPropertyId(value string) (err error) {
	return instance.SetProperty("Id", value)
}

// GetId gets the value of Id for the instance
func (instance *MSCluster_ClusterDisk) GetPropertyId() (value string, err error) {
	retValue, err := instance.GetProperty("Id")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumber sets the value of Number for the instance
func (instance *MSCluster_ClusterDisk) SetPropertyNumber(value uint32) (err error) {
	return instance.SetProperty("Number", value)
}

// GetNumber gets the value of Number for the instance
func (instance *MSCluster_ClusterDisk) GetPropertyNumber() (value uint32, err error) {
	retValue, err := instance.GetProperty("Number")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetScsiBus sets the value of ScsiBus for the instance
func (instance *MSCluster_ClusterDisk) SetPropertyScsiBus(value uint32) (err error) {
	return instance.SetProperty("ScsiBus", value)
}

// GetScsiBus gets the value of ScsiBus for the instance
func (instance *MSCluster_ClusterDisk) GetPropertyScsiBus() (value uint32, err error) {
	retValue, err := instance.GetProperty("ScsiBus")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetScsiLun sets the value of ScsiLun for the instance
func (instance *MSCluster_ClusterDisk) SetPropertyScsiLun(value uint32) (err error) {
	return instance.SetProperty("ScsiLun", value)
}

// GetScsiLun gets the value of ScsiLun for the instance
func (instance *MSCluster_ClusterDisk) GetPropertyScsiLun() (value uint32, err error) {
	retValue, err := instance.GetProperty("ScsiLun")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetScsiPort sets the value of ScsiPort for the instance
func (instance *MSCluster_ClusterDisk) SetPropertyScsiPort(value uint32) (err error) {
	return instance.SetProperty("ScsiPort", value)
}

// GetScsiPort gets the value of ScsiPort for the instance
func (instance *MSCluster_ClusterDisk) GetPropertyScsiPort() (value uint32, err error) {
	retValue, err := instance.GetProperty("ScsiPort")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetScsiTargetId sets the value of ScsiTargetId for the instance
func (instance *MSCluster_ClusterDisk) SetPropertyScsiTargetId(value uint32) (err error) {
	return instance.SetProperty("ScsiTargetId", value)
}

// GetScsiTargetId gets the value of ScsiTargetId for the instance
func (instance *MSCluster_ClusterDisk) GetPropertyScsiTargetId() (value uint32, err error) {
	retValue, err := instance.GetProperty("ScsiTargetId")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSignature sets the value of Signature for the instance
func (instance *MSCluster_ClusterDisk) SetPropertySignature(value uint32) (err error) {
	return instance.SetProperty("Signature", value)
}

// GetSignature gets the value of Signature for the instance
func (instance *MSCluster_ClusterDisk) GetPropertySignature() (value uint32, err error) {
	retValue, err := instance.GetProperty("Signature")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSize sets the value of Size for the instance
func (instance *MSCluster_ClusterDisk) SetPropertySize(value uint64) (err error) {
	return instance.SetProperty("Size", value)
}

// GetSize gets the value of Size for the instance
func (instance *MSCluster_ClusterDisk) GetPropertySize() (value uint64, err error) {
	retValue, err := instance.GetProperty("Size")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStoragePoolId sets the value of StoragePoolId for the instance
func (instance *MSCluster_ClusterDisk) SetPropertyStoragePoolId(value string) (err error) {
	return instance.SetProperty("StoragePoolId", value)
}

// GetStoragePoolId gets the value of StoragePoolId for the instance
func (instance *MSCluster_ClusterDisk) GetPropertyStoragePoolId() (value string, err error) {
	retValue, err := instance.GetProperty("StoragePoolId")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVirtualDiskId sets the value of VirtualDiskId for the instance
func (instance *MSCluster_ClusterDisk) SetPropertyVirtualDiskId(value string) (err error) {
	return instance.SetProperty("VirtualDiskId", value)
}

// GetVirtualDiskId gets the value of VirtualDiskId for the instance
func (instance *MSCluster_ClusterDisk) GetPropertyVirtualDiskId() (value string, err error) {
	retValue, err := instance.GetProperty("VirtualDiskId")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
