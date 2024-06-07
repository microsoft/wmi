// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.MSCluster
//////////////////////////////////////////////
package mscluster
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSCluster_EventObjectRemove struct
type MSCluster_EventObjectRemove struct { 
	*MSCluster_Event
}

	func NewMSCluster_EventObjectRemoveEx1(instance *cim.WmiInstance) (newInstance *MSCluster_EventObjectRemove, err error) {tmp, err := NewMSCluster_EventEx1(instance)
		
	if err != nil { return }
	newInstance = &MSCluster_EventObjectRemove {
	MSCluster_Event: tmp,
	}
	return
	}
	

	func NewMSCluster_EventObjectRemoveEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSCluster_EventObjectRemove, err error) {tmp, err := NewMSCluster_EventEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSCluster_EventObjectRemove {
	MSCluster_Event: tmp,
	}
	return
	}
	

