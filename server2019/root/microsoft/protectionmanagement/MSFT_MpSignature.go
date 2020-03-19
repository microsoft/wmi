// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.Microsoft.protectionManagement
//////////////////////////////////////////////
package protectionmanagement

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_MpSignature struct
type MSFT_MpSignature struct {
	*cim.WmiInstance
}

func NewMSFT_MpSignatureEx1(instance *cim.WmiInstance) (newInstance *MSFT_MpSignature, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_MpSignature{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_MpSignatureEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_MpSignature, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_MpSignature{
		WmiInstance: tmp,
	}
	return
}

//

// <param name="UpdateSource" type="uint8 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_MpSignature) Update( /* IN */ UpdateSource uint8) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Update", UpdateSource)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
