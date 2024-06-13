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

// WNFCallback struct
type WNFCallback struct {
	*WNFTrace

	//
	Callback uint32

	//
	ChangeStamp uint32

	//
	DeliveryFlags uint32

	//
	NameSub uint32

	//
	Return uint32

	//
	StateName uint64

	//
	Subscription uint32
}

func NewWNFCallbackEx1(instance *cim.WmiInstance) (newInstance *WNFCallback, err error) {
	tmp, err := NewWNFTraceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &WNFCallback{
		WNFTrace: tmp,
	}
	return
}

func NewWNFCallbackEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *WNFCallback, err error) {
	tmp, err := NewWNFTraceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &WNFCallback{
		WNFTrace: tmp,
	}
	return
}

// SetCallback sets the value of Callback for the instance
func (instance *WNFCallback) SetPropertyCallback(value uint32) (err error) {
	return instance.SetProperty("Callback", (value))
}

// GetCallback gets the value of Callback for the instance
func (instance *WNFCallback) GetPropertyCallback() (value uint32, err error) {
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

// SetChangeStamp sets the value of ChangeStamp for the instance
func (instance *WNFCallback) SetPropertyChangeStamp(value uint32) (err error) {
	return instance.SetProperty("ChangeStamp", (value))
}

// GetChangeStamp gets the value of ChangeStamp for the instance
func (instance *WNFCallback) GetPropertyChangeStamp() (value uint32, err error) {
	retValue, err := instance.GetProperty("ChangeStamp")
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
func (instance *WNFCallback) SetPropertyDeliveryFlags(value uint32) (err error) {
	return instance.SetProperty("DeliveryFlags", (value))
}

// GetDeliveryFlags gets the value of DeliveryFlags for the instance
func (instance *WNFCallback) GetPropertyDeliveryFlags() (value uint32, err error) {
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
func (instance *WNFCallback) SetPropertyNameSub(value uint32) (err error) {
	return instance.SetProperty("NameSub", (value))
}

// GetNameSub gets the value of NameSub for the instance
func (instance *WNFCallback) GetPropertyNameSub() (value uint32, err error) {
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

// SetReturn sets the value of Return for the instance
func (instance *WNFCallback) SetPropertyReturn(value uint32) (err error) {
	return instance.SetProperty("Return", (value))
}

// GetReturn gets the value of Return for the instance
func (instance *WNFCallback) GetPropertyReturn() (value uint32, err error) {
	retValue, err := instance.GetProperty("Return")
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
func (instance *WNFCallback) SetPropertyStateName(value uint64) (err error) {
	return instance.SetProperty("StateName", (value))
}

// GetStateName gets the value of StateName for the instance
func (instance *WNFCallback) GetPropertyStateName() (value uint64, err error) {
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
func (instance *WNFCallback) SetPropertySubscription(value uint32) (err error) {
	return instance.SetProperty("Subscription", (value))
}

// GetSubscription gets the value of Subscription for the instance
func (instance *WNFCallback) GetPropertySubscription() (value uint32, err error) {
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
