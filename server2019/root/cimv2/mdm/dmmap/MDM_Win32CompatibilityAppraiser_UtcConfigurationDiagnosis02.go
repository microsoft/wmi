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

// MDM_Win32CompatibilityAppraiser_UtcConfigurationDiagnosis02 struct
type MDM_Win32CompatibilityAppraiser_UtcConfigurationDiagnosis02 struct {
	cim.WmiInstance

	//
	CommercialDataOptIn int32

	//
	DiagTrackServiceRunning bool

	//
	InstanceID string

	//
	InternetExplorerTelemetryOptIn int32

	//
	MsaServiceEnabled bool

	//
	ParentID string

	//
	TelemetryOptIn int32
}

// SetCommercialDataOptIn sets the value of CommercialDataOptIn for the instance
func (instance *MDM_Win32CompatibilityAppraiser_UtcConfigurationDiagnosis02) SetPropertyCommercialDataOptIn(value int32) (err error) {
	return instance.SetProperty("CommercialDataOptIn", value)
}

// GetCommercialDataOptIn gets the value of CommercialDataOptIn for the instance
func (instance *MDM_Win32CompatibilityAppraiser_UtcConfigurationDiagnosis02) GetPropertyCommercialDataOptIn() (value int32, err error) {
	retValue, err := instance.GetProperty("CommercialDataOptIn")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDiagTrackServiceRunning sets the value of DiagTrackServiceRunning for the instance
func (instance *MDM_Win32CompatibilityAppraiser_UtcConfigurationDiagnosis02) SetPropertyDiagTrackServiceRunning(value bool) (err error) {
	return instance.SetProperty("DiagTrackServiceRunning", value)
}

// GetDiagTrackServiceRunning gets the value of DiagTrackServiceRunning for the instance
func (instance *MDM_Win32CompatibilityAppraiser_UtcConfigurationDiagnosis02) GetPropertyDiagTrackServiceRunning() (value bool, err error) {
	retValue, err := instance.GetProperty("DiagTrackServiceRunning")
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
func (instance *MDM_Win32CompatibilityAppraiser_UtcConfigurationDiagnosis02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Win32CompatibilityAppraiser_UtcConfigurationDiagnosis02) GetPropertyInstanceID() (value string, err error) {
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

// SetInternetExplorerTelemetryOptIn sets the value of InternetExplorerTelemetryOptIn for the instance
func (instance *MDM_Win32CompatibilityAppraiser_UtcConfigurationDiagnosis02) SetPropertyInternetExplorerTelemetryOptIn(value int32) (err error) {
	return instance.SetProperty("InternetExplorerTelemetryOptIn", value)
}

// GetInternetExplorerTelemetryOptIn gets the value of InternetExplorerTelemetryOptIn for the instance
func (instance *MDM_Win32CompatibilityAppraiser_UtcConfigurationDiagnosis02) GetPropertyInternetExplorerTelemetryOptIn() (value int32, err error) {
	retValue, err := instance.GetProperty("InternetExplorerTelemetryOptIn")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMsaServiceEnabled sets the value of MsaServiceEnabled for the instance
func (instance *MDM_Win32CompatibilityAppraiser_UtcConfigurationDiagnosis02) SetPropertyMsaServiceEnabled(value bool) (err error) {
	return instance.SetProperty("MsaServiceEnabled", value)
}

// GetMsaServiceEnabled gets the value of MsaServiceEnabled for the instance
func (instance *MDM_Win32CompatibilityAppraiser_UtcConfigurationDiagnosis02) GetPropertyMsaServiceEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("MsaServiceEnabled")
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
func (instance *MDM_Win32CompatibilityAppraiser_UtcConfigurationDiagnosis02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Win32CompatibilityAppraiser_UtcConfigurationDiagnosis02) GetPropertyParentID() (value string, err error) {
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

// SetTelemetryOptIn sets the value of TelemetryOptIn for the instance
func (instance *MDM_Win32CompatibilityAppraiser_UtcConfigurationDiagnosis02) SetPropertyTelemetryOptIn(value int32) (err error) {
	return instance.SetProperty("TelemetryOptIn", value)
}

// GetTelemetryOptIn gets the value of TelemetryOptIn for the instance
func (instance *MDM_Win32CompatibilityAppraiser_UtcConfigurationDiagnosis02) GetPropertyTelemetryOptIn() (value int32, err error) {
	retValue, err := instance.GetProperty("TelemetryOptIn")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}
