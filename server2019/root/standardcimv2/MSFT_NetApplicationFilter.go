// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// MSFT_NetApplicationFilter struct
type MSFT_NetApplicationFilter struct { 
	*CIM_FilterEntryBase

	// 
	AppPath string

	// 
	Package string
}

	func NewMSFT_NetApplicationFilterEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetApplicationFilter, err error) {tmp, err := NewCIM_FilterEntryBaseEx1(instance)
		
	if err != nil { return }
	newInstance = &MSFT_NetApplicationFilter {
	CIM_FilterEntryBase: tmp,
	}
	return
	}
	

	func NewMSFT_NetApplicationFilterEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSFT_NetApplicationFilter, err error) {tmp, err := NewCIM_FilterEntryBaseEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSFT_NetApplicationFilter {
	CIM_FilterEntryBase: tmp,
	}
	return
	}
	

// SetAppPath sets the value of AppPath for the instance
func (instance *MSFT_NetApplicationFilter) SetPropertyAppPath(value string) (err error) { 
    return instance.SetProperty("AppPath", (value))
}

// GetAppPath gets the value of AppPath for the instance
func (instance *MSFT_NetApplicationFilter) GetPropertyAppPath()(value string, err error) { 
    retValue, err := instance.GetProperty("AppPath")
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

// SetPackage sets the value of Package for the instance
func (instance *MSFT_NetApplicationFilter) SetPropertyPackage(value string) (err error) { 
    return instance.SetProperty("Package", (value))
}

// GetPackage gets the value of Package for the instance
func (instance *MSFT_NetApplicationFilter) GetPropertyPackage()(value string, err error) { 
    retValue, err := instance.GetProperty("Package")
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

