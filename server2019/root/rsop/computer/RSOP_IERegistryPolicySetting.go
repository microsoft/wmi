// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.RSOP.Computer
//////////////////////////////////////////////
package computer

// RSOP_IERegistryPolicySetting struct
type RSOP_IERegistryPolicySetting struct {
	RSOP_PolicySetting

	//
	command string

	//
	currentUser bool

	//
	deleted bool

	//
	registryKey string

	//
	value []uint8

	//
	valueName string

	//
	valueType uint32
}

// Setcommand sets the value of command for the instance
func (instance *RSOP_IERegistryPolicySetting) SetPropertycommand(value string) (err error) {
	return instance.SetProperty("command", value)
}

// Getcommand gets the value of command for the instance
func (instance *RSOP_IERegistryPolicySetting) GetPropertycommand() (value string, err error) {
	retValue, err := instance.GetProperty("command")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetcurrentUser sets the value of currentUser for the instance
func (instance *RSOP_IERegistryPolicySetting) SetPropertycurrentUser(value bool) (err error) {
	return instance.SetProperty("currentUser", value)
}

// GetcurrentUser gets the value of currentUser for the instance
func (instance *RSOP_IERegistryPolicySetting) GetPropertycurrentUser() (value bool, err error) {
	retValue, err := instance.GetProperty("currentUser")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// Setdeleted sets the value of deleted for the instance
func (instance *RSOP_IERegistryPolicySetting) SetPropertydeleted(value bool) (err error) {
	return instance.SetProperty("deleted", value)
}

// Getdeleted gets the value of deleted for the instance
func (instance *RSOP_IERegistryPolicySetting) GetPropertydeleted() (value bool, err error) {
	retValue, err := instance.GetProperty("deleted")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetregistryKey sets the value of registryKey for the instance
func (instance *RSOP_IERegistryPolicySetting) SetPropertyregistryKey(value string) (err error) {
	return instance.SetProperty("registryKey", value)
}

// GetregistryKey gets the value of registryKey for the instance
func (instance *RSOP_IERegistryPolicySetting) GetPropertyregistryKey() (value string, err error) {
	retValue, err := instance.GetProperty("registryKey")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// Setvalue sets the value of value for the instance
func (instance *RSOP_IERegistryPolicySetting) SetPropertyvalue(value []uint8) (err error) {
	return instance.SetProperty("value", value)
}

// Getvalue gets the value of value for the instance
func (instance *RSOP_IERegistryPolicySetting) GetPropertyvalue() (value []uint8, err error) {
	retValue, err := instance.GetProperty("value")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetvalueName sets the value of valueName for the instance
func (instance *RSOP_IERegistryPolicySetting) SetPropertyvalueName(value string) (err error) {
	return instance.SetProperty("valueName", value)
}

// GetvalueName gets the value of valueName for the instance
func (instance *RSOP_IERegistryPolicySetting) GetPropertyvalueName() (value string, err error) {
	retValue, err := instance.GetProperty("valueName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetvalueType sets the value of valueType for the instance
func (instance *RSOP_IERegistryPolicySetting) SetPropertyvalueType(value uint32) (err error) {
	return instance.SetProperty("valueType", value)
}

// GetvalueType gets the value of valueType for the instance
func (instance *RSOP_IERegistryPolicySetting) GetPropertyvalueType() (value uint32, err error) {
	retValue, err := instance.GetProperty("valueType")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
