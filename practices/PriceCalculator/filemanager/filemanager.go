package filemanager

import (
  "os"
  "bufio"
  "errors"
  "encoding/json"
)

type FileManager struct {
  InputFilePath string
  OutputFilePath string
}

func (fm FileManager) ReadLines() ([]string, error) {
  file, err := os.Open(fm.InputFilePath)

  if err != nil {
    return nil, errors.New("Failed to open the file.")
  }

  defer file.Close()

  var lines []string

  scanner := bufio.NewScanner(file)

  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }

  err = scanner.Err()

  if err != nil {
    return nil, errors.New("Failed to read file.")
  }

  return lines, nil
}

func (fm FileManager) WriteResult(data interface{}) error {
  file, err := os.Create(fm.OutputFilePath)

  if err != nil {
    return errors.New("Failed to create file.")
  }

  defer file.Close()

  encoder := json.NewEncoder(file)
  err = encoder.Encode(data)

  if err != nil {
    return errors.New("Failed to convert data to JSON.")
  }

  return nil
}

func New(inputPath, outputPath string) FileManager {
  return FileManager{
    InputFilePath: inputPath,
    OutputFilePath: outputPath,
  }
}
