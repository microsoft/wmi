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

// SamShutdownSamSrv_End struct
type SamShutdownSamSrv_End struct {
	*SamShutdownSamSrv
}

func NewSamShutdownSamSrv_EndEx1(instance *cim.WmiInstance) (newInstance *SamShutdownSamSrv_End, err error) {
	tmp, err := NewSamShutdownSamSrvEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SamShutdownSamSrv_End{
		SamShutdownSamSrv: tmp,
	}
	return
}

func NewSamShutdownSamSrv_EndEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SamShutdownSamSrv_End, err error) {
	tmp, err := NewSamShutdownSamSrvEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SamShutdownSamSrv_End{
		SamShutdownSamSrv: tmp,
	}
	return
}
