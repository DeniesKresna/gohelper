package utstruct

import "encoding/json"

// sourceStruct: source of the struct
//
// destinationStruct: struct where the struct will be copied to
//
// WARNING: if destination struct has value, it will be replaced by the source struct
func InjectStructValue[T any](sourceStruct interface{}, destinationStruct *T) error {
	var dataDst map[string]interface{}
	byteDst, err := json.Marshal(destinationStruct)
	if err != nil {
		return err
	}
	err = json.Unmarshal(byteDst, &dataDst)
	if err != nil {
		return err
	}

	var dataSrc map[string]interface{}
	byteSrc, err := json.Marshal(sourceStruct)
	if err != nil {
		return err
	}
	err = json.Unmarshal(byteSrc, &dataSrc)
	if err != nil {
		return err
	}

	for dstField, v := range dataDst {
		_ = v
		dSrcV, ok := dataSrc[dstField]
		if ok {
			if dSrcV != "" && dSrcV != 0 && dSrcV != nil {
				dataDst[dstField] = dataSrc[dstField]
			}
		}
	}

	byteNewDst, err := json.Marshal(dataDst)
	if err != nil {
		return err
	}

	err = json.Unmarshal(byteNewDst, destinationStruct)
	if err != nil {
		return err
	}
	return nil
}
