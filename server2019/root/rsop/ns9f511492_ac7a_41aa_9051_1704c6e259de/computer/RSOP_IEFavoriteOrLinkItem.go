// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.RSOP.NS9F511492_AC7A_41AA_9051_1704C6E259DE.Computer
//
// ////////////////////////////////////////////
package computer

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// RSOP_IEFavoriteOrLinkItem struct
type RSOP_IEFavoriteOrLinkItem struct {
	*cim.WmiInstance

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

func NewRSOP_IEFavoriteOrLinkItemEx1(instance *cim.WmiInstance) (newInstance *RSOP_IEFavoriteOrLinkItem, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &RSOP_IEFavoriteOrLinkItem{
		WmiInstance: tmp,
	}
	return
}

func NewRSOP_IEFavoriteOrLinkItemEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RSOP_IEFavoriteOrLinkItem, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RSOP_IEFavoriteOrLinkItem{
		WmiInstance: tmp,
	}
	return
}

// SeticonPath sets the value of iconPath for the instance
func (instance *RSOP_IEFavoriteOrLinkItem) SetPropertyiconPath(value string) (err error) {
	return instance.SetProperty("iconPath", (value))
}

// GeticonPath gets the value of iconPath for the instance
func (instance *RSOP_IEFavoriteOrLinkItem) GetPropertyiconPath() (value string, err error) {
	retValue, err := instance.GetProperty("iconPath")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetmakeAvailableOffline sets the value of makeAvailableOffline for the instance
func (instance *RSOP_IEFavoriteOrLinkItem) SetPropertymakeAvailableOffline(value bool) (err error) {
	return instance.SetProperty("makeAvailableOffline", (value))
}

// GetmakeAvailableOffline gets the value of makeAvailableOffline for the instance
func (instance *RSOP_IEFavoriteOrLinkItem) GetPropertymakeAvailableOffline() (value bool, err error) {
	retValue, err := instance.GetProperty("makeAvailableOffline")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// Setname sets the value of name for the instance
func (instance *RSOP_IEFavoriteOrLinkItem) SetPropertyname(value string) (err error) {
	return instance.SetProperty("name", (value))
}

// Getname gets the value of name for the instance
func (instance *RSOP_IEFavoriteOrLinkItem) GetPropertyname() (value string, err error) {
	retValue, err := instance.GetProperty("name")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// Setorder sets the value of order for the instance
func (instance *RSOP_IEFavoriteOrLinkItem) SetPropertyorder(value int32) (err error) {
	return instance.SetProperty("order", (value))
}

// Getorder gets the value of order for the instance
func (instance *RSOP_IEFavoriteOrLinkItem) GetPropertyorder() (value int32, err error) {
	retValue, err := instance.GetProperty("order")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int32(valuetmp)

	return
}

// Seturl sets the value of url for the instance
func (instance *RSOP_IEFavoriteOrLinkItem) SetPropertyurl(value string) (err error) {
	return instance.SetProperty("url", (value))
}

// Geturl gets the value of url for the instance
func (instance *RSOP_IEFavoriteOrLinkItem) GetPropertyurl() (value string, err error) {
	retValue, err := instance.GetProperty("url")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}
