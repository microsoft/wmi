// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.Microsoft.Windows.Defender
//////////////////////////////////////////////
package defender
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/base/instance"
)

// MSFT_MpHeartBeat struct
type MSFT_MpHeartBeat struct { 
	*cim.WmiInstance
}

	func NewMSFT_MpHeartBeatEx1(instance *cim.WmiInstance) (newInstance *MSFT_MpHeartBeat, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &MSFT_MpHeartBeat {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewMSFT_MpHeartBeatEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSFT_MpHeartBeat, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSFT_MpHeartBeat {
	WmiInstance: tmp,
	}
	return
	}
	

// 

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_MpHeartBeat) Send() (result uint32, err error) {retVal, err := instance.InvokeMethodWithReturn("Send" );
	if err != nil { return }
	result = uint32(retVal)
	return
	
}


