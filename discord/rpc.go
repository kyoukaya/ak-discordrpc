// Discord RPC changes uses the discord rich presence feature to display
// a logged in user's region, name, and status.
// The status is limited to combat (practice, auto, fighting) and non-combat (idling only)

package discord

import (
	"fmt"
	"sync"
	"time"

	"github.com/elazarl/goproxy"
	"github.com/kyoukaya/rhine/proxy"
	"github.com/kyoukaya/rhine/utils/gamedata"
	"github.com/kyoukaya/rhine/utils/gamedata/stagetable"
	discord "github.com/kyoukaya/rich-go/client"
	"github.com/tidwall/gjson"
)

const (
	modName          = "Discord RPC"
	appID            = "674504097076084747"
	loggingInText    = "Logging in"
	largeImageID     = "arknights"
	largeImageText   = "Arknights"
	largeImageTextJP = "アークナイツ"
	largeImageTextKR = "명일방주"
	idleText         = "Idling"
	practiceText     = "Practicing "
	autoplayText     = "Autoing "
	battleText       = "Fighting "
)

type modState struct {
	mutex      sync.Mutex
	activity   *discord.Activity
	stageTable *stagetable.StageTable
	*gamedata.GameData
	*proxy.RhineModule
}

func (mod *modState) updateActivity() {
	err := discord.SetActivity(*mod.activity)
	if err != nil {
		mod.Warnln(err)
	}
}

func (mod *modState) battleFinishHandler(_ string, data []byte, _ *goproxy.ProxyCtx) []byte {
	mod.mutex.Lock()
	defer mod.mutex.Unlock()
	mod.activity.State = idleText
	mod.activity.Timestamps = nil
	go mod.updateActivity()
	return data
}

func (mod *modState) battleStart(data []byte) {
	mod.mutex.Lock()
	defer mod.mutex.Unlock()
	isPractice := gjson.GetBytes(data, "usePracticeTicket").Bool()
	if isPractice {
		mod.activity.State = practiceText
	} else {
		if gjson.GetBytes(data, "isReplay").Bool() {
			mod.activity.State = autoplayText
		} else {
			mod.activity.State = battleText
		}
	}
	stageID := gjson.GetBytes(data, "stageId").String()
	stage := mod.stageTable.Stages[stageID]
	if stage.Name == nil {
		mod.activity.State += stage.Code
	} else {
		mod.activity.State += fmt.Sprintf("%s - %s", stage.Code, *stage.Name)
	}
	now := time.Now()
	mod.activity.Timestamps = &discord.Timestamps{
		Start: &now,
	}
	mod.updateActivity()
}

func (mod *modState) battleStartHandler(op string, data []byte, ctx *goproxy.ProxyCtx) []byte {
	go mod.battleStart(proxy.GetRequestContext(ctx).RequestData)
	return data
}

func (mod *modState) syncData() {
	mod.mutex.Lock()
	defer mod.mutex.Unlock()
	mod.stageTable = mod.GetStageInfo()
	s := mod.GetGameState().Status
	mod.activity.Details = fmt.Sprintf("[%s] %s#%s", mod.Region, s.NickName, s.NickNumber)
	if mod.Region == "JP" {
		mod.activity.LargeText = largeImageTextJP
	} else if mod.Region == "KR" {
		mod.activity.LargeText = largeImageTextKR
	}
	mod.updateActivity()
}

func (mod *modState) syncDataHandler(op string, data []byte, ctx *goproxy.ProxyCtx) []byte {
	go mod.syncData()
	return data
}

func initFunc(mod *proxy.RhineModule) {
	gd, err := gamedata.New(mod.Region, mod.Logger)
	if err != nil {
		mod.Warnln(err)
	}
	state := modState{
		activity: &discord.Activity{
			State:      idleText,
			Details:    loggingInText,
			LargeImage: largeImageID,
			LargeText:  largeImageText,
		},
		RhineModule: mod,
		GameData:    gd,
	}
	// Handshake may take awhile, so we'll do that in a goroutine
	go func() {
		state.mutex.Lock()
		defer state.mutex.Unlock()
		err := discord.Login(appID)
		if err != nil {
			mod.Warnln(err)
		}
		state.updateActivity()
	}()
	mod.Hook("S/account/syncData", 0, state.syncDataHandler)
	mod.Hook("S/quest/battleFinish", 0, state.battleFinishHandler)
	mod.Hook("S/campaign/battleFinish", 0, state.battleFinishHandler)
	mod.Hook("S/quest/battleStart", 0, state.battleStartHandler)
	mod.Hook("S/campaign/battleStart", 0, state.battleStartHandler)
}

func init() {
	proxy.RegisterInitFunc(modName, initFunc)
}
