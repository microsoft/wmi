// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.StandardCimv2.mlnx
//////////////////////////////////////////////
package mlnx

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// CIM_StatisticalData struct
type CIM_StatisticalData struct {
	*CIM_ManagedElement

	//
	SampleInterval string

	//
	StartStatisticTime string

	//
	StatisticTime string
}

func NewCIM_StatisticalDataEx1(instance *cim.WmiInstance) (newInstance *CIM_StatisticalData, err error) {
	tmp, err := NewCIM_ManagedElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_StatisticalData{
		CIM_ManagedElement: tmp,
	}
	return
}

func NewCIM_StatisticalDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_StatisticalData, err error) {
	tmp, err := NewCIM_ManagedElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_StatisticalData{
		CIM_ManagedElement: tmp,
	}
	return
}

// SetSampleInterval sets the value of SampleInterval for the instance
func (instance *CIM_StatisticalData) SetPropertySampleInterval(value string) (err error) {
	return instance.SetProperty("SampleInterval", (value))
}

// GetSampleInterval gets the value of SampleInterval for the instance
func (instance *CIM_StatisticalData) GetPropertySampleInterval() (value string, err error) {
	retValue, err := instance.GetProperty("SampleInterval")
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

// SetStartStatisticTime sets the value of StartStatisticTime for the instance
func (instance *CIM_StatisticalData) SetPropertyStartStatisticTime(value string) (err error) {
	return instance.SetProperty("StartStatisticTime", (value))
}

// GetStartStatisticTime gets the value of StartStatisticTime for the instance
func (instance *CIM_StatisticalData) GetPropertyStartStatisticTime() (value string, err error) {
	retValue, err := instance.GetProperty("StartStatisticTime")
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

// SetStatisticTime sets the value of StatisticTime for the instance
func (instance *CIM_StatisticalData) SetPropertyStatisticTime(value string) (err error) {
	return instance.SetProperty("StatisticTime", (value))
}

// GetStatisticTime gets the value of StatisticTime for the instance
func (instance *CIM_StatisticalData) GetPropertyStatisticTime() (value string, err error) {
	retValue, err := instance.GetProperty("StatisticTime")
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

//

// <param name="SelectedStatistics" type="string []"></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_StatisticalData) ResetSelectedStats( /* IN */ SelectedStatistics []string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("ResetSelectedStats", SelectedStatistics)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
