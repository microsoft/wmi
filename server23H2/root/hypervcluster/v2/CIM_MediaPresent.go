// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.HyperVCluster.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// CIM_MediaPresent struct
type CIM_MediaPresent struct {
	*CIM_Dependency

	// Boolean indicating that the accessed StorageExtent is fixed in the MediaAccessDevice and can not be ejected.
	FixedMedia bool
}

func NewCIM_MediaPresentEx1(instance *cim.WmiInstance) (newInstance *CIM_MediaPresent, err error) {
	tmp, err := NewCIM_DependencyEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_MediaPresent{
		CIM_Dependency: tmp,
	}
	return
}

func NewCIM_MediaPresentEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_MediaPresent, err error) {
	tmp, err := NewCIM_DependencyEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_MediaPresent{
		CIM_Dependency: tmp,
	}
	return
}

// SetFixedMedia sets the value of FixedMedia for the instance
func (instance *CIM_MediaPresent) SetPropertyFixedMedia(value bool) (err error) {
	return instance.SetProperty("FixedMedia", (value))
}

// GetFixedMedia gets the value of FixedMedia for the instance
func (instance *CIM_MediaPresent) GetPropertyFixedMedia() (value bool, err error) {
	retValue, err := instance.GetProperty("FixedMedia")
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
