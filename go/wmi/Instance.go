package wmi

type Instance interface {
	GetInstance() string
	GetProperty() string
	ResetProperty() string
	Class() *Class
	EmbeddedInstance() string
	InstanceManager() *InstanceManager
	Equals(*Instance) bool
	Refresh()
	Commit()
	Modify()
	Delete()
	InstancePath() string
	InvokeMethod(methodName string, propertyValues map[string]string) (string, error)
	GetRelated(resultClassName string) *[]Instance
	GetRelatedEx(resultClassName, associatedClassName, resultRole, sourceRole string) *[]Instance
	GetAssociated(resultClassName, associatedClassName, resultRole, sourceRole string) *[]Instance
	EnumerateReferencingInstances(associatedClassName, sourceRole string) *[]Instance
	Dispose()
}
