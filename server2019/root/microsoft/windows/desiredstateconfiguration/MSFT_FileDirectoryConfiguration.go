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

// MSFT_FileDirectoryConfiguration struct
type MSFT_FileDirectoryConfiguration struct { 
	*OMI_BaseResource

	// 
	Attributes []string

	// 
	Checksum string

	// 
	Contents string

	// 
	CreatedDate string

	// 
	Credential MSFT_Credential

	// 
	DestinationPath string

	// 
	Ensure string

	// 
	Force bool

	// 
	MatchSource bool

	// 
	ModifiedDate string

	// 
	Recurse bool

	// 
	Size uint64

	// 
	SourcePath string

	// 
	SubItems []MSFT_FileDirectoryConfiguration

	// 
	Type string
}

	func NewMSFT_FileDirectoryConfigurationEx1(instance *cim.WmiInstance) (newInstance *MSFT_FileDirectoryConfiguration, err error) {tmp, err := NewOMI_BaseResourceEx1(instance)
		
	if err != nil { return }
	newInstance = &MSFT_FileDirectoryConfiguration {
	OMI_BaseResource: tmp,
	}
	return
	}
	

	func NewMSFT_FileDirectoryConfigurationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSFT_FileDirectoryConfiguration, err error) {tmp, err := NewOMI_BaseResourceEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSFT_FileDirectoryConfiguration {
	OMI_BaseResource: tmp,
	}
	return
	}
	

// SetAttributes sets the value of Attributes for the instance
func (instance *MSFT_FileDirectoryConfiguration) SetPropertyAttributes(value []string) (err error) { 
    return instance.SetProperty("Attributes", (value))
}

// GetAttributes gets the value of Attributes for the instance
func (instance *MSFT_FileDirectoryConfiguration) GetPropertyAttributes()(value []string, err error) { 
    retValue, err := instance.GetProperty("Attributes")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(string); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, string(valuetmp))
    }

    return
}

// SetChecksum sets the value of Checksum for the instance
func (instance *MSFT_FileDirectoryConfiguration) SetPropertyChecksum(value string) (err error) { 
    return instance.SetProperty("Checksum", (value))
}

