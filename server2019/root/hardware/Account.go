// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.Hardware
//////////////////////////////////////////////
package hardware
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Account struct
type Account struct { 
	*CIM_Account
}

	func NewAccountEx1(instance *cim.WmiInstance) (newInstance *Account, err error) {tmp, err := NewCIM_AccountEx1(instance)
		
	if err != nil { return }
	newInstance = &Account {
	CIM_Account: tmp,
	}
	return
	}
	

	func NewAccountEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *Account, err error) {tmp, err := NewCIM_AccountEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &Account {
	CIM_Account: tmp,
	}
	return
	}
	

