package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"cleur-barcelona-2018/web/ciscolive.service.content/utils"

	"github.com/valyala/fasthttp"
)

func SearchCatalog(ctx *fasthttp.RequestCtx) utils.Result {
	form := url.Values{}
	form.Add("showMyInterest", "false")
	form.Add("showEnrolled", "false")
	form.Add("search", string(ctx.QueryArgs().Peek("search")))
	form.Add("type", "session")

	req, err := http.NewRequest("POST", "https://events.rainfocus.com/api/search", strings.NewReader(form.Encode()))
	if err != nil {
		return utils.Error(fmt.Errorf("http error: " + err.Error()))
	}

	headers := http.Header{}
	headers.Add("rfApiProfileId", "GyJydfgvjNclemea18")
	headers.Add("rfWidgetId", "9mho0zaabykve2m8iLtBtpck53hL0WBt")
	headers.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header = headers
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return utils.Error(fmt.Errorf("http error: " + err.Error()))
	}

	// we use defer to close the http handles
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return utils.Error(fmt.Errorf("http error: " + err.Error()))
	}

	searchResponse := &SearchResponse{}
	json.Unmarshal(body, searchResponse)

	type Speaker struct {
		CompanyName string `json:"companyName"`
		FullName    string `json:"fullName"`
		JobTitle    string `json:"jobTitle"`
	}
	type CustomResponses struct {
		SessionName string    `json:"sessionName"`
		SessionCode string    `json:"sessionCode"`
		Speakers    []Speaker `json:"speakers"`
	}

	var customResponses []CustomResponses

	for _, section := range searchResponse.SectionList {
		for _, item := range section.Items {
			customResponse := CustomResponses{
				SessionName: item.Title,
				SessionCode: item.Code,
			}
			for _, participant := range item.Participants {
				speaker := Speaker{
					CompanyName: participant.CompanyName,
					FullName:    participant.FullName,
					JobTitle:    participant.JobTitle,
				}
				customResponse.Speakers = append(customResponse.Speakers, speaker)
			}
			customResponses = append(customResponses, customResponse)
		}
	}
	return utils.Ok(customResponses)
	// return utils.Ok(searchResponse)
}
