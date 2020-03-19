// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.Microsoft.Windows.Storage
//////////////////////////////////////////////
package storage

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_DiskToPartition struct
type MSFT_DiskToPartition struct {
	*cim.WmiInstance

	//
	Disk MSFT_Disk

	//
	Partition MSFT_Partition
}

func NewMSFT_DiskToPartitionEx1(instance *cim.WmiInstance) (newInstance *MSFT_DiskToPartition, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_DiskToPartition{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_DiskToPartitionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_DiskToPartition, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_DiskToPartition{
		WmiInstance: tmp,
	}
	return
}

// SetDisk sets the value of Disk for the instance
func (instance *MSFT_DiskToPartition) SetPropertyDisk(value MSFT_Disk) (err error) {
	return instance.SetProperty("Disk", value)
}

// GetDisk gets the value of Disk for the instance
func (instance *MSFT_DiskToPartition) GetPropertyDisk() (value MSFT_Disk, err error) {
	retValue, err := instance.GetProperty("Disk")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_Disk)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPartition sets the value of Partition for the instance
func (instance *MSFT_DiskToPartition) SetPropertyPartition(value MSFT_Partition) (err error) {
	return instance.SetProperty("Partition", value)
}

// GetPartition gets the value of Partition for the instance
func (instance *MSFT_DiskToPartition) GetPropertyPartition() (value MSFT_Partition, err error) {
	retValue, err := instance.GetProperty("Partition")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_Partition)
	if !ok {
		// TODO: Set an error
	}
	return
}
