// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

// CIM_SettingsDefineCapabilities struct
type CIM_SettingsDefineCapabilities struct {
	CIM_Component

	// PropertyPolicy defines whether or not the non-null, non-key properties of the associated SettingData instance are treated independently or as a correlated set. For instance, an independent set of maximum properties might be defined, when there is no relationship between each property. On the other hand, several correlated sets of maximum properties might need to be defined when the maximum values of each are dependent on some of the others.
	PropertyPolicy SettingsDefineCapabilities_PropertyPolicy

	// The ValueRange property indicates further semantics on the interpretation of all non-null, non-key properties of the Component SettingData.
	///"Point" indicates that this SettingData instance provides a single set of values.
	///"Minimums", "Maximums", and "Increments", are only evaluated against non-null, non-key, non-enumerated, non-boolean, numeric properties of the SettingData instance. Each property of that set shall be mathematically comparable to other instances of that property.
	///"Minimums" indicates that this SettingData instance provides minimum values for evaluated properties. When used with PropertyPolicy = "Independent", only one such setting per particular SettingData instance shall be specified for any Capabilities. Unless restricted by a Maximums on the same set of properties, all values that compare higher than the specified values are also considered to be supported by the associated capabilities instance.
	///"Maximums" indicates that this SettingData instance provides maximum values for evaluated properties. When used with PropertyPolicy = "Independent", only one such setting per particular SettingData instance shall be specified for any Capabilities. Unless restricted by a Minimums on the same set of properties, all values that compare lower than the specified values are also considered to be supported by the associated capabilities instance.
	///"Increments" indicates that this SettingData instance provides increment values for evaluated properties. For the associated Capabilities, if an evaluated property currently has no corresponding minimums or maximums, then the property has no affect. Otherwise, for each evaluated property: its value x shall be between the minimum and maximum, inclusively, and shall have the property that both the result of maximum minus x and the result of x minus minimum are each an integer multiple of the increment. If either minimum or maximum is not specified and the other is, then the missing value shall be respectively assumed to be the lowest or highest supported value for the property's data-type. Additionally, if both a minimum and a maximum are specified for an evaluated property, then the result of maximum minus minimum shall be an integer multiple of the increment.
	ValueRange SettingsDefineCapabilities_ValueRange

	// The ValueRole property indicates further semantics on the interpretation of the non-null, non-key properties of the Component SettingData.
	///"Default" indicates that property values of the component SettingData instance will be used as default values, when a new SettingData instance is created for elements whose capabilities are defined by the associated Capabilities instance.
	///Across instances of settingdata, for particular properties having the same semantic purpose, at most one such settingdata instance shall be specified as a default.
	///"Optimal" indicates that the SettingData instance represents optimal setting values for elements associated with the associated capabilities instance. Multiple component SettingData instances may be declared as optimal."Mean" indicates that the non-null, non-key, non-enumerated, non-boolean, numeric properties of the associated SettingData instance represents an average point along some dimension. For different combinations of SettingData properties, multiple component SettingData instances may be declared as "Mean". "Supported" indicates that the non-null, non-key properties of the Component SettingData instance represents a set of supported property values that are not otherwise qualified.
	ValueRole SettingsDefineCapabilities_ValueRole
}

// SetPropertyPolicy sets the value of PropertyPolicy for the instance
func (instance *CIM_SettingsDefineCapabilities) SetPropertyPropertyPolicy(value SettingsDefineCapabilities_PropertyPolicy) (err error) {
	return instance.SetProperty("PropertyPolicy", value)
}

// GetPropertyPolicy gets the value of PropertyPolicy for the instance
func (instance *CIM_SettingsDefineCapabilities) GetPropertyPropertyPolicy() (value SettingsDefineCapabilities_PropertyPolicy, err error) {
	retValue, err := instance.GetProperty("PropertyPolicy")
	if err != nil {
		return
	}
	value, ok := retValue.(SettingsDefineCapabilities_PropertyPolicy)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetValueRange sets the value of ValueRange for the instance
func (instance *CIM_SettingsDefineCapabilities) SetPropertyValueRange(value SettingsDefineCapabilities_ValueRange) (err error) {
	return instance.SetProperty("ValueRange", value)
}

// GetValueRange gets the value of ValueRange for the instance
func (instance *CIM_SettingsDefineCapabilities) GetPropertyValueRange() (value SettingsDefineCapabilities_ValueRange, err error) {
	retValue, err := instance.GetProperty("ValueRange")
	if err != nil {
		return
	}
	value, ok := retValue.(SettingsDefineCapabilities_ValueRange)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetValueRole sets the value of ValueRole for the instance
func (instance *CIM_SettingsDefineCapabilities) SetPropertyValueRole(value SettingsDefineCapabilities_ValueRole) (err error) {
	return instance.SetProperty("ValueRole", value)
}

// GetValueRole gets the value of ValueRole for the instance
func (instance *CIM_SettingsDefineCapabilities) GetPropertyValueRole() (value SettingsDefineCapabilities_ValueRole, err error) {
	retValue, err := instance.GetProperty("ValueRole")
	if err != nil {
		return
	}
	value, ok := retValue.(SettingsDefineCapabilities_ValueRole)
	if !ok {
		// TODO: Set an error
	}
	return
}
