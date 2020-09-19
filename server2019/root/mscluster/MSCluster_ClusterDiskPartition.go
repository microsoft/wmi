// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.MSCluster
//////////////////////////////////////////////
package mscluster

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSCluster_ClusterDiskPartition struct
type MSCluster_ClusterDiskPartition struct {
	*MSCluster_LogicalElement

	//
	FileSystem string

	//
	FileSystemFlags uint32

	//
	FreeSpace uint32

	//
	MaximumComponentLength uint32

	//
	MountPoints []string

	//
	PartitionNumber uint32

	//
	Path string

	//
	SerialNumber uint32

	//
	TotalSize uint32

	//
	VolumeGuid string

	//
	VolumeLabel string
}

func NewMSCluster_ClusterDiskPartitionEx1(instance *cim.WmiInstance) (newInstance *MSCluster_ClusterDiskPartition, err error) {
	tmp, err := NewMSCluster_LogicalElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSCluster_ClusterDiskPartition{
		MSCluster_LogicalElement: tmp,
	}
	return
}

func NewMSCluster_ClusterDiskPartitionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSCluster_ClusterDiskPartition, err error) {
	tmp, err := NewMSCluster_LogicalElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSCluster_ClusterDiskPartition{
		MSCluster_LogicalElement: tmp,
	}
	return
}

// SetFileSystem sets the value of FileSystem for the instance
func (instance *MSCluster_ClusterDiskPartition) SetPropertyFileSystem(value string) (err error) {
	return instance.SetProperty("FileSystem", (value))
}

// GetFileSystem gets the value of FileSystem for the instance
func (instance *MSCluster_ClusterDiskPartition) GetPropertyFileSystem() (value string, err error) {
	retValue, err := instance.GetProperty("FileSystem")
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

// SetFileSystemFlags sets the value of FileSystemFlags for the instance
func (instance *MSCluster_ClusterDiskPartition) SetPropertyFileSystemFlags(value uint32) (err error) {
	return instance.SetProperty("FileSystemFlags", (value))
}

// GetFileSystemFlags gets the value of FileSystemFlags for the instance
func (instance *MSCluster_ClusterDiskPartition) GetPropertyFileSystemFlags() (value uint32, err error) {
	retValue, err := instance.GetProperty("FileSystemFlags")
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

// SetFreeSpace sets the value of FreeSpace for the instance
func (instance *MSCluster_ClusterDiskPartition) SetPropertyFreeSpace(value uint32) (err error) {
	return instance.SetProperty("FreeSpace", (value))
}

// GetFreeSpace gets the value of FreeSpace for the instance
func (instance *MSCluster_ClusterDiskPartition) GetPropertyFreeSpace() (value uint32, err error) {
	retValue, err := instance.GetProperty("FreeSpace")
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

// SetMaximumComponentLength sets the value of MaximumComponentLength for the instance
func (instance *MSCluster_ClusterDiskPartition) SetPropertyMaximumComponentLength(value uint32) (err error) {
	return instance.SetProperty("MaximumComponentLength", (value))
}

// GetMaximumComponentLength gets the value of MaximumComponentLength for the instance
func (instance *MSCluster_ClusterDiskPartition) GetPropertyMaximumComponentLength() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaximumComponentLength")
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

// SetMountPoints sets the value of MountPoints for the instance
func (instance *MSCluster_ClusterDiskPartition) SetPropertyMountPoints(value []string) (err error) {
	return instance.SetProperty("MountPoints", (value))
}

// GetMountPoints gets the value of MountPoints for the instance
func (instance *MSCluster_ClusterDiskPartition) GetPropertyMountPoints() (value []string, err error) {
	retValue, err := instance.GetProperty("MountPoints")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}

// SetPartitionNumber sets the value of PartitionNumber for the instance
func (instance *MSCluster_ClusterDiskPartition) SetPropertyPartitionNumber(value uint32) (err error) {
	return instance.SetProperty("PartitionNumber", (value))
}

// GetPartitionNumber gets the value of PartitionNumber for the instance
func (instance *MSCluster_ClusterDiskPartition) GetPropertyPartitionNumber() (value uint32, err error) {
	retValue, err := instance.GetProperty("PartitionNumber")
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

// SetPath sets the value of Path for the instance
func (instance *MSCluster_ClusterDiskPartition) SetPropertyPath(value string) (err error) {
	return instance.SetProperty("Path", (value))
}

// GetPath gets the value of Path for the instance
func (instance *MSCluster_ClusterDiskPartition) GetPropertyPath() (value string, err error) {
	retValue, err := instance.GetProperty("Path")
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

// SetSerialNumber sets the value of SerialNumber for the instance
func (instance *MSCluster_ClusterDiskPartition) SetPropertySerialNumber(value uint32) (err error) {
	return instance.SetProperty("SerialNumber", (value))
}

// GetSerialNumber gets the value of SerialNumber for the instance
func (instance *MSCluster_ClusterDiskPartition) GetPropertySerialNumber() (value uint32, err error) {
	retValue, err := instance.GetProperty("SerialNumber")
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

// SetTotalSize sets the value of TotalSize for the instance
func (instance *MSCluster_ClusterDiskPartition) SetPropertyTotalSize(value uint32) (err error) {
	return instance.SetProperty("TotalSize", (value))
}

// GetTotalSize gets the value of TotalSize for the instance
func (instance *MSCluster_ClusterDiskPartition) GetPropertyTotalSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("TotalSize")
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

// SetVolumeGuid sets the value of VolumeGuid for the instance
func (instance *MSCluster_ClusterDiskPartition) SetPropertyVolumeGuid(value string) (err error) {
	return instance.SetProperty("VolumeGuid", (value))
}

// GetVolumeGuid gets the value of VolumeGuid for the instance
func (instance *MSCluster_ClusterDiskPartition) GetPropertyVolumeGuid() (value string, err error) {
	retValue, err := instance.GetProperty("VolumeGuid")
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

// SetVolumeLabel sets the value of VolumeLabel for the instance
func (instance *MSCluster_ClusterDiskPartition) SetPropertyVolumeLabel(value string) (err error) {
	return instance.SetProperty("VolumeLabel", (value))
}

// GetVolumeLabel gets the value of VolumeLabel for the instance
func (instance *MSCluster_ClusterDiskPartition) GetPropertyVolumeLabel() (value string, err error) {
	retValue, err := instance.GetProperty("VolumeLabel")
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
