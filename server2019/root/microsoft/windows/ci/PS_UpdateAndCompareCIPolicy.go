// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.Microsoft.Windows.CI
//////////////////////////////////////////////
package ci

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// PS_UpdateAndCompareCIPolicy struct
type PS_UpdateAndCompareCIPolicy struct {
	*cim.WmiInstance
}

func NewPS_UpdateAndCompareCIPolicyEx1(instance *cim.WmiInstance) (newInstance *PS_UpdateAndCompareCIPolicy, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &PS_UpdateAndCompareCIPolicy{
		WmiInstance: tmp,
	}
	return
}

func NewPS_UpdateAndCompareCIPolicyEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *PS_UpdateAndCompareCIPolicy, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &PS_UpdateAndCompareCIPolicy{
		WmiInstance: tmp,
	}
	return
}

//

// <param name="FilePath" type="string "></param>

// <param name="cmdletOutput" type="uint64 "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_UpdateAndCompareCIPolicy) Compare( /* IN */ FilePath string,
	/* OUT */ cmdletOutput uint64) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Compare", FilePath)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="FilePath" type="string "></param>

// <param name="cmdletOutput" type="uint64 "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_UpdateAndCompareCIPolicy) Update( /* IN */ FilePath string,
	/* OUT */ cmdletOutput uint64) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Update", FilePath)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="cmdletOutput" type="uint64 "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_UpdateAndCompareCIPolicy) Delete( /* OUT */ cmdletOutput uint64) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Delete")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
