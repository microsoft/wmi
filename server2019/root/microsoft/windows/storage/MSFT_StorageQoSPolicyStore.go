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

// MSFT_StorageQoSPolicyStore struct
type MSFT_StorageQoSPolicyStore struct {
	cim.WmiInstance

	//
	Id string

	//
	IOPSNormalizationSize uint32
}

// SetId sets the value of Id for the instance
func (instance *MSFT_StorageQoSPolicyStore) SetPropertyId(value string) (err error) {
	return instance.SetProperty("Id", value)
}

// GetId gets the value of Id for the instance
func (instance *MSFT_StorageQoSPolicyStore) GetPropertyId() (value string, err error) {
	retValue, err := instance.GetProperty("Id")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIOPSNormalizationSize sets the value of IOPSNormalizationSize for the instance
func (instance *MSFT_StorageQoSPolicyStore) SetPropertyIOPSNormalizationSize(value uint32) (err error) {
	return instance.SetProperty("IOPSNormalizationSize", value)
}

// GetIOPSNormalizationSize gets the value of IOPSNormalizationSize for the instance
func (instance *MSFT_StorageQoSPolicyStore) GetPropertyIOPSNormalizationSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("IOPSNormalizationSize")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

//

// <param name="Policy" type="MSFT_StorageQoSPolicy "></param>

// <param name="Policy" type="MSFT_StorageQoSPolicy "></param>
// <param name="ReturnValue" type="int32 "></param>
func (instance *MSFT_StorageQoSPolicyStore) CreatePolicy( /* IN/OUT */ Policy MSFT_StorageQoSPolicy) (result int32, err error) {
	retVal, err := instance.InvokeMethod("CreatePolicy")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = int32(retValue)
	return

}

//

// <param name="IOPSNormalizationSize" type="uint32 "></param>

// <param name="ReturnValue" type="int32 "></param>
func (instance *MSFT_StorageQoSPolicyStore) SetAttributes( /* IN */ IOPSNormalizationSize uint32) (result int32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetAttributes", IOPSNormalizationSize)
	if err != nil {
		return
	}
	result = int32(retVal)
	return

}
