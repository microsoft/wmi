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

// RSOP_IEESCPrivacySettings struct
type RSOP_IEESCPrivacySettings struct {
	*RSOP_IEPrivacySettings

	//
	EscEnabled bool
}

func NewRSOP_IEESCPrivacySettingsEx1(instance *cim.WmiInstance) (newInstance *RSOP_IEESCPrivacySettings, err error) {
	tmp, err := NewRSOP_IEPrivacySettingsEx1(instance)

	if err != nil {
		return
	}
	newInstance = &RSOP_IEESCPrivacySettings{
		RSOP_IEPrivacySettings: tmp,
	}
	return
}

func NewRSOP_IEESCPrivacySettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RSOP_IEESCPrivacySettings, err error) {
	tmp, err := NewRSOP_IEPrivacySettingsEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RSOP_IEESCPrivacySettings{
		RSOP_IEPrivacySettings: tmp,
	}
	return
}

// SetEscEnabled sets the value of EscEnabled for the instance
func (instance *RSOP_IEESCPrivacySettings) SetPropertyEscEnabled(value bool) (err error) {
	return instance.SetProperty("EscEnabled", value)
}

// GetEscEnabled gets the value of EscEnabled for the instance
func (instance *RSOP_IEESCPrivacySettings) GetPropertyEscEnabled() (value bool, err error) {
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
