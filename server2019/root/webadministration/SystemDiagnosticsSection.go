// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.WebAdministration
//////////////////////////////////////////////
package webadministration
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// SystemDiagnosticsSection struct
type SystemDiagnosticsSection struct { 
	*ConfigurationSectionWithCollection

	// 
	Assert AssertSettings

	// 
	PerformanceCounters PerformanceCounterSettings

	// 
	SharedListeners ListenerSettings

	// 
	Sources SourceSettings

	// 
	Switches SwitchSettings

	// 
	Trace TraceSettings
}

	func NewSystemDiagnosticsSectionEx1(instance *cim.WmiInstance) (newInstance *SystemDiagnosticsSection, err error) {tmp, err := NewConfigurationSectionWithCollectionEx1(instance)
		
	if err != nil { return }
	newInstance = &SystemDiagnosticsSection {
	ConfigurationSectionWithCollection: tmp,
	}
	return
	}
	

	func NewSystemDiagnosticsSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *SystemDiagnosticsSection, err error) {tmp, err := NewConfigurationSectionWithCollectionEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &SystemDiagnosticsSection {
	ConfigurationSectionWithCollection: tmp,
	}
	return
	}
	

// SetAssert sets the value of Assert for the instance
func (instance *SystemDiagnosticsSection) SetPropertyAssert(value AssertSettings) (err error) { 
    return instance.SetProperty("Assert", (value))
}

// GetAssert gets the value of Assert for the instance
func (instance *SystemDiagnosticsSection) GetPropertyAssert()(value AssertSettings, err error) { 
    retValue, err := instance.GetProperty("Assert")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(AssertSettings); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " AssertSettings is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = AssertSettings(valuetmp)

    return
}

// SetPerformanceCounters sets the value of PerformanceCounters for the instance
func (instance *SystemDiagnosticsSection) SetPropertyPerformanceCounters(value PerformanceCounterSettings) (err error) { 
    return instance.SetProperty("PerformanceCounters", (value))
}

// GetPerformanceCounters gets the value of PerformanceCounters for the instance
func (instance *SystemDiagnosticsSection) GetPropertyPerformanceCounters()(value PerformanceCounterSettings, err error) { 
    retValue, err := instance.GetProperty("PerformanceCounters")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(PerformanceCounterSettings); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " PerformanceCounterSettings is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = PerformanceCounterSettings(valuetmp)

    return
}

// SetSharedListeners sets the value of SharedListeners for the instance
func (instance *SystemDiagnosticsSection) SetPropertySharedListeners(value ListenerSettings) (err error) { 
    return instance.SetProperty("SharedListeners", (value))
}

// GetSharedListeners gets the value of SharedListeners for the instance
func (instance *SystemDiagnosticsSection) GetPropertySharedListeners()(value ListenerSettings, err error) { 
    retValue, err := instance.GetProperty("SharedListeners")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(ListenerSettings); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " ListenerSettings is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = ListenerSettings(valuetmp)

    return
}

// SetSources sets the value of Sources for the instance
func (instance *SystemDiagnosticsSection) SetPropertySources(value SourceSettings) (err error) { 
    return instance.SetProperty("Sources", (value))
}

// GetSources gets the value of Sources for the instance
func (instance *SystemDiagnosticsSection) GetPropertySources()(value SourceSettings, err error) { 
    retValue, err := instance.GetProperty("Sources")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(SourceSettings); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " SourceSettings is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = SourceSettings(valuetmp)

    return
}

// SetSwitches sets the value of Switches for the instance
func (instance *SystemDiagnosticsSection) SetPropertySwitches(value SwitchSettings) (err error) { 
    return instance.SetProperty("Switches", (value))
}

// GetSwitches gets the value of Switches for the instance
func (instance *SystemDiagnosticsSection) GetPropertySwitches()(value SwitchSettings, err error) { 
    retValue, err := instance.GetProperty("Switches")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(SwitchSettings); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " SwitchSettings is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = SwitchSettings(valuetmp)

    return
}

// SetTrace sets the value of Trace for the instance
func (instance *SystemDiagnosticsSection) SetPropertyTrace(value TraceSettings) (err error) { 
    return instance.SetProperty("Trace", (value))
}

// GetTrace gets the value of Trace for the instance
func (instance *SystemDiagnosticsSection) GetPropertyTrace()(value TraceSettings, err error) { 
    retValue, err := instance.GetProperty("Trace")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(TraceSettings); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " TraceSettings is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = TraceSettings(valuetmp)

    return
}

