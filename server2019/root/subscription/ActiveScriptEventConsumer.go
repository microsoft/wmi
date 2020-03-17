// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.subscription
//////////////////////////////////////////////
package subscription

// ActiveScriptEventConsumer struct
type ActiveScriptEventConsumer struct {
	__EventConsumer

	//
	KillTimeout uint32

	//
	Name string

	//
	ScriptFilename string

	//
	ScriptingEngine string

	//
	ScriptText string
}

// SetKillTimeout sets the value of KillTimeout for the instance
func (instance *ActiveScriptEventConsumer) SetPropertyKillTimeout(value uint32) (err error) {
	return instance.SetProperty("KillTimeout", value)
}

// GetKillTimeout gets the value of KillTimeout for the instance
func (instance *ActiveScriptEventConsumer) GetPropertyKillTimeout() (value uint32, err error) {
	retValue, err := instance.GetProperty("KillTimeout")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetName sets the value of Name for the instance
func (instance *ActiveScriptEventConsumer) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", value)
}

// GetName gets the value of Name for the instance
func (instance *ActiveScriptEventConsumer) GetPropertyName() (value string, err error) {
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

// SetScriptFilename sets the value of ScriptFilename for the instance
func (instance *ActiveScriptEventConsumer) SetPropertyScriptFilename(value string) (err error) {
	return instance.SetProperty("ScriptFilename", value)
}

// GetScriptFilename gets the value of ScriptFilename for the instance
func (instance *ActiveScriptEventConsumer) GetPropertyScriptFilename() (value string, err error) {
	retValue, err := instance.GetProperty("ScriptFilename")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetScriptingEngine sets the value of ScriptingEngine for the instance
func (instance *ActiveScriptEventConsumer) SetPropertyScriptingEngine(value string) (err error) {
	return instance.SetProperty("ScriptingEngine", value)
}

// GetScriptingEngine gets the value of ScriptingEngine for the instance
func (instance *ActiveScriptEventConsumer) GetPropertyScriptingEngine() (value string, err error) {
	retValue, err := instance.GetProperty("ScriptingEngine")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetScriptText sets the value of ScriptText for the instance
func (instance *ActiveScriptEventConsumer) SetPropertyScriptText(value string) (err error) {
	return instance.SetProperty("ScriptText", value)
}

// GetScriptText gets the value of ScriptText for the instance
func (instance *ActiveScriptEventConsumer) GetPropertyScriptText() (value string, err error) {
	retValue, err := instance.GetProperty("ScriptText")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
