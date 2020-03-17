// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_PrinterNfcTagTasks struct
type MSFT_PrinterNfcTagTasks struct {
	cim.WmiInstance
}

//

// <param name="Lock" type="bool "></param>
// <param name="SharePath" type="string []"></param>
// <param name="WsdAddress" type="string []"></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_PrinterNfcTagTasks) WriteByManualSpecification( /* IN */ SharePath []string,
	/* IN */ WsdAddress []string,
	/* IN */ Lock bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("WriteByManualSpecification", SharePath, WsdAddress, Lock)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="InputObject" type="MSFT_PrinterNfcTag "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_PrinterNfcTagTasks) WriteByPrinterNfcTag( /* IN */ InputObject MSFT_PrinterNfcTag) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("WriteByPrinterNfcTag", InputObject)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="cmdletOutput" type="MSFT_PrinterNfcTag "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_PrinterNfcTagTasks) Read( /* OUT */ cmdletOutput MSFT_PrinterNfcTag) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Read")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
