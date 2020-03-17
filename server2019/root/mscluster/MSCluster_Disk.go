// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.MSCluster
//////////////////////////////////////////////
package mscluster

// MSCluster_Disk struct
type MSCluster_Disk struct {
	MSCluster_ClusterDisk

	//
	UniqueId string

	//
	UniqueIdFormat uint16
}

// SetUniqueId sets the value of UniqueId for the instance
func (instance *MSCluster_Disk) SetPropertyUniqueId(value string) (err error) {
	return instance.SetProperty("UniqueId", value)
}

// GetUniqueId gets the value of UniqueId for the instance
func (instance *MSCluster_Disk) GetPropertyUniqueId() (value string, err error) {
	retValue, err := instance.GetProperty("UniqueId")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUniqueIdFormat sets the value of UniqueIdFormat for the instance
func (instance *MSCluster_Disk) SetPropertyUniqueIdFormat(value uint16) (err error) {
	return instance.SetProperty("UniqueIdFormat", value)
}

// GetUniqueIdFormat gets the value of UniqueIdFormat for the instance
func (instance *MSCluster_Disk) GetPropertyUniqueIdFormat() (value uint16, err error) {
	retValue, err := instance.GetProperty("UniqueIdFormat")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}
