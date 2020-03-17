// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.SECURITY
//////////////////////////////////////////////
package security

// __Subject struct
type __Subject struct {
	__SecurityRelatedClass

	//
	Authority string

	//
	EditSecurity bool

	//
	Enabled bool

	//
	ExecuteMethods bool

	//
	Name string

	//
	Permissions int32
}

// SetAuthority sets the value of Authority for the instance
func (instance *__Subject) SetPropertyAuthority(value string) (err error) {
	return instance.SetProperty("Authority", value)
}

// GetAuthority gets the value of Authority for the instance
func (instance *__Subject) GetPropertyAuthority() (value string, err error) {
	retValue, err := instance.GetProperty("Authority")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEditSecurity sets the value of EditSecurity for the instance
func (instance *__Subject) SetPropertyEditSecurity(value bool) (err error) {
	return instance.SetProperty("EditSecurity", value)
}

// GetEditSecurity gets the value of EditSecurity for the instance
func (instance *__Subject) GetPropertyEditSecurity() (value bool, err error) {
	retValue, err := instance.GetProperty("EditSecurity")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEnabled sets the value of Enabled for the instance
func (instance *__Subject) SetPropertyEnabled(value bool) (err error) {
	return instance.SetProperty("Enabled", value)
}

// GetEnabled gets the value of Enabled for the instance
func (instance *__Subject) GetPropertyEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("Enabled")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetExecuteMethods sets the value of ExecuteMethods for the instance
func (instance *__Subject) SetPropertyExecuteMethods(value bool) (err error) {
	return instance.SetProperty("ExecuteMethods", value)
}

// GetExecuteMethods gets the value of ExecuteMethods for the instance
func (instance *__Subject) GetPropertyExecuteMethods() (value bool, err error) {
	retValue, err := instance.GetProperty("ExecuteMethods")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetName sets the value of Name for the instance
func (instance *__Subject) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", value)
}

// GetName gets the value of Name for the instance
func (instance *__Subject) GetPropertyName() (value string, err error) {
	retValue, err := instance.GetProperty("Name")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPermissions sets the value of Permissions for the instance
func (instance *__Subject) SetPropertyPermissions(value int32) (err error) {
	return instance.SetProperty("Permissions", value)
}

// GetPermissions gets the value of Permissions for the instance
func (instance *__Subject) GetPropertyPermissions() (value int32, err error) {
	retValue, err := instance.GetProperty("Permissions")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}
