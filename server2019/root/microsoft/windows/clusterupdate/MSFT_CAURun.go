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

// MSFT_CAURun struct
type MSFT_CAURun struct {
	cim.WmiInstance

	//
	OrchestratorGuid string
}

// SetOrchestratorGuid sets the value of OrchestratorGuid for the instance
func (instance *MSFT_CAURun) SetPropertyOrchestratorGuid(value string) (err error) {
	return instance.SetProperty("OrchestratorGuid", value)
}

// GetOrchestratorGuid gets the value of OrchestratorGuid for the instance
func (instance *MSFT_CAURun) GetPropertyOrchestratorGuid() (value string, err error) {
	retValue, err := instance.GetProperty("OrchestratorGuid")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
