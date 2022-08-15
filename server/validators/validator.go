package validators

import (
	"bytes"
	"encoding/binary"
	"io"
	"io/ioutil"
)

type IParam interface{}

func ConvToStruct[T IParam](body *io.ReadCloser) (T, error) {
	var params T
	bin, err := ioutil.ReadAll(*body)
	reader := bytes.NewReader(bin)
	binary.Read(reader, binary.LittleEndian, params)
	return params, err
}
