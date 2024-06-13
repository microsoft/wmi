// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.SDDC.Management
//////////////////////////////////////////////
package management

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// SDDC_Cluster struct
type SDDC_Cluster struct {
	*cim.WmiInstance

	//
	Alerts []SDDC_Alert

	//
	IsStretch bool

	//
	Jobs []SDDC_Job

	//
	Name string

	//
	StoragePools []string
}

func NewSDDC_ClusterEx1(instance *cim.WmiInstance) (newInstance *SDDC_Cluster, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &SDDC_Cluster{
		WmiInstance: tmp,
	}
	return
}

func NewSDDC_ClusterEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SDDC_Cluster, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SDDC_Cluster{
		WmiInstance: tmp,
	}
	return
}

// SetAlerts sets the value of Alerts for the instance
func (instance *SDDC_Cluster) SetPropertyAlerts(value []SDDC_Alert) (err error) {
	return instance.SetProperty("Alerts", (value))
}

// GetAlerts gets the value of Alerts for the instance
func (instance *SDDC_Cluster) GetPropertyAlerts() (value []SDDC_Alert, err error) {
	retValue, err := instance.GetProperty("Alerts")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(SDDC_Alert)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " SDDC_Alert is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, SDDC_Alert(valuetmp))
	}

	return
}

// SetIsStretch sets the value of IsStretch for the instance
func (instance *SDDC_Cluster) SetPropertyIsStretch(value bool) (err error) {
	return instance.SetProperty("IsStretch", (value))
}

// GetIsStretch gets the value of IsStretch for the instance
func (instance *SDDC_Cluster) GetPropertyIsStretch() (value bool, err error) {
	retValue, err := instance.GetProperty("IsStretch")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetJobs sets the value of Jobs for the instance
func (instance *SDDC_Cluster) SetPropertyJobs(value []SDDC_Job) (err error) {
	return instance.SetProperty("Jobs", (value))
}

// GetJobs gets the value of Jobs for the instance
func (instance *SDDC_Cluster) GetPropertyJobs() (value []SDDC_Job, err error) {
	retValue, err := instance.GetProperty("Jobs")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(SDDC_Job)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " SDDC_Job is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, SDDC_Job(valuetmp))
	}

	return
}

// SetName sets the value of Name for the instance
func (instance *SDDC_Cluster) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *SDDC_Cluster) GetPropertyName() (value string, err error) {
	retValue, err := instance.GetProperty("Name")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetStoragePools sets the value of StoragePools for the instance
func (instance *SDDC_Cluster) SetPropertyStoragePools(value []string) (err error) {
	return instance.SetProperty("StoragePools", (value))
}

// GetStoragePools gets the value of StoragePools for the instance
func (instance *SDDC_Cluster) GetPropertyStoragePools() (value []string, err error) {
	retValue, err := instance.GetProperty("StoragePools")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}

//

// <param name="SeriesName" type="string "></param>
// <param name="TimeFrame" type="uint16 "></param>

// <param name="Metric" type="SDDC_Metric "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *SDDC_Cluster) GetMetrics( /* IN */ SeriesName string,
	/* IN */ TimeFrame uint16,
	/* OUT */ Metric SDDC_Metric) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetMetrics", SeriesName, TimeFrame)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
