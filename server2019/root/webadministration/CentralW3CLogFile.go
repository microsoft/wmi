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

// CentralW3CLogFile struct
type CentralW3CLogFile struct { 
	*EmbeddedObject

	// 
	Directory string

	// 
	Enabled bool

	// 
	LocalTimeRollover bool

	// 
	LogExtFileFlags int32

	// 
	Period int32

	// 
	TruncateSize string
}

	func NewCentralW3CLogFileEx1(instance *cim.WmiInstance) (newInstance *CentralW3CLogFile, err error) {tmp, err := NewEmbeddedObjectEx1(instance)
		
	if err != nil { return }
	newInstance = &CentralW3CLogFile {
	EmbeddedObject: tmp,
	}
	return
	}
	

	func NewCentralW3CLogFileEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *CentralW3CLogFile, err error) {tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &CentralW3CLogFile {
	EmbeddedObject: tmp,
	}
	return
	}
	

// SetDirectory sets the value of Directory for the instance
func (instance *CentralW3CLogFile) SetPropertyDirectory(value string) (err error) { 
    return instance.SetProperty("Directory", (value))
}

// GetDirectory gets the value of Directory for the instance
func (instance *CentralW3CLogFile) GetPropertyDirectory()(value string, err error) { 
    retValue, err := instance.GetProperty("Directory")
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

// SetEnabled sets the value of Enabled for the instance
func (instance *CentralW3CLogFile) SetPropertyEnabled(value bool) (err error) { 
    return instance.SetProperty("Enabled", (value))
}

// GetEnabled gets the value of Enabled for the instance
func (instance *CentralW3CLogFile) GetPropertyEnabled()(value bool, err error) { 
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

// SetLocalTimeRollover sets the value of LocalTimeRollover for the instance
func (instance *CentralW3CLogFile) SetPropertyLocalTimeRollover(value bool) (err error) { 
    return instance.SetProperty("LocalTimeRollover", (value))
}

// GetLocalTimeRollover gets the value of LocalTimeRollover for the instance
func (instance *CentralW3CLogFile) GetPropertyLocalTimeRollover()(value bool, err error) { 
    retValue, err := instance.GetProperty("LocalTimeRollover")
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

// SetLogExtFileFlags sets the value of LogExtFileFlags for the instance
func (instance *CentralW3CLogFile) SetPropertyLogExtFileFlags(value int32) (err error) { 
    return instance.SetProperty("LogExtFileFlags", (value))
}

// GetLogExtFileFlags gets the value of LogExtFileFlags for the instance
func (instance *CentralW3CLogFile) GetPropertyLogExtFileFlags()(value int32, err error) { 
    retValue, err := instance.GetProperty("LogExtFileFlags")
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

// SetPeriod sets the value of Period for the instance
func (instance *CentralW3CLogFile) SetPropertyPeriod(value int32) (err error) { 
    return instance.SetProperty("Period", (value))
}

// GetPeriod gets the value of Period for the instance
func (instance *CentralW3CLogFile) GetPropertyPeriod()(value int32, err error) { 
    retValue, err := instance.GetProperty("Period")
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

// SetTruncateSize sets the value of TruncateSize for the instance
func (instance *CentralW3CLogFile) SetPropertyTruncateSize(value string) (err error) { 
    return instance.SetProperty("TruncateSize", (value))
}

// GetTruncateSize gets the value of TruncateSize for the instance
func (instance *CentralW3CLogFile) GetPropertyTruncateSize()(value string, err error) { 
    retValue, err := instance.GetProperty("TruncateSize")
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

