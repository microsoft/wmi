// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.Microsoft.Windows.DesiredStateConfiguration
//////////////////////////////////////////////
package desiredstateconfiguration
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// MSFT_DSCResource struct
type MSFT_DSCResource struct { 
	*OMI_BaseResource

	// 
	DurationInSeconds float64

	// 
	Error string

	// 
	FinalState string

	// 
	InDesiredState bool

	// 
	InitialState string

	// 
	InstanceName string

	// 
	RebootRequested bool

	// 
	ResourceName string

	// 
	StartDate string

	// 
	StateChanged bool
}

	func NewMSFT_DSCResourceEx1(instance *cim.WmiInstance) (newInstance *MSFT_DSCResource, err error) {tmp, err := NewOMI_BaseResourceEx1(instance)
		
	if err != nil { return }
	newInstance = &MSFT_DSCResource {
	OMI_BaseResource: tmp,
	}
	return
	}
	

	func NewMSFT_DSCResourceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSFT_DSCResource, err error) {tmp, err := NewOMI_BaseResourceEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSFT_DSCResource {
	OMI_BaseResource: tmp,
	}
	return
	}
	

// SetDurationInSeconds sets the value of DurationInSeconds for the instance
func (instance *MSFT_DSCResource) SetPropertyDurationInSeconds(value float64) (err error) { 
    return instance.SetProperty("DurationInSeconds", (value))
}

// GetDurationInSeconds gets the value of DurationInSeconds for the instance
func (instance *MSFT_DSCResource) GetPropertyDurationInSeconds()(value float64, err error) { 
    retValue, err := instance.GetProperty("DurationInSeconds")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(float64); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " float64 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = float64(valuetmp)

    return
}

// SetError sets the value of Error for the instance
func (instance *MSFT_DSCResource) SetPropertyError(value string) (err error) { 
    return instance.SetProperty("Error", (value))
}

// GetError gets the value of Error for the instance
func (instance *MSFT_DSCResource) GetPropertyError()(value string, err error) { 
    retValue, err := instance.GetProperty("Error")
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

// SetFinalState sets the value of FinalState for the instance
func (instance *MSFT_DSCResource) SetPropertyFinalState(value string) (err error) { 
    return instance.SetProperty("FinalState", (value))
}

// GetFinalState gets the value of FinalState for the instance
func (instance *MSFT_DSCResource) GetPropertyFinalState()(value string, err error) { 
    retValue, err := instance.GetProperty("FinalState")
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

// SetInDesiredState sets the value of InDesiredState for the instance
func (instance *MSFT_DSCResource) SetPropertyInDesiredState(value bool) (err error) { 
    return instance.SetProperty("InDesiredState", (value))
}

// GetInDesiredState gets the value of InDesiredState for the instance
func (instance *MSFT_DSCResource) GetPropertyInDesiredState()(value bool, err error) { 
    retValue, err := instance.GetProperty("InDesiredState")
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

// SetInitialState sets the value of InitialState for the instance
func (instance *MSFT_DSCResource) SetPropertyInitialState(value string) (err error) { 
    return instance.SetProperty("InitialState", (value))
}

// GetInitialState gets the value of InitialState for the instance
func (instance *MSFT_DSCResource) GetPropertyInitialState()(value string, err error) { 
    retValue, err := instance.GetProperty("InitialState")
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

// SetInstanceName sets the value of InstanceName for the instance
func (instance *MSFT_DSCResource) SetPropertyInstanceName(value string) (err error) { 
    return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSFT_DSCResource) GetPropertyInstanceName()(value string, err error) { 
    retValue, err := instance.GetProperty("InstanceName")
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

// SetRebootRequested sets the value of RebootRequested for the instance
func (instance *MSFT_DSCResource) SetPropertyRebootRequested(value bool) (err error) { 
    return instance.SetProperty("RebootRequested", (value))
}

// GetRebootRequested gets the value of RebootRequested for the instance
func (instance *MSFT_DSCResource) GetPropertyRebootRequested()(value bool, err error) { 
    retValue, err := instance.GetProperty("RebootRequested")
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

// SetResourceName sets the value of ResourceName for the instance
func (instance *MSFT_DSCResource) SetPropertyResourceName(value string) (err error) { 
    return instance.SetProperty("ResourceName", (value))
}

// GetResourceName gets the value of ResourceName for the instance
func (instance *MSFT_DSCResource) GetPropertyResourceName()(value string, err error) { 
    retValue, err := instance.GetProperty("ResourceName")
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

// SetStartDate sets the value of StartDate for the instance
func (instance *MSFT_DSCResource) SetPropertyStartDate(value string) (err error) { 
    return instance.SetProperty("StartDate", (value))
}

// GetStartDate gets the value of StartDate for the instance
func (instance *MSFT_DSCResource) GetPropertyStartDate()(value string, err error) { 
    retValue, err := instance.GetProperty("StartDate")
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

// SetStateChanged sets the value of StateChanged for the instance
func (instance *MSFT_DSCResource) SetPropertyStateChanged(value bool) (err error) { 
    return instance.SetProperty("StateChanged", (value))
}

// GetStateChanged gets the value of StateChanged for the instance
func (instance *MSFT_DSCResource) GetPropertyStateChanged()(value bool, err error) { 
    retValue, err := instance.GetProperty("StateChanged")
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

