// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.MSCluster
//////////////////////////////////////////////
package mscluster

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSCluster_MetricRecord struct
type MSCluster_MetricRecord struct {
	*cim.WmiInstance

	//
	TimeStamp string
}

func NewMSCluster_MetricRecordEx1(instance *cim.WmiInstance) (newInstance *MSCluster_MetricRecord, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSCluster_MetricRecord{
		WmiInstance: tmp,
	}
	return
}

func NewMSCluster_MetricRecordEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSCluster_MetricRecord, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSCluster_MetricRecord{
		WmiInstance: tmp,
	}
	return
}

// SetTimeStamp sets the value of TimeStamp for the instance
func (instance *MSCluster_MetricRecord) SetPropertyTimeStamp(value string) (err error) {
	return instance.SetProperty("TimeStamp", value)
}

// GetTimeStamp gets the value of TimeStamp for the instance
func (instance *MSCluster_MetricRecord) GetPropertyTimeStamp() (value string, err error) {
	retValue, err := instance.GetProperty("TimeStamp")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
