// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.RSOP.Computer
//////////////////////////////////////////////
package computer

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// RSOP_IESecurityZoneSettings struct
type RSOP_IESecurityZoneSettings struct {
	*cim.WmiInstance

	//
	actionValues []string

	//
	currentTemplateLevel uint32

	//
	description string

	//
	displayName string

	//
	flags uint32

	//
	iconPath string

	//
	minimumTemplateLevel uint32

	//
	recommendedTemplateLevel uint32

	//
	rsopID string

	//
	rsopPrecedence int32

	//
	useHKLM bool

	//
	zoneIndex uint32

	//
	zoneMappings []string
}

func NewRSOP_IESecurityZoneSettingsEx1(instance *cim.WmiInstance) (newInstance *RSOP_IESecurityZoneSettings, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &RSOP_IESecurityZoneSettings{
		WmiInstance: tmp,
	}
	return
}

func NewRSOP_IESecurityZoneSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RSOP_IESecurityZoneSettings, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RSOP_IESecurityZoneSettings{
		WmiInstance: tmp,
	}
	return
}

// SetactionValues sets the value of actionValues for the instance
func (instance *RSOP_IESecurityZoneSettings) SetPropertyactionValues(value []string) (err error) {
	return instance.SetProperty("actionValues", value)
}

// GetactionValues gets the value of actionValues for the instance
func (instance *RSOP_IESecurityZoneSettings) GetPropertyactionValues() (value []string, err error) {
	retValue, err := instance.GetProperty("actionValues")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetcurrentTemplateLevel sets the value of currentTemplateLevel for the instance
func (instance *RSOP_IESecurityZoneSettings) SetPropertycurrentTemplateLevel(value uint32) (err error) {
	return instance.SetProperty("currentTemplateLevel", value)
}

// GetcurrentTemplateLevel gets the value of currentTemplateLevel for the instance
func (instance *RSOP_IESecurityZoneSettings) GetPropertycurrentTemplateLevel() (value uint32, err error) {
	retValue, err := instance.GetProperty("currentTemplateLevel")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// Setdescription sets the value of description for the instance
func (instance *RSOP_IESecurityZoneSettings) SetPropertydescription(value string) (err error) {
	return instance.SetProperty("description", value)
}

// Getdescription gets the value of description for the instance
func (instance *RSOP_IESecurityZoneSettings) GetPropertydescription() (value string, err error) {
	retValue, err := instance.GetProperty("description")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetdisplayName sets the value of displayName for the instance
func (instance *RSOP_IESecurityZoneSettings) SetPropertydisplayName(value string) (err error) {
	return instance.SetProperty("displayName", value)
}

// GetdisplayName gets the value of displayName for the instance
func (instance *RSOP_IESecurityZoneSettings) GetPropertydisplayName() (value string, err error) {
	retValue, err := instance.GetProperty("displayName")
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
func (instance *RSOP_IESecurityZoneSettings) SetPropertyflags(value uint32) (err error) {
	return instance.SetProperty("flags", value)
}

// Getflags gets the value of flags for the instance
func (instance *RSOP_IESecurityZoneSettings) GetPropertyflags() (value uint32, err error) {
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

// SeticonPath sets the value of iconPath for the instance
func (instance *RSOP_IESecurityZoneSettings) SetPropertyiconPath(value string) (err error) {
	return instance.SetProperty("iconPath", value)
}

// GeticonPath gets the value of iconPath for the instance
func (instance *RSOP_IESecurityZoneSettings) GetPropertyiconPath() (value string, err error) {
	retValue, err := instance.GetProperty("iconPath")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetminimumTemplateLevel sets the value of minimumTemplateLevel for the instance
func (instance *RSOP_IESecurityZoneSettings) SetPropertyminimumTemplateLevel(value uint32) (err error) {
	return instance.SetProperty("minimumTemplateLevel", value)
}

// GetminimumTemplateLevel gets the value of minimumTemplateLevel for the instance
func (instance *RSOP_IESecurityZoneSettings) GetPropertyminimumTemplateLevel() (value uint32, err error) {
	retValue, err := instance.GetProperty("minimumTemplateLevel")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetrecommendedTemplateLevel sets the value of recommendedTemplateLevel for the instance
func (instance *RSOP_IESecurityZoneSettings) SetPropertyrecommendedTemplateLevel(value uint32) (err error) {
	return instance.SetProperty("recommendedTemplateLevel", value)
}

// GetrecommendedTemplateLevel gets the value of recommendedTemplateLevel for the instance
func (instance *RSOP_IESecurityZoneSettings) GetPropertyrecommendedTemplateLevel() (value uint32, err error) {
	retValue, err := instance.GetProperty("recommendedTemplateLevel")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetrsopID sets the value of rsopID for the instance
func (instance *RSOP_IESecurityZoneSettings) SetPropertyrsopID(value string) (err error) {
	return instance.SetProperty("rsopID", value)
}

// GetrsopID gets the value of rsopID for the instance
func (instance *RSOP_IESecurityZoneSettings) GetPropertyrsopID() (value string, err error) {
	retValue, err := instance.GetProperty("rsopID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetrsopPrecedence sets the value of rsopPrecedence for the instance
func (instance *RSOP_IESecurityZoneSettings) SetPropertyrsopPrecedence(value int32) (err error) {
	return instance.SetProperty("rsopPrecedence", value)
}

// GetrsopPrecedence gets the value of rsopPrecedence for the instance
func (instance *RSOP_IESecurityZoneSettings) GetPropertyrsopPrecedence() (value int32, err error) {
	retValue, err := instance.GetProperty("rsopPrecedence")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetuseHKLM sets the value of useHKLM for the instance
func (instance *RSOP_IESecurityZoneSettings) SetPropertyuseHKLM(value bool) (err error) {
	return instance.SetProperty("useHKLM", value)
}

// GetuseHKLM gets the value of useHKLM for the instance
func (instance *RSOP_IESecurityZoneSettings) GetPropertyuseHKLM() (value bool, err error) {
	retValue, err := instance.GetProperty("useHKLM")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetzoneIndex sets the value of zoneIndex for the instance
func (instance *RSOP_IESecurityZoneSettings) SetPropertyzoneIndex(value uint32) (err error) {
	return instance.SetProperty("zoneIndex", value)
}

// GetzoneIndex gets the value of zoneIndex for the instance
func (instance *RSOP_IESecurityZoneSettings) GetPropertyzoneIndex() (value uint32, err error) {
	retValue, err := instance.GetProperty("zoneIndex")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetzoneMappings sets the value of zoneMappings for the instance
func (instance *RSOP_IESecurityZoneSettings) SetPropertyzoneMappings(value []string) (err error) {
	return instance.SetProperty("zoneMappings", value)
}

// GetzoneMappings gets the value of zoneMappings for the instance
func (instance *RSOP_IESecurityZoneSettings) GetPropertyzoneMappings() (value []string, err error) {
	retValue, err := instance.GetProperty("zoneMappings")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}
