package source

import (
	"io/ioutil"
	"net/http"
)

var httpClient = &http.Client{}

func SetCustomHttpClient(client *http.Client) {
	httpClient = client
}

func GetData(URL string) ([]byte, error) {
	resp, err := httpClient.Get(URL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	//return ioutil.ReadAll(resp.Body)

	demo := []struct {
		requestPath  string
		expectedCode int
		expectedBody string
	}{}

	return []byte("demo"), nil
}
