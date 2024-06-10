// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.
//
// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.CIMV2.power
//
// ////////////////////////////////////////////
package power

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_PowerMeter struct
type CIM_PowerMeter struct {
	*cim.WmiInstance
}

func NewCIM_PowerMeterEx1(instance *cim.WmiInstance) (newInstance *CIM_PowerMeter, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &CIM_PowerMeter{
		WmiInstance: tmp,
	}
	return
}

func NewCIM_PowerMeterEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_PowerMeter, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_PowerMeter{
		WmiInstance: tmp,
	}
	return
}
