// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.Microsoft.Windows.Wdac
//////////////////////////////////////////////
package wdac

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_OdbcDriverTask struct
type MSFT_OdbcDriverTask struct {
	*cim.WmiInstance
}

func NewMSFT_OdbcDriverTaskEx1(instance *cim.WmiInstance) (newInstance *MSFT_OdbcDriverTask, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_OdbcDriverTask{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_OdbcDriverTaskEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_OdbcDriverTask, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_OdbcDriverTask{
		WmiInstance: tmp,
	}
	return
}

//

// <param name="Name" type="string "></param>
// <param name="Platform" type="string "></param>

// <param name="cmdletOutput" type="MSFT_OdbcDriver []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_OdbcDriverTask) Get( /* IN */ Name string,
	/* IN */ Platform string,
	/* OUT */ cmdletOutput []MSFT_OdbcDriver) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Get", Name, Platform)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="InputObject" type="MSFT_OdbcDriver []"></param>
// <param name="PassThru" type="bool "></param>
// <param name="RemovePropertyValue" type="string []"></param>
// <param name="SetPropertyValue" type="string []"></param>

// <param name="cmdletOutput" type="MSFT_OdbcDriver []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_OdbcDriverTask) SetByInputObject( /* IN */ PassThru bool,
	/* IN */ SetPropertyValue []string,
	/* IN */ RemovePropertyValue []string,
	/* IN */ InputObject []MSFT_OdbcDriver,
	/* OUT */ cmdletOutput []MSFT_OdbcDriver) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("SetByInputObject", PassThru, SetPropertyValue, RemovePropertyValue, InputObject)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Name" type="string "></param>
// <param name="PassThru" type="bool "></param>
// <param name="Platform" type="string "></param>
// <param name="RemovePropertyValue" type="string []"></param>
// <param name="SetPropertyValue" type="string []"></param>

// <param name="cmdletOutput" type="MSFT_OdbcDriver []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_OdbcDriverTask) SetByName( /* IN */ PassThru bool,
	/* IN */ SetPropertyValue []string,
	/* IN */ RemovePropertyValue []string,
	/* IN */ Name string,
	/* IN */ Platform string,
	/* OUT */ cmdletOutput []MSFT_OdbcDriver) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("SetByName", PassThru, SetPropertyValue, RemovePropertyValue, Name, Platform)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
