// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSNdis_80211_WLanBssId struct
type MSNdis_80211_WLanBssId struct {
	*MSNdis

	//
	Ndis80211Configuration MSNdis_80211_ConfigurationInfo

	//
	Ndis80211InfrastructureMode MSNdis_80211_NetworkInfrastructure

	//
	Ndis80211MacAddress []uint8

	//
	Ndis80211NetworkTypeInUse MSNdis_80211_NetworkType

	//
	Ndis80211Privacy uint32

	//
	Ndis80211Rssi uint32

	//
	Ndis80211SsId []uint8

	//
	Ndis80211SsIdLength uint32

	//
	Ndis80211SupportedRate []uint8

	//
	Ndis80211WLanBssIdLength uint32

	//
	Reserved uint16
}

func NewMSNdis_80211_WLanBssIdEx1(instance *cim.WmiInstance) (newInstance *MSNdis_80211_WLanBssId, err error) {
	tmp, err := NewMSNdisEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSNdis_80211_WLanBssId{
		MSNdis: tmp,
	}
	return
}

func NewMSNdis_80211_WLanBssIdEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSNdis_80211_WLanBssId, err error) {
	tmp, err := NewMSNdisEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSNdis_80211_WLanBssId{
		MSNdis: tmp,
	}
	return
}

// SetNdis80211Configuration sets the value of Ndis80211Configuration for the instance
func (instance *MSNdis_80211_WLanBssId) SetPropertyNdis80211Configuration(value MSNdis_80211_ConfigurationInfo) (err error) {
	return instance.SetProperty("Ndis80211Configuration", (value))
}

// GetNdis80211Configuration gets the value of Ndis80211Configuration for the instance
func (instance *MSNdis_80211_WLanBssId) GetPropertyNdis80211Configuration() (value MSNdis_80211_ConfigurationInfo, err error) {
	retValue, err := instance.GetProperty("Ndis80211Configuration")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(MSNdis_80211_ConfigurationInfo)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " MSNdis_80211_ConfigurationInfo is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = MSNdis_80211_ConfigurationInfo(valuetmp)

	return
}

// SetNdis80211InfrastructureMode sets the value of Ndis80211InfrastructureMode for the instance
func (instance *MSNdis_80211_WLanBssId) SetPropertyNdis80211InfrastructureMode(value MSNdis_80211_NetworkInfrastructure) (err error) {
	return instance.SetProperty("Ndis80211InfrastructureMode", (value))
}

// GetNdis80211InfrastructureMode gets the value of Ndis80211InfrastructureMode for the instance
func (instance *MSNdis_80211_WLanBssId) GetPropertyNdis80211InfrastructureMode() (value MSNdis_80211_NetworkInfrastructure, err error) {
	retValue, err := instance.GetProperty("Ndis80211InfrastructureMode")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(MSNdis_80211_NetworkInfrastructure)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " MSNdis_80211_NetworkInfrastructure is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = MSNdis_80211_NetworkInfrastructure(valuetmp)

	return
}

// SetNdis80211MacAddress sets the value of Ndis80211MacAddress for the instance
func (instance *MSNdis_80211_WLanBssId) SetPropertyNdis80211MacAddress(value []uint8) (err error) {
	return instance.SetProperty("Ndis80211MacAddress", (value))
}

// GetNdis80211MacAddress gets the value of Ndis80211MacAddress for the instance
func (instance *MSNdis_80211_WLanBssId) GetPropertyNdis80211MacAddress() (value []uint8, err error) {
	retValue, err := instance.GetProperty("Ndis80211MacAddress")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint8)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint8(valuetmp))
	}

	return
}

// SetNdis80211NetworkTypeInUse sets the value of Ndis80211NetworkTypeInUse for the instance
func (instance *MSNdis_80211_WLanBssId) SetPropertyNdis80211NetworkTypeInUse(value MSNdis_80211_NetworkType) (err error) {
	return instance.SetProperty("Ndis80211NetworkTypeInUse", (value))
}

// GetNdis80211NetworkTypeInUse gets the value of Ndis80211NetworkTypeInUse for the instance
func (instance *MSNdis_80211_WLanBssId) GetPropertyNdis80211NetworkTypeInUse() (value MSNdis_80211_NetworkType, err error) {
	retValue, err := instance.GetProperty("Ndis80211NetworkTypeInUse")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(MSNdis_80211_NetworkType)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " MSNdis_80211_NetworkType is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = MSNdis_80211_NetworkType(valuetmp)

	return
}

