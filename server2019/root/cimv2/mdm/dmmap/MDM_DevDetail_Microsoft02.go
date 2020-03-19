// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.CIMV2.mdm.dmmap
//////////////////////////////////////////////
package dmmap

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MDM_DevDetail_Microsoft02 struct
type MDM_DevDetail_Microsoft02 struct {
	*cim.WmiInstance

	//
	CommercializationOperator string

	//
	DeviceName string

	//
	InstanceID string

	//
	LocalTime string

	//
	MobileID string

	//
	OSPlatform string

	//
	ParentID string

	//
	ProcessorArchitecture int32

	//
	ProcessorType int32

	//
	RadioSwV string

	//
	Resolution string
}

func NewMDM_DevDetail_Microsoft02Ex1(instance *cim.WmiInstance) (newInstance *MDM_DevDetail_Microsoft02, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_DevDetail_Microsoft02{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_DevDetail_Microsoft02Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_DevDetail_Microsoft02, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_DevDetail_Microsoft02{
		WmiInstance: tmp,
	}
	return
}

// SetCommercializationOperator sets the value of CommercializationOperator for the instance
func (instance *MDM_DevDetail_Microsoft02) SetPropertyCommercializationOperator(value string) (err error) {
	return instance.SetProperty("CommercializationOperator", value)
}

// GetCommercializationOperator gets the value of CommercializationOperator for the instance
func (instance *MDM_DevDetail_Microsoft02) GetPropertyCommercializationOperator() (value string, err error) {
	retValue, err := instance.GetProperty("CommercializationOperator")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDeviceName sets the value of DeviceName for the instance
func (instance *MDM_DevDetail_Microsoft02) SetPropertyDeviceName(value string) (err error) {
	return instance.SetProperty("DeviceName", value)
}

// GetDeviceName gets the value of DeviceName for the instance
func (instance *MDM_DevDetail_Microsoft02) GetPropertyDeviceName() (value string, err error) {
	retValue, err := instance.GetProperty("DeviceName")
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
func (instance *MDM_DevDetail_Microsoft02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_DevDetail_Microsoft02) GetPropertyInstanceID() (value string, err error) {
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

// SetLocalTime sets the value of LocalTime for the instance
func (instance *MDM_DevDetail_Microsoft02) SetPropertyLocalTime(value string) (err error) {
	return instance.SetProperty("LocalTime", value)
}

// GetLocalTime gets the value of LocalTime for the instance
func (instance *MDM_DevDetail_Microsoft02) GetPropertyLocalTime() (value string, err error) {
	retValue, err := instance.GetProperty("LocalTime")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMobileID sets the value of MobileID for the instance
func (instance *MDM_DevDetail_Microsoft02) SetPropertyMobileID(value string) (err error) {
	return instance.SetProperty("MobileID", value)
}

// GetMobileID gets the value of MobileID for the instance
func (instance *MDM_DevDetail_Microsoft02) GetPropertyMobileID() (value string, err error) {
	retValue, err := instance.GetProperty("MobileID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOSPlatform sets the value of OSPlatform for the instance
func (instance *MDM_DevDetail_Microsoft02) SetPropertyOSPlatform(value string) (err error) {
	return instance.SetProperty("OSPlatform", value)
}

// GetOSPlatform gets the value of OSPlatform for the instance
func (instance *MDM_DevDetail_Microsoft02) GetPropertyOSPlatform() (value string, err error) {
	retValue, err := instance.GetProperty("OSPlatform")
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
func (instance *MDM_DevDetail_Microsoft02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_DevDetail_Microsoft02) GetPropertyParentID() (value string, err error) {
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

// SetProcessorArchitecture sets the value of ProcessorArchitecture for the instance
func (instance *MDM_DevDetail_Microsoft02) SetPropertyProcessorArchitecture(value int32) (err error) {
	return instance.SetProperty("ProcessorArchitecture", value)
}

// GetProcessorArchitecture gets the value of ProcessorArchitecture for the instance
func (instance *MDM_DevDetail_Microsoft02) GetPropertyProcessorArchitecture() (value int32, err error) {
	retValue, err := instance.GetProperty("ProcessorArchitecture")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetProcessorType sets the value of ProcessorType for the instance
func (instance *MDM_DevDetail_Microsoft02) SetPropertyProcessorType(value int32) (err error) {
	return instance.SetProperty("ProcessorType", value)
}

// GetProcessorType gets the value of ProcessorType for the instance
func (instance *MDM_DevDetail_Microsoft02) GetPropertyProcessorType() (value int32, err error) {
	retValue, err := instance.GetProperty("ProcessorType")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRadioSwV sets the value of RadioSwV for the instance
func (instance *MDM_DevDetail_Microsoft02) SetPropertyRadioSwV(value string) (err error) {
	return instance.SetProperty("RadioSwV", value)
}

// GetRadioSwV gets the value of RadioSwV for the instance
func (instance *MDM_DevDetail_Microsoft02) GetPropertyRadioSwV() (value string, err error) {
	retValue, err := instance.GetProperty("RadioSwV")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetResolution sets the value of Resolution for the instance
func (instance *MDM_DevDetail_Microsoft02) SetPropertyResolution(value string) (err error) {
	return instance.SetProperty("Resolution", value)
}

// GetResolution gets the value of Resolution for the instance
func (instance *MDM_DevDetail_Microsoft02) GetPropertyResolution() (value string, err error) {
	retValue, err := instance.GetProperty("Resolution")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
