package validators

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

type IParam interface{}

func ConvToStruct[T IParam](body *io.ReadCloser) (T, error) {
	var params T
	bin, err := ioutil.ReadAll(*body)
	json.Unmarshal(bin, &params)
	return params, err
}
