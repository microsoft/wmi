// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.RSOP.NSC47E2BF0_19DB_4433_AD42_977167D2A3EF.User
//////////////////////////////////////////////
package user

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// RSOP_IEImportedProgramSettings struct
type RSOP_IEImportedProgramSettings struct {
	*cim.WmiInstance

	//
	policySetting RSOP_IEAKPolicySetting

	//
	programSettings RSOP_IEProgramSettings
}

func NewRSOP_IEImportedProgramSettingsEx1(instance *cim.WmiInstance) (newInstance *RSOP_IEImportedProgramSettings, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &RSOP_IEImportedProgramSettings{
		WmiInstance: tmp,
	}
	return
}

func NewRSOP_IEImportedProgramSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RSOP_IEImportedProgramSettings, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RSOP_IEImportedProgramSettings{
		WmiInstance: tmp,
	}
	return
}

// SetpolicySetting sets the value of policySetting for the instance
func (instance *RSOP_IEImportedProgramSettings) SetPropertypolicySetting(value RSOP_IEAKPolicySetting) (err error) {
	return instance.SetProperty("policySetting", (value))
}

// GetpolicySetting gets the value of policySetting for the instance
func (instance *RSOP_IEImportedProgramSettings) GetPropertypolicySetting() (value RSOP_IEAKPolicySetting, err error) {
	retValue, err := instance.GetProperty("policySetting")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(RSOP_IEAKPolicySetting)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " RSOP_IEAKPolicySetting is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = RSOP_IEAKPolicySetting(valuetmp)

	return
}

// SetprogramSettings sets the value of programSettings for the instance
func (instance *RSOP_IEImportedProgramSettings) SetPropertyprogramSettings(value RSOP_IEProgramSettings) (err error) {
	return instance.SetProperty("programSettings", (value))
}

// GetprogramSettings gets the value of programSettings for the instance
func (instance *RSOP_IEImportedProgramSettings) GetPropertyprogramSettings() (value RSOP_IEProgramSettings, err error) {
	retValue, err := instance.GetProperty("programSettings")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(RSOP_IEProgramSettings)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " RSOP_IEProgramSettings is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = RSOP_IEProgramSettings(valuetmp)

	return
}
