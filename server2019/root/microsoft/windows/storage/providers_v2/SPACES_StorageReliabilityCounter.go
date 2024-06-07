// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.Microsoft.Windows.Storage.Providers_v2
//////////////////////////////////////////////
package providers_v2
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// SPACES_StorageReliabilityCounter struct
type SPACES_StorageReliabilityCounter struct { 
	*MSFT_StorageReliabilityCounter
}

	func NewSPACES_StorageReliabilityCounterEx1(instance *cim.WmiInstance) (newInstance *SPACES_StorageReliabilityCounter, err error) {tmp, err := NewMSFT_StorageReliabilityCounterEx1(instance)
		
	if err != nil { return }
	newInstance = &SPACES_StorageReliabilityCounter {
	MSFT_StorageReliabilityCounter: tmp,
	}
	return
	}
	

	func NewSPACES_StorageReliabilityCounterEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *SPACES_StorageReliabilityCounter, err error) {tmp, err := NewMSFT_StorageReliabilityCounterEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &SPACES_StorageReliabilityCounter {
	MSFT_StorageReliabilityCounter: tmp,
	}
	return
	}
	

// 

// <param name="ReturnValue" type="uint32 "></param>
func (instance *SPACES_StorageReliabilityCounter) Reset() (result uint32, err error) {retVal, err := instance.InvokeMethodWithReturn("Reset" );
	if err != nil { return }
	result = uint32(retVal)
	return
	
}


