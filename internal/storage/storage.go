package storage

import (
	"errors"
	"fmt"
	"github.com/AlexMinsk2017/Lesson_1/internal/file"
	"github.com/google/uuid"
)

type Storage struct {
	Files map[uuid.UUID]*file.File
}

func NewStorage() *Storage {
	return &Storage{
		Files: make(map[uuid.UUID]*file.File),
	}
}

func (s *Storage) Upload(filename string, blob []byte) (*file.File, error) {
	newFile, err := file.NewFila(filename, blob)
	s.Files[newFile.ID] = newFile
	return newFile, err
}

func (s *Storage) GetByID(fileID uuid.UUID) (*file.File, error) {
	fileThis, ok := s.Files[fileID]
	if !ok {
		return nil, errors.New(fmt.Sprintf("File %s not found", fileID))
	}
	return fileThis, nil
}
