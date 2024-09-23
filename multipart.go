package sdkgen

import (
	"bytes"
	"io"
	"mime/multipart"
)

type Multipart struct {
	files       []FilePart
	fields      []FieldPart
	contentType string
}

func (multi *Multipart) AddFile(name string, fileName string, reader io.Reader) {
	multi.files = append(multi.files, FilePart{name, fileName, reader})
}

func (multi *Multipart) AddField(name string, reader io.Reader) {
	multi.fields = append(multi.fields, FieldPart{name, reader})
}

func (multi *Multipart) GetContentType() string {
	return multi.contentType
}

func (multi *Multipart) Build() *bytes.Buffer {
	var reqBody = &bytes.Buffer{}
	writer := multipart.NewWriter(reqBody)

	multi.contentType = writer.FormDataContentType()

	for _, file := range multi.files {
		part, _ := writer.CreateFormFile(file.Name, file.FileName)
		io.Copy(part, file.Reader)
	}

	for _, field := range multi.fields {
		part, _ := writer.CreateFormField(field.Name)
		io.Copy(part, field.Reader)
	}

	writer.Close()

	return reqBody
}

type FilePart struct {
	Name     string
	FileName string
	Reader   io.Reader
}

type FieldPart struct {
	Name   string
	Reader io.Reader
}
