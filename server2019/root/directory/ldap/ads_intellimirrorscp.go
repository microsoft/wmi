// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.directory.LDAP
//////////////////////////////////////////////
package ldap
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// ads_intellimirrorscp struct
type ads_intellimirrorscp struct { 
	*ads_serviceadministrationpoint

	// 
	DS_netbootAllowNewClients bool

	// 
	DS_netbootAnswerOnlyValidClients bool

	// 
	DS_netbootAnswerRequests bool

	// 
	DS_netbootCurrentClientCount int32

	// 
	DS_netbootIntelliMirrorOSes []string

	// 
	DS_netbootLimitClients bool

	// 
	DS_netbootLocallyInstalledOSes []string

	// 
	DS_netbootMachineFilePath string

	// 
	DS_netbootMaxClients int32

	// 
	DS_netbootNewMachineNamingPolicy []string

	// 
	DS_netbootNewMachineOU string

	// 
	DS_netbootServer string

	// 
	DS_netbootTools []string
}

	func Newads_intellimirrorscpEx1(instance *cim.WmiInstance) (newInstance *ads_intellimirrorscp, err error) {tmp, err := Newads_serviceadministrationpointEx1(instance)
		
	if err != nil { return }
	newInstance = &ads_intellimirrorscp {
	ads_serviceadministrationpoint: tmp,
	}
	return
	}
	

	func Newads_intellimirrorscpEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ads_intellimirrorscp, err error) {tmp, err := Newads_serviceadministrationpointEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ads_intellimirrorscp {
	ads_serviceadministrationpoint: tmp,
	}
	return
	}
	

// SetDS_netbootAllowNewClients sets the value of DS_netbootAllowNewClients for the instance
func (instance *ads_intellimirrorscp) SetPropertyDS_netbootAllowNewClients(value bool) (err error) { 
    return instance.SetProperty("DS_netbootAllowNewClients", (value))
}

