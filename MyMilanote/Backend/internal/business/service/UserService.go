package service

import Domains "Backend/internal/business/domains"

type UserService struct{
	//TODO
}

func (service *UserService) GetObjects(projectID string) ([]Domains.Object, error) {
	//
}
CreateObject(object *Object) error
UpdateObject(object *Object) error
DeleteObject(object *Object) error
ConnectObjects(connection *ObjectConnection) error
