package main

import (
	"github.com/yenkeia/mirgo/common"
	"github.com/yenkeia/mirgo/proto/server"
)

type ServerMessage struct{}

func (ServerMessage) SetConcentration() *server.SetConcentration {
	sc := new(server.SetConcentration)
	sc.ObjectID = 66432
	sc.Enabled = false
	sc.Interrupted = false
	return sc
}

func (ServerMessage) ObjectPlayer(o IMapObject) (res *server.ObjectPlayer) {
	return o.GetInfo().(*server.ObjectPlayer)
}

func (ServerMessage) ObjectMonster(o IMapObject) *server.ObjectMonster {
	return o.GetInfo().(*server.ObjectMonster)
}

func (ServerMessage) MapInformation(info *common.MapInfo) *server.MapInformation {
	mi := new(server.MapInformation)
	mi.FileName = info.Filename
	mi.Title = info.Title
	mi.MiniMap = uint16(info.MiniMap)
	mi.BigMap = uint16(info.BigMap)
	mi.Music = uint16(info.Music)
	mi.Lights = common.LightSetting(info.Light)
	mi.Lightning = true
	mi.MapDarkLight = 0
	return mi
}

func (ServerMessage) StartGame(result, resolution int) *server.StartGame {
	/*
	 * 0: Disabled.
	 * 1: Not logged in
	 * 2: Character not found.
	 * 3: Start Game Error
	 * 4: Success
	 * */
	sg := new(server.StartGame)
	sg.Result = uint8(result)
	sg.Resolution = int32(resolution)
	return sg
}

func (ServerMessage) UserInformation(p *Player) *server.UserInformation {
	ui := new(server.UserInformation)
	ui.ObjectID = p.GetID()
	ui.RealID = p.GetID()
	ui.Name = p.Name
	ui.GuildName = p.GuildName
	ui.GuildRank = p.GuildRankName
	ui.NameColor = common.Color{R: 255, G: 255, B: 255}.ToInt32()
	ui.Class = p.Class
	ui.Gender = p.Gender
	ui.Level = p.Level
	ui.Location = p.CurrentLocation
	ui.Direction = p.CurrentDirection
	ui.Hair = p.Hair
	ui.HP = p.HP
	ui.MP = p.MP
	ui.Experience = p.Experience             // TODO
	ui.MaxExperience = 100                   // TODO
	ui.LevelEffect = common.LevelEffectsNone // TODO
	ui.Gold = uint32(p.Gold)
	ui.Credit = 100 // TODO
	ui.Inventory = p.Inventory
	ui.Equipment = p.Equipment
	ui.QuestInventory = p.QuestInventory
	return ui
}

func (ServerMessage) UserLocation(p *Player) *server.UserLocation {
	return &server.UserLocation{
		Location:  p.Point(),
		Direction: p.CurrentDirection,
	}
}

func (ServerMessage) ObjectTurn(o IMapObject) *server.ObjectTurn {
	return &server.ObjectTurn{
		ObjectID:  o.GetID(),
		Location:  o.GetPoint(),
		Direction: o.GetDirection(),
	}
}

func (ServerMessage) ObjectWalk(o IMapObject) *server.ObjectWalk {
	return &server.ObjectWalk{
		ObjectID:  o.GetID(),
		Location:  o.GetPoint(),
		Direction: o.GetDirection(),
	}
}

func (ServerMessage) ObjectRun(o IMapObject) *server.ObjectRun {
	return &server.ObjectRun{
		ObjectID:  o.GetID(),
		Location:  o.GetPoint(),
		Direction: o.GetDirection(),
	}
}

func (ServerMessage) ObjectRemove(o IMapObject) *server.ObjectRemove {
	return &server.ObjectRemove{ObjectID: o.GetID()}
}

func (ServerMessage) ObjectChat(p *Player, message string, chatType common.ChatType) *server.ObjectChat {
	text := p.Name + ":" + message
	return &server.ObjectChat{
		ObjectID: p.ID,
		Text:     text,
		Type:     chatType,
	}
}

func (ServerMessage) ObjectNPC(o IMapObject) *server.ObjectNPC {
	return o.GetInfo().(*server.ObjectNPC)

}

func (ServerMessage) NewItemInfo(item *common.ItemInfo) *server.NewItemInfo {
	if item == nil {
		panic("new item info, item = nil !!!")
	}
	return &server.NewItemInfo{Info: *item}
}

func (ServerMessage) PlayerInspect(p *Player) *server.PlayerInspect {
	return &server.PlayerInspect{
		Name:      p.Name,
		GuildName: p.GuildName,
		GuildRank: p.GuildRankName,
		Equipment: p.Equipment,
		Class:     p.Class,
		Gender:    p.Gender,
		Hair:      p.Hair,
		Level:     p.Level,
		LoverName: "",
	}
}

func (ServerMessage) Login(result int) *server.Login {
	/*
	 * 0: Disabled
	 * 1: Bad AccountID
	 * 2: Bad Password
	 * 3: Account Not Exist
	 * 4: Wrong Password
	 */
	return &server.Login{Result: uint8(result)}
}

func (ServerMessage) NewCharacter(result int) interface{} {
	/*
	 * 0: Disabled.
	 * 1: Bad Character Name
	 * 2: Bad Gender
	 * 3: Bad Class
	 * 4: Max Characters
	 * 5: Character Exists.
	 * */
	return &server.NewCharacter{Result: uint8(result)}
}

func (ServerMessage) NewCharacterSuccess(g *Game, AccountID int, name string, class common.MirClass, gender common.MirGender) *server.NewCharacterSuccess {
	c := new(common.Character)
	c.Name = name
	c.Level = 8
	c.Class = class
	c.Gender = gender
	c.Hair = 1
	c.CurrentMapID = 1
	c.CurrentLocationX = 284
	c.CurrentLocationY = 608
	c.Direction = common.MirDirectionDown
	c.HP = 15
	c.MP = 17
	c.Experience = 0
	c.AttackMode = common.AttackModeAll
	c.PetMode = common.PetModeBoth
	g.DB.Table("character").Create(c)
	g.DB.Table("character").Where("name = ?", name).Last(c)
	ac := new(common.AccountCharacter)
	ac.AccountID = AccountID
	ac.CharacterID = int(c.ID)
	g.DB.Table("account_character").Create(ac)
	res := new(server.NewCharacterSuccess)
	res.CharInfo.Index = uint32(c.ID)
	res.CharInfo.Name = name
	res.CharInfo.Class = class
	res.CharInfo.Gender = gender
	return res
}

func (ServerMessage) LogOutSuccess(characters []common.SelectInfo) *server.LogOutSuccess {
	return &server.LogOutSuccess{Characters: characters}
}

func (ServerMessage) TimeOfDay(light common.LightSetting) *server.TimeOfDay {
	return &server.TimeOfDay{Lights: light}
}

func (ServerMessage) NPCResponse(page []string) *server.NPCResponse {
	return &server.NPCResponse{Page: page}
}
