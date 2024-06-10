// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.Microsoft.Windows.RemoteAccess.Server
//
// ////////////////////////////////////////////
package server

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// RemoteAccessConnectionLocal struct
type RemoteAccessConnectionLocal struct {
	*cim.WmiInstance

	//
	AuthMethod string

	//
	ClientExternalAddress string

	//
	ClientIPv4Address string

	//
	ClientIPv6Address string

	//
	ConnectionDuration uint32

	//
	ConnectionStartTime string

	//
	ConnectionType string

	//
	HealthStatus string

	//
	HostName string

	//
	TotalBytesIn uint64

	//
	TotalBytesOut uint64

	//
	TransitionTechnology string

	//
	TunnelType string
}

func NewRemoteAccessConnectionLocalEx1(instance *cim.WmiInstance) (newInstance *RemoteAccessConnectionLocal, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &RemoteAccessConnectionLocal{
		WmiInstance: tmp,
	}
	return
}

func NewRemoteAccessConnectionLocalEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RemoteAccessConnectionLocal, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RemoteAccessConnectionLocal{
		WmiInstance: tmp,
	}
	return
}

// SetAuthMethod sets the value of AuthMethod for the instance
func (instance *RemoteAccessConnectionLocal) SetPropertyAuthMethod(value string) (err error) {
	return instance.SetProperty("AuthMethod", (value))
}

// GetAuthMethod gets the value of AuthMethod for the instance
func (instance *RemoteAccessConnectionLocal) GetPropertyAuthMethod() (value string, err error) {
	retValue, err := instance.GetProperty("AuthMethod")
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

// SetClientExternalAddress sets the value of ClientExternalAddress for the instance
func (instance *RemoteAccessConnectionLocal) SetPropertyClientExternalAddress(value string) (err error) {
	return instance.SetProperty("ClientExternalAddress", (value))
}

// GetClientExternalAddress gets the value of ClientExternalAddress for the instance
func (instance *RemoteAccessConnectionLocal) GetPropertyClientExternalAddress() (value string, err error) {
	retValue, err := instance.GetProperty("ClientExternalAddress")
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

// SetClientIPv4Address sets the value of ClientIPv4Address for the instance
func (instance *RemoteAccessConnectionLocal) SetPropertyClientIPv4Address(value string) (err error) {
	return instance.SetProperty("ClientIPv4Address", (value))
}

// GetClientIPv4Address gets the value of ClientIPv4Address for the instance
func (instance *RemoteAccessConnectionLocal) GetPropertyClientIPv4Address() (value string, err error) {
	retValue, err := instance.GetProperty("ClientIPv4Address")
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

// SetClientIPv6Address sets the value of ClientIPv6Address for the instance
func (instance *RemoteAccessConnectionLocal) SetPropertyClientIPv6Address(value string) (err error) {
	return instance.SetProperty("ClientIPv6Address", (value))
}

// GetClientIPv6Address gets the value of ClientIPv6Address for the instance
func (instance *RemoteAccessConnectionLocal) GetPropertyClientIPv6Address() (value string, err error) {
	retValue, err := instance.GetProperty("ClientIPv6Address")
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

// SetConnectionDuration sets the value of ConnectionDuration for the instance
func (instance *RemoteAccessConnectionLocal) SetPropertyConnectionDuration(value uint32) (err error) {
	return instance.SetProperty("ConnectionDuration", (value))
}

// GetConnectionDuration gets the value of ConnectionDuration for the instance
func (instance *RemoteAccessConnectionLocal) GetPropertyConnectionDuration() (value uint32, err error) {
	retValue, err := instance.GetProperty("ConnectionDuration")
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

// SetConnectionStartTime sets the value of ConnectionStartTime for the instance
func (instance *RemoteAccessConnectionLocal) SetPropertyConnectionStartTime(value string) (err error) {
	return instance.SetProperty("ConnectionStartTime", (value))
}

// GetConnectionStartTime gets the value of ConnectionStartTime for the instance
func (instance *RemoteAccessConnectionLocal) GetPropertyConnectionStartTime() (value string, err error) {
	retValue, err := instance.GetProperty("ConnectionStartTime")
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

// SetConnectionType sets the value of ConnectionType for the instance
func (instance *RemoteAccessConnectionLocal) SetPropertyConnectionType(value string) (err error) {
	return instance.SetProperty("ConnectionType", (value))
}

// GetConnectionType gets the value of ConnectionType for the instance
func (instance *RemoteAccessConnectionLocal) GetPropertyConnectionType() (value string, err error) {
	retValue, err := instance.GetProperty("ConnectionType")
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

// SetHealthStatus sets the value of HealthStatus for the instance
func (instance *RemoteAccessConnectionLocal) SetPropertyHealthStatus(value string) (err error) {
	return instance.SetProperty("HealthStatus", (value))
}

// GetHealthStatus gets the value of HealthStatus for the instance
func (instance *RemoteAccessConnectionLocal) GetPropertyHealthStatus() (value string, err error) {
	retValue, err := instance.GetProperty("HealthStatus")
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

// SetHostName sets the value of HostName for the instance
func (instance *RemoteAccessConnectionLocal) SetPropertyHostName(value string) (err error) {
	return instance.SetProperty("HostName", (value))
}

// GetHostName gets the value of HostName for the instance
func (instance *RemoteAccessConnectionLocal) GetPropertyHostName() (value string, err error) {
	retValue, err := instance.GetProperty("HostName")
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

// SetTotalBytesIn sets the value of TotalBytesIn for the instance
func (instance *RemoteAccessConnectionLocal) SetPropertyTotalBytesIn(value uint64) (err error) {
	return instance.SetProperty("TotalBytesIn", (value))
}

// GetTotalBytesIn gets the value of TotalBytesIn for the instance
func (instance *RemoteAccessConnectionLocal) GetPropertyTotalBytesIn() (value uint64, err error) {
	retValue, err := instance.GetProperty("TotalBytesIn")
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

// SetTotalBytesOut sets the value of TotalBytesOut for the instance
func (instance *RemoteAccessConnectionLocal) SetPropertyTotalBytesOut(value uint64) (err error) {
	return instance.SetProperty("TotalBytesOut", (value))
}

// GetTotalBytesOut gets the value of TotalBytesOut for the instance
func (instance *RemoteAccessConnectionLocal) GetPropertyTotalBytesOut() (value uint64, err error) {
	retValue, err := instance.GetProperty("TotalBytesOut")
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

// SetTransitionTechnology sets the value of TransitionTechnology for the instance
func (instance *RemoteAccessConnectionLocal) SetPropertyTransitionTechnology(value string) (err error) {
	return instance.SetProperty("TransitionTechnology", (value))
}

// GetTransitionTechnology gets the value of TransitionTechnology for the instance
func (instance *RemoteAccessConnectionLocal) GetPropertyTransitionTechnology() (value string, err error) {
	retValue, err := instance.GetProperty("TransitionTechnology")
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

// SetTunnelType sets the value of TunnelType for the instance
func (instance *RemoteAccessConnectionLocal) SetPropertyTunnelType(value string) (err error) {
	return instance.SetProperty("TunnelType", (value))
}

// GetTunnelType gets the value of TunnelType for the instance
func (instance *RemoteAccessConnectionLocal) GetPropertyTunnelType() (value string, err error) {
	retValue, err := instance.GetProperty("TunnelType")
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
