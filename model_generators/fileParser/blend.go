package fileparser

import (
	"os"
)

type BlendParser struct {

}

func (this BlendParser) Write(model *Model, name string) {
	file, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE, 0644);
	if err != nil {
		//TODO: Верни ошибку
	}
	file.Write([]byte("BLENDER-v402"));
	//TODO: 
	sizeof_bhead := 24;

}
