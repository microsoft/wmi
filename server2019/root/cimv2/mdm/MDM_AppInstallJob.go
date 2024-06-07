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

// MDM_AppInstallJob struct
type MDM_AppInstallJob struct { 
	*cim.WmiInstance

	// 
	ActionType uint32

	// 
	CreationTime string

	// 
	Dependencies []string

	// 
	DependencyUrlLists []string

	// 
	DeploymentOptions uint32

	// 
	DownloadUrlList []string

	// 
	JobID string

	// 
	JobType uint32

	// 
	LastError uint32

	// 
	LicenseXml string

	// 
	PackageFullName string

	// 
	Progress uint32

	// 
	Status uint32
}

	func NewMDM_AppInstallJobEx1(instance *cim.WmiInstance) (newInstance *MDM_AppInstallJob, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &MDM_AppInstallJob {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewMDM_AppInstallJobEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MDM_AppInstallJob, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MDM_AppInstallJob {
	WmiInstance: tmp,
	}
	return
	}
	

// SetActionType sets the value of ActionType for the instance
func (instance *MDM_AppInstallJob) SetPropertyActionType(value uint32) (err error) { 
    return instance.SetProperty("ActionType", (value))
}

// GetActionType gets the value of ActionType for the instance
func (instance *MDM_AppInstallJob) GetPropertyActionType()(value uint32, err error) { 
    retValue, err := instance.GetProperty("ActionType")
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

// SetCreationTime sets the value of CreationTime for the instance
func (instance *MDM_AppInstallJob) SetPropertyCreationTime(value string) (err error) { 
    return instance.SetProperty("CreationTime", (value))
}

// GetCreationTime gets the value of CreationTime for the instance
func (instance *MDM_AppInstallJob) GetPropertyCreationTime()(value string, err error) { 
    retValue, err := instance.GetProperty("CreationTime")
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

// SetDependencies sets the value of Dependencies for the instance
func (instance *MDM_AppInstallJob) SetPropertyDependencies(value []string) (err error) { 
    return instance.SetProperty("Dependencies", (value))
}

// GetDependencies gets the value of Dependencies for the instance
func (instance *MDM_AppInstallJob) GetPropertyDependencies()(value []string, err error) { 
    retValue, err := instance.GetProperty("Dependencies")
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

// SetDependencyUrlLists sets the value of DependencyUrlLists for the instance
func (instance *MDM_AppInstallJob) SetPropertyDependencyUrlLists(value []string) (err error) { 
    return instance.SetProperty("DependencyUrlLists", (value))
}

// GetDependencyUrlLists gets the value of DependencyUrlLists for the instance
func (instance *MDM_AppInstallJob) GetPropertyDependencyUrlLists()(value []string, err error) { 
    retValue, err := instance.GetProperty("DependencyUrlLists")
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

// SetDeploymentOptions sets the value of DeploymentOptions for the instance
func (instance *MDM_AppInstallJob) SetPropertyDeploymentOptions(value uint32) (err error) { 
    return instance.SetProperty("DeploymentOptions", (value))
}

// GetDeploymentOptions gets the value of DeploymentOptions for the instance
func (instance *MDM_AppInstallJob) GetPropertyDeploymentOptions()(value uint32, err error) { 
    retValue, err := instance.GetProperty("DeploymentOptions")
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

// SetDownloadUrlList sets the value of DownloadUrlList for the instance
func (instance *MDM_AppInstallJob) SetPropertyDownloadUrlList(value []string) (err error) { 
    return instance.SetProperty("DownloadUrlList", (value))
}

// GetDownloadUrlList gets the value of DownloadUrlList for the instance
func (instance *MDM_AppInstallJob) GetPropertyDownloadUrlList()(value []string, err error) { 
    retValue, err := instance.GetProperty("DownloadUrlList")
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

// SetJobID sets the value of JobID for the instance
func (instance *MDM_AppInstallJob) SetPropertyJobID(value string) (err error) { 
    return instance.SetProperty("JobID", (value))
}

// GetJobID gets the value of JobID for the instance
func (instance *MDM_AppInstallJob) GetPropertyJobID()(value string, err error) { 
    retValue, err := instance.GetProperty("JobID")
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

// SetJobType sets the value of JobType for the instance
func (instance *MDM_AppInstallJob) SetPropertyJobType(value uint32) (err error) { 
    return instance.SetProperty("JobType", (value))
}

// GetJobType gets the value of JobType for the instance
func (instance *MDM_AppInstallJob) GetPropertyJobType()(value uint32, err error) { 
    retValue, err := instance.GetProperty("JobType")
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

// SetLastError sets the value of LastError for the instance
func (instance *MDM_AppInstallJob) SetPropertyLastError(value uint32) (err error) { 
    return instance.SetProperty("LastError", (value))
}

// GetLastError gets the value of LastError for the instance
func (instance *MDM_AppInstallJob) GetPropertyLastError()(value uint32, err error) { 
    retValue, err := instance.GetProperty("LastError")
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

// SetLicenseXml sets the value of LicenseXml for the instance
func (instance *MDM_AppInstallJob) SetPropertyLicenseXml(value string) (err error) { 
    return instance.SetProperty("LicenseXml", (value))
}

// GetLicenseXml gets the value of LicenseXml for the instance
func (instance *MDM_AppInstallJob) GetPropertyLicenseXml()(value string, err error) { 
    retValue, err := instance.GetProperty("LicenseXml")
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
func (instance *MDM_AppInstallJob) SetPropertyPackageFullName(value string) (err error) { 
    return instance.SetProperty("PackageFullName", (value))
}

// GetPackageFullName gets the value of PackageFullName for the instance
func (instance *MDM_AppInstallJob) GetPropertyPackageFullName()(value string, err error) { 
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

// SetProgress sets the value of Progress for the instance
func (instance *MDM_AppInstallJob) SetPropertyProgress(value uint32) (err error) { 
    return instance.SetProperty("Progress", (value))
}

// GetProgress gets the value of Progress for the instance
func (instance *MDM_AppInstallJob) GetPropertyProgress()(value uint32, err error) { 
    retValue, err := instance.GetProperty("Progress")
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

// SetStatus sets the value of Status for the instance
func (instance *MDM_AppInstallJob) SetPropertyStatus(value uint32) (err error) { 
    return instance.SetProperty("Status", (value))
}

// GetStatus gets the value of Status for the instance
func (instance *MDM_AppInstallJob) GetPropertyStatus()(value uint32, err error) { 
    retValue, err := instance.GetProperty("Status")
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

// 

// <param name="JobData" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MDM_AppInstallJob) CreateJob( /* IN */ JobData string) (result uint32, err error) {retVal, err := instance.InvokeMethodWithReturn("CreateJob" , JobData);
	if err != nil { return }
	result = uint32(retVal)
	return
	
}


