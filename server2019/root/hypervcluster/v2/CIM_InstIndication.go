// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.HyperVCluster.v2
//////////////////////////////////////////////
package v2

// CIM_InstIndication struct
type CIM_InstIndication struct {
	CIM_Indication

	// A copy of the instance that changed to generate the Indication. SourceInstance contains the current values of the properties selected by the Indication Filter's Query. In the case of CIM_InstDeletion, the property values are copied before the instance is deleted.
	SourceInstance interface{}

	// The host name or IP address of the SourceInstance.
	SourceInstanceHost string

	// The Model Path of the SourceInstance. The following format MUST be used to encode the Model Path:
	///<NamespacePath>:<ClassName>.<Prop1>="<Value1>",
	///<Prop2>="<Value2>", ...
	SourceInstanceModelPath string
}

// SetSourceInstance sets the value of SourceInstance for the instance
func (instance *CIM_InstIndication) SetPropertySourceInstance(value interface{}) (err error) {
	return instance.SetProperty("SourceInstance", value)
}

// GetSourceInstance gets the value of SourceInstance for the instance
func (instance *CIM_InstIndication) GetPropertySourceInstance() (value interface{}, err error) {
	retValue, err := instance.GetProperty("SourceInstance")
	if err != nil {
		return
	}
	value, ok := retValue.(interface{})
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSourceInstanceHost sets the value of SourceInstanceHost for the instance
func (instance *CIM_InstIndication) SetPropertySourceInstanceHost(value string) (err error) {
	return instance.SetProperty("SourceInstanceHost", value)
}

// GetSourceInstanceHost gets the value of SourceInstanceHost for the instance
func (instance *CIM_InstIndication) GetPropertySourceInstanceHost() (value string, err error) {
	retValue, err := instance.GetProperty("SourceInstanceHost")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSourceInstanceModelPath sets the value of SourceInstanceModelPath for the instance
func (instance *CIM_InstIndication) SetPropertySourceInstanceModelPath(value string) (err error) {
	return instance.SetProperty("SourceInstanceModelPath", value)
}

// GetSourceInstanceModelPath gets the value of SourceInstanceModelPath for the instance
func (instance *CIM_InstIndication) GetPropertySourceInstanceModelPath() (value string, err error) {
	retValue, err := instance.GetProperty("SourceInstanceModelPath")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
