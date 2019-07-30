package helpers

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//IsWord validates if the combination is a word or not
func IsWord(wr string) (string, bool) {
	url := fmt.Sprintf("%s%s%s", "https://api.datamuse.com/words?sp=", wr, "&md=d")

	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	if len(body) > 50 {
		return string(body), true
	} else {
		return "", false
	}
}

