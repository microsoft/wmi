// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.ServiceModel
//////////////////////////////////////////////
package servicemodel
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// HttpsTransportBindingElement struct
type HttpsTransportBindingElement struct { 
	*HttpTransportBindingElement

	// A value that indicates whether SSL client authentication is required.
	RequireClientCertificate bool
}

	func NewHttpsTransportBindingElementEx1(instance *cim.WmiInstance) (newInstance *HttpsTransportBindingElement, err error) {tmp, err := NewHttpTransportBindingElementEx1(instance)
		
	if err != nil { return }
	newInstance = &HttpsTransportBindingElement {
	HttpTransportBindingElement: tmp,
	}
	return
	}
	

	func NewHttpsTransportBindingElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *HttpsTransportBindingElement, err error) {tmp, err := NewHttpTransportBindingElementEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &HttpsTransportBindingElement {
	HttpTransportBindingElement: tmp,
	}
	return
	}
	

// SetRequireClientCertificate sets the value of RequireClientCertificate for the instance
func (instance *HttpsTransportBindingElement) SetPropertyRequireClientCertificate(value bool) (err error) { 
    return instance.SetProperty("RequireClientCertificate", (value))
}

// GetRequireClientCertificate gets the value of RequireClientCertificate for the instance
func (instance *HttpsTransportBindingElement) GetPropertyRequireClientCertificate()(value bool, err error) { 
    retValue, err := instance.GetProperty("RequireClientCertificate")
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

