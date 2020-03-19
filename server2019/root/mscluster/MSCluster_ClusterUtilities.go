// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.MSCluster
//////////////////////////////////////////////
package mscluster

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSCluster_ClusterUtilities struct
type MSCluster_ClusterUtilities struct {
	*cim.WmiInstance

	//
	Fqdn string

	//
	HasSystemAccess bool
}

func NewMSCluster_ClusterUtilitiesEx1(instance *cim.WmiInstance) (newInstance *MSCluster_ClusterUtilities, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSCluster_ClusterUtilities{
		WmiInstance: tmp,
	}
	return
}

func NewMSCluster_ClusterUtilitiesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSCluster_ClusterUtilities, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSCluster_ClusterUtilities{
		WmiInstance: tmp,
	}
	return
}

// SetFqdn sets the value of Fqdn for the instance
func (instance *MSCluster_ClusterUtilities) SetPropertyFqdn(value string) (err error) {
	return instance.SetProperty("Fqdn", value)
}

// GetFqdn gets the value of Fqdn for the instance
func (instance *MSCluster_ClusterUtilities) GetPropertyFqdn() (value string, err error) {
	retValue, err := instance.GetProperty("Fqdn")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetHasSystemAccess sets the value of HasSystemAccess for the instance
func (instance *MSCluster_ClusterUtilities) SetPropertyHasSystemAccess(value bool) (err error) {
	return instance.SetProperty("HasSystemAccess", value)
}

// GetHasSystemAccess gets the value of HasSystemAccess for the instance
func (instance *MSCluster_ClusterUtilities) GetPropertyHasSystemAccess() (value bool, err error) {
	retValue, err := instance.GetProperty("HasSystemAccess")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

//

// <param name="ReturnValue" type="bool "></param>
func (instance *MSCluster_ClusterUtilities) IsClusterSupported() (result bool, err error) {
	retVal, err := instance.InvokeMethodWithReturn("IsClusterSupported")
	if err != nil {
		return
	}
	result = (retVal > 0)
	return

}

//

// <param name="ReturnValue" type="bool "></param>
func (instance *MSCluster_ClusterUtilities) IsStorageSpacesDirectSupported() (result bool, err error) {
	retVal, err := instance.InvokeMethodWithReturn("IsStorageSpacesDirectSupported")
	if err != nil {
		return
	}
	result = (retVal > 0)
	return

}

//

// <param name="ReturnValue" type="bool "></param>
func (instance *MSCluster_ClusterUtilities) IsStorageSpacesDirectCacheSupported() (result bool, err error) {
	retVal, err := instance.InvokeMethodWithReturn("IsStorageSpacesDirectCacheSupported")
	if err != nil {
		return
	}
	result = (retVal > 0)
	return

}
