package discord

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

var (
	defaultAppID          = "674504097076084747"
	defaultLoggingInText  = "Logging in"
	defaultLargeImageID   = "arknights"
	defaultLargeImageText = "Arknights"
	defaultSmallImageID   = "rhine_labs"
	defaultSmallImageText = "Rhine Labs"
	defaultIdleText       = "Idling"
	defaultPracticeText   = "Practicing "
	defaultAutoplayText   = "Autoing "
	defaultBattleText     = "Fighting "
)

func getConfig(fileName string) *config {
	var r config
	defer fillDefaults(&r)
	f, err := os.Open(fileName)
	if err != nil {
		return &r
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return &r
	}
	err = json.Unmarshal(b, &r)
	if err != nil {
		return &r
	}
	return &r
}

func fillDefaults(c *config) {
	// Don't @ me
	if c.AppID == nil {
		c.AppID = &defaultAppID
	}
	if c.LoggingInText == nil {
		c.LoggingInText = &defaultLoggingInText
	}
	if c.LargeImageID == nil {
		c.LargeImageID = &defaultLargeImageID
	}
	if c.LargeImageText == nil {
		c.LargeImageText = &defaultLargeImageText
	}
	if c.SmallImageID == nil {
		c.SmallImageID = &defaultSmallImageID
	}
	if c.SmallImageText == nil {
		c.SmallImageText = &defaultSmallImageText
	}
	if c.IdleText == nil {
		c.IdleText = &defaultIdleText
	}
	if c.PracticeText == nil {
		c.PracticeText = &defaultPracticeText
	}
	if c.AutoplayText == nil {
		c.AutoplayText = &defaultAutoplayText
	}
	if c.BattleText == nil {
		c.BattleText = &defaultBattleText
	}
}

type config struct {
	AppID          *string `json:"appID"`
	LoggingInText  *string `json:"loggingInText"`
	LargeImageID   *string `json:"largeImageID"`
	LargeImageText *string `json:"largeImageText"`
	SmallImageID   *string `json:"smallImageID"`
	SmallImageText *string `json:"smallImageText"`
	IdleText       *string `json:"idleText"`
	PracticeText   *string `json:"practiceText"`
	AutoplayText   *string `json:"autoplayText"`
	BattleText     *string `json:"battleText"`
}
