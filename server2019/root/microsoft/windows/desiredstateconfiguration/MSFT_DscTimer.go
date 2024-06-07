// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.Microsoft.Windows.DesiredStateConfiguration
//////////////////////////////////////////////
package desiredstateconfiguration
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/base/instance"
)

// MSFT_DscTimer struct
type MSFT_DscTimer struct { 
	*cim.WmiInstance
}

	func NewMSFT_DscTimerEx1(instance *cim.WmiInstance) (newInstance *MSFT_DscTimer, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &MSFT_DscTimer {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewMSFT_DscTimerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSFT_DscTimer, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSFT_DscTimer {
	WmiInstance: tmp,
	}
	return
	}
	

// 

// <param name="ConsistencyTimerValue" type="uint32 "></param>
// <param name="RefreshTimerValue" type="uint32 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_DscTimer) StartDscTimer( /* IN */ ConsistencyTimerValue uint32,
 /* IN */ RefreshTimerValue uint32) (result uint32, err error) {retVal, err := instance.InvokeMethodWithReturn("StartDscTimer" , ConsistencyTimerValue, RefreshTimerValue);
	if err != nil { return }
	result = uint32(retVal)
	return
	
}


// 

// <param name="ReportingTimerValue" type="uint32 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_DscTimer) StartDscReportingTimer( /* IN */ ReportingTimerValue uint32) (result uint32, err error) {retVal, err := instance.InvokeMethodWithReturn("StartDscReportingTimer" , ReportingTimerValue);
	if err != nil { return }
	result = uint32(retVal)
	return
	
}


