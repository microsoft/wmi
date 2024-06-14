// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// WNFSubUnsub struct
type WNFSubUnsub struct {
	*WNFTrace

	//
	Callback uint32

	//
	DeliveryFlags uint32

	//
	NameSub uint32

	//
	RefCount uint32

	//
	StateName uint64

	//
	Subscription uint32
}

func NewWNFSubUnsubEx1(instance *cim.WmiInstance) (newInstance *WNFSubUnsub, err error) {
	tmp, err := NewWNFTraceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &WNFSubUnsub{
		WNFTrace: tmp,
	}
	return
}

func NewWNFSubUnsubEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *WNFSubUnsub, err error) {
	tmp, err := NewWNFTraceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &WNFSubUnsub{
		WNFTrace: tmp,
	}
	return
}

// SetCallback sets the value of Callback for the instance
func (instance *WNFSubUnsub) SetPropertyCallback(value uint32) (err error) {
	return instance.SetProperty("Callback", (value))
}

// GetCallback gets the value of Callback for the instance
func (instance *WNFSubUnsub) GetPropertyCallback() (value uint32, err error) {
	retValue, err := instance.GetProperty("Callback")
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

// SetDeliveryFlags sets the value of DeliveryFlags for the instance
func (instance *WNFSubUnsub) SetPropertyDeliveryFlags(value uint32) (err error) {
	return instance.SetProperty("DeliveryFlags", (value))
}

// GetDeliveryFlags gets the value of DeliveryFlags for the instance
func (instance *WNFSubUnsub) GetPropertyDeliveryFlags() (value uint32, err error) {
	retValue, err := instance.GetProperty("DeliveryFlags")
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

// SetNameSub sets the value of NameSub for the instance
func (instance *WNFSubUnsub) SetPropertyNameSub(value uint32) (err error) {
	return instance.SetProperty("NameSub", (value))
}

// GetNameSub gets the value of NameSub for the instance
func (instance *WNFSubUnsub) GetPropertyNameSub() (value uint32, err error) {
	retValue, err := instance.GetProperty("NameSub")
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

// SetRefCount sets the value of RefCount for the instance
func (instance *WNFSubUnsub) SetPropertyRefCount(value uint32) (err error) {
	return instance.SetProperty("RefCount", (value))
}

// GetRefCount gets the value of RefCount for the instance
func (instance *WNFSubUnsub) GetPropertyRefCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("RefCount")
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

// SetStateName sets the value of StateName for the instance
func (instance *WNFSubUnsub) SetPropertyStateName(value uint64) (err error) {
	return instance.SetProperty("StateName", (value))
}

// GetStateName gets the value of StateName for the instance
func (instance *WNFSubUnsub) GetPropertyStateName() (value uint64, err error) {
	retValue, err := instance.GetProperty("StateName")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetSubscription sets the value of Subscription for the instance
func (instance *WNFSubUnsub) SetPropertySubscription(value uint32) (err error) {
	return instance.SetProperty("Subscription", (value))
}

// GetSubscription gets the value of Subscription for the instance
func (instance *WNFSubUnsub) GetPropertySubscription() (value uint32, err error) {
	retValue, err := instance.GetProperty("Subscription")
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
