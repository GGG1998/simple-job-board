package data

import "encoding/json"

func RemapStructJSON[T any, U any](t *T) (*U, error) {
	jsonData, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}

	var u U
	err = json.Unmarshal(jsonData, &u)
	if err != nil {
		return nil, err
	}

	return &u, nil

}

func RemapArrayStructJSON[T any, U any](t *[]T) (*[]U, error) {
	jsonData, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}

	var u []U
	err = json.Unmarshal(jsonData, &u)
	if err != nil {
		return nil, err
	}

	return &u, nil
}
