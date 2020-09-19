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

// MSFT_WUOperations struct
type MSFT_WUOperations struct {
	*cim.WmiInstance
}

func NewMSFT_WUOperationsEx1(instance *cim.WmiInstance) (newInstance *MSFT_WUOperations, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_WUOperations{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_WUOperationsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_WUOperations, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_WUOperations{
		WmiInstance: tmp,
	}
	return
}

//

// <param name="SearchCriteria" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
// <param name="Updates" type="MSFT_WUUpdate []"></param>
func (instance *MSFT_WUOperations) ScanForUpdates( /* IN */ SearchCriteria string,
	/* OUT */ Updates []MSFT_WUUpdate) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("ScanForUpdates", SearchCriteria)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="DownloadOnly" type="bool "></param>
// <param name="Updates" type="MSFT_WUUpdate []"></param>

// <param name="RebootRequired" type="bool "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_WUOperations) InstallUpdates( /* IN */ Updates []MSFT_WUUpdate,
	/* IN */ DownloadOnly bool,
	/* OUT */ RebootRequired bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("InstallUpdates", Updates, DownloadOnly)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
