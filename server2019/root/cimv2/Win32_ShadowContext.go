// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_ShadowContext struct
type Win32_ShadowContext struct {
	CIM_Setting

	//
	ClientAccessible bool

	//
	Differential bool

	//
	ExposedLocally bool

	//
	ExposedRemotely bool

	//
	HardwareAssisted bool

	//
	Imported bool

	//
	Name string

	//
	NoAutoRelease bool

	//
	NotSurfaced bool

	//
	NoWriters bool

	//
	Persistent bool

	//
	Plex bool

	//
	Transportable bool
}

// SetClientAccessible sets the value of ClientAccessible for the instance
func (instance *Win32_ShadowContext) SetPropertyClientAccessible(value bool) (err error) {
	return instance.SetProperty("ClientAccessible", value)
}

// GetClientAccessible gets the value of ClientAccessible for the instance
func (instance *Win32_ShadowContext) GetPropertyClientAccessible() (value bool, err error) {
	retValue, err := instance.GetProperty("ClientAccessible")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDifferential sets the value of Differential for the instance
func (instance *Win32_ShadowContext) SetPropertyDifferential(value bool) (err error) {
	return instance.SetProperty("Differential", value)
}

// GetDifferential gets the value of Differential for the instance
func (instance *Win32_ShadowContext) GetPropertyDifferential() (value bool, err error) {
	retValue, err := instance.GetProperty("Differential")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetExposedLocally sets the value of ExposedLocally for the instance
func (instance *Win32_ShadowContext) SetPropertyExposedLocally(value bool) (err error) {
	return instance.SetProperty("ExposedLocally", value)
}

// GetExposedLocally gets the value of ExposedLocally for the instance
func (instance *Win32_ShadowContext) GetPropertyExposedLocally() (value bool, err error) {
	retValue, err := instance.GetProperty("ExposedLocally")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetExposedRemotely sets the value of ExposedRemotely for the instance
func (instance *Win32_ShadowContext) SetPropertyExposedRemotely(value bool) (err error) {
	return instance.SetProperty("ExposedRemotely", value)
}

// GetExposedRemotely gets the value of ExposedRemotely for the instance
func (instance *Win32_ShadowContext) GetPropertyExposedRemotely() (value bool, err error) {
	retValue, err := instance.GetProperty("ExposedRemotely")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetHardwareAssisted sets the value of HardwareAssisted for the instance
func (instance *Win32_ShadowContext) SetPropertyHardwareAssisted(value bool) (err error) {
	return instance.SetProperty("HardwareAssisted", value)
}

// GetHardwareAssisted gets the value of HardwareAssisted for the instance
func (instance *Win32_ShadowContext) GetPropertyHardwareAssisted() (value bool, err error) {
	retValue, err := instance.GetProperty("HardwareAssisted")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetImported sets the value of Imported for the instance
func (instance *Win32_ShadowContext) SetPropertyImported(value bool) (err error) {
	return instance.SetProperty("Imported", value)
}

// GetImported gets the value of Imported for the instance
func (instance *Win32_ShadowContext) GetPropertyImported() (value bool, err error) {
	retValue, err := instance.GetProperty("Imported")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetName sets the value of Name for the instance
func (instance *Win32_ShadowContext) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", value)
}

// GetName gets the value of Name for the instance
func (instance *Win32_ShadowContext) GetPropertyName() (value string, err error) {
	retValue, err := instance.GetProperty("Name")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNoAutoRelease sets the value of NoAutoRelease for the instance
func (instance *Win32_ShadowContext) SetPropertyNoAutoRelease(value bool) (err error) {
	return instance.SetProperty("NoAutoRelease", value)
}

// GetNoAutoRelease gets the value of NoAutoRelease for the instance
func (instance *Win32_ShadowContext) GetPropertyNoAutoRelease() (value bool, err error) {
	retValue, err := instance.GetProperty("NoAutoRelease")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNotSurfaced sets the value of NotSurfaced for the instance
func (instance *Win32_ShadowContext) SetPropertyNotSurfaced(value bool) (err error) {
	return instance.SetProperty("NotSurfaced", value)
}

// GetNotSurfaced gets the value of NotSurfaced for the instance
func (instance *Win32_ShadowContext) GetPropertyNotSurfaced() (value bool, err error) {
	retValue, err := instance.GetProperty("NotSurfaced")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNoWriters sets the value of NoWriters for the instance
func (instance *Win32_ShadowContext) SetPropertyNoWriters(value bool) (err error) {
	return instance.SetProperty("NoWriters", value)
}

// GetNoWriters gets the value of NoWriters for the instance
func (instance *Win32_ShadowContext) GetPropertyNoWriters() (value bool, err error) {
	retValue, err := instance.GetProperty("NoWriters")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPersistent sets the value of Persistent for the instance
func (instance *Win32_ShadowContext) SetPropertyPersistent(value bool) (err error) {
	return instance.SetProperty("Persistent", value)
}

// GetPersistent gets the value of Persistent for the instance
func (instance *Win32_ShadowContext) GetPropertyPersistent() (value bool, err error) {
	retValue, err := instance.GetProperty("Persistent")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPlex sets the value of Plex for the instance
func (instance *Win32_ShadowContext) SetPropertyPlex(value bool) (err error) {
	return instance.SetProperty("Plex", value)
}

// GetPlex gets the value of Plex for the instance
func (instance *Win32_ShadowContext) GetPropertyPlex() (value bool, err error) {
	retValue, err := instance.GetProperty("Plex")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTransportable sets the value of Transportable for the instance
func (instance *Win32_ShadowContext) SetPropertyTransportable(value bool) (err error) {
	return instance.SetProperty("Transportable", value)
}

// GetTransportable gets the value of Transportable for the instance
func (instance *Win32_ShadowContext) GetPropertyTransportable() (value bool, err error) {
	retValue, err := instance.GetProperty("Transportable")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}
