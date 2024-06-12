// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// ThreadPoolTrace_V2 struct
type ThreadPoolTrace_V2 struct {
	*ThreadPool
}

func NewThreadPoolTrace_V2Ex1(instance *cim.WmiInstance) (newInstance *ThreadPoolTrace_V2, err error) {
	tmp, err := NewThreadPoolEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ThreadPoolTrace_V2{
		ThreadPool: tmp,
	}
	return
}

func NewThreadPoolTrace_V2Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ThreadPoolTrace_V2, err error) {
	tmp, err := NewThreadPoolEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ThreadPoolTrace_V2{
		ThreadPool: tmp,
	}
	return
}
