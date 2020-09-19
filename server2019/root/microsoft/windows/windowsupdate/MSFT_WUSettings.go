// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.Microsoft.Windows.WindowsUpdate
//////////////////////////////////////////////
package windowsupdate

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_WUSettings struct
type MSFT_WUSettings struct {
	*cim.WmiInstance
}

func NewMSFT_WUSettingsEx1(instance *cim.WmiInstance) (newInstance *MSFT_WUSettings, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_WUSettings{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_WUSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_WUSettings, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_WUSettings{
		WmiInstance: tmp,
	}
	return
}

//

// <param name="ReturnValue" type="uint32 "></param>
// <param name="WUAVersion" type="string "></param>
func (instance *MSFT_WUSettings) GetWUAVersion( /* OUT */ WUAVersion string) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetWUAVersion")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="LastUpdateInstallationDate" type="string "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_WUSettings) GetLastUpdateInstallationDate( /* OUT */ LastUpdateInstallationDate string) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetLastUpdateInstallationDate")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="LastScanSuccessDate" type="string "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_WUSettings) GetLastScanSuccessDate( /* OUT */ LastScanSuccessDate string) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetLastScanSuccessDate")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="PendingReboot" type="bool "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_WUSettings) IsPendingReboot( /* OUT */ PendingReboot bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("IsPendingReboot")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