// SetNdis80211Privacy sets the value of Ndis80211Privacy for the instance
func (instance *MSNdis_80211_WLanBssId) SetPropertyNdis80211Privacy(value uint32) (err error) {
	return instance.SetProperty("Ndis80211Privacy", (value))
}

// GetNdis80211Privacy gets the value of Ndis80211Privacy for the instance
func (instance *MSNdis_80211_WLanBssId) GetPropertyNdis80211Privacy() (value uint32, err error) {
	retValue, err := instance.GetProperty("Ndis80211Privacy")
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

// SetNdis80211Rssi sets the value of Ndis80211Rssi for the instance
func (instance *MSNdis_80211_WLanBssId) SetPropertyNdis80211Rssi(value uint32) (err error) {
	return instance.SetProperty("Ndis80211Rssi", (value))
}

// GetNdis80211Rssi gets the value of Ndis80211Rssi for the instance
func (instance *MSNdis_80211_WLanBssId) GetPropertyNdis80211Rssi() (value uint32, err error) {
	retValue, err := instance.GetProperty("Ndis80211Rssi")
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

// SetNdis80211SsId sets the value of Ndis80211SsId for the instance
func (instance *MSNdis_80211_WLanBssId) SetPropertyNdis80211SsId(value []uint8) (err error) {
	return instance.SetProperty("Ndis80211SsId", (value))
}

// GetNdis80211SsId gets the value of Ndis80211SsId for the instance
func (instance *MSNdis_80211_WLanBssId) GetPropertyNdis80211SsId() (value []uint8, err error) {
	retValue, err := instance.GetProperty("Ndis80211SsId")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint8)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint8(valuetmp))
	}

	return
}

// SetNdis80211SsIdLength sets the value of Ndis80211SsIdLength for the instance
func (instance *MSNdis_80211_WLanBssId) SetPropertyNdis80211SsIdLength(value uint32) (err error) {
	return instance.SetProperty("Ndis80211SsIdLength", (value))
}

// GetNdis80211SsIdLength gets the value of Ndis80211SsIdLength for the instance
func (instance *MSNdis_80211_WLanBssId) GetPropertyNdis80211SsIdLength() (value uint32, err error) {
	retValue, err := instance.GetProperty("Ndis80211SsIdLength")
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

// SetNdis80211SupportedRate sets the value of Ndis80211SupportedRate for the instance
func (instance *MSNdis_80211_WLanBssId) SetPropertyNdis80211SupportedRate(value []uint8) (err error) {
	return instance.SetProperty("Ndis80211SupportedRate", (value))
}

// GetNdis80211SupportedRate gets the value of Ndis80211SupportedRate for the instance
func (instance *MSNdis_80211_WLanBssId) GetPropertyNdis80211SupportedRate() (value []uint8, err error) {
	retValue, err := instance.GetProperty("Ndis80211SupportedRate")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint8)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint8(valuetmp))
	}

	return
}

// SetNdis80211WLanBssIdLength sets the value of Ndis80211WLanBssIdLength for the instance
func (instance *MSNdis_80211_WLanBssId) SetPropertyNdis80211WLanBssIdLength(value uint32) (err error) {
	return instance.SetProperty("Ndis80211WLanBssIdLength", (value))
}

// GetNdis80211WLanBssIdLength gets the value of Ndis80211WLanBssIdLength for the instance
func (instance *MSNdis_80211_WLanBssId) GetPropertyNdis80211WLanBssIdLength() (value uint32, err error) {
	retValue, err := instance.GetProperty("Ndis80211WLanBssIdLength")
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

// SetReserved sets the value of Reserved for the instance
func (instance *MSNdis_80211_WLanBssId) SetPropertyReserved(value uint16) (err error) {
	return instance.SetProperty("Reserved", (value))
}

// GetReserved gets the value of Reserved for the instance
func (instance *MSNdis_80211_WLanBssId) GetPropertyReserved() (value uint16, err error) {
	retValue, err := instance.GetProperty("Reserved")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}
