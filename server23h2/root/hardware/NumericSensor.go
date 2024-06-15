// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.Hardware
//////////////////////////////////////////////
package hardware

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// NumericSensor struct
type NumericSensor struct {
	*CIM_NumericSensor
}

func NewNumericSensorEx1(instance *cim.WmiInstance) (newInstance *NumericSensor, err error) {
	tmp, err := NewCIM_NumericSensorEx1(instance)

	if err != nil {
		return
	}
	newInstance = &NumericSensor{
		CIM_NumericSensor: tmp,
	}
	return
}

func NewNumericSensorEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *NumericSensor, err error) {
	tmp, err := NewCIM_NumericSensorEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &NumericSensor{
		CIM_NumericSensor: tmp,
	}
	return
}
