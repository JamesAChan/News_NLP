package google

import (
	"net/http"
	"encoding/json"
	"strings"
	"io/ioutil"
	"public-data/logger"
	"time"
)

// declare constants
var Category = [...]string {"fiscal", "monetary", "financial"}
const (
	BaseURL = "https://newsapi.org/v2/everything?"
	Method  = "GET"
)


// Query data from Googlenews api.
// See help doc: [https://newsapi.org/s/google-news-api]

func Query(duration time.Duration, api string) []MyArticle {
	// query: from start date to now [every duration, get queried]
	start := time.Now().Add(-duration*time.Hour).Format("2006-01-02")
	// aggregate results of three categories
	var artls []MyArticle
	for _, cate := range Category {
		url := BaseURL + "q=" + cate + "&from=" + start
		var res string = httpDo(Method, url, nil, api)
		var googlejs GoogleJS
		err := json.Unmarshal([]byte(res), &googlejs)
		if err != nil {
			logger.Error().Err(err).Msg("Unmarshal Error for GoogleNews")
		}
		var myart MyArticle
		for _, artl := range googlejs.Articles {
			myart.Sources = artl.Sources
			myart.Author = artl.Author
			myart.Title = artl.Title
			myart.Description = artl.Description
			myart.URL = artl.URL
			myart.URL2image = artl.URL2image
			myart.Time = artl.Time
			myart.Content = artl.Content
			myart.Category = cate
			artls = append(artls, myart)
		}
	}
	return artls
}


/**
*  http request
 */
func httpDo(method string, url string, param map[string]string, api string) string {
	// initialize a client
	client := &http.Client{}
	// set parameters, if any
	jsonparam := ""
	if param != nil {
		bytesParams, _ := json.Marshal(param)
		jsonparam = string(bytesParams)
	}
	req, err := http.NewRequest(method, url, strings.NewReader(jsonparam))
	if err != nil {
		panic("http.NewRequest Error")
	}
	// add API key into header
	req.Header.Add("Content-Type", "application/json")
	req.Header.Set("x-api-key", api)
	// get results
	resp, err := client.Do(req)
	if err != nil {
		logger.Error().Err(err).Msg("httpDo has an error")
		return ""
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error().Err(err).Msg("ioutil.ReadAll(resp.Body) has error")
	}
	return string(body)
}
