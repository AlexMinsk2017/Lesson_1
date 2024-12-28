package main

import (
	"Lesson_1/internal/storage"
	"fmt"
)

func main() {

	st := storage.NewStorage()
	file, err := st.Upload("text.txt", []byte("hello"))
	if err != nil {
		return
	}

	file, err = st.GetByID(file.ID)
	if err != nil {
		return
	}

	fmt.Println(st.Files)

	fmt.Println("restored ", *file)
}
