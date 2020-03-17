// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.Storage
//////////////////////////////////////////////
package storage

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_PartitionToVolume struct
type MSFT_PartitionToVolume struct {
	cim.WmiInstance

	//
	Partition MSFT_Partition

	//
	Volume MSFT_Volume
}

// SetPartition sets the value of Partition for the instance
func (instance *MSFT_PartitionToVolume) SetPropertyPartition(value MSFT_Partition) (err error) {
	return instance.SetProperty("Partition", value)
}

// GetPartition gets the value of Partition for the instance
func (instance *MSFT_PartitionToVolume) GetPropertyPartition() (value MSFT_Partition, err error) {
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

// SetVolume sets the value of Volume for the instance
func (instance *MSFT_PartitionToVolume) SetPropertyVolume(value MSFT_Volume) (err error) {
	return instance.SetProperty("Volume", value)
}

// GetVolume gets the value of Volume for the instance
func (instance *MSFT_PartitionToVolume) GetPropertyVolume() (value MSFT_Volume, err error) {
	retValue, err := instance.GetProperty("Volume")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_Volume)
	if !ok {
		// TODO: Set an error
	}
	return
}
