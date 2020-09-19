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

// MSCluster_ExtendedStatus struct
type MSCluster_ExtendedStatus struct {
	*__ExtendedStatus

	//
	ErrorType uint32
}

func NewMSCluster_ExtendedStatusEx1(instance *cim.WmiInstance) (newInstance *MSCluster_ExtendedStatus, err error) {
	tmp, err := New__ExtendedStatusEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSCluster_ExtendedStatus{
		__ExtendedStatus: tmp,
	}
	return
}

func NewMSCluster_ExtendedStatusEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSCluster_ExtendedStatus, err error) {
	tmp, err := New__ExtendedStatusEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSCluster_ExtendedStatus{
		__ExtendedStatus: tmp,
	}
	return
}

// SetErrorType sets the value of ErrorType for the instance
func (instance *MSCluster_ExtendedStatus) SetPropertyErrorType(value uint32) (err error) {
	return instance.SetProperty("ErrorType", (value))
}

// GetErrorType gets the value of ErrorType for the instance
func (instance *MSCluster_ExtendedStatus) GetPropertyErrorType() (value uint32, err error) {
	retValue, err := instance.GetProperty("ErrorType")
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
