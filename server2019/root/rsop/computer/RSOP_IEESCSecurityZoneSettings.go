// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.RSOP.Computer
//////////////////////////////////////////////
package computer

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// RSOP_IEESCSecurityZoneSettings struct
type RSOP_IEESCSecurityZoneSettings struct {
	*RSOP_IESecurityZoneSettings

	//
	EscEnabled bool
}

func NewRSOP_IEESCSecurityZoneSettingsEx1(instance *cim.WmiInstance) (newInstance *RSOP_IEESCSecurityZoneSettings, err error) {
	tmp, err := NewRSOP_IESecurityZoneSettingsEx1(instance)

	if err != nil {
		return
	}
	newInstance = &RSOP_IEESCSecurityZoneSettings{
		RSOP_IESecurityZoneSettings: tmp,
	}
	return
}

func NewRSOP_IEESCSecurityZoneSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RSOP_IEESCSecurityZoneSettings, err error) {
	tmp, err := NewRSOP_IESecurityZoneSettingsEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RSOP_IEESCSecurityZoneSettings{
		RSOP_IESecurityZoneSettings: tmp,
	}
	return
}

// SetEscEnabled sets the value of EscEnabled for the instance
func (instance *RSOP_IEESCSecurityZoneSettings) SetPropertyEscEnabled(value bool) (err error) {
	return instance.SetProperty("EscEnabled", value)
}

// GetEscEnabled gets the value of EscEnabled for the instance
func (instance *RSOP_IEESCSecurityZoneSettings) GetPropertyEscEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("EscEnabled")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}
