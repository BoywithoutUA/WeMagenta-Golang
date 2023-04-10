package api

import (
	"fmt"
	"time"
)

type JsonTime time.Time

func (j JsonTime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(j).Format("2006-01-02 03:04:05"))
	return []byte(stamp), nil
}

type CompositionResponse struct {
	Id              uint32   `json:"-"`
	CompositionName string   `json:"compositionName"`
	ForName         string   `json:"forName"`
	ForType         string   `json:"forType"`
	FilePath        string   `json:"filePath"`
	SubmitTime      JsonTime `json:"submitTime"`
}
