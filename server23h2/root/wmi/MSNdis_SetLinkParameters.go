// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSNdis_SetLinkParameters struct
type MSNdis_SetLinkParameters struct {
	*MSNdis

	//
	Active bool

	//
	InstanceName string
}

func NewMSNdis_SetLinkParametersEx1(instance *cim.WmiInstance) (newInstance *MSNdis_SetLinkParameters, err error) {
	tmp, err := NewMSNdisEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSNdis_SetLinkParameters{
		MSNdis: tmp,
	}
	return
}

func NewMSNdis_SetLinkParametersEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSNdis_SetLinkParameters, err error) {
	tmp, err := NewMSNdisEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSNdis_SetLinkParameters{
		MSNdis: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *MSNdis_SetLinkParameters) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSNdis_SetLinkParameters) GetPropertyActive() (value bool, err error) {
	retValue, err := instance.GetProperty("Active")
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

// SetInstanceName sets the value of InstanceName for the instance
func (instance *MSNdis_SetLinkParameters) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSNdis_SetLinkParameters) GetPropertyInstanceName() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceName")
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

// <param name="LinkParameters" type="MSNdis_LinkParameters "></param>
// <param name="MethodHeader" type="MSNdis_WmiMethodHeader "></param>

// <param name="OutputInfo" type="MSNdis_WmiOutputInfo "></param>
func (instance *MSNdis_SetLinkParameters) WmiSetLinkParameters( /* IN */ MethodHeader MSNdis_WmiMethodHeader,
	/* IN */ LinkParameters MSNdis_LinkParameters,
	/* OUT */ OutputInfo MSNdis_WmiOutputInfo) (err error) {
	_, err = instance.InvokeMethod("WmiSetLinkParameters", MethodHeader, LinkParameters)
	if err != nil {
		return
	}
	return

}
