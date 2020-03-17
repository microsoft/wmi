// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2.mdm.dmmap
//////////////////////////////////////////////
package dmmap

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MDM_EnterpriseAPN_01 struct
type MDM_EnterpriseAPN_01 struct {
	cim.WmiInstance

	//
	AlwaysOn bool

	//
	APNName string

	//
	AuthType string

	//
	ClassId string

	//
	Enabled bool

	//
	IccId string

	//
	InstanceID string

	//
	IPType string

	//
	IsAttachAPN bool

	//
	ParentID string

	//
	Password string

	//
	Roaming string

	//
	UserName string
}

// SetAlwaysOn sets the value of AlwaysOn for the instance
func (instance *MDM_EnterpriseAPN_01) SetPropertyAlwaysOn(value bool) (err error) {
	return instance.SetProperty("AlwaysOn", value)
}

// GetAlwaysOn gets the value of AlwaysOn for the instance
func (instance *MDM_EnterpriseAPN_01) GetPropertyAlwaysOn() (value bool, err error) {
	retValue, err := instance.GetProperty("AlwaysOn")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAPNName sets the value of APNName for the instance
func (instance *MDM_EnterpriseAPN_01) SetPropertyAPNName(value string) (err error) {
	return instance.SetProperty("APNName", value)
}

// GetAPNName gets the value of APNName for the instance
func (instance *MDM_EnterpriseAPN_01) GetPropertyAPNName() (value string, err error) {
	retValue, err := instance.GetProperty("APNName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAuthType sets the value of AuthType for the instance
func (instance *MDM_EnterpriseAPN_01) SetPropertyAuthType(value string) (err error) {
	return instance.SetProperty("AuthType", value)
}

// GetAuthType gets the value of AuthType for the instance
func (instance *MDM_EnterpriseAPN_01) GetPropertyAuthType() (value string, err error) {
	retValue, err := instance.GetProperty("AuthType")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetClassId sets the value of ClassId for the instance
func (instance *MDM_EnterpriseAPN_01) SetPropertyClassId(value string) (err error) {
	return instance.SetProperty("ClassId", value)
}

// GetClassId gets the value of ClassId for the instance
func (instance *MDM_EnterpriseAPN_01) GetPropertyClassId() (value string, err error) {
	retValue, err := instance.GetProperty("ClassId")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEnabled sets the value of Enabled for the instance
func (instance *MDM_EnterpriseAPN_01) SetPropertyEnabled(value bool) (err error) {
	return instance.SetProperty("Enabled", value)
}

// GetEnabled gets the value of Enabled for the instance
func (instance *MDM_EnterpriseAPN_01) GetPropertyEnabled() (value bool, err error) {
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

// SetIccId sets the value of IccId for the instance
func (instance *MDM_EnterpriseAPN_01) SetPropertyIccId(value string) (err error) {
	return instance.SetProperty("IccId", value)
}

// GetIccId gets the value of IccId for the instance
func (instance *MDM_EnterpriseAPN_01) GetPropertyIccId() (value string, err error) {
	retValue, err := instance.GetProperty("IccId")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_EnterpriseAPN_01) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_EnterpriseAPN_01) GetPropertyInstanceID() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIPType sets the value of IPType for the instance
func (instance *MDM_EnterpriseAPN_01) SetPropertyIPType(value string) (err error) {
	return instance.SetProperty("IPType", value)
}

// GetIPType gets the value of IPType for the instance
func (instance *MDM_EnterpriseAPN_01) GetPropertyIPType() (value string, err error) {
	retValue, err := instance.GetProperty("IPType")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIsAttachAPN sets the value of IsAttachAPN for the instance
func (instance *MDM_EnterpriseAPN_01) SetPropertyIsAttachAPN(value bool) (err error) {
	return instance.SetProperty("IsAttachAPN", value)
}

// GetIsAttachAPN gets the value of IsAttachAPN for the instance
func (instance *MDM_EnterpriseAPN_01) GetPropertyIsAttachAPN() (value bool, err error) {
	retValue, err := instance.GetProperty("IsAttachAPN")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetParentID sets the value of ParentID for the instance
func (instance *MDM_EnterpriseAPN_01) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_EnterpriseAPN_01) GetPropertyParentID() (value string, err error) {
	retValue, err := instance.GetProperty("ParentID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPassword sets the value of Password for the instance
func (instance *MDM_EnterpriseAPN_01) SetPropertyPassword(value string) (err error) {
	return instance.SetProperty("Password", value)
}

// GetPassword gets the value of Password for the instance
func (instance *MDM_EnterpriseAPN_01) GetPropertyPassword() (value string, err error) {
	retValue, err := instance.GetProperty("Password")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRoaming sets the value of Roaming for the instance
func (instance *MDM_EnterpriseAPN_01) SetPropertyRoaming(value string) (err error) {
	return instance.SetProperty("Roaming", value)
}

// GetRoaming gets the value of Roaming for the instance
func (instance *MDM_EnterpriseAPN_01) GetPropertyRoaming() (value string, err error) {
	retValue, err := instance.GetProperty("Roaming")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUserName sets the value of UserName for the instance
func (instance *MDM_EnterpriseAPN_01) SetPropertyUserName(value string) (err error) {
	return instance.SetProperty("UserName", value)
}

// GetUserName gets the value of UserName for the instance
func (instance *MDM_EnterpriseAPN_01) GetPropertyUserName() (value string, err error) {
	retValue, err := instance.GetProperty("UserName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
