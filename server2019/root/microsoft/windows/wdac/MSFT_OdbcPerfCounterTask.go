// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.Microsoft.Windows.Wdac
//////////////////////////////////////////////
package wdac

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_OdbcPerfCounterTask struct
type MSFT_OdbcPerfCounterTask struct {
	*cim.WmiInstance
}

func NewMSFT_OdbcPerfCounterTaskEx1(instance *cim.WmiInstance) (newInstance *MSFT_OdbcPerfCounterTask, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_OdbcPerfCounterTask{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_OdbcPerfCounterTaskEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_OdbcPerfCounterTask, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_OdbcPerfCounterTask{
		WmiInstance: tmp,
	}
	return
}

//

// <param name="InputObject" type="MSFT_OdbcPerfCounter []"></param>
// <param name="PassThru" type="bool "></param>

// <param name="cmdletOutput" type="MSFT_OdbcPerfCounter []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_OdbcPerfCounterTask) EnableByInputObject( /* IN */ PassThru bool,
	/* IN */ InputObject []MSFT_OdbcPerfCounter,
	/* OUT */ cmdletOutput []MSFT_OdbcPerfCounter) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("EnableByInputObject", PassThru, InputObject)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="PassThru" type="bool "></param>
// <param name="Platform" type="string "></param>

// <param name="cmdletOutput" type="MSFT_OdbcPerfCounter []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_OdbcPerfCounterTask) EnableByPlatform( /* IN */ PassThru bool,
	/* IN */ Platform string,
	/* OUT */ cmdletOutput []MSFT_OdbcPerfCounter) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("EnableByPlatform", PassThru, Platform)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="InputObject" type="MSFT_OdbcPerfCounter []"></param>
// <param name="PassThru" type="bool "></param>

// <param name="cmdletOutput" type="MSFT_OdbcPerfCounter []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_OdbcPerfCounterTask) DisableByInputObject( /* IN */ PassThru bool,
	/* IN */ InputObject []MSFT_OdbcPerfCounter,
	/* OUT */ cmdletOutput []MSFT_OdbcPerfCounter) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("DisableByInputObject", PassThru, InputObject)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="PassThru" type="bool "></param>
// <param name="Platform" type="string "></param>

// <param name="cmdletOutput" type="MSFT_OdbcPerfCounter []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_OdbcPerfCounterTask) DisableByPlatform( /* IN */ PassThru bool,
	/* IN */ Platform string,
	/* OUT */ cmdletOutput []MSFT_OdbcPerfCounter) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("DisableByPlatform", PassThru, Platform)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Platform" type="string "></param>

// <param name="cmdletOutput" type="MSFT_OdbcPerfCounter []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_OdbcPerfCounterTask) Get( /* IN */ Platform string,
	/* OUT */ cmdletOutput []MSFT_OdbcPerfCounter) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Get", Platform)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
