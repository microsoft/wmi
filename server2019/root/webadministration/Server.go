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

// Server struct
type Server struct { 
	*ConfiguredObject

	// 
	ApplicationDefaults ApplicationElementDefaults

	// 
	ApplicationPoolDefaults ApplicationPoolElementDefaults

	// 
	SiteDefaults SiteElementDefaults

	// 
	VirtualDirectoryDefaults VirtualDirectoryElementDefaults
}

	func NewServerEx1(instance *cim.WmiInstance) (newInstance *Server, err error) {tmp, err := NewConfiguredObjectEx1(instance)
		
	if err != nil { return }
	newInstance = &Server {
	ConfiguredObject: tmp,
	}
	return
	}
	

	func NewServerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *Server, err error) {tmp, err := NewConfiguredObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &Server {
	ConfiguredObject: tmp,
	}
	return
	}
	

// SetApplicationDefaults sets the value of ApplicationDefaults for the instance
func (instance *Server) SetPropertyApplicationDefaults(value ApplicationElementDefaults) (err error) { 
    return instance.SetProperty("ApplicationDefaults", (value))
}

// GetApplicationDefaults gets the value of ApplicationDefaults for the instance
func (instance *Server) GetPropertyApplicationDefaults()(value ApplicationElementDefaults, err error) { 
    retValue, err := instance.GetProperty("ApplicationDefaults")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(ApplicationElementDefaults); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " ApplicationElementDefaults is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = ApplicationElementDefaults(valuetmp)

    return
}

// SetApplicationPoolDefaults sets the value of ApplicationPoolDefaults for the instance
func (instance *Server) SetPropertyApplicationPoolDefaults(value ApplicationPoolElementDefaults) (err error) { 
    return instance.SetProperty("ApplicationPoolDefaults", (value))
}

// GetApplicationPoolDefaults gets the value of ApplicationPoolDefaults for the instance
func (instance *Server) GetPropertyApplicationPoolDefaults()(value ApplicationPoolElementDefaults, err error) { 
    retValue, err := instance.GetProperty("ApplicationPoolDefaults")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(ApplicationPoolElementDefaults); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " ApplicationPoolElementDefaults is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = ApplicationPoolElementDefaults(valuetmp)

    return
}

// SetSiteDefaults sets the value of SiteDefaults for the instance
func (instance *Server) SetPropertySiteDefaults(value SiteElementDefaults) (err error) { 
    return instance.SetProperty("SiteDefaults", (value))
}

// GetSiteDefaults gets the value of SiteDefaults for the instance
func (instance *Server) GetPropertySiteDefaults()(value SiteElementDefaults, err error) { 
    retValue, err := instance.GetProperty("SiteDefaults")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(SiteElementDefaults); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " SiteElementDefaults is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = SiteElementDefaults(valuetmp)

    return
}

// SetVirtualDirectoryDefaults sets the value of VirtualDirectoryDefaults for the instance
func (instance *Server) SetPropertyVirtualDirectoryDefaults(value VirtualDirectoryElementDefaults) (err error) { 
    return instance.SetProperty("VirtualDirectoryDefaults", (value))
}

// GetVirtualDirectoryDefaults gets the value of VirtualDirectoryDefaults for the instance
func (instance *Server) GetPropertyVirtualDirectoryDefaults()(value VirtualDirectoryElementDefaults, err error) { 
    retValue, err := instance.GetProperty("VirtualDirectoryDefaults")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(VirtualDirectoryElementDefaults); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " VirtualDirectoryElementDefaults is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = VirtualDirectoryElementDefaults(valuetmp)

    return
}

// 

// <param name="Path" type="string "></param>
func (instance *Server) BeginUpdateBatch( /* IN */ Path string) (err error) {_, err = instance.InvokeMethodWithReturn("BeginUpdateBatch" , Path);
	if err != nil { return }
	return
	
}


// 

// <param name="DoCommitChanges" type="bool "></param>
func (instance *Server) EndUpdateBatch( /* OPTIONAL IN */ DoCommitChanges bool) (err error) {_, err = instance.InvokeMethodWithReturn("EndUpdateBatch" , DoCommitChanges);
	if err != nil { return }
	return
	
}


// 

// <param name="Name" type="string "></param>

// <param name="ReturnValue" type="string "></param>
func (instance *Server) GetMetadata( /* IN */ Name string) (result string, err error) {retVal, err := instance.InvokeMethodWithReturn("GetMetadata" , Name);
	if err != nil { return }
	result = string(retVal)
	return
	
}


// 

// <param name="Name" type="string "></param>
// <param name="Value" type="string "></param>
func (instance *Server) SetMetadata( /* IN */ Name string,
 /* IN */ Value string) (err error) {_, err = instance.InvokeMethodWithReturn("SetMetadata" , Name, Value);
	if err != nil { return }
	return
	
}


