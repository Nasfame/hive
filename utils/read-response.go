package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func ReadResonse(resp *http.Response) (res json.RawMessage, err error) {
	res, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		err = fmt.Errorf("failed to read the response -%v", res)
	}
	return
}

func ConvertJSONToGoType(res json.RawMessage, customType interface{}) {
	err := json.Unmarshal(res, customType)

	if err != nil {
		log.Errorln("failed to unmarshall due to ", err.Error())
	}
}
