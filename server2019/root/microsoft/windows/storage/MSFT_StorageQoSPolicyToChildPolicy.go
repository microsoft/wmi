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

// MSFT_StorageQoSPolicyToChildPolicy struct
type MSFT_StorageQoSPolicyToChildPolicy struct {
	cim.WmiInstance

	//
	ChildPolicy MSFT_StorageQoSPolicy

	//
	ParentPolicy MSFT_StorageQoSPolicy
}

// SetChildPolicy sets the value of ChildPolicy for the instance
func (instance *MSFT_StorageQoSPolicyToChildPolicy) SetPropertyChildPolicy(value MSFT_StorageQoSPolicy) (err error) {
	return instance.SetProperty("ChildPolicy", value)
}

// GetChildPolicy gets the value of ChildPolicy for the instance
func (instance *MSFT_StorageQoSPolicyToChildPolicy) GetPropertyChildPolicy() (value MSFT_StorageQoSPolicy, err error) {
	retValue, err := instance.GetProperty("ChildPolicy")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_StorageQoSPolicy)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetParentPolicy sets the value of ParentPolicy for the instance
func (instance *MSFT_StorageQoSPolicyToChildPolicy) SetPropertyParentPolicy(value MSFT_StorageQoSPolicy) (err error) {
	return instance.SetProperty("ParentPolicy", value)
}

// GetParentPolicy gets the value of ParentPolicy for the instance
func (instance *MSFT_StorageQoSPolicyToChildPolicy) GetPropertyParentPolicy() (value MSFT_StorageQoSPolicy, err error) {
	retValue, err := instance.GetProperty("ParentPolicy")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_StorageQoSPolicy)
	if !ok {
		// TODO: Set an error
	}
	return
}
