// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
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

// MSCluster_ValidationStatus struct
type MSCluster_ValidationStatus struct {
	*cim.WmiInstance

	//
	Id string
}

func NewMSCluster_ValidationStatusEx1(instance *cim.WmiInstance) (newInstance *MSCluster_ValidationStatus, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSCluster_ValidationStatus{
		WmiInstance: tmp,
	}
	return
}

func NewMSCluster_ValidationStatusEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSCluster_ValidationStatus, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSCluster_ValidationStatus{
		WmiInstance: tmp,
	}
	return
}

// SetId sets the value of Id for the instance
func (instance *MSCluster_ValidationStatus) SetPropertyId(value string) (err error) {
	return instance.SetProperty("Id", (value))
}

// GetId gets the value of Id for the instance
func (instance *MSCluster_ValidationStatus) GetPropertyId() (value string, err error) {
	retValue, err := instance.GetProperty("Id")
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

//

// <param name="Status" type="uint32 "></param>
func (instance *MSCluster_ValidationStatus) GetStatus( /* OUT */ Status uint32) (err error) {
	_, err = instance.InvokeMethod("GetStatus")
	if err != nil {
		return
	}
	return

}

//

// <param name="ReturnValue" type="bool "></param>
func (instance *MSCluster_ValidationStatus) IsValidationSuccessful() (result bool, err error) {
	retVal, err := instance.InvokeMethodWithReturn("IsValidationSuccessful")
	if err != nil {
		return
	}
	result = (retVal > 0)
	return

}

//

// <param name="NodeNames" type="string []"></param>

// <param name="Status" type="uint32 "></param>
func (instance *MSCluster_ValidationStatus) GetNodeStatus( /* IN */ NodeNames []string,
	/* OUT */ Status uint32) (err error) {
	_, err = instance.InvokeMethod("GetNodeStatus", NodeNames)
	if err != nil {
		return
	}
	return

}
