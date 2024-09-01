// Copyright 2022 Prometheus Team
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package rocketchat

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/go-kit/log"
	"github.com/prometheus/alertmanager/config"
	"github.com/prometheus/alertmanager/notify"
	commoncfg "github.com/prometheus/common/config"

	"github.com/prometheus/alertmanager/template"
	"github.com/prometheus/alertmanager/types"
)

// Notifier implements a Notifier for RocketChat notifications.
type Notifier struct {
	conf   *config.RocketchatConfig
	tmpl   *template.Template
	logger log.Logger
}

// New returns a new RocketChat notification handler.
func New(conf *config.RocketchatConfig, t *template.Template, l log.Logger, httpOpts ...commoncfg.HTTPClientOption) (*Notifier, error) {
	return &Notifier{
		conf:   conf,
		tmpl:   t,
		logger: l,
	}, nil
}

type requestTokenStruct struct {
	LoginId  string `json:"user"`
	Password string `json:"password"`
}

type responseToken struct {
	Data struct {
		UserId string `json:"userId"`
		Token  string `json:"authToken"`
	} `json:"data"`
}

func obtainToken(server_url string, user_name string, password string) (string, string, error) {
	request := requestTokenStruct{
		LoginId:  user_name,
		Password: password,
	}
	contentType := "application/json"
	jsonBody, err := json.Marshal(request)
	if err != nil {
		return "", "", err
	}
	url := fmt.Sprintf("%s/api/v1/login", server_url)
	r, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", "", err
	}
	client := &http.Client{}
	r.Header.Add("Content-Type", contentType)
	res, err := client.Do(r)
	if err != nil {
		return "", "", err
	}
	if res.StatusCode != http.StatusOK {
		return "", "", fmt.Errorf("status code %d", res.StatusCode)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", "", err
	}
	res.Body.Close()
	data := responseToken{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return "", "", err
	}
	return data.Data.UserId, data.Data.Token, nil
}

type postMessageStruct struct {
	Channel string `json:"channel"`
	Message string `json:"msg"`
}

// sendMessage posts a message to the specified RocketChat channel.
func sendMessage(server_url string, userId string, token string, channel string, text string) error {
	request := postMessageStruct{
		Channel: channel,
		Message: text,
	}
	contentType := "application/json"
	jsonBody, err := json.Marshal(request)
	if err != nil {
		return err
	}
	url := fmt.Sprintf("%s/api/v1/chat.postMessage", server_url)
	r, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))

	if err != nil {
		return err
	}
	r.Header.Add("X-User-Id", userId)
	r.Header.Add("X-Auth-Token", token)
	r.Header.Add("Content-Type", contentType)
	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusCreated {
		return fmt.Errorf("status code %d", res.StatusCode)
	}
	return nil
}

// FormatMessage formats the alert into a string message for RocketChat.
func FormatMessage(alert *types.Alert) string {
	status := strings.ToUpper(string(alert.Status()))
	startTime := alert.StartsAt.Local().String()
	endTime := alert.EndsAt.Local().String()
	severity := strings.ToUpper(string(alert.Labels["severity"]))
	summary := string(alert.Annotations["summary"])
	description := string(alert.Annotations["description"])

	emoji := ":white_check_mark:"
	if status == "FIRING" {
		emoji = ":bangbang:"
	}

	tim := ""
	if status == "FIRING" {
		tim = fmt.Sprintf("The alert fired at **%s**", startTime)
	} else {
		tim = fmt.Sprintf("The alert which fired at **%s** was resolved at **%s**", startTime, endTime)
	}

	resultMessage := fmt.Sprintf("%s **%s** - %s %s\n%s\n", emoji, status, summary, emoji, tim) +
		fmt.Sprintf("**Severity: ** %s\n", severity) + fmt.Sprintf("**Description: ** %s\n", description)

	return resultMessage
}

// Notify sends notifications for the given alerts to Mattermost.
func (n *Notifier) Notify(ctx context.Context, alert ...*types.Alert) (bool, error) {
	userId, token, err := obtainToken(n.conf.ServerUrl, n.conf.UserName, n.conf.Password)
	if err != nil {
		fmt.Printf("Error obtaining token %v\n", err)
	} else {
		fmt.Printf("Token: %s\n", token)

		var (
			data     = notify.GetTemplateData(ctx, n.tmpl, alert, n.logger)
			tmplText = notify.TmplText(n.tmpl, data, &err)
		)
		err = sendMessage(n.conf.ServerUrl, userId, token, n.conf.ChannelId, tmplText(FormatMessage(alert[0])))
		if err != nil {
			fmt.Printf("Error while sending message: %v\n", err)
		}
	}
	return false, nil
}
