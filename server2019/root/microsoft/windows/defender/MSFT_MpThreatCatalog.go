// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.Microsoft.Windows.Defender
//////////////////////////////////////////////
package defender

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_MpThreatCatalog struct
type MSFT_MpThreatCatalog struct {
	*BaseStatus

	//
	CategoryID uint8

	//
	SeverityID uint8

	//
	ThreatID int64

	//
	ThreatName string

	//
	TypeID uint8
}

func NewMSFT_MpThreatCatalogEx1(instance *cim.WmiInstance) (newInstance *MSFT_MpThreatCatalog, err error) {
	tmp, err := NewBaseStatusEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_MpThreatCatalog{
		BaseStatus: tmp,
	}
	return
}

func NewMSFT_MpThreatCatalogEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_MpThreatCatalog, err error) {
	tmp, err := NewBaseStatusEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_MpThreatCatalog{
		BaseStatus: tmp,
	}
	return
}

// SetCategoryID sets the value of CategoryID for the instance
func (instance *MSFT_MpThreatCatalog) SetPropertyCategoryID(value uint8) (err error) {
	return instance.SetProperty("CategoryID", (value))
}

// GetCategoryID gets the value of CategoryID for the instance
func (instance *MSFT_MpThreatCatalog) GetPropertyCategoryID() (value uint8, err error) {
	retValue, err := instance.GetProperty("CategoryID")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetSeverityID sets the value of SeverityID for the instance
func (instance *MSFT_MpThreatCatalog) SetPropertySeverityID(value uint8) (err error) {
	return instance.SetProperty("SeverityID", (value))
}

// GetSeverityID gets the value of SeverityID for the instance
func (instance *MSFT_MpThreatCatalog) GetPropertySeverityID() (value uint8, err error) {
	retValue, err := instance.GetProperty("SeverityID")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetThreatID sets the value of ThreatID for the instance
func (instance *MSFT_MpThreatCatalog) SetPropertyThreatID(value int64) (err error) {
	return instance.SetProperty("ThreatID", (value))
}

// GetThreatID gets the value of ThreatID for the instance
func (instance *MSFT_MpThreatCatalog) GetPropertyThreatID() (value int64, err error) {
	retValue, err := instance.GetProperty("ThreatID")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int64(valuetmp)

	return
}

// SetThreatName sets the value of ThreatName for the instance
func (instance *MSFT_MpThreatCatalog) SetPropertyThreatName(value string) (err error) {
	return instance.SetProperty("ThreatName", (value))
}

// GetThreatName gets the value of ThreatName for the instance
func (instance *MSFT_MpThreatCatalog) GetPropertyThreatName() (value string, err error) {
	retValue, err := instance.GetProperty("ThreatName")
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

// SetTypeID sets the value of TypeID for the instance
func (instance *MSFT_MpThreatCatalog) SetPropertyTypeID(value uint8) (err error) {
	return instance.SetProperty("TypeID", (value))
}

// GetTypeID gets the value of TypeID for the instance
func (instance *MSFT_MpThreatCatalog) GetPropertyTypeID() (value uint8, err error) {
	retValue, err := instance.GetProperty("TypeID")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}
