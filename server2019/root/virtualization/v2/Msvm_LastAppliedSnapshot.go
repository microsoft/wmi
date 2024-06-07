// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_LastAppliedSnapshot struct
type Msvm_LastAppliedSnapshot struct { 
	*CIM_LastAppliedSnapshot
}

	func NewMsvm_LastAppliedSnapshotEx1(instance *cim.WmiInstance) (newInstance *Msvm_LastAppliedSnapshot, err error) {tmp, err := NewCIM_LastAppliedSnapshotEx1(instance)
		
	if err != nil { return }
	newInstance = &Msvm_LastAppliedSnapshot {
	CIM_LastAppliedSnapshot: tmp,
	}
	return
	}
	

	func NewMsvm_LastAppliedSnapshotEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *Msvm_LastAppliedSnapshot, err error) {tmp, err := NewCIM_LastAppliedSnapshotEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &Msvm_LastAppliedSnapshot {
	CIM_LastAppliedSnapshot: tmp,
	}
	return
	}
	

