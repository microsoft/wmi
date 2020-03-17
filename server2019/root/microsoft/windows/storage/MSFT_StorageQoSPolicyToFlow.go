// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.Storage
//////////////////////////////////////////////
package storage

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_StorageQoSPolicyToFlow struct
type MSFT_StorageQoSPolicyToFlow struct {
	cim.WmiInstance

	//
	Flow MSFT_StorageQoSFlow

	//
	Policy MSFT_StorageQoSPolicy
}

// SetFlow sets the value of Flow for the instance
func (instance *MSFT_StorageQoSPolicyToFlow) SetPropertyFlow(value MSFT_StorageQoSFlow) (err error) {
	return instance.SetProperty("Flow", value)
}

// GetFlow gets the value of Flow for the instance
func (instance *MSFT_StorageQoSPolicyToFlow) GetPropertyFlow() (value MSFT_StorageQoSFlow, err error) {
	retValue, err := instance.GetProperty("Flow")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_StorageQoSFlow)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPolicy sets the value of Policy for the instance
func (instance *MSFT_StorageQoSPolicyToFlow) SetPropertyPolicy(value MSFT_StorageQoSPolicy) (err error) {
	return instance.SetProperty("Policy", value)
}

// GetPolicy gets the value of Policy for the instance
func (instance *MSFT_StorageQoSPolicyToFlow) GetPropertyPolicy() (value MSFT_StorageQoSPolicy, err error) {
	retValue, err := instance.GetProperty("Policy")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_StorageQoSPolicy)
	if !ok {
		// TODO: Set an error
	}
	return
}
