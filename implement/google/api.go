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
var Member  = [...]string {
	"Afghanistan", "Albania", "Algeria", "Angola", "Anguilla", "Antigua and Barbuda", "Argentina", "Armenia", "Aruba",
	"Australia", "Austria", "Azerbaijan", "Bahamas", "Bahrain", "Bangladesh", "Barbados", "Belarus", "Belgium", "Belize",
	"Benin", "Bhutan", "Bolivia", "Bonaire", "Bosnia and Herzegovina", "Botswana", "Brazil", "Brunei Darussalam",
	"Bulgaria", "Burkina Faso", "Burundi", "Cabo Verde", "Cambodia", "Cameroon", "Canada", "Central African",
	"Republic", "Chad", "Chile", "China", "Colombia", "Comoros", "Congo", "Costa Rica", "Côted'Ivoire", "Croatia", "Curaçao",
	"Cyprus", "Czech Republic", "Denmark", "Djibouti", "Dominica", "Dominican Republic", "Ecuador", "Egypt", "El Salvador", "Equatorial Guinea",
	"Eritrea", "Estonia", "Ethiopia", "Fiji", "Finland", "Former Yugoslav Republic of Macedonia", "France", "Gabon", "Gambia",
	"Georgia", "Germany", "Ghana", "Greece", "Grenada", "Guatemala", "Guinea", "Guinea-Bissau", "Guyana", "Haiti",
	"Honduras", "Hong Kong", "Hungary", "Iceland", "India", "Indonesia", "Iran", "Iraq", "Ireland", "Israel", "Italy","Jamaica", "Japan", "Jordan",
	"Kazakhstan", "Kenya", "Kiribati", "Korea", "Kosovo", "Kuwait", "Kyrgyz Republic", "Lao People's Democratic Republic", "Latvia", "Lebanon", "Lesotho",
	"Liberia", "Libya", "Lithuania", "Luxembourg", "Macau", "Madagascar", "Malawi", "Malaysia", "Maldives", "Mali", "Malta", "Marshall Islands", "Mauritania",
	"Mauritius", "Mexico", "Micronesia", "Moldova", "Mongolia", "Montenegro", "Montserrat", "Morocco", "Mozambique", "Myanmar", "Namibia", "Nauru", "Nepal",
	"Netherlands", "New Zealand", "Nicaragua", "Niger","Nigeria", "Norway", "Oman", "Pakistan", "Palau", "Panama", "Papua New Guinea", "Paraguay", "Peru",
	"Philippines","Poland", "Portugal", "Qatar", "Romania", "Russia", "Rwanda", "Saba", "Saint Eustatius", "Saint Kitts and Nevis","Saint Lucia", "Saint Vincent and the Grenadines",
	"Samoa", "San Marino", "São Tomé and Príncipe", "Saudi Arabia", "Senegal", "Serbia", "Seychelles", "Sierra Leone", "Singapore", "Sint Maarten", "Slovak Republic", "Slovenia",
	"Solomon Islands", "Somalia", "South Africa", "South Sudan", "Spain", "Sri Lanka", "Sudan", "Suriname", "Swaziland", "Sweden", "Switzerland", "Syrian Arab Republic", "Tajikistan",
	"Tanzania", "Thailand", "TimorLeste", "Togo", "Tonga", "Trinidad and Tobago", "Tunisia", "Turkey", "Turkmenistan", "Tuvalu", "Uganda", "Ukraine", "United Arab Emirates", "United Kingdom",
	"United States", "Uruguay", "Uzbekistan", "Vanuatu", "Venezuela", "Viet Nam", "Yemen", "Zambia", "Zimbabwe"}
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
		for _, cty := range Member {
			url := BaseURL + "q=" + cate + ", " + cty + "&from=" + start
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
				myart.Country = cty
				artls = append(artls, myart)
			}
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
