// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_NetIKEP1AuthSet struct
type MSFT_NetIKEP1AuthSet struct {
	*MSFT_NetIKEAuthSet
}

func NewMSFT_NetIKEP1AuthSetEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetIKEP1AuthSet, err error) {
	tmp, err := NewMSFT_NetIKEAuthSetEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetIKEP1AuthSet{
		MSFT_NetIKEAuthSet: tmp,
	}
	return
}

func NewMSFT_NetIKEP1AuthSetEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetIKEP1AuthSet, err error) {
	tmp, err := NewMSFT_NetIKEAuthSetEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetIKEP1AuthSet{
		MSFT_NetIKEAuthSet: tmp,
	}
	return
}

//

// <param name="NewName" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetIKEP1AuthSet) Rename( /* IN */ NewName string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Rename", NewName)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="NewGPOSession" type="string "></param>
// <param name="NewID" type="string "></param>
// <param name="NewName" type="string "></param>
// <param name="NewPolicyStore" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetIKEP1AuthSet) CloneObject( /* IN */ NewName string,
	/* IN */ NewID string,
	/* IN */ NewPolicyStore string,
	/* IN */ NewGPOSession string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("CloneObject", NewName, NewID, NewPolicyStore, NewGPOSession)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
