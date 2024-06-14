// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.RSOP.Computer
//////////////////////////////////////////////
package computer

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// RSOP_IEFavoriteItem struct
type RSOP_IEFavoriteItem struct {
	*RSOP_IEFavoriteOrLinkItem

	//
	folderItem bool

	//
	parentPath string

	//
	rsopID string

	//
	rsopPrecedence uint32

	//
	shortName string
}

func NewRSOP_IEFavoriteItemEx1(instance *cim.WmiInstance) (newInstance *RSOP_IEFavoriteItem, err error) {
	tmp, err := NewRSOP_IEFavoriteOrLinkItemEx1(instance)

	if err != nil {
		return
	}
	newInstance = &RSOP_IEFavoriteItem{
		RSOP_IEFavoriteOrLinkItem: tmp,
	}
	return
}

func NewRSOP_IEFavoriteItemEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RSOP_IEFavoriteItem, err error) {
	tmp, err := NewRSOP_IEFavoriteOrLinkItemEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RSOP_IEFavoriteItem{
		RSOP_IEFavoriteOrLinkItem: tmp,
	}
	return
}

// SetfolderItem sets the value of folderItem for the instance
func (instance *RSOP_IEFavoriteItem) SetPropertyfolderItem(value bool) (err error) {
	return instance.SetProperty("folderItem", (value))
}

// GetfolderItem gets the value of folderItem for the instance
func (instance *RSOP_IEFavoriteItem) GetPropertyfolderItem() (value bool, err error) {
	retValue, err := instance.GetProperty("folderItem")
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

// SetparentPath sets the value of parentPath for the instance
func (instance *RSOP_IEFavoriteItem) SetPropertyparentPath(value string) (err error) {
	return instance.SetProperty("parentPath", (value))
}

// GetparentPath gets the value of parentPath for the instance
func (instance *RSOP_IEFavoriteItem) GetPropertyparentPath() (value string, err error) {
	retValue, err := instance.GetProperty("parentPath")
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

// SetrsopID sets the value of rsopID for the instance
func (instance *RSOP_IEFavoriteItem) SetPropertyrsopID(value string) (err error) {
	return instance.SetProperty("rsopID", (value))
}

// GetrsopID gets the value of rsopID for the instance
func (instance *RSOP_IEFavoriteItem) GetPropertyrsopID() (value string, err error) {
	retValue, err := instance.GetProperty("rsopID")
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

// SetrsopPrecedence sets the value of rsopPrecedence for the instance
func (instance *RSOP_IEFavoriteItem) SetPropertyrsopPrecedence(value uint32) (err error) {
	return instance.SetProperty("rsopPrecedence", (value))
}

// GetrsopPrecedence gets the value of rsopPrecedence for the instance
func (instance *RSOP_IEFavoriteItem) GetPropertyrsopPrecedence() (value uint32, err error) {
	retValue, err := instance.GetProperty("rsopPrecedence")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetshortName sets the value of shortName for the instance
func (instance *RSOP_IEFavoriteItem) SetPropertyshortName(value string) (err error) {
	return instance.SetProperty("shortName", (value))
}

// GetshortName gets the value of shortName for the instance
func (instance *RSOP_IEFavoriteItem) GetPropertyshortName() (value string, err error) {
	retValue, err := instance.GetProperty("shortName")
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
