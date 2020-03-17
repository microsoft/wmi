// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.Wdac
//////////////////////////////////////////////
package wdac

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_OdbcPerfCounterTask struct
type MSFT_OdbcPerfCounterTask struct {
	cim.WmiInstance
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
