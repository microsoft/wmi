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

// MMCSS_TypeGroup struct
type MMCSS_TypeGroup struct {
	*MMCSSTrace
}

func NewMMCSS_TypeGroupEx1(instance *cim.WmiInstance) (newInstance *MMCSS_TypeGroup, err error) {
	tmp, err := NewMMCSSTraceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MMCSS_TypeGroup{
		MMCSSTrace: tmp,
	}
	return
}

func NewMMCSS_TypeGroupEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MMCSS_TypeGroup, err error) {
	tmp, err := NewMMCSSTraceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MMCSS_TypeGroup{
		MMCSSTrace: tmp,
	}
	return
}
