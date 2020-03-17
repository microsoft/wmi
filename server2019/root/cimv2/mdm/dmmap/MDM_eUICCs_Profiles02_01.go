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

// MDM_eUICCs_Profiles02_01 struct
type MDM_eUICCs_Profiles02_01 struct {
	cim.WmiInstance

	//
	ErrorDetail int32

	//
	InstanceID string

	//
	IsEnabled bool

	//
	MatchingID string

	//
	ParentID string

	//
	PPR1Set bool

	//
	PPR2Set bool

	//
	ServerName string

	//
	State int32
}

// SetErrorDetail sets the value of ErrorDetail for the instance
func (instance *MDM_eUICCs_Profiles02_01) SetPropertyErrorDetail(value int32) (err error) {
	return instance.SetProperty("ErrorDetail", value)
}

// GetErrorDetail gets the value of ErrorDetail for the instance
func (instance *MDM_eUICCs_Profiles02_01) GetPropertyErrorDetail() (value int32, err error) {
	retValue, err := instance.GetProperty("ErrorDetail")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_eUICCs_Profiles02_01) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_eUICCs_Profiles02_01) GetPropertyInstanceID() (value string, err error) {
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

// SetIsEnabled sets the value of IsEnabled for the instance
func (instance *MDM_eUICCs_Profiles02_01) SetPropertyIsEnabled(value bool) (err error) {
	return instance.SetProperty("IsEnabled", value)
}

// GetIsEnabled gets the value of IsEnabled for the instance
func (instance *MDM_eUICCs_Profiles02_01) GetPropertyIsEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("IsEnabled")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMatchingID sets the value of MatchingID for the instance
func (instance *MDM_eUICCs_Profiles02_01) SetPropertyMatchingID(value string) (err error) {
	return instance.SetProperty("MatchingID", value)
}

// GetMatchingID gets the value of MatchingID for the instance
func (instance *MDM_eUICCs_Profiles02_01) GetPropertyMatchingID() (value string, err error) {
	retValue, err := instance.GetProperty("MatchingID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetParentID sets the value of ParentID for the instance
func (instance *MDM_eUICCs_Profiles02_01) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_eUICCs_Profiles02_01) GetPropertyParentID() (value string, err error) {
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

// SetPPR1Set sets the value of PPR1Set for the instance
func (instance *MDM_eUICCs_Profiles02_01) SetPropertyPPR1Set(value bool) (err error) {
	return instance.SetProperty("PPR1Set", value)
}

// GetPPR1Set gets the value of PPR1Set for the instance
func (instance *MDM_eUICCs_Profiles02_01) GetPropertyPPR1Set() (value bool, err error) {
	retValue, err := instance.GetProperty("PPR1Set")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPPR2Set sets the value of PPR2Set for the instance
func (instance *MDM_eUICCs_Profiles02_01) SetPropertyPPR2Set(value bool) (err error) {
	return instance.SetProperty("PPR2Set", value)
}

// GetPPR2Set gets the value of PPR2Set for the instance
func (instance *MDM_eUICCs_Profiles02_01) GetPropertyPPR2Set() (value bool, err error) {
	retValue, err := instance.GetProperty("PPR2Set")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetServerName sets the value of ServerName for the instance
func (instance *MDM_eUICCs_Profiles02_01) SetPropertyServerName(value string) (err error) {
	return instance.SetProperty("ServerName", value)
}

// GetServerName gets the value of ServerName for the instance
func (instance *MDM_eUICCs_Profiles02_01) GetPropertyServerName() (value string, err error) {
	retValue, err := instance.GetProperty("ServerName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetState sets the value of State for the instance
func (instance *MDM_eUICCs_Profiles02_01) SetPropertyState(value int32) (err error) {
	return instance.SetProperty("State", value)
}

// GetState gets the value of State for the instance
func (instance *MDM_eUICCs_Profiles02_01) GetPropertyState() (value int32, err error) {
	retValue, err := instance.GetProperty("State")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}
