package wmi;

// WmiInstanceManager interface
type WmiInstanceManager interface {
	ServerName() string;
	Namespace() string;
	Credentials() *WmiCredentials;
	EnumerateInstances(className string ) (*[]WmiInstance, error);
	QueryInstances(query string )(*[]WmiInstance , error);
	QueryInstances( query WmiQuery) (*[]WmiInstance, error);

	CreateInstance(className string, propertyValues map[string]string ) (*WmiInstance, error);
	GetInstance( className string, propertyValues map[string]string )(*WmiInstance, error);
	GetClass( className string)(*WmiClass, error);
	EnumerateClasses()(*[]WmiClass, error);
	GetInstancesFromPaths(pathArray string[] )(*[]WmiInstance, error);
}