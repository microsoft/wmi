// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.DesiredStateConfiguration
//////////////////////////////////////////////
package desiredstateconfiguration

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_DscTimer struct
type MSFT_DscTimer struct {
	cim.WmiInstance
}

//

// <param name="ConsistencyTimerValue" type="uint32 "></param>
// <param name="RefreshTimerValue" type="uint32 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_DscTimer) StartDscTimer( /* IN */ ConsistencyTimerValue uint32,
	/* IN */ RefreshTimerValue uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("StartDscTimer", ConsistencyTimerValue, RefreshTimerValue)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ReportingTimerValue" type="uint32 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_DscTimer) StartDscReportingTimer( /* IN */ ReportingTimerValue uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("StartDscReportingTimer", ReportingTimerValue)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