// GetDS_netbootAllowNewClients gets the value of DS_netbootAllowNewClients for the instance
func (instance *ads_intellimirrorscp) GetPropertyDS_netbootAllowNewClients()(value bool, err error) { 
    retValue, err := instance.GetProperty("DS_netbootAllowNewClients")
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

// SetDS_netbootAnswerOnlyValidClients sets the value of DS_netbootAnswerOnlyValidClients for the instance
func (instance *ads_intellimirrorscp) SetPropertyDS_netbootAnswerOnlyValidClients(value bool) (err error) { 
    return instance.SetProperty("DS_netbootAnswerOnlyValidClients", (value))
}

// GetDS_netbootAnswerOnlyValidClients gets the value of DS_netbootAnswerOnlyValidClients for the instance
func (instance *ads_intellimirrorscp) GetPropertyDS_netbootAnswerOnlyValidClients()(value bool, err error) { 
    retValue, err := instance.GetProperty("DS_netbootAnswerOnlyValidClients")
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

// SetDS_netbootAnswerRequests sets the value of DS_netbootAnswerRequests for the instance
func (instance *ads_intellimirrorscp) SetPropertyDS_netbootAnswerRequests(value bool) (err error) { 
    return instance.SetProperty("DS_netbootAnswerRequests", (value))
}

// GetDS_netbootAnswerRequests gets the value of DS_netbootAnswerRequests for the instance
func (instance *ads_intellimirrorscp) GetPropertyDS_netbootAnswerRequests()(value bool, err error) { 
    retValue, err := instance.GetProperty("DS_netbootAnswerRequests")
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

// SetDS_netbootCurrentClientCount sets the value of DS_netbootCurrentClientCount for the instance
func (instance *ads_intellimirrorscp) SetPropertyDS_netbootCurrentClientCount(value int32) (err error) { 
    return instance.SetProperty("DS_netbootCurrentClientCount", (value))
}

// GetDS_netbootCurrentClientCount gets the value of DS_netbootCurrentClientCount for the instance
func (instance *ads_intellimirrorscp) GetPropertyDS_netbootCurrentClientCount()(value int32, err error) { 
    retValue, err := instance.GetProperty("DS_netbootCurrentClientCount")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(int32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = int32(valuetmp)

    return
}

// SetDS_netbootIntelliMirrorOSes sets the value of DS_netbootIntelliMirrorOSes for the instance
func (instance *ads_intellimirrorscp) SetPropertyDS_netbootIntelliMirrorOSes(value []string) (err error) { 
    return instance.SetProperty("DS_netbootIntelliMirrorOSes", (value))
}

// GetDS_netbootIntelliMirrorOSes gets the value of DS_netbootIntelliMirrorOSes for the instance
func (instance *ads_intellimirrorscp) GetPropertyDS_netbootIntelliMirrorOSes()(value []string, err error) { 
    retValue, err := instance.GetProperty("DS_netbootIntelliMirrorOSes")
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

// SetDS_netbootLimitClients sets the value of DS_netbootLimitClients for the instance
func (instance *ads_intellimirrorscp) SetPropertyDS_netbootLimitClients(value bool) (err error) { 
    return instance.SetProperty("DS_netbootLimitClients", (value))
}

// GetDS_netbootLimitClients gets the value of DS_netbootLimitClients for the instance
func (instance *ads_intellimirrorscp) GetPropertyDS_netbootLimitClients()(value bool, err error) { 
    retValue, err := instance.GetProperty("DS_netbootLimitClients")
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

// SetDS_netbootLocallyInstalledOSes sets the value of DS_netbootLocallyInstalledOSes for the instance
func (instance *ads_intellimirrorscp) SetPropertyDS_netbootLocallyInstalledOSes(value []string) (err error) { 
    return instance.SetProperty("DS_netbootLocallyInstalledOSes", (value))
}

// GetDS_netbootLocallyInstalledOSes gets the value of DS_netbootLocallyInstalledOSes for the instance
func (instance *ads_intellimirrorscp) GetPropertyDS_netbootLocallyInstalledOSes()(value []string, err error) { 
    retValue, err := instance.GetProperty("DS_netbootLocallyInstalledOSes")
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

// SetDS_netbootMachineFilePath sets the value of DS_netbootMachineFilePath for the instance
func (instance *ads_intellimirrorscp) SetPropertyDS_netbootMachineFilePath(value string) (err error) { 
    return instance.SetProperty("DS_netbootMachineFilePath", (value))
}

// GetDS_netbootMachineFilePath gets the value of DS_netbootMachineFilePath for the instance
func (instance *ads_intellimirrorscp) GetPropertyDS_netbootMachineFilePath()(value string, err error) { 
    retValue, err := instance.GetProperty("DS_netbootMachineFilePath")
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

// SetDS_netbootMaxClients sets the value of DS_netbootMaxClients for the instance
func (instance *ads_intellimirrorscp) SetPropertyDS_netbootMaxClients(value int32) (err error) { 
    return instance.SetProperty("DS_netbootMaxClients", (value))
}

// GetDS_netbootMaxClients gets the value of DS_netbootMaxClients for the instance
func (instance *ads_intellimirrorscp) GetPropertyDS_netbootMaxClients()(value int32, err error) { 
    retValue, err := instance.GetProperty("DS_netbootMaxClients")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(int32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = int32(valuetmp)

    return
}

// SetDS_netbootNewMachineNamingPolicy sets the value of DS_netbootNewMachineNamingPolicy for the instance
func (instance *ads_intellimirrorscp) SetPropertyDS_netbootNewMachineNamingPolicy(value []string) (err error) { 
    return instance.SetProperty("DS_netbootNewMachineNamingPolicy", (value))
}

// GetDS_netbootNewMachineNamingPolicy gets the value of DS_netbootNewMachineNamingPolicy for the instance
func (instance *ads_intellimirrorscp) GetPropertyDS_netbootNewMachineNamingPolicy()(value []string, err error) { 
    retValue, err := instance.GetProperty("DS_netbootNewMachineNamingPolicy")
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

// SetDS_netbootNewMachineOU sets the value of DS_netbootNewMachineOU for the instance
func (instance *ads_intellimirrorscp) SetPropertyDS_netbootNewMachineOU(value string) (err error) { 
    return instance.SetProperty("DS_netbootNewMachineOU", (value))
}

// GetDS_netbootNewMachineOU gets the value of DS_netbootNewMachineOU for the instance
func (instance *ads_intellimirrorscp) GetPropertyDS_netbootNewMachineOU()(value string, err error) { 
    retValue, err := instance.GetProperty("DS_netbootNewMachineOU")
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

// SetDS_netbootServer sets the value of DS_netbootServer for the instance
func (instance *ads_intellimirrorscp) SetPropertyDS_netbootServer(value string) (err error) { 
    return instance.SetProperty("DS_netbootServer", (value))
}

// GetDS_netbootServer gets the value of DS_netbootServer for the instance
func (instance *ads_intellimirrorscp) GetPropertyDS_netbootServer()(value string, err error) { 
    retValue, err := instance.GetProperty("DS_netbootServer")
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

// SetDS_netbootTools sets the value of DS_netbootTools for the instance
func (instance *ads_intellimirrorscp) SetPropertyDS_netbootTools(value []string) (err error) { 
    return instance.SetProperty("DS_netbootTools", (value))
}

// GetDS_netbootTools gets the value of DS_netbootTools for the instance
func (instance *ads_intellimirrorscp) GetPropertyDS_netbootTools()(value []string, err error) { 
    retValue, err := instance.GetProperty("DS_netbootTools")
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

