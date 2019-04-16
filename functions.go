package gotrends

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

func InitializeTrends(keyword, country, timeOption, property, category string) DataHouse {
	url, _ := url.Parse("https://trends.google.com/trends/api/explore?hl=en-US&tz=-180&req=%7B%22comparisonItem%22:%5B%7B%22keyword%22:%22" + keyword + "%22,%22geo%22:%22" + country + "%22,%22time%22:%22" + timeOption + "%22%7D%5D,%22category%22:" + category + ",%22property%22:%22" + property + "%22%7D&tz=-180")

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		log.Fatal(err)
	}

	setHeaders(req)

	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	file, _ := os.Create("Trends.json")

	bodyByte, _ := ioutil.ReadAll(res.Body)

	body := string(bodyByte)

	_, errF := file.WriteString(body[4:])
	if errF != nil {
		fmt.Println("error writing to file")
	}

	var data Trends

	byte, _ := ioutil.ReadFile("Trends.json")

	err = json.Unmarshal(byte, &data)
	if err != nil {
		log.Fatal(err)
	}

	var widg []Widget = data.Widgets

	dataHouse := DataHouse{
		widg,
		country,
		keyword,
		property,
	}

	return dataHouse
}

func InterestOverTime(data DataHouse) []byte {

	time := *data.Trends[0].Request.Time
	cat := data.Trends[0].Request.RequestOptions.Category
	category := strconv.FormatInt(cat, 10)
	time = strings.Replace(time, "\\", "\\\\", -1)
	url, _ := url.Parse("https://trends.google.com/trends/api/widgetdata/multiline?hl=en-US&tz=-180&req=%7B%22time%22:%22" + time + "%22,%22resolution%22:%22" + *data.Trends[0].Request.Resolution + "%22,%22locale%22:%22" + *data.Trends[0].Request.Locale + "%22,%22comparisonItem%22:%5B%7B%22geo%22:%7B%22country%22:%22" + data.Country + "%22%7D,%22complexKeywordsRestriction%22:%7B%22keyword%22:%5B%7B%22type%22:%22BROAD%22,%22value%22:%22" + data.Keyword + "%22%7D%5D%7D%7D%5D,%22requestOptions%22:%7B%22property%22:%22" + data.Property + "%22,%22backend%22:%22" + data.Trends[0].Request.RequestOptions.Backend + "%22,%22category%22:" + category + "%7D%7D&tz=-180")
	u := url.Query()
	u.Add("token", data.Trends[0].Token)
	url.RawQuery = u.Encode()

	req := setRequest(url)

	newclient := &http.Client{}

	res, _ := newclient.Do(req)

	bodyByte, _ := ioutil.ReadAll(res.Body)

	return bodyByte

}

func InterestBySubregion(data DataHouse, resolution string) []byte {
	time := *data.Trends[1].Request.ComparisonItem[0].Time
	cat := data.Trends[1].Request.RequestOptions.Category
	category := strconv.FormatInt(cat, 10)
	time = strings.Replace(time, "\\", "\\\\", -1)
	url, _ := url.Parse("https://trends.google.com.tr/trends/api/widgetdata/comparedgeo?hl=en-US&tz=-180&req=%7B%22geo%22:%7B%22country%22:%22" + data.Country + "%22%7D,%22comparisonItem%22:%5B%7B%22time%22:%22" + time + "%22,%22complexKeywordsRestriction%22:%7B%22keyword%22:%5B%7B%22type%22:%22BROAD%22,%22value%22:%22" + data.Keyword + "%22%7D%5D%7D%7D%5D,%22resolution%22:%22" + resolution + "%22,%22locale%22:%22" + *data.Trends[1].Request.Locale + "%22,%22requestOptions%22:%7B%22property%22:%22" + data.Property + "%22,%22backend%22:%22" + data.Trends[1].Request.RequestOptions.Backend + "%22,%22category%22:" + category + "%7D%7D")
	u := url.Query()
	u.Add("token", data.Trends[1].Token)
	url.RawQuery = u.Encode()

	req := setRequest(url)

	newclient := &http.Client{}

	res, _ := newclient.Do(req)

	bodyByte, _ := ioutil.ReadAll(res.Body)

	return bodyByte
}

func RelatedTopics(data DataHouse) []byte {
	*data.Trends[2].Request.Restriction.Time = strings.Replace(*data.Trends[2].Request.Restriction.Time, "\\", "\\\\", -1)
	data.Trends[2].Request.TrendinessSettings.CompareTime = strings.Replace(data.Trends[2].Request.TrendinessSettings.CompareTime, "\\", "\\\\", -1)
	cat := data.Trends[1].Request.RequestOptions.Category
	category := strconv.FormatInt(cat, 10)
	url, _ := url.Parse("https://trends.google.com.tr/trends/api/widgetdata/relatedsearches?hl=en-US&tz=-180&req=%7B%22restriction%22:%7B%22geo%22:%7B%22country%22:%22" + data.Country + "%22%7D,%22time%22:%22" + *data.Trends[2].Request.Restriction.Time + "%22,%22originalTimeRangeForExploreUrl%22:%22" + *data.Trends[2].Request.Restriction.OriginalTimeRangeForExploreURL + "%22,%22complexKeywordsRestriction%22:%7B%22keyword%22:%5B%7B%22type%22:%22BROAD%22,%22value%22:%22" + data.Keyword + "%22%7D%5D%7D%7D,%22keywordType%22:%22ENTITY%22,%22metric%22:%5B%22TOP%22,%22RISING%22%5D,%22trendinessSettings%22:%7B%22compareTime%22:%22" + data.Trends[2].Request.TrendinessSettings.CompareTime + "%22%7D,%22requestOptions%22:%7B%22property%22:%22" + data.Property + "%22,%22backend%22:%22" + data.Trends[2].Request.RequestOptions.Backend + "%22,%22category%22:" + category + "%7D,%22language%22:%22en%22%7D")
	u := url.Query()
	u.Add("token", data.Trends[2].Token)
	url.RawQuery = u.Encode()

	req := setRequest(url)

	newclient := &http.Client{}

	res, _ := newclient.Do(req)

	bodyByte, _ := ioutil.ReadAll(res.Body)

	return bodyByte
}

