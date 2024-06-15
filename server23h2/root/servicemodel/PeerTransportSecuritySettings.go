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

// PeerTransportSecuritySettings struct
type PeerTransportSecuritySettings struct {
	*cim.WmiInstance

	// The transport credential type of the peer security element.
	CredentialType string
}

func NewPeerTransportSecuritySettingsEx1(instance *cim.WmiInstance) (newInstance *PeerTransportSecuritySettings, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &PeerTransportSecuritySettings{
		WmiInstance: tmp,
	}
	return
}

func NewPeerTransportSecuritySettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *PeerTransportSecuritySettings, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &PeerTransportSecuritySettings{
		WmiInstance: tmp,
	}
	return
}

// SetCredentialType sets the value of CredentialType for the instance
func (instance *PeerTransportSecuritySettings) SetPropertyCredentialType(value string) (err error) {
	return instance.SetProperty("CredentialType", (value))
}

// GetCredentialType gets the value of CredentialType for the instance
func (instance *PeerTransportSecuritySettings) GetPropertyCredentialType() (value string, err error) {
	retValue, err := instance.GetProperty("CredentialType")
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
