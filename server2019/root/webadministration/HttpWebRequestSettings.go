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

// HttpWebRequestSettings struct
type HttpWebRequestSettings struct { 
	*EmbeddedObject

	// 
	MaximumErrorResponseLength int32

	// 
	MaximumResponseHeadersLength int32

	// 
	MaximumUnauthorizedUploadLength int32

	// 
	UseUnsafeHeaderParsing bool
}

	func NewHttpWebRequestSettingsEx1(instance *cim.WmiInstance) (newInstance *HttpWebRequestSettings, err error) {tmp, err := NewEmbeddedObjectEx1(instance)
		
	if err != nil { return }
	newInstance = &HttpWebRequestSettings {
	EmbeddedObject: tmp,
	}
	return
	}
	

	func NewHttpWebRequestSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *HttpWebRequestSettings, err error) {tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &HttpWebRequestSettings {
	EmbeddedObject: tmp,
	}
	return
	}
	

// SetMaximumErrorResponseLength sets the value of MaximumErrorResponseLength for the instance
func (instance *HttpWebRequestSettings) SetPropertyMaximumErrorResponseLength(value int32) (err error) { 
    return instance.SetProperty("MaximumErrorResponseLength", (value))
}

// GetMaximumErrorResponseLength gets the value of MaximumErrorResponseLength for the instance
func (instance *HttpWebRequestSettings) GetPropertyMaximumErrorResponseLength()(value int32, err error) { 
    retValue, err := instance.GetProperty("MaximumErrorResponseLength")
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

// SetMaximumResponseHeadersLength sets the value of MaximumResponseHeadersLength for the instance
func (instance *HttpWebRequestSettings) SetPropertyMaximumResponseHeadersLength(value int32) (err error) { 
    return instance.SetProperty("MaximumResponseHeadersLength", (value))
}

// GetMaximumResponseHeadersLength gets the value of MaximumResponseHeadersLength for the instance
func (instance *HttpWebRequestSettings) GetPropertyMaximumResponseHeadersLength()(value int32, err error) { 
    retValue, err := instance.GetProperty("MaximumResponseHeadersLength")
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

// SetMaximumUnauthorizedUploadLength sets the value of MaximumUnauthorizedUploadLength for the instance
func (instance *HttpWebRequestSettings) SetPropertyMaximumUnauthorizedUploadLength(value int32) (err error) { 
    return instance.SetProperty("MaximumUnauthorizedUploadLength", (value))
}

// GetMaximumUnauthorizedUploadLength gets the value of MaximumUnauthorizedUploadLength for the instance
func (instance *HttpWebRequestSettings) GetPropertyMaximumUnauthorizedUploadLength()(value int32, err error) { 
    retValue, err := instance.GetProperty("MaximumUnauthorizedUploadLength")
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

// SetUseUnsafeHeaderParsing sets the value of UseUnsafeHeaderParsing for the instance
func (instance *HttpWebRequestSettings) SetPropertyUseUnsafeHeaderParsing(value bool) (err error) { 
    return instance.SetProperty("UseUnsafeHeaderParsing", (value))
}

// GetUseUnsafeHeaderParsing gets the value of UseUnsafeHeaderParsing for the instance
func (instance *HttpWebRequestSettings) GetPropertyUseUnsafeHeaderParsing()(value bool, err error) { 
    retValue, err := instance.GetProperty("UseUnsafeHeaderParsing")
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

