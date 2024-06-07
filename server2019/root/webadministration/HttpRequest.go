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

// HttpRequest struct
type HttpRequest struct { 
	*Object

	// 
	ClientIPAddress string

	// 
	ConnectionId string

	// 
	CurrentModule string

	// 
	GUID string

	// 
	HostName string

	// 
	LocalIPAddress string

	// 
	LocalPort uint32

	// 
	PipelineState uint32

	// 
	SiteId uint32

	// 
	TimeElapsed uint32

	// 
	TimeInModule uint32

	// 
	TimeInState uint32

	// 
	Url string

	// 
	Verb string
}

	func NewHttpRequestEx1(instance *cim.WmiInstance) (newInstance *HttpRequest, err error) {tmp, err := NewObjectEx1(instance)
		
	if err != nil { return }
	newInstance = &HttpRequest {
	Object: tmp,
	}
	return
	}
	

	func NewHttpRequestEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *HttpRequest, err error) {tmp, err := NewObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &HttpRequest {
	Object: tmp,
	}
	return
	}
	

// SetClientIPAddress sets the value of ClientIPAddress for the instance
func (instance *HttpRequest) SetPropertyClientIPAddress(value string) (err error) { 
    return instance.SetProperty("ClientIPAddress", (value))
}

// GetClientIPAddress gets the value of ClientIPAddress for the instance
func (instance *HttpRequest) GetPropertyClientIPAddress()(value string, err error) { 
    retValue, err := instance.GetProperty("ClientIPAddress")
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

// SetConnectionId sets the value of ConnectionId for the instance
func (instance *HttpRequest) SetPropertyConnectionId(value string) (err error) { 
    return instance.SetProperty("ConnectionId", (value))
}

// GetConnectionId gets the value of ConnectionId for the instance
func (instance *HttpRequest) GetPropertyConnectionId()(value string, err error) { 
    retValue, err := instance.GetProperty("ConnectionId")
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

// SetCurrentModule sets the value of CurrentModule for the instance
func (instance *HttpRequest) SetPropertyCurrentModule(value string) (err error) { 
    return instance.SetProperty("CurrentModule", (value))
}

// GetCurrentModule gets the value of CurrentModule for the instance
func (instance *HttpRequest) GetPropertyCurrentModule()(value string, err error) { 
    retValue, err := instance.GetProperty("CurrentModule")
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

// SetGUID sets the value of GUID for the instance
func (instance *HttpRequest) SetPropertyGUID(value string) (err error) { 
    return instance.SetProperty("GUID", (value))
}

// GetGUID gets the value of GUID for the instance
func (instance *HttpRequest) GetPropertyGUID()(value string, err error) { 
    retValue, err := instance.GetProperty("GUID")
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

// SetHostName sets the value of HostName for the instance
func (instance *HttpRequest) SetPropertyHostName(value string) (err error) { 
    return instance.SetProperty("HostName", (value))
}

// GetHostName gets the value of HostName for the instance
func (instance *HttpRequest) GetPropertyHostName()(value string, err error) { 
    retValue, err := instance.GetProperty("HostName")
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

// SetLocalIPAddress sets the value of LocalIPAddress for the instance
func (instance *HttpRequest) SetPropertyLocalIPAddress(value string) (err error) { 
    return instance.SetProperty("LocalIPAddress", (value))
}

// GetLocalIPAddress gets the value of LocalIPAddress for the instance
func (instance *HttpRequest) GetPropertyLocalIPAddress()(value string, err error) { 
    retValue, err := instance.GetProperty("LocalIPAddress")
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

// SetLocalPort sets the value of LocalPort for the instance
func (instance *HttpRequest) SetPropertyLocalPort(value uint32) (err error) { 
    return instance.SetProperty("LocalPort", (value))
}

// GetLocalPort gets the value of LocalPort for the instance
func (instance *HttpRequest) GetPropertyLocalPort()(value uint32, err error) { 
    retValue, err := instance.GetProperty("LocalPort")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetPipelineState sets the value of PipelineState for the instance
func (instance *HttpRequest) SetPropertyPipelineState(value uint32) (err error) { 
    return instance.SetProperty("PipelineState", (value))
}

// GetPipelineState gets the value of PipelineState for the instance
func (instance *HttpRequest) GetPropertyPipelineState()(value uint32, err error) { 
    retValue, err := instance.GetProperty("PipelineState")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetSiteId sets the value of SiteId for the instance
func (instance *HttpRequest) SetPropertySiteId(value uint32) (err error) { 
    return instance.SetProperty("SiteId", (value))
}

// GetSiteId gets the value of SiteId for the instance
func (instance *HttpRequest) GetPropertySiteId()(value uint32, err error) { 
    retValue, err := instance.GetProperty("SiteId")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetTimeElapsed sets the value of TimeElapsed for the instance
func (instance *HttpRequest) SetPropertyTimeElapsed(value uint32) (err error) { 
    return instance.SetProperty("TimeElapsed", (value))
}

// GetTimeElapsed gets the value of TimeElapsed for the instance
func (instance *HttpRequest) GetPropertyTimeElapsed()(value uint32, err error) { 
    retValue, err := instance.GetProperty("TimeElapsed")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetTimeInModule sets the value of TimeInModule for the instance
func (instance *HttpRequest) SetPropertyTimeInModule(value uint32) (err error) { 
    return instance.SetProperty("TimeInModule", (value))
}

// GetTimeInModule gets the value of TimeInModule for the instance
func (instance *HttpRequest) GetPropertyTimeInModule()(value uint32, err error) { 
    retValue, err := instance.GetProperty("TimeInModule")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetTimeInState sets the value of TimeInState for the instance
func (instance *HttpRequest) SetPropertyTimeInState(value uint32) (err error) { 
    return instance.SetProperty("TimeInState", (value))
}

// GetTimeInState gets the value of TimeInState for the instance
func (instance *HttpRequest) GetPropertyTimeInState()(value uint32, err error) { 
    retValue, err := instance.GetProperty("TimeInState")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetUrl sets the value of Url for the instance
func (instance *HttpRequest) SetPropertyUrl(value string) (err error) { 
    return instance.SetProperty("Url", (value))
}

// GetUrl gets the value of Url for the instance
func (instance *HttpRequest) GetPropertyUrl()(value string, err error) { 
    retValue, err := instance.GetProperty("Url")
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

// SetVerb sets the value of Verb for the instance
func (instance *HttpRequest) SetPropertyVerb(value string) (err error) { 
    return instance.SetProperty("Verb", (value))
}

// GetVerb gets the value of Verb for the instance
func (instance *HttpRequest) GetPropertyVerb()(value string, err error) { 
    retValue, err := instance.GetProperty("Verb")
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

