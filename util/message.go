package util

import (
	"bytes"
	"encoding/json"
	"net/http"
)

const (
	MessageTypePanic = "panic"
	MessageTypeAlert = "alert"
)

type Message struct {
	zone        string
	messageType string
	card        int
	section     int

	Cards []Card `json:"cards"`
}

type Card struct {
	Header   Header    `json:"header"`
	Sections []Section `json:"sections"`
}

type Header struct {
	Title    string `json:"title"`
	SubTitle string `json:"subtitle,omitempty"`
}

type Section struct {
	Widgets []map[string]interface{} `json:"widgets"`
}

func NewMessage() *Message {
	return &Message{
		card:    0,
		section: 0,
		Cards: []Card{{
			Sections: []Section{{
				Widgets: []map[string]interface{}{},
			}},
		}},
	}
}

func (m *Message) SetZone(zone string) *Message {
	m.zone = zone
	return m
}

func (m *Message) SetMessageType(messageType string) *Message {
	m.messageType = messageType
	return m
}

func (m *Message) SetTitle(title string) *Message {
	m.Cards[0].Header.Title = title
	return m
}

func (m *Message) SetSubtitle(subtitle string) *Message {
	m.Cards[0].Header.SubTitle = subtitle
	return m
}

func (m *Message) EndSection() *Message {
	m.section++
	m.Cards[m.card].Sections = append(m.Cards[m.card].Sections, Section{
		Widgets: []map[string]interface{}{},
	})
	return m
}

func (m *Message) AddKeyValueWidget(key, value string) *Message {
	if value == "" {
		value = "Empty"
	}

	keyValue := struct {
		TopLabel string `json:"topLabel"`
		Content  string `json:"content"`
	}{
		TopLabel: key,
		Content:  value,
	}

	widget := make(map[string]interface{})
	widget["keyValue"] = keyValue

	m.Cards[m.card].Sections[m.section].Widgets = append(m.Cards[m.card].Sections[m.section].Widgets, widget)
	return m
}

func (m *Message) AddTextParagraphWidget(text string) *Message {
	if text == "" {
		text = "Empty"
	}

	textParagraph := struct {
		Text string `json:"text"`
	}{
		Text: text,
	}

	widget := make(map[string]interface{})
	widget["textParagraph"] = textParagraph

	m.Cards[m.card].Sections[m.section].Widgets = append(m.Cards[0].Sections[m.section].Widgets, widget)
	return m
}

func (m *Message) String() string {
	b, err := json.Marshal(m)
	if err != nil {
		return ""
	}
	return string(b)
}

func (m *Message) SendMessage() {
	go func() {
		defer func() {
			recover()
		}()

		buff := bytes.NewBuffer([]byte(m.String()))

		switch m.messageType {
		case "alert":
			switch m.zone {
			case "prod":
				_, _ = http.Post("https://chat.googleapis.com/v1/spaces/AAAAglqF_Uc/messages?key=AIzaSyDdI0hCZtE6vySjMm-WEfRq3CPzqKqqsHI&token=3py5jQYhy1yzXcTePY7wPRjScFO2Rgyvl59JACyImUE%3D", "application/json", buff)
			case "dq":
				_, _ = http.Post("https://chat.googleapis.com/v1/spaces/AAAA6Iv71g8/messages?key=AIzaSyDdI0hCZtE6vySjMm-WEfRq3CPzqKqqsHI&token=mKvqmUFsj-p2BlUc1_7CeNkyAmwYs3LB-M0ie7e-Td0%3D", "application/json", buff)
			default:
				_, _ = http.Post("https://chat.googleapis.com/v1/spaces/AAAAO332jpM/messages?key=AIzaSyDdI0hCZtE6vySjMm-WEfRq3CPzqKqqsHI&token=zo9cKj9iS7I9JMi14acgOqfUilgwYnXGmb4Fmoxbadw%3D", "application/json", buff)
			}

		default:
			switch m.zone {
			case "dq", "prod":
				_, _ = http.Post("https://chat.googleapis.com/v1/spaces/AAAAI_KvlCo/messages?key=AIzaSyDdI0hCZtE6vySjMm-WEfRq3CPzqKqqsHI&token=Zft9DC4dYyyRu-c04QTr4bmMuXk9c09uUS3xtPafIOg%3D", "application/json", buff)
			default:
				_, _ = http.Post("https://chat.googleapis.com/v1/spaces/AAAA-I74ML0/messages?key=AIzaSyDdI0hCZtE6vySjMm-WEfRq3CPzqKqqsHI&token=H1hVEntNTha6WCF9DjdlDvZvzypSec1Bhefbqp-ayIk%3D", "application/json", buff)
			}
		}
	}()
}
