// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.MSCluster
//////////////////////////////////////////////
package mscluster

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSCluster_Disk struct
type MSCluster_Disk struct {
	*MSCluster_ClusterDisk

	//
	UniqueId string

	//
	UniqueIdFormat uint16
}

func NewMSCluster_DiskEx1(instance *cim.WmiInstance) (newInstance *MSCluster_Disk, err error) {
	tmp, err := NewMSCluster_ClusterDiskEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSCluster_Disk{
		MSCluster_ClusterDisk: tmp,
	}
	return
}

func NewMSCluster_DiskEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSCluster_Disk, err error) {
	tmp, err := NewMSCluster_ClusterDiskEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSCluster_Disk{
		MSCluster_ClusterDisk: tmp,
	}
	return
}

// SetUniqueId sets the value of UniqueId for the instance
func (instance *MSCluster_Disk) SetPropertyUniqueId(value string) (err error) {
	return instance.SetProperty("UniqueId", (value))
}

// GetUniqueId gets the value of UniqueId for the instance
func (instance *MSCluster_Disk) GetPropertyUniqueId() (value string, err error) {
	retValue, err := instance.GetProperty("UniqueId")
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

// SetUniqueIdFormat sets the value of UniqueIdFormat for the instance
func (instance *MSCluster_Disk) SetPropertyUniqueIdFormat(value uint16) (err error) {
	return instance.SetProperty("UniqueIdFormat", (value))
}

// GetUniqueIdFormat gets the value of UniqueIdFormat for the instance
func (instance *MSCluster_Disk) GetPropertyUniqueIdFormat() (value uint16, err error) {
	retValue, err := instance.GetProperty("UniqueIdFormat")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}
