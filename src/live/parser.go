package live

import (
	"cmdline"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"regexp"
	"sports"
)

type InitData struct {
	LiveUpdatePath   string `json:"liveUpdatePath"`
	SiteStyle        string `json:"siteStyle"`
	LiveUpdateDomain string `json:"liveUpdateDomain"`
	Region           string `json:"region"`
	Updated          int    `json:"updated"`
	OddsType         string `json: "oddsType"`
	Stompf           Stomp  `json:"stomp"`
}

type Stomp struct {
	QueueName string `json:"queueName"`
}

type ReactData struct {
	LiveMenuEvents     LiveMenuEvent `json:"liveMenuEvents"`
	AnimationWidgetUrl string        `json:"animationWidgetUrl"`
	StatisticWidgetUrl string        `json:"statisticWidgetUrl"`
}

func (rdata *ReactData) LiveFootballMatchids() []int {
	matchids := make([]int, 50)
	for _, sportEvent := range rdata.LiveMenuEvents.Childs {
		if sportEvent.Uid == sports.FootballID {
			for _, levent := range sportEvent.Childs {
				for _, mevent := range levent.Childs {
					matchids = append(matchids, mevent.Uid)
				}
			}
		}
	}
	return matchids
}

type LiveMenuEvent struct {
	Uid    int           `json:"uid,string"`
	Childs []SportsEvent `json:"childs"`
}

type SportsEvent struct {
	Uid    int           `json:"uid,string"`
	Type   string        `json:"type"`
	Label  string        `json:"label"`
	Href   string        `json:"href"`
	Childs []LeagueEvent `json:"childs"`
}

type LeagueEvent struct {
	Uid    int          `json:"uid,string"`
	Type   string       `json:"type"`
	Childs []MatchEvent `json:"childs"`
}

type MatchEvent struct {
	Uid   int    `json:"uid,string"`
	Type  string `json:"type"`
	Label string `json:"label"`
	Href  string `json:"href"`
}

var initDataReg = regexp.MustCompile(`initData\s*=\s*(\{.*\})\s*;\s*reactData`)
var reactDataReg = regexp.MustCompile(`reactData\s*=\s*(\{.*\})\s*;\s*//\]\]>>`)

func Parse(content io.Reader) (*InitData, *ReactData, error) {
	bs, err := ioutil.ReadAll(content)
	if err != nil {
		return nil, nil, err
	}
	submatches := initDataReg.FindSubmatch(bs)
	if submatches == nil {
		return nil, nil, fmt.Errorf("grep initData failed\n")
	}
	if cmdline.Debug {
		fmt.Printf("submatches: %s\n", string(submatches[1]))
	}
	var initData InitData
	var reactData ReactData
	err = json.Unmarshal(submatches[1], &initData)
	if err != nil {
		return nil, nil, fmt.Errorf("Unmarshal initData failed: %v\n", err)
	}
	submatches = reactDataReg.FindSubmatch(bs)
	if cmdline.Debug {
		fmt.Printf("submatches: %s\n", string(submatches[1]))
	}
	if submatches == nil {
		return nil, nil, fmt.Errorf("grep reactData faield\n")
	}
	err = json.Unmarshal(submatches[1], &reactData)
	if err != nil {
		return nil, nil, fmt.Errorf("Unmarshal reactData failed: %v\n", err)
	}
	return &initData, &reactData, nil
}
