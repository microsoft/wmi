// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.protectionManagement
//////////////////////////////////////////////
package protectionmanagement

// __Event struct
type __Event struct {
	__IndicationRelated

	//
	SECURITY_DESCRIPTOR []uint8

	//
	TIME_CREATED uint64
}

// SetSECURITY_DESCRIPTOR sets the value of SECURITY_DESCRIPTOR for the instance
func (instance *__Event) SetPropertySECURITY_DESCRIPTOR(value []uint8) (err error) {
	return instance.SetProperty("SECURITY_DESCRIPTOR", value)
}

// GetSECURITY_DESCRIPTOR gets the value of SECURITY_DESCRIPTOR for the instance
func (instance *__Event) GetPropertySECURITY_DESCRIPTOR() (value []uint8, err error) {
	retValue, err := instance.GetProperty("SECURITY_DESCRIPTOR")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTIME_CREATED sets the value of TIME_CREATED for the instance
func (instance *__Event) SetPropertyTIME_CREATED(value uint64) (err error) {
	return instance.SetProperty("TIME_CREATED", value)
}

// GetTIME_CREATED gets the value of TIME_CREATED for the instance
func (instance *__Event) GetPropertyTIME_CREATED() (value uint64, err error) {
	retValue, err := instance.GetProperty("TIME_CREATED")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}
