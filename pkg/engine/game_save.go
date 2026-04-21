package engine

import (
	"encoding/json"
	"os"
)

type GameSave interface {
	Save() error
	Load() error
	Get(key string) any
	Set(key string, value any)
}

type JSONSaveFile struct {
	Values   map[string]any
	filename string
}

func NewJSONSaveFile(filename string) *JSONSaveFile {
	return &JSONSaveFile{
		Values:   make(map[string]any),
		filename: filename,
	}
}

func (j *JSONSaveFile) Save() error {
	save, err := json.Marshal(j.Values)
	if err != nil {
		return err
	}

	return os.WriteFile(j.filename, save, 0644)
}

func (j *JSONSaveFile) Load() error {
	save, err := os.ReadFile(j.filename)
	if err != nil {
		return err
	}

	return json.Unmarshal(save, &j.Values)
}

func (j *JSONSaveFile) Get(key string) any {
	return j.Values[key]
}

func (j *JSONSaveFile) Set(key string, value any) {
	j.Values[key] = value
}

var _ GameSave = &JSONSaveFile{}
