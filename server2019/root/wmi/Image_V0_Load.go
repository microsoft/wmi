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

// Image_V0_Load struct
type Image_V0_Load struct { 
	*Image_V0

	// 
	BaseAddress uint32

	// 
	ImageFileName string

	// 
	ModuleSize uint32
}

	func NewImage_V0_LoadEx1(instance *cim.WmiInstance) (newInstance *Image_V0_Load, err error) {tmp, err := NewImage_V0Ex1(instance)
		
	if err != nil { return }
	newInstance = &Image_V0_Load {
	Image_V0: tmp,
	}
	return
	}
	

	func NewImage_V0_LoadEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *Image_V0_Load, err error) {tmp, err := NewImage_V0Ex6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &Image_V0_Load {
	Image_V0: tmp,
	}
	return
	}
	

// SetBaseAddress sets the value of BaseAddress for the instance
func (instance *Image_V0_Load) SetPropertyBaseAddress(value uint32) (err error) { 
    return instance.SetProperty("BaseAddress", (value))
}

// GetBaseAddress gets the value of BaseAddress for the instance
func (instance *Image_V0_Load) GetPropertyBaseAddress()(value uint32, err error) { 
    retValue, err := instance.GetProperty("BaseAddress")
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

// SetImageFileName sets the value of ImageFileName for the instance
func (instance *Image_V0_Load) SetPropertyImageFileName(value string) (err error) { 
    return instance.SetProperty("ImageFileName", (value))
}

// GetImageFileName gets the value of ImageFileName for the instance
func (instance *Image_V0_Load) GetPropertyImageFileName()(value string, err error) { 
    retValue, err := instance.GetProperty("ImageFileName")
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

// SetModuleSize sets the value of ModuleSize for the instance
func (instance *Image_V0_Load) SetPropertyModuleSize(value uint32) (err error) { 
    return instance.SetProperty("ModuleSize", (value))
}

// GetModuleSize gets the value of ModuleSize for the instance
func (instance *Image_V0_Load) GetPropertyModuleSize()(value uint32, err error) { 
    retValue, err := instance.GetProperty("ModuleSize")
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