func RelatedQueries(data DataHouse) []byte {
	*data.Trends[3].Request.Restriction.Time = strings.Replace(*data.Trends[3].Request.Restriction.Time, "\\", "\\\\", -1)
	data.Trends[3].Request.TrendinessSettings.CompareTime = strings.Replace(data.Trends[3].Request.TrendinessSettings.CompareTime, "\\", "\\\\", -1)
	cat := data.Trends[1].Request.RequestOptions.Category
	category := strconv.FormatInt(cat, 10)
	url, _ := url.Parse("https://trends.google.com.tr/trends/api/widgetdata/relatedsearches?hl=en-US&tz=-180&req=%7B%22restriction%22:%7B%22geo%22:%7B%22country%22:%22" + data.Country + "%22%7D,%22time%22:%22" + *data.Trends[3].Request.Restriction.Time + "%22,%22originalTimeRangeForExploreUrl%22:%22" + *data.Trends[3].Request.Restriction.OriginalTimeRangeForExploreURL + "%22,%22complexKeywordsRestriction%22:%7B%22keyword%22:%5B%7B%22type%22:%22BROAD%22,%22value%22:%22" + data.Keyword + "%22%7D%5D%7D%7D,%22keywordType%22:%22QUERY%22,%22metric%22:%5B%22TOP%22,%22RISING%22%5D,%22trendinessSettings%22:%7B%22compareTime%22:%22" + data.Trends[3].Request.TrendinessSettings.CompareTime + "%22%7D,%22requestOptions%22:%7B%22property%22:%22" + data.Property + "%22,%22backend%22:%22" + data.Trends[3].Request.RequestOptions.Backend + "%22,%22category%22:" + category + "%7D,%22language%22:%22en%22%7D")
	u := url.Query()
	u.Add("token", data.Trends[3].Token)
	url.RawQuery = u.Encode()

	req := setRequest(url)

	newclient := &http.Client{}

	res, _ := newclient.Do(req)

	bodyByte, _ := ioutil.ReadAll(res.Body)

	return bodyByte
}

func setHeaders(req *http.Request) {
	req.Header.Set("accept-language", "en-US,en;q=0.9,tr-TR;q=0.8,tr;q=0.7")
	req.Header.Set("user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.86 Safari/537.36")
	req.Header.Set("accept", "application/json, text/plain, /")
	req.Header.Set("referer", "https://trends.google.com.tr/trends/explore?q=seo&geo=TR")
	req.Header.Set("authority", "trends.google.com.tr")
	req.Header.Set("cookie", "__utma=10102256.1714845982.1553849901.1554383869.1554454359.8; __utmc=10102256; __utmz=10102256.1554454359.8.7.utmcsr=google|utmccn=(organic)|utmcmd=organic|utmctr=(not%20provided); __utmt=1; __utmb=10102256.11.9.1554455265495; CONSENT=YES+TR.en+20161030-06-0; OGPC=19007347-1:; ANID=AHWqTUlEGxeLOO4BknJq4Afe_2xUg1NV8jnEpDkpFosBVIB6Qz9_QOwyHZUx-six; SID=QQfyBHW7nyB1yotHupAlTv0UYHeTxMwPWLpeTGnl7fNOv0refD81Lfg_dyZRplijV1LvHA.; HSID=AgoTcoa6fP2vUhrxp; SSID=Adu5aLZOA7buuMWyJ; APISID=a-8INBlBcUTInJCE/AfC1NEOmozqUrmign; SAPISID=yo1Q0nmSr7rnpFqO/AfyAZGvKCS0uJEDBG; NID=180=K4wWbY2CRVOMeDLRG-bcp0HVxdrJ6RbTqx7ciWLHBV0Fcb-iyQAWzQSAJi5p3oz3nBalENs7FJsXrHdDHzveTwevD5wFfGLFu2oeD0KfGsi6epqbZ4OOBFVSdmZ8zoXlIQt0TMoUOds2VyLmov4WC2nZJUC4fw9u2XD8V-M1pZz8nJn6HlzVvfIW7LGDbyS0AV5X-9eb6NQDNLKkUUFPwzW5T-N3cjsQEVOXNYdAURQt-CmuuybjBq_K6ZRt-k5671DrRw; 1P_JAR=2019-4-5-9; SIDCC=AN0-TYtOTxXzDiNdQnzUzfDPMq2Dl6dwBYMgCHewmTb2u1yd51X1WMURD1ZYyrwu-MkSHqD5gCI")
	req.Header.Set("x-client-data", "CIy2yQEIprbJAQipncoBCKijygEIsafKAQjiqMoBCK+sygE=")
}

func setRequest(url *url.URL) *http.Request {
	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		log.Fatal(err)
	}

	setHeaders(req)

	return req
}
