// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.MSCluster
//////////////////////////////////////////////
package mscluster

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSCluster_HealthMetric struct
type MSCluster_HealthMetric struct {
	*cim.WmiInstance

	//
	MetricId string

	//
	Records []MSCluster_MetricRecord
}

func NewMSCluster_HealthMetricEx1(instance *cim.WmiInstance) (newInstance *MSCluster_HealthMetric, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSCluster_HealthMetric{
		WmiInstance: tmp,
	}
	return
}

func NewMSCluster_HealthMetricEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSCluster_HealthMetric, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSCluster_HealthMetric{
		WmiInstance: tmp,
	}
	return
}

// SetMetricId sets the value of MetricId for the instance
func (instance *MSCluster_HealthMetric) SetPropertyMetricId(value string) (err error) {
	return instance.SetProperty("MetricId", (value))
}

// GetMetricId gets the value of MetricId for the instance
func (instance *MSCluster_HealthMetric) GetPropertyMetricId() (value string, err error) {
	retValue, err := instance.GetProperty("MetricId")
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

// SetRecords sets the value of Records for the instance
func (instance *MSCluster_HealthMetric) SetPropertyRecords(value []MSCluster_MetricRecord) (err error) {
	return instance.SetProperty("Records", (value))
}

// GetRecords gets the value of Records for the instance
func (instance *MSCluster_HealthMetric) GetPropertyRecords() (value []MSCluster_MetricRecord, err error) {
	retValue, err := instance.GetProperty("Records")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(MSCluster_MetricRecord)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " MSCluster_MetricRecord is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, MSCluster_MetricRecord(valuetmp))
	}

	return
}
