// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

// CIM_NetworkService struct
type CIM_NetworkService struct {
	CIM_Service

	// This is a free-form array of strings that provide descriptive words and phrases that can be used in queries. To-date, this property has not been implemented, since it is not standardized. Also, if this was a necessary query construct, then it would be required higher in the inheritance hierarchy. The latter has not proven necessary. Therefore, the property is deprecated.
	Keywords []string

	// This is a URL that provides the protocol, network location, and other service-specific information required in order to access the service. It is deprecated with the recommendation that ServiceAccessURI be instantiated instead. This new class correctly positions the semantics of the service access, and clarifies the format of the information.
	ServiceURL string

	// This is a free-form array of strings that specify any specific pre-conditions that must be met in order for this service to start correctly. It was expected that subclasses would refine the inherited StartService() method to suit their specific needs. To-date, this refinement has not been necessary. Also, the property is not very useful, since it is not standardized. If this was a necessary construct, then it would be required higher in the inheritance hierarchy (on Service). The latter has not proven true. Therefore, the property is deprecated.
	StartupConditions []string

	// This is a free-form array of strings that specify any specific parameters that must be supplied to the StartService() method in order for this service to start correctly. It was expected that subclasses would refine the inherited StartService() methods to suit their specific needs. To-date, this refinement has not been necessary. If indeed the method were refined, then its parameters would more formally convey this information. Therefore, the property is deprecated.
	StartupParameters []string
}

// SetKeywords sets the value of Keywords for the instance
func (instance *CIM_NetworkService) SetPropertyKeywords(value []string) (err error) {
	return instance.SetProperty("Keywords", value)
}

// GetKeywords gets the value of Keywords for the instance
func (instance *CIM_NetworkService) GetPropertyKeywords() (value []string, err error) {
	retValue, err := instance.GetProperty("Keywords")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetServiceURL sets the value of ServiceURL for the instance
func (instance *CIM_NetworkService) SetPropertyServiceURL(value string) (err error) {
	return instance.SetProperty("ServiceURL", value)
}

// GetServiceURL gets the value of ServiceURL for the instance
func (instance *CIM_NetworkService) GetPropertyServiceURL() (value string, err error) {
	retValue, err := instance.GetProperty("ServiceURL")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStartupConditions sets the value of StartupConditions for the instance
func (instance *CIM_NetworkService) SetPropertyStartupConditions(value []string) (err error) {
	return instance.SetProperty("StartupConditions", value)
}

// GetStartupConditions gets the value of StartupConditions for the instance
func (instance *CIM_NetworkService) GetPropertyStartupConditions() (value []string, err error) {
	retValue, err := instance.GetProperty("StartupConditions")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStartupParameters sets the value of StartupParameters for the instance
func (instance *CIM_NetworkService) SetPropertyStartupParameters(value []string) (err error) {
	return instance.SetProperty("StartupParameters", value)
}

// GetStartupParameters gets the value of StartupParameters for the instance
func (instance *CIM_NetworkService) GetPropertyStartupParameters() (value []string, err error) {
	retValue, err := instance.GetProperty("StartupParameters")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}
