// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// LoaderDllSearchResults struct
type LoaderDllSearchResults struct { 
	*Image

	// 
	FullDllName string

	// 
	LdrLoadFlags uint32

	// 
	LdrSearchFlags uint32

	// 
	LoadReason uint32

	// 
	SearchInfo uint32
}

	func NewLoaderDllSearchResultsEx1(instance *cim.WmiInstance) (newInstance *LoaderDllSearchResults, err error) {tmp, err := NewImageEx1(instance)
		
	if err != nil { return }
	newInstance = &LoaderDllSearchResults {
	Image: tmp,
	}
	return
	}
	

	func NewLoaderDllSearchResultsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *LoaderDllSearchResults, err error) {tmp, err := NewImageEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &LoaderDllSearchResults {
	Image: tmp,
	}
	return
	}
	

// SetFullDllName sets the value of FullDllName for the instance
func (instance *LoaderDllSearchResults) SetPropertyFullDllName(value string) (err error) { 
    return instance.SetProperty("FullDllName", (value))
}

// GetFullDllName gets the value of FullDllName for the instance
func (instance *LoaderDllSearchResults) GetPropertyFullDllName()(value string, err error) { 
    retValue, err := instance.GetProperty("FullDllName")
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

// SetLdrLoadFlags sets the value of LdrLoadFlags for the instance
func (instance *LoaderDllSearchResults) SetPropertyLdrLoadFlags(value uint32) (err error) { 
    return instance.SetProperty("LdrLoadFlags", (value))
}

// GetLdrLoadFlags gets the value of LdrLoadFlags for the instance
func (instance *LoaderDllSearchResults) GetPropertyLdrLoadFlags()(value uint32, err error) { 
    retValue, err := instance.GetProperty("LdrLoadFlags")
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

// SetLdrSearchFlags sets the value of LdrSearchFlags for the instance
func (instance *LoaderDllSearchResults) SetPropertyLdrSearchFlags(value uint32) (err error) { 
    return instance.SetProperty("LdrSearchFlags", (value))
}

// GetLdrSearchFlags gets the value of LdrSearchFlags for the instance
func (instance *LoaderDllSearchResults) GetPropertyLdrSearchFlags()(value uint32, err error) { 
    retValue, err := instance.GetProperty("LdrSearchFlags")
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

// SetLoadReason sets the value of LoadReason for the instance
func (instance *LoaderDllSearchResults) SetPropertyLoadReason(value uint32) (err error) { 
    return instance.SetProperty("LoadReason", (value))
}

// GetLoadReason gets the value of LoadReason for the instance
func (instance *LoaderDllSearchResults) GetPropertyLoadReason()(value uint32, err error) { 
    retValue, err := instance.GetProperty("LoadReason")
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

// SetSearchInfo sets the value of SearchInfo for the instance
func (instance *LoaderDllSearchResults) SetPropertySearchInfo(value uint32) (err error) { 
    return instance.SetProperty("SearchInfo", (value))
}

// GetSearchInfo gets the value of SearchInfo for the instance
func (instance *LoaderDllSearchResults) GetPropertySearchInfo()(value uint32, err error) { 
    retValue, err := instance.GetProperty("SearchInfo")
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

