// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// ThreadPoolTrace struct
type ThreadPoolTrace struct {
	*ThreadPool
}

func NewThreadPoolTraceEx1(instance *cim.WmiInstance) (newInstance *ThreadPoolTrace, err error) {
	tmp, err := NewThreadPoolEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ThreadPoolTrace{
		ThreadPool: tmp,
	}
	return
}

func NewThreadPoolTraceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ThreadPoolTrace, err error) {
	tmp, err := NewThreadPoolEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ThreadPoolTrace{
		ThreadPool: tmp,
	}
	return
}
