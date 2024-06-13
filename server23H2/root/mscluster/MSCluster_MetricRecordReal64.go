// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.MSCluster
//////////////////////////////////////////////
package mscluster

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSCluster_MetricRecordReal64 struct
type MSCluster_MetricRecordReal64 struct {
	*MSCluster_MetricRecord

	//
	Value float64
}

func NewMSCluster_MetricRecordReal64Ex1(instance *cim.WmiInstance) (newInstance *MSCluster_MetricRecordReal64, err error) {
	tmp, err := NewMSCluster_MetricRecordEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSCluster_MetricRecordReal64{
		MSCluster_MetricRecord: tmp,
	}
	return
}

func NewMSCluster_MetricRecordReal64Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSCluster_MetricRecordReal64, err error) {
	tmp, err := NewMSCluster_MetricRecordEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSCluster_MetricRecordReal64{
		MSCluster_MetricRecord: tmp,
	}
	return
}

// SetValue sets the value of Value for the instance
func (instance *MSCluster_MetricRecordReal64) SetPropertyValue(value float64) (err error) {
	return instance.SetProperty("Value", (value))
}

// GetValue gets the value of Value for the instance
func (instance *MSCluster_MetricRecordReal64) GetPropertyValue() (value float64, err error) {
	retValue, err := instance.GetProperty("Value")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(float64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " float64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = float64(valuetmp)

	return
}
