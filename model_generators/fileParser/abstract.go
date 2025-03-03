package fileparser


//TODO: Удалить и перенести в отдельный файл
type Model struct {

}

type IOParser interface {
	Write(model *Model, name string);
	Read(name string) model *Model;
}
