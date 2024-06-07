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

// TraceSection struct
type TraceSection struct { 
	*ConfigurationSection

	// 
	Enabled bool

	// 
	LocalOnly bool

	// 
	MostRecent bool

	// 
	PageOutput bool

	// 
	RequestLimit int32

	// 
	TraceMode int32

	// 
	WriteToDiagnosticsTrace bool
}

	func NewTraceSectionEx1(instance *cim.WmiInstance) (newInstance *TraceSection, err error) {tmp, err := NewConfigurationSectionEx1(instance)
		
	if err != nil { return }
	newInstance = &TraceSection {
	ConfigurationSection: tmp,
	}
	return
	}
	

	func NewTraceSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *TraceSection, err error) {tmp, err := NewConfigurationSectionEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &TraceSection {
	ConfigurationSection: tmp,
	}
	return
	}
	

// SetEnabled sets the value of Enabled for the instance
func (instance *TraceSection) SetPropertyEnabled(value bool) (err error) { 
    return instance.SetProperty("Enabled", (value))
}

// GetEnabled gets the value of Enabled for the instance
func (instance *TraceSection) GetPropertyEnabled()(value bool, err error) { 
    retValue, err := instance.GetProperty("Enabled")
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

// SetLocalOnly sets the value of LocalOnly for the instance
func (instance *TraceSection) SetPropertyLocalOnly(value bool) (err error) { 
    return instance.SetProperty("LocalOnly", (value))
}

// GetLocalOnly gets the value of LocalOnly for the instance
func (instance *TraceSection) GetPropertyLocalOnly()(value bool, err error) { 
    retValue, err := instance.GetProperty("LocalOnly")
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

// SetMostRecent sets the value of MostRecent for the instance
func (instance *TraceSection) SetPropertyMostRecent(value bool) (err error) { 
    return instance.SetProperty("MostRecent", (value))
}

// GetMostRecent gets the value of MostRecent for the instance
func (instance *TraceSection) GetPropertyMostRecent()(value bool, err error) { 
    retValue, err := instance.GetProperty("MostRecent")
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

// SetPageOutput sets the value of PageOutput for the instance
func (instance *TraceSection) SetPropertyPageOutput(value bool) (err error) { 
    return instance.SetProperty("PageOutput", (value))
}

// GetPageOutput gets the value of PageOutput for the instance
func (instance *TraceSection) GetPropertyPageOutput()(value bool, err error) { 
    retValue, err := instance.GetProperty("PageOutput")
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

// SetRequestLimit sets the value of RequestLimit for the instance
func (instance *TraceSection) SetPropertyRequestLimit(value int32) (err error) { 
    return instance.SetProperty("RequestLimit", (value))
}

// GetRequestLimit gets the value of RequestLimit for the instance
func (instance *TraceSection) GetPropertyRequestLimit()(value int32, err error) { 
    retValue, err := instance.GetProperty("RequestLimit")
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

// SetTraceMode sets the value of TraceMode for the instance
func (instance *TraceSection) SetPropertyTraceMode(value int32) (err error) { 
    return instance.SetProperty("TraceMode", (value))
}

// GetTraceMode gets the value of TraceMode for the instance
func (instance *TraceSection) GetPropertyTraceMode()(value int32, err error) { 
    retValue, err := instance.GetProperty("TraceMode")
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

// SetWriteToDiagnosticsTrace sets the value of WriteToDiagnosticsTrace for the instance
func (instance *TraceSection) SetPropertyWriteToDiagnosticsTrace(value bool) (err error) { 
    return instance.SetProperty("WriteToDiagnosticsTrace", (value))
}

// GetWriteToDiagnosticsTrace gets the value of WriteToDiagnosticsTrace for the instance
func (instance *TraceSection) GetPropertyWriteToDiagnosticsTrace()(value bool, err error) { 
    retValue, err := instance.GetProperty("WriteToDiagnosticsTrace")
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

