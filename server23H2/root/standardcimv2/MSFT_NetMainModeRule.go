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
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_NetMainModeRule struct
type MSFT_NetMainModeRule struct {
	*MSFT_NetSARule
}

func NewMSFT_NetMainModeRuleEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetMainModeRule, err error) {
	tmp, err := NewMSFT_NetSARuleEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetMainModeRule{
		MSFT_NetSARule: tmp,
	}
	return
}

func NewMSFT_NetMainModeRuleEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetMainModeRule, err error) {
	tmp, err := NewMSFT_NetSARuleEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetMainModeRule{
		MSFT_NetSARule: tmp,
	}
	return
}

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetMainModeRule) Enable() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Enable")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetMainModeRule) Disable() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Disable")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="NewName" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetMainModeRule) Rename( /* IN */ NewName string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Rename", NewName)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="NewGPOSession" type="string "></param>
// <param name="NewName" type="string "></param>
// <param name="NewPolicyStore" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetMainModeRule) CloneObject( /* IN */ NewName string,
	/* IN */ NewPolicyStore string,
	/* IN */ NewGPOSession string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("CloneObject", NewName, NewPolicyStore, NewGPOSession)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
