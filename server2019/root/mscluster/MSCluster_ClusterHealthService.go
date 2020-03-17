// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.MSCluster
//////////////////////////////////////////////
package mscluster

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSCluster_ClusterHealthService struct
type MSCluster_ClusterHealthService struct {
	cim.WmiInstance

	//
	ElementName string
}

// SetElementName sets the value of ElementName for the instance
func (instance *MSCluster_ClusterHealthService) SetPropertyElementName(value string) (err error) {
	return instance.SetProperty("ElementName", value)
}

// GetElementName gets the value of ElementName for the instance
func (instance *MSCluster_ClusterHealthService) GetPropertyElementName() (value string, err error) {
	retValue, err := instance.GetProperty("ElementName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
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
