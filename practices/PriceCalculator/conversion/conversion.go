package conversion

import (
  "errors"
  "strconv"
)

func StringsToFloats(strings []string) ([]float64, error) {
  var floats []float64

  for _, str := range strings {
    floatValue, err := strconv.ParseFloat(str, 64)

    if err != nil {
      return nil, errors.New("Failed to convert string to float.")
    }

    floats = append(floats, floatValue)
  }

  return floats, nil
}
