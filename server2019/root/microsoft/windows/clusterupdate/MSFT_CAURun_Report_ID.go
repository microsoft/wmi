// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.ClusterUpdate
//////////////////////////////////////////////
package clusterupdate

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_CAURun_Report_ID struct
type MSFT_CAURun_Report_ID struct {
	cim.WmiInstance

	//
	Timestamp string
}

// SetTimestamp sets the value of Timestamp for the instance
func (instance *MSFT_CAURun_Report_ID) SetPropertyTimestamp(value string) (err error) {
	return instance.SetProperty("Timestamp", value)
}

// GetTimestamp gets the value of Timestamp for the instance
func (instance *MSFT_CAURun_Report_ID) GetPropertyTimestamp() (value string, err error) {
	retValue, err := instance.GetProperty("Timestamp")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
