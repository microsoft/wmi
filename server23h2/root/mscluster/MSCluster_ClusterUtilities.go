// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.MSCluster
//////////////////////////////////////////////
package mscluster

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
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
	return instance.SetProperty("Fqdn", (value))
}

// GetFqdn gets the value of Fqdn for the instance
func (instance *MSCluster_ClusterUtilities) GetPropertyFqdn() (value string, err error) {
	retValue, err := instance.GetProperty("Fqdn")
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

// SetHasSystemAccess sets the value of HasSystemAccess for the instance
func (instance *MSCluster_ClusterUtilities) SetPropertyHasSystemAccess(value bool) (err error) {
	return instance.SetProperty("HasSystemAccess", (value))
}

// GetHasSystemAccess gets the value of HasSystemAccess for the instance
func (instance *MSCluster_ClusterUtilities) GetPropertyHasSystemAccess() (value bool, err error) {
	retValue, err := instance.GetProperty("HasSystemAccess")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

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