// GetChecksum gets the value of Checksum for the instance
func (instance *MSFT_FileDirectoryConfiguration) GetPropertyChecksum()(value string, err error) { 
    retValue, err := instance.GetProperty("Checksum")
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

// SetContents sets the value of Contents for the instance
func (instance *MSFT_FileDirectoryConfiguration) SetPropertyContents(value string) (err error) { 
    return instance.SetProperty("Contents", (value))
}

// GetContents gets the value of Contents for the instance
func (instance *MSFT_FileDirectoryConfiguration) GetPropertyContents()(value string, err error) { 
    retValue, err := instance.GetProperty("Contents")
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

// SetCreatedDate sets the value of CreatedDate for the instance
func (instance *MSFT_FileDirectoryConfiguration) SetPropertyCreatedDate(value string) (err error) { 
    return instance.SetProperty("CreatedDate", (value))
}

// GetCreatedDate gets the value of CreatedDate for the instance
func (instance *MSFT_FileDirectoryConfiguration) GetPropertyCreatedDate()(value string, err error) { 
    retValue, err := instance.GetProperty("CreatedDate")
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

// SetCredential sets the value of Credential for the instance
func (instance *MSFT_FileDirectoryConfiguration) SetPropertyCredential(value MSFT_Credential) (err error) { 
    return instance.SetProperty("Credential", (value))
}

// GetCredential gets the value of Credential for the instance
func (instance *MSFT_FileDirectoryConfiguration) GetPropertyCredential()(value MSFT_Credential, err error) { 
    retValue, err := instance.GetProperty("Credential")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(MSFT_Credential); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " MSFT_Credential is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = MSFT_Credential(valuetmp)

    return
}

// SetDestinationPath sets the value of DestinationPath for the instance
func (instance *MSFT_FileDirectoryConfiguration) SetPropertyDestinationPath(value string) (err error) { 
    return instance.SetProperty("DestinationPath", (value))
}

// GetDestinationPath gets the value of DestinationPath for the instance
func (instance *MSFT_FileDirectoryConfiguration) GetPropertyDestinationPath()(value string, err error) { 
    retValue, err := instance.GetProperty("DestinationPath")
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

// SetEnsure sets the value of Ensure for the instance
func (instance *MSFT_FileDirectoryConfiguration) SetPropertyEnsure(value string) (err error) { 
    return instance.SetProperty("Ensure", (value))
}

// GetEnsure gets the value of Ensure for the instance
func (instance *MSFT_FileDirectoryConfiguration) GetPropertyEnsure()(value string, err error) { 
    retValue, err := instance.GetProperty("Ensure")
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

// SetForce sets the value of Force for the instance
func (instance *MSFT_FileDirectoryConfiguration) SetPropertyForce(value bool) (err error) { 
    return instance.SetProperty("Force", (value))
}

// GetForce gets the value of Force for the instance
func (instance *MSFT_FileDirectoryConfiguration) GetPropertyForce()(value bool, err error) { 
    retValue, err := instance.GetProperty("Force")
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

// SetMatchSource sets the value of MatchSource for the instance
func (instance *MSFT_FileDirectoryConfiguration) SetPropertyMatchSource(value bool) (err error) { 
    return instance.SetProperty("MatchSource", (value))
}

// GetMatchSource gets the value of MatchSource for the instance
func (instance *MSFT_FileDirectoryConfiguration) GetPropertyMatchSource()(value bool, err error) { 
    retValue, err := instance.GetProperty("MatchSource")
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

// SetModifiedDate sets the value of ModifiedDate for the instance
func (instance *MSFT_FileDirectoryConfiguration) SetPropertyModifiedDate(value string) (err error) { 
    return instance.SetProperty("ModifiedDate", (value))
}

// GetModifiedDate gets the value of ModifiedDate for the instance
func (instance *MSFT_FileDirectoryConfiguration) GetPropertyModifiedDate()(value string, err error) { 
    retValue, err := instance.GetProperty("ModifiedDate")
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

// SetRecurse sets the value of Recurse for the instance
func (instance *MSFT_FileDirectoryConfiguration) SetPropertyRecurse(value bool) (err error) { 
    return instance.SetProperty("Recurse", (value))
}

// GetRecurse gets the value of Recurse for the instance
func (instance *MSFT_FileDirectoryConfiguration) GetPropertyRecurse()(value bool, err error) { 
    retValue, err := instance.GetProperty("Recurse")
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

// SetSize sets the value of Size for the instance
func (instance *MSFT_FileDirectoryConfiguration) SetPropertySize(value uint64) (err error) { 
    return instance.SetProperty("Size", (value))
}

// GetSize gets the value of Size for the instance
func (instance *MSFT_FileDirectoryConfiguration) GetPropertySize()(value uint64, err error) { 
    retValue, err := instance.GetProperty("Size")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint64); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint64(valuetmp)

    return
}

// SetSourcePath sets the value of SourcePath for the instance
func (instance *MSFT_FileDirectoryConfiguration) SetPropertySourcePath(value string) (err error) { 
    return instance.SetProperty("SourcePath", (value))
}

// GetSourcePath gets the value of SourcePath for the instance
func (instance *MSFT_FileDirectoryConfiguration) GetPropertySourcePath()(value string, err error) { 
    retValue, err := instance.GetProperty("SourcePath")
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

// SetSubItems sets the value of SubItems for the instance
func (instance *MSFT_FileDirectoryConfiguration) SetPropertySubItems(value []MSFT_FileDirectoryConfiguration) (err error) { 
    return instance.SetProperty("SubItems", (value))
}

// GetSubItems gets the value of SubItems for the instance
func (instance *MSFT_FileDirectoryConfiguration) GetPropertySubItems()(value []MSFT_FileDirectoryConfiguration, err error) { 
    retValue, err := instance.GetProperty("SubItems")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(MSFT_FileDirectoryConfiguration); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " MSFT_FileDirectoryConfiguration is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, MSFT_FileDirectoryConfiguration(valuetmp))
    }

    return
}

// SetType sets the value of Type for the instance
func (instance *MSFT_FileDirectoryConfiguration) SetPropertyType(value string) (err error) { 
    return instance.SetProperty("Type", (value))
}

// GetType gets the value of Type for the instance
func (instance *MSFT_FileDirectoryConfiguration) GetPropertyType()(value string, err error) { 
    retValue, err := instance.GetProperty("Type")
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

// <param name="Flags" type="uint32 "></param>
// <param name="InputResource" type="MSFT_FileDirectoryConfiguration "></param>

// <param name="OutputResource" type="MSFT_FileDirectoryConfiguration "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_FileDirectoryConfiguration) GetTargetResource( /* IN */ InputResource MSFT_FileDirectoryConfiguration,
 /* IN */ Flags uint32,
 /* OUT */ OutputResource MSFT_FileDirectoryConfiguration) (result uint32, err error) {retVal, err := instance.InvokeMethod("GetTargetResource" , InputResource, Flags)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="Flags" type="uint32 "></param>
// <param name="InputResource" type="MSFT_FileDirectoryConfiguration "></param>

// <param name="ProviderContext" type="uint64 "></param>
// <param name="Result" type="bool "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_FileDirectoryConfiguration) TestTargetResource( /* IN */ InputResource MSFT_FileDirectoryConfiguration,
 /* IN */ Flags uint32,
 /* OUT */ Result bool,
 /* OUT */ ProviderContext uint64) (result uint32, err error) {retVal, err := instance.InvokeMethod("TestTargetResource" , InputResource, Flags)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="Flags" type="uint32 "></param>
// <param name="InputResource" type="MSFT_FileDirectoryConfiguration "></param>
// <param name="ProviderContext" type="uint64 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_FileDirectoryConfiguration) SetTargetResource( /* IN */ InputResource MSFT_FileDirectoryConfiguration,
 /* IN */ ProviderContext uint64,
 /* IN */ Flags uint32) (result uint32, err error) {retVal, err := instance.InvokeMethodWithReturn("SetTargetResource" , InputResource, ProviderContext, Flags);
	if err != nil { return }
	result = uint32(retVal)
	return
	
}


