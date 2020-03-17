// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.RSOP.Computer
//////////////////////////////////////////////
package computer

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// RSOP_IEFavoriteOrLinkItem struct
type RSOP_IEFavoriteOrLinkItem struct {
	cim.WmiInstance

	//
	iconPath string

	//
	makeAvailableOffline bool

	//
	name string

	//
	order int32

	//
	url string
}

// SeticonPath sets the value of iconPath for the instance
func (instance *RSOP_IEFavoriteOrLinkItem) SetPropertyiconPath(value string) (err error) {
	return instance.SetProperty("iconPath", value)
}

// GeticonPath gets the value of iconPath for the instance
func (instance *RSOP_IEFavoriteOrLinkItem) GetPropertyiconPath() (value string, err error) {
	retValue, err := instance.GetProperty("iconPath")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetmakeAvailableOffline sets the value of makeAvailableOffline for the instance
func (instance *RSOP_IEFavoriteOrLinkItem) SetPropertymakeAvailableOffline(value bool) (err error) {
	return instance.SetProperty("makeAvailableOffline", value)
}

// GetmakeAvailableOffline gets the value of makeAvailableOffline for the instance
func (instance *RSOP_IEFavoriteOrLinkItem) GetPropertymakeAvailableOffline() (value bool, err error) {
	retValue, err := instance.GetProperty("makeAvailableOffline")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// Setname sets the value of name for the instance
func (instance *RSOP_IEFavoriteOrLinkItem) SetPropertyname(value string) (err error) {
	return instance.SetProperty("name", value)
}

// Getname gets the value of name for the instance
func (instance *RSOP_IEFavoriteOrLinkItem) GetPropertyname() (value string, err error) {
	retValue, err := instance.GetProperty("name")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// Setorder sets the value of order for the instance
func (instance *RSOP_IEFavoriteOrLinkItem) SetPropertyorder(value int32) (err error) {
	return instance.SetProperty("order", value)
}

// Getorder gets the value of order for the instance
func (instance *RSOP_IEFavoriteOrLinkItem) GetPropertyorder() (value int32, err error) {
	retValue, err := instance.GetProperty("order")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// Seturl sets the value of url for the instance
func (instance *RSOP_IEFavoriteOrLinkItem) SetPropertyurl(value string) (err error) {
	return instance.SetProperty("url", value)
}

// Geturl gets the value of url for the instance
func (instance *RSOP_IEFavoriteOrLinkItem) GetPropertyurl() (value string, err error) {
	retValue, err := instance.GetProperty("url")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
