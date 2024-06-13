// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_NetFirewallHyperVVMCreator struct
type MSFT_NetFirewallHyperVVMCreator struct {
	*CIM_ManagedElement

	//
	FriendlyName string

	//
	VMCreatorId string
}

func NewMSFT_NetFirewallHyperVVMCreatorEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetFirewallHyperVVMCreator, err error) {
	tmp, err := NewCIM_ManagedElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetFirewallHyperVVMCreator{
		CIM_ManagedElement: tmp,
	}
	return
}

func NewMSFT_NetFirewallHyperVVMCreatorEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetFirewallHyperVVMCreator, err error) {
	tmp, err := NewCIM_ManagedElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetFirewallHyperVVMCreator{
		CIM_ManagedElement: tmp,
	}
	return
}

// SetFriendlyName sets the value of FriendlyName for the instance
func (instance *MSFT_NetFirewallHyperVVMCreator) SetPropertyFriendlyName(value string) (err error) {
	return instance.SetProperty("FriendlyName", (value))
}

// GetFriendlyName gets the value of FriendlyName for the instance
func (instance *MSFT_NetFirewallHyperVVMCreator) GetPropertyFriendlyName() (value string, err error) {
	retValue, err := instance.GetProperty("FriendlyName")
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

// SetVMCreatorId sets the value of VMCreatorId for the instance
func (instance *MSFT_NetFirewallHyperVVMCreator) SetPropertyVMCreatorId(value string) (err error) {
	return instance.SetProperty("VMCreatorId", (value))
}

// GetVMCreatorId gets the value of VMCreatorId for the instance
func (instance *MSFT_NetFirewallHyperVVMCreator) GetPropertyVMCreatorId() (value string, err error) {
	retValue, err := instance.GetProperty("VMCreatorId")
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

// <param name="FriendlyName" type="string "></param>
// <param name="VMCreatorId" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetFirewallHyperVVMCreator) RegisterHyperVVMCreator( /* IN */ VMCreatorId string,
	/* IN */ FriendlyName string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("RegisterHyperVVMCreator", VMCreatorId, FriendlyName)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="VMCreatorId" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetFirewallHyperVVMCreator) UnregisterHyperVVMCreator( /* IN */ VMCreatorId string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("UnregisterHyperVVMCreator", VMCreatorId)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
