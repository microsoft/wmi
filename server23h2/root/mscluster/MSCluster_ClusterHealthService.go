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

// MSCluster_ClusterHealthService struct
type MSCluster_ClusterHealthService struct {
	*cim.WmiInstance

	//
	ElementName string
}

func NewMSCluster_ClusterHealthServiceEx1(instance *cim.WmiInstance) (newInstance *MSCluster_ClusterHealthService, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSCluster_ClusterHealthService{
		WmiInstance: tmp,
	}
	return
}

func NewMSCluster_ClusterHealthServiceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSCluster_ClusterHealthService, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSCluster_ClusterHealthService{
		WmiInstance: tmp,
	}
	return
}

// SetElementName sets the value of ElementName for the instance
func (instance *MSCluster_ClusterHealthService) SetPropertyElementName(value string) (err error) {
	return instance.SetProperty("ElementName", (value))
}

// GetElementName gets the value of ElementName for the instance
func (instance *MSCluster_ClusterHealthService) GetPropertyElementName() (value string, err error) {
	retValue, err := instance.GetProperty("ElementName")
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

// <param name="Flags" type="uint32 "></param>
// <param name="MetricName" type="string []"></param>
// <param name="StreamName" type="string "></param>

// <param name="Datapoints" type="MSCluster_HealthMetric []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSCluster_ClusterHealthService) GetMetric( /* IN */ MetricName []string,
	/* IN */ StreamName string,
	/* IN */ Flags uint32,
	/* OUT */ Datapoints []MSCluster_HealthMetric) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetMetric", MetricName, StreamName, Flags)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ReportingKey" type="string "></param>
// <param name="ReportingType" type="string "></param>

// <param name="Faults" type="MSCluster_HealthFault []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSCluster_ClusterHealthService) GetFault( /* IN */ ReportingKey string,
	/* IN */ ReportingType string,
	/* OUT */ Faults []MSCluster_HealthFault) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetFault", ReportingKey, ReportingType)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Flags" type="uint32 "></param>
// <param name="Name" type="string "></param>
// <param name="Parameters" type="string []"></param>
// <param name="Timeout" type="uint32 "></param>

// <param name="CommandId" type="string "></param>
// <param name="ReturnValue" type="uint32 "></param>
// <param name="Status" type="int32 "></param>
func (instance *MSCluster_ClusterHealthService) InvokeCommand( /* IN */ Name string,
	/* IN */ Flags uint32,
	/* IN */ Parameters []string,
	/* IN */ Timeout uint32,
	/* OUT */ CommandId string,
	/* OUT */ Status int32) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("InvokeCommand", Name, Flags, Parameters, Timeout)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
