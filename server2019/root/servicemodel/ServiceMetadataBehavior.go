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

// ServiceMetadataBehavior struct
type ServiceMetadataBehavior struct { 
	*Behavior

	// Sets the location to which the service redirects metadata requests.
	ExternalMetadataLocation string

	// Controls the binding for metadata retrieval using HTTP.
	HttpGetBinding Binding

	// Controls whether the service publishes its WSDL at the address controlled by the HttpGetUrl attribute.
	HttpGetEnabled bool

	// Sets the location at which the service WSDL is published for retrieval using HTTP.
	HttpGetUrl string

	// Controls the binding for metadata retrieval using HTTPS.
	HttpsGetBinding Binding

	// Controls whether the service publishes its WSDL over HTTPS at the address controlled by the HttpsGetUrl attribute.
	HttpsGetEnabled bool

	// Sets the location at which the service WSDL is published for retrieval using HTTPS.
	HttpsGetUrl string

	// The component responsible for metadata generation associated with this service.
	MetadataExportInfo MetadataExporter
}

	func NewServiceMetadataBehaviorEx1(instance *cim.WmiInstance) (newInstance *ServiceMetadataBehavior, err error) {tmp, err := NewBehaviorEx1(instance)
		
	if err != nil { return }
	newInstance = &ServiceMetadataBehavior {
	Behavior: tmp,
	}
	return
	}
	

	func NewServiceMetadataBehaviorEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ServiceMetadataBehavior, err error) {tmp, err := NewBehaviorEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ServiceMetadataBehavior {
	Behavior: tmp,
	}
	return
	}
	

// SetExternalMetadataLocation sets the value of ExternalMetadataLocation for the instance
func (instance *ServiceMetadataBehavior) SetPropertyExternalMetadataLocation(value string) (err error) { 
    return instance.SetProperty("ExternalMetadataLocation", (value))
}

// GetExternalMetadataLocation gets the value of ExternalMetadataLocation for the instance
func (instance *ServiceMetadataBehavior) GetPropertyExternalMetadataLocation()(value string, err error) { 
    retValue, err := instance.GetProperty("ExternalMetadataLocation")
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

// SetHttpGetBinding sets the value of HttpGetBinding for the instance
func (instance *ServiceMetadataBehavior) SetPropertyHttpGetBinding(value Binding) (err error) { 
    return instance.SetProperty("HttpGetBinding", (value))
}

// GetHttpGetBinding gets the value of HttpGetBinding for the instance
func (instance *ServiceMetadataBehavior) GetPropertyHttpGetBinding()(value Binding, err error) { 
    retValue, err := instance.GetProperty("HttpGetBinding")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(Binding); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " Binding is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = Binding(valuetmp)

    return
}

// SetHttpGetEnabled sets the value of HttpGetEnabled for the instance
func (instance *ServiceMetadataBehavior) SetPropertyHttpGetEnabled(value bool) (err error) { 
    return instance.SetProperty("HttpGetEnabled", (value))
}

// GetHttpGetEnabled gets the value of HttpGetEnabled for the instance
func (instance *ServiceMetadataBehavior) GetPropertyHttpGetEnabled()(value bool, err error) { 
    retValue, err := instance.GetProperty("HttpGetEnabled")
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

// SetHttpGetUrl sets the value of HttpGetUrl for the instance
func (instance *ServiceMetadataBehavior) SetPropertyHttpGetUrl(value string) (err error) { 
    return instance.SetProperty("HttpGetUrl", (value))
}

// GetHttpGetUrl gets the value of HttpGetUrl for the instance
func (instance *ServiceMetadataBehavior) GetPropertyHttpGetUrl()(value string, err error) { 
    retValue, err := instance.GetProperty("HttpGetUrl")
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

// SetHttpsGetBinding sets the value of HttpsGetBinding for the instance
func (instance *ServiceMetadataBehavior) SetPropertyHttpsGetBinding(value Binding) (err error) { 
    return instance.SetProperty("HttpsGetBinding", (value))
}

// GetHttpsGetBinding gets the value of HttpsGetBinding for the instance
func (instance *ServiceMetadataBehavior) GetPropertyHttpsGetBinding()(value Binding, err error) { 
    retValue, err := instance.GetProperty("HttpsGetBinding")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(Binding); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " Binding is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = Binding(valuetmp)

    return
}

// SetHttpsGetEnabled sets the value of HttpsGetEnabled for the instance
func (instance *ServiceMetadataBehavior) SetPropertyHttpsGetEnabled(value bool) (err error) { 
    return instance.SetProperty("HttpsGetEnabled", (value))
}

// GetHttpsGetEnabled gets the value of HttpsGetEnabled for the instance
func (instance *ServiceMetadataBehavior) GetPropertyHttpsGetEnabled()(value bool, err error) { 
    retValue, err := instance.GetProperty("HttpsGetEnabled")
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

// SetHttpsGetUrl sets the value of HttpsGetUrl for the instance
func (instance *ServiceMetadataBehavior) SetPropertyHttpsGetUrl(value string) (err error) { 
    return instance.SetProperty("HttpsGetUrl", (value))
}

// GetHttpsGetUrl gets the value of HttpsGetUrl for the instance
func (instance *ServiceMetadataBehavior) GetPropertyHttpsGetUrl()(value string, err error) { 
    retValue, err := instance.GetProperty("HttpsGetUrl")
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

// SetMetadataExportInfo sets the value of MetadataExportInfo for the instance
func (instance *ServiceMetadataBehavior) SetPropertyMetadataExportInfo(value MetadataExporter) (err error) { 
    return instance.SetProperty("MetadataExportInfo", (value))
}

// GetMetadataExportInfo gets the value of MetadataExportInfo for the instance
func (instance *ServiceMetadataBehavior) GetPropertyMetadataExportInfo()(value MetadataExporter, err error) { 
    retValue, err := instance.GetProperty("MetadataExportInfo")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(MetadataExporter); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " MetadataExporter is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = MetadataExporter(valuetmp)

    return
}

