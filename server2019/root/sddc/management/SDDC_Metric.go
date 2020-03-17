// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.SDDC.Management
//////////////////////////////////////////////
package management

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// SDDC_Metric struct
type SDDC_Metric struct {
	cim.WmiInstance

	//
	CurrentTime string

	//
	Datapoints []SDDC_Datapoint
}

// SetCurrentTime sets the value of CurrentTime for the instance
func (instance *SDDC_Metric) SetPropertyCurrentTime(value string) (err error) {
	return instance.SetProperty("CurrentTime", value)
}

// GetCurrentTime gets the value of CurrentTime for the instance
func (instance *SDDC_Metric) GetPropertyCurrentTime() (value string, err error) {
	retValue, err := instance.GetProperty("CurrentTime")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDatapoints sets the value of Datapoints for the instance
func (instance *SDDC_Metric) SetPropertyDatapoints(value []SDDC_Datapoint) (err error) {
	return instance.SetProperty("Datapoints", value)
}

// GetDatapoints gets the value of Datapoints for the instance
func (instance *SDDC_Metric) GetPropertyDatapoints() (value []SDDC_Datapoint, err error) {
	retValue, err := instance.GetProperty("Datapoints")
	if err != nil {
		return
	}
	value, ok := retValue.([]SDDC_Datapoint)
	if !ok {
		// TODO: Set an error
	}
	return
}
