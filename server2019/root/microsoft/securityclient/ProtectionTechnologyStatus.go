// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.Microsoft.SecurityClient
//////////////////////////////////////////////
package securityclient
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// ProtectionTechnologyStatus struct
type ProtectionTechnologyStatus struct { 
	*SerializableToXml

	// Is protection technology enabled
	Enabled bool

	// Protection technology name
	Name string

	// Protection technology version (major, minor, build, revision)
	Version string
}

	func NewProtectionTechnologyStatusEx1(instance *cim.WmiInstance) (newInstance *ProtectionTechnologyStatus, err error) {tmp, err := NewSerializableToXmlEx1(instance)
		
	if err != nil { return }
	newInstance = &ProtectionTechnologyStatus {
	SerializableToXml: tmp,
	}
	return
	}
	

	func NewProtectionTechnologyStatusEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ProtectionTechnologyStatus, err error) {tmp, err := NewSerializableToXmlEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ProtectionTechnologyStatus {
	SerializableToXml: tmp,
	}
	return
	}
	

// SetEnabled sets the value of Enabled for the instance
func (instance *ProtectionTechnologyStatus) SetPropertyEnabled(value bool) (err error) { 
    return instance.SetProperty("Enabled", (value))
}

// GetEnabled gets the value of Enabled for the instance
func (instance *ProtectionTechnologyStatus) GetPropertyEnabled()(value bool, err error) { 
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

// SetName sets the value of Name for the instance
func (instance *ProtectionTechnologyStatus) SetPropertyName(value string) (err error) { 
    return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *ProtectionTechnologyStatus) GetPropertyName()(value string, err error) { 
    retValue, err := instance.GetProperty("Name")
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

// SetVersion sets the value of Version for the instance
func (instance *ProtectionTechnologyStatus) SetPropertyVersion(value string) (err error) { 
    return instance.SetProperty("Version", (value))
}

// GetVersion gets the value of Version for the instance
func (instance *ProtectionTechnologyStatus) GetPropertyVersion()(value string, err error) { 
    retValue, err := instance.GetProperty("Version")
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

