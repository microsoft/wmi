// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_SAProposal struct
type CIM_SAProposal struct { 
	*CIM_ScopedSettingData
}

	func NewCIM_SAProposalEx1(instance *cim.WmiInstance) (newInstance *CIM_SAProposal, err error) {tmp, err := NewCIM_ScopedSettingDataEx1(instance)
		
	if err != nil { return }
	newInstance = &CIM_SAProposal {
	CIM_ScopedSettingData: tmp,
	}
	return
	}
	

	func NewCIM_SAProposalEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *CIM_SAProposal, err error) {tmp, err := NewCIM_ScopedSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &CIM_SAProposal {
	CIM_ScopedSettingData: tmp,
	}
	return
	}
	

