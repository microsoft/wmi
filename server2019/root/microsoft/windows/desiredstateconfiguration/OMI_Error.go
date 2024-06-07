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

// OMI_Error struct
type OMI_Error struct { 
	*CIM_Error

	// 
	error_Category uint16

	// 
	error_Code uint32

	// 
	error_Type string
}

	func NewOMI_ErrorEx1(instance *cim.WmiInstance) (newInstance *OMI_Error, err error) {tmp, err := NewCIM_ErrorEx1(instance)
		
	if err != nil { return }
	newInstance = &OMI_Error {
	CIM_Error: tmp,
	}
	return
	}
	

	func NewOMI_ErrorEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *OMI_Error, err error) {tmp, err := NewCIM_ErrorEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &OMI_Error {
	CIM_Error: tmp,
	}
	return
	}
	

// Seterror_Category sets the value of error_Category for the instance
func (instance *OMI_Error) SetPropertyerror_Category(value uint16) (err error) { 
    return instance.SetProperty("error_Category", (value))
}

// Geterror_Category gets the value of error_Category for the instance
func (instance *OMI_Error) GetPropertyerror_Category()(value uint16, err error) { 
    retValue, err := instance.GetProperty("error_Category")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint16); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint16(valuetmp)

    return
}

// Seterror_Code sets the value of error_Code for the instance
func (instance *OMI_Error) SetPropertyerror_Code(value uint32) (err error) { 
    return instance.SetProperty("error_Code", (value))
}

// Geterror_Code gets the value of error_Code for the instance
func (instance *OMI_Error) GetPropertyerror_Code()(value uint32, err error) { 
    retValue, err := instance.GetProperty("error_Code")
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

// Seterror_Type sets the value of error_Type for the instance
func (instance *OMI_Error) SetPropertyerror_Type(value string) (err error) { 
    return instance.SetProperty("error_Type", (value))
}

// Geterror_Type gets the value of error_Type for the instance
func (instance *OMI_Error) GetPropertyerror_Type()(value string, err error) { 
    retValue, err := instance.GetProperty("error_Type")
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

