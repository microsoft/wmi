// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.Microsoft.Windows.SMB
//////////////////////////////////////////////
package smb

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_SmbSession struct
type MSFT_SmbSession struct {
	*cim.WmiInstance

	//
	ClientComputerName string

	//
	ClientUserName string

	//
	ClusterNodeName string

	//
	Dialect string

	//
	NumOpens uint64

	//
	ScopeName string

	//
	SecondsExists uint32

	//
	SecondsIdle uint32

	//
	SessionId uint64

	//
	SmbInstance SmbSession_SmbInstance

	//
	TransportName string
}

func NewMSFT_SmbSessionEx1(instance *cim.WmiInstance) (newInstance *MSFT_SmbSession, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_SmbSession{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_SmbSessionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_SmbSession, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_SmbSession{
		WmiInstance: tmp,
	}
	return
}

// SetClientComputerName sets the value of ClientComputerName for the instance
func (instance *MSFT_SmbSession) SetPropertyClientComputerName(value string) (err error) {
	return instance.SetProperty("ClientComputerName", (value))
}

// GetClientComputerName gets the value of ClientComputerName for the instance
func (instance *MSFT_SmbSession) GetPropertyClientComputerName() (value string, err error) {
	retValue, err := instance.GetProperty("ClientComputerName")
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

// SetClientUserName sets the value of ClientUserName for the instance
func (instance *MSFT_SmbSession) SetPropertyClientUserName(value string) (err error) {
	return instance.SetProperty("ClientUserName", (value))
}

// GetClientUserName gets the value of ClientUserName for the instance
func (instance *MSFT_SmbSession) GetPropertyClientUserName() (value string, err error) {
	retValue, err := instance.GetProperty("ClientUserName")
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

// SetClusterNodeName sets the value of ClusterNodeName for the instance
func (instance *MSFT_SmbSession) SetPropertyClusterNodeName(value string) (err error) {
	return instance.SetProperty("ClusterNodeName", (value))
}

// GetClusterNodeName gets the value of ClusterNodeName for the instance
func (instance *MSFT_SmbSession) GetPropertyClusterNodeName() (value string, err error) {
	retValue, err := instance.GetProperty("ClusterNodeName")
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

// SetDialect sets the value of Dialect for the instance
func (instance *MSFT_SmbSession) SetPropertyDialect(value string) (err error) {
	return instance.SetProperty("Dialect", (value))
}

// GetDialect gets the value of Dialect for the instance
func (instance *MSFT_SmbSession) GetPropertyDialect() (value string, err error) {
	retValue, err := instance.GetProperty("Dialect")
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

// SetNumOpens sets the value of NumOpens for the instance
func (instance *MSFT_SmbSession) SetPropertyNumOpens(value uint64) (err error) {
	return instance.SetProperty("NumOpens", (value))
}

// GetNumOpens gets the value of NumOpens for the instance
func (instance *MSFT_SmbSession) GetPropertyNumOpens() (value uint64, err error) {
	retValue, err := instance.GetProperty("NumOpens")
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

// SetScopeName sets the value of ScopeName for the instance
func (instance *MSFT_SmbSession) SetPropertyScopeName(value string) (err error) {
	return instance.SetProperty("ScopeName", (value))
}

// GetScopeName gets the value of ScopeName for the instance
func (instance *MSFT_SmbSession) GetPropertyScopeName() (value string, err error) {
	retValue, err := instance.GetProperty("ScopeName")
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

// SetSecondsExists sets the value of SecondsExists for the instance
func (instance *MSFT_SmbSession) SetPropertySecondsExists(value uint32) (err error) {
	return instance.SetProperty("SecondsExists", (value))
}

// GetSecondsExists gets the value of SecondsExists for the instance
func (instance *MSFT_SmbSession) GetPropertySecondsExists() (value uint32, err error) {
	retValue, err := instance.GetProperty("SecondsExists")
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

// SetSecondsIdle sets the value of SecondsIdle for the instance
func (instance *MSFT_SmbSession) SetPropertySecondsIdle(value uint32) (err error) {
	return instance.SetProperty("SecondsIdle", (value))
}

// GetSecondsIdle gets the value of SecondsIdle for the instance
func (instance *MSFT_SmbSession) GetPropertySecondsIdle() (value uint32, err error) {
	retValue, err := instance.GetProperty("SecondsIdle")
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

// SetSessionId sets the value of SessionId for the instance
func (instance *MSFT_SmbSession) SetPropertySessionId(value uint64) (err error) {
	return instance.SetProperty("SessionId", (value))
}

// GetSessionId gets the value of SessionId for the instance
func (instance *MSFT_SmbSession) GetPropertySessionId() (value uint64, err error) {
	retValue, err := instance.GetProperty("SessionId")
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

// SetSmbInstance sets the value of SmbInstance for the instance
func (instance *MSFT_SmbSession) SetPropertySmbInstance(value SmbSession_SmbInstance) (err error) {
	return instance.SetProperty("SmbInstance", (value))
}

// GetSmbInstance gets the value of SmbInstance for the instance
func (instance *MSFT_SmbSession) GetPropertySmbInstance() (value SmbSession_SmbInstance, err error) {
	retValue, err := instance.GetProperty("SmbInstance")
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

	value = SmbSession_SmbInstance(valuetmp)

	return
}

// SetTransportName sets the value of TransportName for the instance
func (instance *MSFT_SmbSession) SetPropertyTransportName(value string) (err error) {
	return instance.SetProperty("TransportName", (value))
}

// GetTransportName gets the value of TransportName for the instance
func (instance *MSFT_SmbSession) GetPropertyTransportName() (value string, err error) {
	retValue, err := instance.GetProperty("TransportName")
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

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_SmbSession) ForceClose() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("ForceClose")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
