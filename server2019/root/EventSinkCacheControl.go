// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root
//////////////////////////////////////////////
package root

// __EventSinkCacheControl struct
type __EventSinkCacheControl struct {
	__CacheControl

	//
	ClearAfter string
}

// SetClearAfter sets the value of ClearAfter for the instance
func (instance *__EventSinkCacheControl) SetPropertyClearAfter(value string) (err error) {
	return instance.SetProperty("ClearAfter", value)
}

// GetClearAfter gets the value of ClearAfter for the instance
func (instance *__EventSinkCacheControl) GetPropertyClearAfter() (value string, err error) {
	retValue, err := instance.GetProperty("ClearAfter")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
