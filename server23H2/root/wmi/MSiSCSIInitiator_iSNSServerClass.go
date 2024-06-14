// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSiSCSIInitiator_iSNSServerClass struct
type MSiSCSIInitiator_iSNSServerClass struct {
	*cim.WmiInstance

	//
	iSNSServerAddress string
}

func NewMSiSCSIInitiator_iSNSServerClassEx1(instance *cim.WmiInstance) (newInstance *MSiSCSIInitiator_iSNSServerClass, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSiSCSIInitiator_iSNSServerClass{
		WmiInstance: tmp,
	}
	return
}

func NewMSiSCSIInitiator_iSNSServerClassEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSiSCSIInitiator_iSNSServerClass, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSiSCSIInitiator_iSNSServerClass{
		WmiInstance: tmp,
	}
	return
}

// SetiSNSServerAddress sets the value of iSNSServerAddress for the instance
func (instance *MSiSCSIInitiator_iSNSServerClass) SetPropertyiSNSServerAddress(value string) (err error) {
	return instance.SetProperty("iSNSServerAddress", (value))
}

// GetiSNSServerAddress gets the value of iSNSServerAddress for the instance
func (instance *MSiSCSIInitiator_iSNSServerClass) GetPropertyiSNSServerAddress() (value string, err error) {
	retValue, err := instance.GetProperty("iSNSServerAddress")
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

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSiSCSIInitiator_iSNSServerClass) Refresh() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Refresh")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
