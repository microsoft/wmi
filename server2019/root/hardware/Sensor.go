// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.Hardware
//////////////////////////////////////////////
package hardware

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Sensor struct
type Sensor struct {
	*CIM_Sensor
}

func NewSensorEx1(instance *cim.WmiInstance) (newInstance *Sensor, err error) {
	tmp, err := NewCIM_SensorEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Sensor{
		CIM_Sensor: tmp,
	}
	return
}

func NewSensorEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Sensor, err error) {
	tmp, err := NewCIM_SensorEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Sensor{
		CIM_Sensor: tmp,
	}
	return
}
