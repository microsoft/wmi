// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.ServiceModel
//
// ////////////////////////////////////////////
package servicemodel

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// TcpTransportBindingElement struct
type TcpTransportBindingElement struct {
	*ConnectionOrientedTransportBindingElement

	// The connection pool settings.
	ConnectionPoolSettings TcpConnectionPoolSettings

	// The extended protection policy used by the server to validate incoming client connections.
	ExtendedProtectionPolicy ExtendedProtectionPolicy

	// The maximum number of queued connection requests that can be pending.
	ListenBacklog int32

	// A boolean value that specifies whether TCP port sharing is enabled for this connection.
	PortSharingEnabled bool

	// A Boolean value that specifies whether Teredo (a technology for addressing clients that are behind firewalls) is enabled.
	TeredoEnabled bool
}

func NewTcpTransportBindingElementEx1(instance *cim.WmiInstance) (newInstance *TcpTransportBindingElement, err error) {
	tmp, err := NewConnectionOrientedTransportBindingElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &TcpTransportBindingElement{
		ConnectionOrientedTransportBindingElement: tmp,
	}
	return
}

func NewTcpTransportBindingElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *TcpTransportBindingElement, err error) {
	tmp, err := NewConnectionOrientedTransportBindingElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &TcpTransportBindingElement{
		ConnectionOrientedTransportBindingElement: tmp,
	}
	return
}

// SetConnectionPoolSettings sets the value of ConnectionPoolSettings for the instance
func (instance *TcpTransportBindingElement) SetPropertyConnectionPoolSettings(value TcpConnectionPoolSettings) (err error) {
	return instance.SetProperty("ConnectionPoolSettings", (value))
}

// GetConnectionPoolSettings gets the value of ConnectionPoolSettings for the instance
func (instance *TcpTransportBindingElement) GetPropertyConnectionPoolSettings() (value TcpConnectionPoolSettings, err error) {
	retValue, err := instance.GetProperty("ConnectionPoolSettings")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(TcpConnectionPoolSettings)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " TcpConnectionPoolSettings is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = TcpConnectionPoolSettings(valuetmp)

	return
}

// SetExtendedProtectionPolicy sets the value of ExtendedProtectionPolicy for the instance
func (instance *TcpTransportBindingElement) SetPropertyExtendedProtectionPolicy(value ExtendedProtectionPolicy) (err error) {
	return instance.SetProperty("ExtendedProtectionPolicy", (value))
}

// GetExtendedProtectionPolicy gets the value of ExtendedProtectionPolicy for the instance
func (instance *TcpTransportBindingElement) GetPropertyExtendedProtectionPolicy() (value ExtendedProtectionPolicy, err error) {
	retValue, err := instance.GetProperty("ExtendedProtectionPolicy")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(ExtendedProtectionPolicy)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " ExtendedProtectionPolicy is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = ExtendedProtectionPolicy(valuetmp)

	return
}

// SetListenBacklog sets the value of ListenBacklog for the instance
func (instance *TcpTransportBindingElement) SetPropertyListenBacklog(value int32) (err error) {
	return instance.SetProperty("ListenBacklog", (value))
}

// GetListenBacklog gets the value of ListenBacklog for the instance
func (instance *TcpTransportBindingElement) GetPropertyListenBacklog() (value int32, err error) {
	retValue, err := instance.GetProperty("ListenBacklog")
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

// SetPortSharingEnabled sets the value of PortSharingEnabled for the instance
func (instance *TcpTransportBindingElement) SetPropertyPortSharingEnabled(value bool) (err error) {
	return instance.SetProperty("PortSharingEnabled", (value))
}

// GetPortSharingEnabled gets the value of PortSharingEnabled for the instance
func (instance *TcpTransportBindingElement) GetPropertyPortSharingEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("PortSharingEnabled")
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

// SetTeredoEnabled sets the value of TeredoEnabled for the instance
func (instance *TcpTransportBindingElement) SetPropertyTeredoEnabled(value bool) (err error) {
	return instance.SetProperty("TeredoEnabled", (value))
}

// GetTeredoEnabled gets the value of TeredoEnabled for the instance
func (instance *TcpTransportBindingElement) GetPropertyTeredoEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("TeredoEnabled")
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
