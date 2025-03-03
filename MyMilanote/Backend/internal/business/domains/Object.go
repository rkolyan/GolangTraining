package domains

type Object struct {
	//TODO
}

type ObjectConnection struct {
	//TODO
}

type IObjectService interface {
	GetObjects(projectID string) ([]Object, error)
	CreateObject(object *Object) error
	UpdateObject(object *Object) error
	DeleteObject(object *Object) error
	ConnectObjects(connection *ObjectConnection) error
}

type IObjectRepository interface {
}
