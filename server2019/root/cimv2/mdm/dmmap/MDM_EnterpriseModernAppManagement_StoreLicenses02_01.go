// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.CIMV2.mdm.dmmap
//////////////////////////////////////////////
package dmmap
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/base/instance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// MDM_EnterpriseModernAppManagement_StoreLicenses02_01 struct
type MDM_EnterpriseModernAppManagement_StoreLicenses02_01 struct { 
	*cim.WmiInstance

	// 
	InstanceID string

	// 
	LicenseCategory string

	// 
	LicenseUsage string

	// 
	ParentID string

	// 
	RequesterID string
}

	func NewMDM_EnterpriseModernAppManagement_StoreLicenses02_01Ex1(instance *cim.WmiInstance) (newInstance *MDM_EnterpriseModernAppManagement_StoreLicenses02_01, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &MDM_EnterpriseModernAppManagement_StoreLicenses02_01 {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewMDM_EnterpriseModernAppManagement_StoreLicenses02_01Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MDM_EnterpriseModernAppManagement_StoreLicenses02_01, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MDM_EnterpriseModernAppManagement_StoreLicenses02_01 {
	WmiInstance: tmp,
	}
	return
	}
	

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_EnterpriseModernAppManagement_StoreLicenses02_01) SetPropertyInstanceID(value string) (err error) { 
    return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_EnterpriseModernAppManagement_StoreLicenses02_01) GetPropertyInstanceID()(value string, err error) { 
    retValue, err := instance.GetProperty("InstanceID")
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

// SetLicenseCategory sets the value of LicenseCategory for the instance
func (instance *MDM_EnterpriseModernAppManagement_StoreLicenses02_01) SetPropertyLicenseCategory(value string) (err error) { 
    return instance.SetProperty("LicenseCategory", (value))
}

// GetLicenseCategory gets the value of LicenseCategory for the instance
func (instance *MDM_EnterpriseModernAppManagement_StoreLicenses02_01) GetPropertyLicenseCategory()(value string, err error) { 
    retValue, err := instance.GetProperty("LicenseCategory")
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

// SetLicenseUsage sets the value of LicenseUsage for the instance
func (instance *MDM_EnterpriseModernAppManagement_StoreLicenses02_01) SetPropertyLicenseUsage(value string) (err error) { 
    return instance.SetProperty("LicenseUsage", (value))
}

// GetLicenseUsage gets the value of LicenseUsage for the instance
func (instance *MDM_EnterpriseModernAppManagement_StoreLicenses02_01) GetPropertyLicenseUsage()(value string, err error) { 
    retValue, err := instance.GetProperty("LicenseUsage")
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

// SetParentID sets the value of ParentID for the instance
func (instance *MDM_EnterpriseModernAppManagement_StoreLicenses02_01) SetPropertyParentID(value string) (err error) { 
    return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_EnterpriseModernAppManagement_StoreLicenses02_01) GetPropertyParentID()(value string, err error) { 
    retValue, err := instance.GetProperty("ParentID")
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

// SetRequesterID sets the value of RequesterID for the instance
func (instance *MDM_EnterpriseModernAppManagement_StoreLicenses02_01) SetPropertyRequesterID(value string) (err error) { 
    return instance.SetProperty("RequesterID", (value))
}

// GetRequesterID gets the value of RequesterID for the instance
func (instance *MDM_EnterpriseModernAppManagement_StoreLicenses02_01) GetPropertyRequesterID()(value string, err error) { 
    retValue, err := instance.GetProperty("RequesterID")
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

// 

// <param name="param" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MDM_EnterpriseModernAppManagement_StoreLicenses02_01) AddLicenseMethod( /* IN */ param string) (result uint32, err error) {retVal, err := instance.InvokeMethodWithReturn("AddLicenseMethod" , param);
	if err != nil { return }
	result = uint32(retVal)
	return
	
}


// 

// <param name="param" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MDM_EnterpriseModernAppManagement_StoreLicenses02_01) GetLicenseFromStoreMethod( /* IN */ param string) (result uint32, err error) {retVal, err := instance.InvokeMethodWithReturn("GetLicenseFromStoreMethod" , param);
	if err != nil { return }
	result = uint32(retVal)
	return
	
}


