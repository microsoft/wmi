// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSNdis_PortChar struct
type MSNdis_PortChar struct {
	*MSNdis

	//
	Direction uint32

	//
	Flags uint32

	//
	Header MSNdis_ObjectHeader

	//
	MediaConnectState uint32

	//
	PortNumber uint32

	//
	RcvAuthorizationState uint32

	//
	RcvControlState uint32

	//
	RcvLinkSpeed uint64

	//
	SendAuthorizationState uint32

	//
	SendControlState uint32

	//
	Type uint32

	//
	XmitLinkSpeed uint64
}

func NewMSNdis_PortCharEx1(instance *cim.WmiInstance) (newInstance *MSNdis_PortChar, err error) {
	tmp, err := NewMSNdisEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSNdis_PortChar{
		MSNdis: tmp,
	}
	return
}

func NewMSNdis_PortCharEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSNdis_PortChar, err error) {
	tmp, err := NewMSNdisEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSNdis_PortChar{
		MSNdis: tmp,
	}
	return
}

// SetDirection sets the value of Direction for the instance
func (instance *MSNdis_PortChar) SetPropertyDirection(value uint32) (err error) {
	return instance.SetProperty("Direction", (value))
}

// GetDirection gets the value of Direction for the instance
func (instance *MSNdis_PortChar) GetPropertyDirection() (value uint32, err error) {
	retValue, err := instance.GetProperty("Direction")
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

// SetFlags sets the value of Flags for the instance
func (instance *MSNdis_PortChar) SetPropertyFlags(value uint32) (err error) {
	return instance.SetProperty("Flags", (value))
}

// GetFlags gets the value of Flags for the instance
func (instance *MSNdis_PortChar) GetPropertyFlags() (value uint32, err error) {
	retValue, err := instance.GetProperty("Flags")
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

// SetHeader sets the value of Header for the instance
func (instance *MSNdis_PortChar) SetPropertyHeader(value MSNdis_ObjectHeader) (err error) {
	return instance.SetProperty("Header", (value))
}

// GetHeader gets the value of Header for the instance
func (instance *MSNdis_PortChar) GetPropertyHeader() (value MSNdis_ObjectHeader, err error) {
	retValue, err := instance.GetProperty("Header")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(MSNdis_ObjectHeader)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " MSNdis_ObjectHeader is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = MSNdis_ObjectHeader(valuetmp)

	return
}

// SetMediaConnectState sets the value of MediaConnectState for the instance
func (instance *MSNdis_PortChar) SetPropertyMediaConnectState(value uint32) (err error) {
	return instance.SetProperty("MediaConnectState", (value))
}

// GetMediaConnectState gets the value of MediaConnectState for the instance
func (instance *MSNdis_PortChar) GetPropertyMediaConnectState() (value uint32, err error) {
	retValue, err := instance.GetProperty("MediaConnectState")
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

// SetPortNumber sets the value of PortNumber for the instance
func (instance *MSNdis_PortChar) SetPropertyPortNumber(value uint32) (err error) {
	return instance.SetProperty("PortNumber", (value))
}

// GetPortNumber gets the value of PortNumber for the instance
func (instance *MSNdis_PortChar) GetPropertyPortNumber() (value uint32, err error) {
	retValue, err := instance.GetProperty("PortNumber")
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

// SetRcvAuthorizationState sets the value of RcvAuthorizationState for the instance
func (instance *MSNdis_PortChar) SetPropertyRcvAuthorizationState(value uint32) (err error) {
	return instance.SetProperty("RcvAuthorizationState", (value))
}

// GetRcvAuthorizationState gets the value of RcvAuthorizationState for the instance
func (instance *MSNdis_PortChar) GetPropertyRcvAuthorizationState() (value uint32, err error) {
	retValue, err := instance.GetProperty("RcvAuthorizationState")
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

// SetRcvControlState sets the value of RcvControlState for the instance
func (instance *MSNdis_PortChar) SetPropertyRcvControlState(value uint32) (err error) {
	return instance.SetProperty("RcvControlState", (value))
}

// GetRcvControlState gets the value of RcvControlState for the instance
func (instance *MSNdis_PortChar) GetPropertyRcvControlState() (value uint32, err error) {
	retValue, err := instance.GetProperty("RcvControlState")
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

// SetRcvLinkSpeed sets the value of RcvLinkSpeed for the instance
func (instance *MSNdis_PortChar) SetPropertyRcvLinkSpeed(value uint64) (err error) {
	return instance.SetProperty("RcvLinkSpeed", (value))
}

// GetRcvLinkSpeed gets the value of RcvLinkSpeed for the instance
func (instance *MSNdis_PortChar) GetPropertyRcvLinkSpeed() (value uint64, err error) {
	retValue, err := instance.GetProperty("RcvLinkSpeed")
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

// SetSendAuthorizationState sets the value of SendAuthorizationState for the instance
func (instance *MSNdis_PortChar) SetPropertySendAuthorizationState(value uint32) (err error) {
	return instance.SetProperty("SendAuthorizationState", (value))
}

// GetSendAuthorizationState gets the value of SendAuthorizationState for the instance
func (instance *MSNdis_PortChar) GetPropertySendAuthorizationState() (value uint32, err error) {
	retValue, err := instance.GetProperty("SendAuthorizationState")
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

// SetSendControlState sets the value of SendControlState for the instance
func (instance *MSNdis_PortChar) SetPropertySendControlState(value uint32) (err error) {
	return instance.SetProperty("SendControlState", (value))
}

// GetSendControlState gets the value of SendControlState for the instance
func (instance *MSNdis_PortChar) GetPropertySendControlState() (value uint32, err error) {
	retValue, err := instance.GetProperty("SendControlState")
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

// SetType sets the value of Type for the instance
func (instance *MSNdis_PortChar) SetPropertyType(value uint32) (err error) {
	return instance.SetProperty("Type", (value))
}

// GetType gets the value of Type for the instance
func (instance *MSNdis_PortChar) GetPropertyType() (value uint32, err error) {
	retValue, err := instance.GetProperty("Type")
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

// SetXmitLinkSpeed sets the value of XmitLinkSpeed for the instance
func (instance *MSNdis_PortChar) SetPropertyXmitLinkSpeed(value uint64) (err error) {
	return instance.SetProperty("XmitLinkSpeed", (value))
}

// GetXmitLinkSpeed gets the value of XmitLinkSpeed for the instance
func (instance *MSNdis_PortChar) GetPropertyXmitLinkSpeed() (value uint64, err error) {
	retValue, err := instance.GetProperty("XmitLinkSpeed")
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
