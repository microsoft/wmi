// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.CIMV2.mdm
//////////////////////////////////////////////
package mdm
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/base/instance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// MDM_ApplicationMinVersion struct
type MDM_ApplicationMinVersion struct { 
	*cim.WmiInstance

	// 
	MinimumPackageVersion string

	// 
	PackageArchitecture string

	// 
	PackageFullName string

	// 
	PackageName string

	// 
	PackagePublisher string

	// 
	PackageVersion string

	// 
	UserSID string
}

	func NewMDM_ApplicationMinVersionEx1(instance *cim.WmiInstance) (newInstance *MDM_ApplicationMinVersion, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &MDM_ApplicationMinVersion {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewMDM_ApplicationMinVersionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MDM_ApplicationMinVersion, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MDM_ApplicationMinVersion {
	WmiInstance: tmp,
	}
	return
	}
	

// SetMinimumPackageVersion sets the value of MinimumPackageVersion for the instance
func (instance *MDM_ApplicationMinVersion) SetPropertyMinimumPackageVersion(value string) (err error) { 
    return instance.SetProperty("MinimumPackageVersion", (value))
}

// GetMinimumPackageVersion gets the value of MinimumPackageVersion for the instance
func (instance *MDM_ApplicationMinVersion) GetPropertyMinimumPackageVersion()(value string, err error) { 
    retValue, err := instance.GetProperty("MinimumPackageVersion")
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

// SetPackageArchitecture sets the value of PackageArchitecture for the instance
func (instance *MDM_ApplicationMinVersion) SetPropertyPackageArchitecture(value string) (err error) { 
    return instance.SetProperty("PackageArchitecture", (value))
}

// GetPackageArchitecture gets the value of PackageArchitecture for the instance
func (instance *MDM_ApplicationMinVersion) GetPropertyPackageArchitecture()(value string, err error) { 
    retValue, err := instance.GetProperty("PackageArchitecture")
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

// SetPackageFullName sets the value of PackageFullName for the instance
func (instance *MDM_ApplicationMinVersion) SetPropertyPackageFullName(value string) (err error) { 
    return instance.SetProperty("PackageFullName", (value))
}

// GetPackageFullName gets the value of PackageFullName for the instance
func (instance *MDM_ApplicationMinVersion) GetPropertyPackageFullName()(value string, err error) { 
    retValue, err := instance.GetProperty("PackageFullName")
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

// SetPackageName sets the value of PackageName for the instance
func (instance *MDM_ApplicationMinVersion) SetPropertyPackageName(value string) (err error) { 
    return instance.SetProperty("PackageName", (value))
}

// GetPackageName gets the value of PackageName for the instance
func (instance *MDM_ApplicationMinVersion) GetPropertyPackageName()(value string, err error) { 
    retValue, err := instance.GetProperty("PackageName")
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

// SetPackagePublisher sets the value of PackagePublisher for the instance
func (instance *MDM_ApplicationMinVersion) SetPropertyPackagePublisher(value string) (err error) { 
    return instance.SetProperty("PackagePublisher", (value))
}

// GetPackagePublisher gets the value of PackagePublisher for the instance
func (instance *MDM_ApplicationMinVersion) GetPropertyPackagePublisher()(value string, err error) { 
    retValue, err := instance.GetProperty("PackagePublisher")
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

// SetPackageVersion sets the value of PackageVersion for the instance
func (instance *MDM_ApplicationMinVersion) SetPropertyPackageVersion(value string) (err error) { 
    return instance.SetProperty("PackageVersion", (value))
}

// GetPackageVersion gets the value of PackageVersion for the instance
func (instance *MDM_ApplicationMinVersion) GetPropertyPackageVersion()(value string, err error) { 
    retValue, err := instance.GetProperty("PackageVersion")
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

// SetUserSID sets the value of UserSID for the instance
func (instance *MDM_ApplicationMinVersion) SetPropertyUserSID(value string) (err error) { 
    return instance.SetProperty("UserSID", (value))
}

// GetUserSID gets the value of UserSID for the instance
func (instance *MDM_ApplicationMinVersion) GetPropertyUserSID()(value string, err error) { 
    retValue, err := instance.GetProperty("UserSID")
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

