// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.RSOP.Computer
//////////////////////////////////////////////
package computer

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// RSOP_IELinkItem struct
type RSOP_IELinkItem struct {
	*RSOP_IEFavoriteOrLinkItem

	//
	rsopID string

	//
	rsopPrecedence uint32
}

func NewRSOP_IELinkItemEx1(instance *cim.WmiInstance) (newInstance *RSOP_IELinkItem, err error) {
	tmp, err := NewRSOP_IEFavoriteOrLinkItemEx1(instance)

	if err != nil {
		return
	}
	newInstance = &RSOP_IELinkItem{
		RSOP_IEFavoriteOrLinkItem: tmp,
	}
	return
}

func NewRSOP_IELinkItemEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RSOP_IELinkItem, err error) {
	tmp, err := NewRSOP_IEFavoriteOrLinkItemEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RSOP_IELinkItem{
		RSOP_IEFavoriteOrLinkItem: tmp,
	}
	return
}

// SetrsopID sets the value of rsopID for the instance
func (instance *RSOP_IELinkItem) SetPropertyrsopID(value string) (err error) {
	return instance.SetProperty("rsopID", value)
}

// GetrsopID gets the value of rsopID for the instance
func (instance *RSOP_IELinkItem) GetPropertyrsopID() (value string, err error) {
	retValue, err := instance.GetProperty("rsopID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetrsopPrecedence sets the value of rsopPrecedence for the instance
func (instance *RSOP_IELinkItem) SetPropertyrsopPrecedence(value uint32) (err error) {
	return instance.SetProperty("rsopPrecedence", value)
}

// GetrsopPrecedence gets the value of rsopPrecedence for the instance
func (instance *RSOP_IELinkItem) GetPropertyrsopPrecedence() (value uint32, err error) {
	retValue, err := instance.GetProperty("rsopPrecedence")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
