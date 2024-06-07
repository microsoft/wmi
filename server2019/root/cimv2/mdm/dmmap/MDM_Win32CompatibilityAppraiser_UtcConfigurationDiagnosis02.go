// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.CIMV2.mdm.dmmap
//////////////////////////////////////////////
package dmmap
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/base/instance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// MDM_Win32CompatibilityAppraiser_UtcConfigurationDiagnosis02 struct
type MDM_Win32CompatibilityAppraiser_UtcConfigurationDiagnosis02 struct { 
	*cim.WmiInstance

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

	func NewMDM_Win32CompatibilityAppraiser_UtcConfigurationDiagnosis02Ex1(instance *cim.WmiInstance) (newInstance *MDM_Win32CompatibilityAppraiser_UtcConfigurationDiagnosis02, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &MDM_Win32CompatibilityAppraiser_UtcConfigurationDiagnosis02 {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewMDM_Win32CompatibilityAppraiser_UtcConfigurationDiagnosis02Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MDM_Win32CompatibilityAppraiser_UtcConfigurationDiagnosis02, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MDM_Win32CompatibilityAppraiser_UtcConfigurationDiagnosis02 {
	WmiInstance: tmp,
	}
	return
	}
	

// SetCommercialDataOptIn sets the value of CommercialDataOptIn for the instance
func (instance *MDM_Win32CompatibilityAppraiser_UtcConfigurationDiagnosis02) SetPropertyCommercialDataOptIn(value int32) (err error) { 
    return instance.SetProperty("CommercialDataOptIn", (value))
}

// GetCommercialDataOptIn gets the value of CommercialDataOptIn for the instance
func (instance *MDM_Win32CompatibilityAppraiser_UtcConfigurationDiagnosis02) GetPropertyCommercialDataOptIn()(value int32, err error) { 
    retValue, err := instance.GetProperty("CommercialDataOptIn")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(int32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = int32(valuetmp)

    return
}

// SetDiagTrackServiceRunning sets the value of DiagTrackServiceRunning for the instance
func (instance *MDM_Win32CompatibilityAppraiser_UtcConfigurationDiagnosis02) SetPropertyDiagTrackServiceRunning(value bool) (err error) { 
    return instance.SetProperty("DiagTrackServiceRunning", (value))
}

// GetDiagTrackServiceRunning gets the value of DiagTrackServiceRunning for the instance
func (instance *MDM_Win32CompatibilityAppraiser_UtcConfigurationDiagnosis02) GetPropertyDiagTrackServiceRunning()(value bool, err error) { 
    retValue, err := instance.GetProperty("DiagTrackServiceRunning")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(bool); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = bool(valuetmp)

    return
}

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_Win32CompatibilityAppraiser_UtcConfigurationDiagnosis02) SetPropertyInstanceID(value string) (err error) { 
    return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Win32CompatibilityAppraiser_UtcConfigurationDiagnosis02) GetPropertyInstanceID()(value string, err error) { 
    retValue, err := instance.GetProperty("InstanceID")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(string); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = string(valuetmp)

    return
}

// SetInternetExplorerTelemetryOptIn sets the value of InternetExplorerTelemetryOptIn for the instance
func (instance *MDM_Win32CompatibilityAppraiser_UtcConfigurationDiagnosis02) SetPropertyInternetExplorerTelemetryOptIn(value int32) (err error) { 
    return instance.SetProperty("InternetExplorerTelemetryOptIn", (value))
}

// GetInternetExplorerTelemetryOptIn gets the value of InternetExplorerTelemetryOptIn for the instance
func (instance *MDM_Win32CompatibilityAppraiser_UtcConfigurationDiagnosis02) GetPropertyInternetExplorerTelemetryOptIn()(value int32, err error) { 
    retValue, err := instance.GetProperty("InternetExplorerTelemetryOptIn")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(int32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = int32(valuetmp)

    return
}

// SetMsaServiceEnabled sets the value of MsaServiceEnabled for the instance
func (instance *MDM_Win32CompatibilityAppraiser_UtcConfigurationDiagnosis02) SetPropertyMsaServiceEnabled(value bool) (err error) { 
    return instance.SetProperty("MsaServiceEnabled", (value))
}

// GetMsaServiceEnabled gets the value of MsaServiceEnabled for the instance
func (instance *MDM_Win32CompatibilityAppraiser_UtcConfigurationDiagnosis02) GetPropertyMsaServiceEnabled()(value bool, err error) { 
    retValue, err := instance.GetProperty("MsaServiceEnabled")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(bool); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = bool(valuetmp)

    return
}

// SetParentID sets the value of ParentID for the instance
func (instance *MDM_Win32CompatibilityAppraiser_UtcConfigurationDiagnosis02) SetPropertyParentID(value string) (err error) { 
    return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Win32CompatibilityAppraiser_UtcConfigurationDiagnosis02) GetPropertyParentID()(value string, err error) { 
    retValue, err := instance.GetProperty("ParentID")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(string); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = string(valuetmp)

    return
}

// SetTelemetryOptIn sets the value of TelemetryOptIn for the instance
func (instance *MDM_Win32CompatibilityAppraiser_UtcConfigurationDiagnosis02) SetPropertyTelemetryOptIn(value int32) (err error) { 
    return instance.SetProperty("TelemetryOptIn", (value))
}

// GetTelemetryOptIn gets the value of TelemetryOptIn for the instance
func (instance *MDM_Win32CompatibilityAppraiser_UtcConfigurationDiagnosis02) GetPropertyTelemetryOptIn()(value int32, err error) { 
    retValue, err := instance.GetProperty("TelemetryOptIn")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(int32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = int32(valuetmp)

    return
}

