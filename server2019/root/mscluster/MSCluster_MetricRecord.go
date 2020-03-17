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

// MSCluster_MetricRecord struct
type MSCluster_MetricRecord struct {
	cim.WmiInstance

	//
	TimeStamp string
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
