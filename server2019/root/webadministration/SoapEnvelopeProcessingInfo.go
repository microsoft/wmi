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

// SoapEnvelopeProcessingInfo struct
type SoapEnvelopeProcessingInfo struct { 
	*EmbeddedObject

	// 
	ReadTimeout int32

	// 
	Strict bool
}

	func NewSoapEnvelopeProcessingInfoEx1(instance *cim.WmiInstance) (newInstance *SoapEnvelopeProcessingInfo, err error) {tmp, err := NewEmbeddedObjectEx1(instance)
		
	if err != nil { return }
	newInstance = &SoapEnvelopeProcessingInfo {
	EmbeddedObject: tmp,
	}
	return
	}
	

	func NewSoapEnvelopeProcessingInfoEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *SoapEnvelopeProcessingInfo, err error) {tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &SoapEnvelopeProcessingInfo {
	EmbeddedObject: tmp,
	}
	return
	}
	

// SetReadTimeout sets the value of ReadTimeout for the instance
func (instance *SoapEnvelopeProcessingInfo) SetPropertyReadTimeout(value int32) (err error) { 
    return instance.SetProperty("ReadTimeout", (value))
}

// GetReadTimeout gets the value of ReadTimeout for the instance
func (instance *SoapEnvelopeProcessingInfo) GetPropertyReadTimeout()(value int32, err error) { 
    retValue, err := instance.GetProperty("ReadTimeout")
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

// SetStrict sets the value of Strict for the instance
func (instance *SoapEnvelopeProcessingInfo) SetPropertyStrict(value bool) (err error) { 
    return instance.SetProperty("Strict", (value))
}

// GetStrict gets the value of Strict for the instance
func (instance *SoapEnvelopeProcessingInfo) GetPropertyStrict()(value bool, err error) { 
    retValue, err := instance.GetProperty("Strict")
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

