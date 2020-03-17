// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.RSOP.Computer
//////////////////////////////////////////////
package computer

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// RSOP_Session struct
type RSOP_Session struct {
	cim.WmiInstance

	//
	creationTime string

	//
	flags uint32

	//
	id string

	//
	SecurityGroups []string

	//
	Site string

	//
	slowLink bool

	//
	SOM string

	//
	targetName string

	//
	ttlMinutes uint32

	//
	version uint32
}

// SetcreationTime sets the value of creationTime for the instance
func (instance *RSOP_Session) SetPropertycreationTime(value string) (err error) {
	return instance.SetProperty("creationTime", value)
}

// GetcreationTime gets the value of creationTime for the instance
func (instance *RSOP_Session) GetPropertycreationTime() (value string, err error) {
	retValue, err := instance.GetProperty("creationTime")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// Setflags sets the value of flags for the instance
func (instance *RSOP_Session) SetPropertyflags(value uint32) (err error) {
	return instance.SetProperty("flags", value)
}

// Getflags gets the value of flags for the instance
func (instance *RSOP_Session) GetPropertyflags() (value uint32, err error) {
	retValue, err := instance.GetProperty("flags")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// Setid sets the value of id for the instance
func (instance *RSOP_Session) SetPropertyid(value string) (err error) {
	return instance.SetProperty("id", value)
}

// Getid gets the value of id for the instance
func (instance *RSOP_Session) GetPropertyid() (value string, err error) {
	retValue, err := instance.GetProperty("id")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSecurityGroups sets the value of SecurityGroups for the instance
func (instance *RSOP_Session) SetPropertySecurityGroups(value []string) (err error) {
	return instance.SetProperty("SecurityGroups", value)
}

// GetSecurityGroups gets the value of SecurityGroups for the instance
func (instance *RSOP_Session) GetPropertySecurityGroups() (value []string, err error) {
	retValue, err := instance.GetProperty("SecurityGroups")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSite sets the value of Site for the instance
func (instance *RSOP_Session) SetPropertySite(value string) (err error) {
	return instance.SetProperty("Site", value)
}

// GetSite gets the value of Site for the instance
func (instance *RSOP_Session) GetPropertySite() (value string, err error) {
	retValue, err := instance.GetProperty("Site")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetslowLink sets the value of slowLink for the instance
func (instance *RSOP_Session) SetPropertyslowLink(value bool) (err error) {
	return instance.SetProperty("slowLink", value)
}

// GetslowLink gets the value of slowLink for the instance
func (instance *RSOP_Session) GetPropertyslowLink() (value bool, err error) {
	retValue, err := instance.GetProperty("slowLink")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSOM sets the value of SOM for the instance
func (instance *RSOP_Session) SetPropertySOM(value string) (err error) {
	return instance.SetProperty("SOM", value)
}

// GetSOM gets the value of SOM for the instance
func (instance *RSOP_Session) GetPropertySOM() (value string, err error) {
	retValue, err := instance.GetProperty("SOM")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SettargetName sets the value of targetName for the instance
func (instance *RSOP_Session) SetPropertytargetName(value string) (err error) {
	return instance.SetProperty("targetName", value)
}

// GettargetName gets the value of targetName for the instance
func (instance *RSOP_Session) GetPropertytargetName() (value string, err error) {
	retValue, err := instance.GetProperty("targetName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetttlMinutes sets the value of ttlMinutes for the instance
func (instance *RSOP_Session) SetPropertyttlMinutes(value uint32) (err error) {
	return instance.SetProperty("ttlMinutes", value)
}

// GetttlMinutes gets the value of ttlMinutes for the instance
func (instance *RSOP_Session) GetPropertyttlMinutes() (value uint32, err error) {
	retValue, err := instance.GetProperty("ttlMinutes")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// Setversion sets the value of version for the instance
func (instance *RSOP_Session) SetPropertyversion(value uint32) (err error) {
	return instance.SetProperty("version", value)
}

// Getversion gets the value of version for the instance
func (instance *RSOP_Session) GetPropertyversion() (value uint32, err error) {
	retValue, err := instance.GetProperty("version")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
