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

// DiagnosticsSettings struct
type DiagnosticsSettings struct { 
	*EmbeddedObject

	// 
	SuppressReturningExceptions bool
}

	func NewDiagnosticsSettingsEx1(instance *cim.WmiInstance) (newInstance *DiagnosticsSettings, err error) {tmp, err := NewEmbeddedObjectEx1(instance)
		
	if err != nil { return }
	newInstance = &DiagnosticsSettings {
	EmbeddedObject: tmp,
	}
	return
	}
	

	func NewDiagnosticsSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *DiagnosticsSettings, err error) {tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &DiagnosticsSettings {
	EmbeddedObject: tmp,
	}
	return
	}
	

// SetSuppressReturningExceptions sets the value of SuppressReturningExceptions for the instance
func (instance *DiagnosticsSettings) SetPropertySuppressReturningExceptions(value bool) (err error) { 
    return instance.SetProperty("SuppressReturningExceptions", (value))
}

// GetSuppressReturningExceptions gets the value of SuppressReturningExceptions for the instance
func (instance *DiagnosticsSettings) GetPropertySuppressReturningExceptions()(value bool, err error) { 
    retValue, err := instance.GetProperty("SuppressReturningExceptions")
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

