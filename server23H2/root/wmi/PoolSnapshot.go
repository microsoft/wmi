// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// PoolSnapshot struct
type PoolSnapshot struct {
	*PoolTrace
}

func NewPoolSnapshotEx1(instance *cim.WmiInstance) (newInstance *PoolSnapshot, err error) {
	tmp, err := NewPoolTraceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &PoolSnapshot{
		PoolTrace: tmp,
	}
	return
}

func NewPoolSnapshotEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *PoolSnapshot, err error) {
	tmp, err := NewPoolTraceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &PoolSnapshot{
		PoolTrace: tmp,
	}
	return
}
