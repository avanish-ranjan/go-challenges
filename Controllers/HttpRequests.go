package Controllers

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

/*http request POST
1. client := &http.Client{}
2. http.NewRequest
3. client.Do()
4. iotil.ReadAll
*/
func MyHttpReq() error {
	url := "http/:online.com"
	method := "POST"
	reqBody := []byte(`"name":"Avanish","mob":"7999042575"`)
	client := &http.Client{}
	req, errInReq := http.NewRequest(method, url, bytes.NewBuffer(reqBody))
	if errInReq != nil {
		return errInReq
	}
	req.Header.Set("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	responseBody, err1 := ioutil.ReadAll(res.Body)
	if err1 != nil {
		return err1
	}
	fmt.Println(string(responseBody))
	return nil
}
