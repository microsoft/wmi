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
 "github.com/microsoft/wmi/pkg/base/instance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// SerializableToXml struct
type SerializableToXml struct { 
	*cim.WmiInstance

	// This is an all-in-one state data that uses an XML format with a standard CIM DTD schema
	PackedXml string

	// Schema version (major, minor, build, revision)
	SchemaVersion string
}

	func NewSerializableToXmlEx1(instance *cim.WmiInstance) (newInstance *SerializableToXml, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &SerializableToXml {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewSerializableToXmlEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *SerializableToXml, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &SerializableToXml {
	WmiInstance: tmp,
	}
	return
	}
	

// SetPackedXml sets the value of PackedXml for the instance
func (instance *SerializableToXml) SetPropertyPackedXml(value string) (err error) { 
    return instance.SetProperty("PackedXml", (value))
}

// GetPackedXml gets the value of PackedXml for the instance
func (instance *SerializableToXml) GetPropertyPackedXml()(value string, err error) { 
    retValue, err := instance.GetProperty("PackedXml")
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

// SetSchemaVersion sets the value of SchemaVersion for the instance
func (instance *SerializableToXml) SetPropertySchemaVersion(value string) (err error) { 
    return instance.SetProperty("SchemaVersion", (value))
}

// GetSchemaVersion gets the value of SchemaVersion for the instance
func (instance *SerializableToXml) GetPropertySchemaVersion()(value string, err error) { 
    retValue, err := instance.GetProperty("SchemaVersion")
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

