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

// MDM_Win32CompatibilityAppraiser_AppraiserConfigurationDiagnosis02 struct
type MDM_Win32CompatibilityAppraiser_AppraiserConfigurationDiagnosis02 struct {
	*cim.WmiInstance

	//
	AllTargetOsVersionsRequested bool

	//
	AppraiserCodeAndDataVersionsAboveMinimum int32

	//
	CommercialId string

	//
	CommercialIdSetAndValid bool

	//
	InstanceID string

	//
	OsSkuIsValidForAppraiser bool

	//
	ParentID string

	//
	RebootPending bool
}

func NewMDM_Win32CompatibilityAppraiser_AppraiserConfigurationDiagnosis02Ex1(instance *cim.WmiInstance) (newInstance *MDM_Win32CompatibilityAppraiser_AppraiserConfigurationDiagnosis02, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_Win32CompatibilityAppraiser_AppraiserConfigurationDiagnosis02{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_Win32CompatibilityAppraiser_AppraiserConfigurationDiagnosis02Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_Win32CompatibilityAppraiser_AppraiserConfigurationDiagnosis02, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_Win32CompatibilityAppraiser_AppraiserConfigurationDiagnosis02{
		WmiInstance: tmp,
	}
	return
}

// SetAllTargetOsVersionsRequested sets the value of AllTargetOsVersionsRequested for the instance
func (instance *MDM_Win32CompatibilityAppraiser_AppraiserConfigurationDiagnosis02) SetPropertyAllTargetOsVersionsRequested(value bool) (err error) {
	return instance.SetProperty("AllTargetOsVersionsRequested", value)
}

// GetAllTargetOsVersionsRequested gets the value of AllTargetOsVersionsRequested for the instance
func (instance *MDM_Win32CompatibilityAppraiser_AppraiserConfigurationDiagnosis02) GetPropertyAllTargetOsVersionsRequested() (value bool, err error) {
	retValue, err := instance.GetProperty("AllTargetOsVersionsRequested")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAppraiserCodeAndDataVersionsAboveMinimum sets the value of AppraiserCodeAndDataVersionsAboveMinimum for the instance
func (instance *MDM_Win32CompatibilityAppraiser_AppraiserConfigurationDiagnosis02) SetPropertyAppraiserCodeAndDataVersionsAboveMinimum(value int32) (err error) {
	return instance.SetProperty("AppraiserCodeAndDataVersionsAboveMinimum", value)
}

// GetAppraiserCodeAndDataVersionsAboveMinimum gets the value of AppraiserCodeAndDataVersionsAboveMinimum for the instance
func (instance *MDM_Win32CompatibilityAppraiser_AppraiserConfigurationDiagnosis02) GetPropertyAppraiserCodeAndDataVersionsAboveMinimum() (value int32, err error) {
	retValue, err := instance.GetProperty("AppraiserCodeAndDataVersionsAboveMinimum")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCommercialId sets the value of CommercialId for the instance
func (instance *MDM_Win32CompatibilityAppraiser_AppraiserConfigurationDiagnosis02) SetPropertyCommercialId(value string) (err error) {
	return instance.SetProperty("CommercialId", value)
}

// GetCommercialId gets the value of CommercialId for the instance
func (instance *MDM_Win32CompatibilityAppraiser_AppraiserConfigurationDiagnosis02) GetPropertyCommercialId() (value string, err error) {
	retValue, err := instance.GetProperty("CommercialId")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCommercialIdSetAndValid sets the value of CommercialIdSetAndValid for the instance
func (instance *MDM_Win32CompatibilityAppraiser_AppraiserConfigurationDiagnosis02) SetPropertyCommercialIdSetAndValid(value bool) (err error) {
	return instance.SetProperty("CommercialIdSetAndValid", value)
}

// GetCommercialIdSetAndValid gets the value of CommercialIdSetAndValid for the instance
func (instance *MDM_Win32CompatibilityAppraiser_AppraiserConfigurationDiagnosis02) GetPropertyCommercialIdSetAndValid() (value bool, err error) {
	retValue, err := instance.GetProperty("CommercialIdSetAndValid")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_Win32CompatibilityAppraiser_AppraiserConfigurationDiagnosis02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Win32CompatibilityAppraiser_AppraiserConfigurationDiagnosis02) GetPropertyInstanceID() (value string, err error) {
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

// SetOsSkuIsValidForAppraiser sets the value of OsSkuIsValidForAppraiser for the instance
func (instance *MDM_Win32CompatibilityAppraiser_AppraiserConfigurationDiagnosis02) SetPropertyOsSkuIsValidForAppraiser(value bool) (err error) {
	return instance.SetProperty("OsSkuIsValidForAppraiser", value)
}

// GetOsSkuIsValidForAppraiser gets the value of OsSkuIsValidForAppraiser for the instance
func (instance *MDM_Win32CompatibilityAppraiser_AppraiserConfigurationDiagnosis02) GetPropertyOsSkuIsValidForAppraiser() (value bool, err error) {
	retValue, err := instance.GetProperty("OsSkuIsValidForAppraiser")
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
func (instance *MDM_Win32CompatibilityAppraiser_AppraiserConfigurationDiagnosis02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Win32CompatibilityAppraiser_AppraiserConfigurationDiagnosis02) GetPropertyParentID() (value string, err error) {
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

// SetRebootPending sets the value of RebootPending for the instance
func (instance *MDM_Win32CompatibilityAppraiser_AppraiserConfigurationDiagnosis02) SetPropertyRebootPending(value bool) (err error) {
	return instance.SetProperty("RebootPending", value)
}

// GetRebootPending gets the value of RebootPending for the instance
func (instance *MDM_Win32CompatibilityAppraiser_AppraiserConfigurationDiagnosis02) GetPropertyRebootPending() (value bool, err error) {
	retValue, err := instance.GetProperty("RebootPending")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}
