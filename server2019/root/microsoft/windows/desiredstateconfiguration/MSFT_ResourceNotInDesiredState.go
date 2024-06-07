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
)

// MSFT_ResourceNotInDesiredState struct
type MSFT_ResourceNotInDesiredState struct { 
	*MSFT_DSCResource
}

	func NewMSFT_ResourceNotInDesiredStateEx1(instance *cim.WmiInstance) (newInstance *MSFT_ResourceNotInDesiredState, err error) {tmp, err := NewMSFT_DSCResourceEx1(instance)
		
	if err != nil { return }
	newInstance = &MSFT_ResourceNotInDesiredState {
	MSFT_DSCResource: tmp,
	}
	return
	}
	

	func NewMSFT_ResourceNotInDesiredStateEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSFT_ResourceNotInDesiredState, err error) {tmp, err := NewMSFT_DSCResourceEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSFT_ResourceNotInDesiredState {
	MSFT_DSCResource: tmp,
	}
	return
	}
	

