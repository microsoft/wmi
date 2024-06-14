// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.ServiceModel
//////////////////////////////////////////////
package servicemodel

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// PeerSecuritySettings struct
type PeerSecuritySettings struct {
	*cim.WmiInstance

	// Whether message-level and transport-level security are used by an endpoint configured with the binding.
	Mode string

	// Transport security settings.
	Transport PeerTransportSecuritySettings
}

func NewPeerSecuritySettingsEx1(instance *cim.WmiInstance) (newInstance *PeerSecuritySettings, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &PeerSecuritySettings{
		WmiInstance: tmp,
	}
	return
}

func NewPeerSecuritySettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *PeerSecuritySettings, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &PeerSecuritySettings{
		WmiInstance: tmp,
	}
	return
}

// SetMode sets the value of Mode for the instance
func (instance *PeerSecuritySettings) SetPropertyMode(value string) (err error) {
	return instance.SetProperty("Mode", (value))
}

// GetMode gets the value of Mode for the instance
func (instance *PeerSecuritySettings) GetPropertyMode() (value string, err error) {
	retValue, err := instance.GetProperty("Mode")
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

// SetTransport sets the value of Transport for the instance
func (instance *PeerSecuritySettings) SetPropertyTransport(value PeerTransportSecuritySettings) (err error) {
	return instance.SetProperty("Transport", (value))
}

// GetTransport gets the value of Transport for the instance
func (instance *PeerSecuritySettings) GetPropertyTransport() (value PeerTransportSecuritySettings, err error) {
	retValue, err := instance.GetProperty("Transport")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(PeerTransportSecuritySettings)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " PeerTransportSecuritySettings is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = PeerTransportSecuritySettings(valuetmp)

	return
}
