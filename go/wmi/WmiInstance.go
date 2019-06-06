package wmi

type WmiInstance interface {
	GetInstance() string;
	GetProperty() string;
	ResetProperty() string;
	Class() *WmiClass;
	EmbeddedInstance() string;
	InstanceManager() *WmiInstanceManager;
	Equals(*WmiInstance) bool;
	Refresh();
	Commit();
	Modify();
	Delete();
	InstancePath() string;
	InvokeMethod(methodName string, propertyValues map[string]string ) (string, error);
	GetRelated(resultClassName string) *[]WmiInstance;
	GetRelated(resultClassName, associatedClassName, resultRole, sourceRole string) *[]WmiInstance;
	GetAssociated(resultClassName, associatedClassName, resultRole, sourceRole string) *[]WmiInstance;
	EnumerateReferencingInstances(associatedClassName, sourceRole string) *[]WmiInstance;
	Dispose();
}