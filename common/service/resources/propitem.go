package resources

import (
	"bufio"
	"go-ff/common/service/resources/reader"
	"strings"
)

var ItemsProp = loadPropItem()

type ElementPropType uint8

const (
	ElementNone ElementPropType = iota
	ElementFire
	ElementWater
	ElementElectricity
	ElementWind
	ElementEarth
)

type ObjProp struct {
	ID   int32
	Name string
	Type uint64
	AI   uint64
	HP   uint64
}

type CtrlProp struct {
	ObjProp

	// Custom fields
	CtrlKind1 uint64
	CtrlKind2 uint64
	CtrlKind3 uint64
	SFXCtrl   uint64
	SndDamage uint64
}

type ItemProp struct {
	CtrlProp

	// Custom fields
	Version              uint16
	Num                  uint64
	PackMax              uint64
	ItemKind1            uint32 // defineKind.h
	ItemKind2            uint32 // defineKind.h
	ItemKind3            uint32 // defineKind.h
	Job                  uint32 // defineJob.h
	Permanence           bool
	Useable              uint64
	Sex                  uint8
	Cost                 uint64
	Endurance            uint64
	Log                  int32
	Abrasion             int32
	MaxRepair            int32
	Handed               uint32 // define.h
	Flag                 uint64
	Parts                uint32 // define.h
	Partsub              uint32 // define.h (Earring etc..)
	PartsFile            bool
	Exclusive            uint32 // define.h
	BasePartsIgnore      uint32 // define.h
	LV                   uint64
	Rare                 uint64
	ShopAble             bool
	ShellQuantity        uint64
	ActiveSkillLv        uint64
	SpellType            uint64
	LinkKindBullet       uint32 // define.h
	LinkKind             uint32 // define.h
	AbilityMin           uint64
	AbilityMax           uint64
	Charged              bool
	ElementType          ElementPropType // (NONE, ELECTRICITY.........)
	ItemEatk             int16
	Parry                uint64
	BlockRating          uint64
	AddSkillMin          int32
	AddSkillMax          int32
	AtkStyle             uint64
	WeaponType           uint32 // define.h
	ItemAtkOrder1        uint32 // define.h
	ItemAtkOrder2        uint32 // define.h
	ItemAtkOrder3        uint32 // define.h
	ItemAtkOrder4        uint32 // define.h
	ContinuousPainTime   uint64
	Recoil               uint64
	LoadingTime          uint64
	AdjHitRate           int64
	AttackSpeed          float64
	DmgShift             uint64
	AttackRange          uint32 // define.h
	Probability          int32
	DestParam            [3]uint32 // define.h
	AdjParamVal          [3]int64
	ChgParamVal          [3]uint64
	DestData1            [3]int64
	ActiveSkill          uint32 // define.h
	ActiveSkillRate      uint64
	ReqMp                uint64
	ReqFp                uint64
	ReqDisLV             uint64
	ReSkill1             uint64
	ReSkillLevel1        uint64
	ReSkill2             uint64
	ReSkillLevel2        uint64
	SkillReadyType       uint64
	SkillReady           uint64
	SkillRange           uint64
	SfxElemental         uint64
	SfxObj               uint32 // define.h
	SfxObj2              uint32 // define.h
	SfxObj3              uint32 // define.h
	SfxObj4              uint32 // define.h
	SfxObj5              uint32 // define.h
	UseMotion            uint32 // define.h
	CircleTime           uint64
	SkillTime            uint64
	ExeTarget            uint64
	UseChance            uint32 // define.h
	SpellRegion          uint64
	ReferStat1           uint32 // define.h
	ReferStat2           uint32 // define.h
	ReferTarget1         uint32 // define.h
	ReferTarget2         uint32 // define.h
	ReferValue1          uint64
	ReferValue2          uint64
	SkillType            uint64
	ItemResistElecricity int32
	ItemResistDark       int32
	ItemResistFire       int32
	ItemResistWind       int32
	ItemResistWater      int32
	ItemResistEarth      int32
	Evildoing            int64
	ExpertLV             uint32
	ExpertMax            uint64
	SubDefine            uint32 // define.h
	Exp                  uint64
	ComboStyle           uint64
	FlightSpeed          float64
	FlightLRAngle        float64
	FlightTBAngle        float64
	FlightLimit          uint64
	FFuelReMax           uint64
	AFuelReMax           uint64
	FuelRe               uint64
	LimitLevel1          uint64
	Reflect              int32
	SndAttack1           uint32 // define.h
	SndAttack2           uint32 // define.h
	QuestID              uint64
	TextFileName         string
}

func fillField(prop *ItemProp, index int, field string, r *reader.Reader) {
	switch index {
	case 0:
		prop.Version = r.GetUInt16(field)
	case 1:
		prop.ID = r.GetInt32(field)
	case 2:
		prop.Name = field // r.GetUInt32(field)
	case 3:
		prop.Num = r.GetUInt64(field)
	case 4:
		prop.PackMax = r.GetUInt64(field)
	case 5:
		prop.ItemKind1 = r.GetUInt32(field)
	case 6:
		prop.ItemKind2 = r.GetUInt32(field)
	case 7:
		prop.ItemKind3 = r.GetUInt32(field)
	case 8:
		prop.Job = r.GetUInt32(field)
	case 9:
		prop.Permanence = r.GetBool(field)
	case 10:
		prop.Useable = r.GetUInt64(field)
	case 11:
		prop.Sex = r.GetUInt8(field)
	case 12:
		prop.Cost = r.GetUInt64(field)
	case 13:
		prop.Endurance = r.GetUInt64(field)
	case 14:
		prop.Abrasion = r.GetInt32(field)
	case 15:
		prop.MaxRepair = r.GetInt32(field)
	case 16:
		prop.Handed = r.GetUInt32(field)
	case 17:
		prop.Flag = r.GetUInt64(field)
	case 18:
		prop.Parts = r.GetUInt32(field)
	case 19:
		prop.Partsub = r.GetUInt32(field)
	case 20:
		prop.PartsFile = r.GetBool(field)
	case 21:
		prop.Exclusive = r.GetUInt32(field)
	case 22:
		prop.BasePartsIgnore = r.GetUInt32(field)
	case 23:
		prop.LV = r.GetUInt64(field)
	case 24:
		prop.Rare = r.GetUInt64(field)
	case 25:
		prop.ShopAble = r.GetBool(field)
	case 26:
		prop.Log = r.GetInt32(field)
	case 27:
		prop.Charged = r.GetBool(field)
	case 28:
		prop.LinkKindBullet = r.GetUInt32(field)
	case 29:
		prop.LinkKind = r.GetUInt32(field)
	case 30:
		prop.AbilityMin = r.GetUInt64(field)
	case 31:
		prop.AbilityMax = r.GetUInt64(field)
	case 32:
		prop.ElementType = (ElementPropType)(r.GetUInt8(field))
	case 33:
		prop.ItemEatk = r.GetInt16(field)
	case 34:
		prop.Parry = r.GetUInt64(field)
	case 35:
		prop.BlockRating = r.GetUInt64(field)
	case 36:
		prop.AddSkillMin = r.GetInt32(field)
	case 37:
		prop.AddSkillMax = r.GetInt32(field)
	case 38:
		prop.AtkStyle = r.GetUInt64(field)
	case 39:
		prop.WeaponType = r.GetUInt32(field)
	case 40:
		prop.ItemAtkOrder1 = r.GetUInt32(field)
	case 41:
		prop.ItemAtkOrder2 = r.GetUInt32(field)
	case 42:
		prop.ItemAtkOrder3 = r.GetUInt32(field)
	case 43:
		prop.ItemAtkOrder4 = r.GetUInt32(field)
	case 44:
		prop.ContinuousPainTime = r.GetUInt64(field)
	case 45:
		prop.ShellQuantity = r.GetUInt64(field)
	case 46:
		prop.Recoil = r.GetUInt64(field)
	case 47:
		prop.LoadingTime = r.GetUInt64(field)
	case 48:
		prop.AdjHitRate = r.GetInt64(field)
	case 49:
		prop.AttackSpeed = r.GetFloat64(field)
	case 50:
		prop.DmgShift = r.GetUInt64(field)
	case 51:
		prop.AttackRange = r.GetUInt32(field)
	case 52:
		prop.Probability = r.GetInt32(field)
	case 53:
		prop.DestParam[0] = r.GetUInt32(field)
	case 54:
		prop.DestParam[1] = r.GetUInt32(field)
	case 55:
		prop.DestParam[2] = r.GetUInt32(field)
	case 56:
		prop.AdjParamVal[0] = r.GetInt64(field)
	case 57:
		prop.AdjParamVal[1] = r.GetInt64(field)
	case 58:
		prop.AdjParamVal[2] = r.GetInt64(field)
	case 59:
		prop.ChgParamVal[0] = r.GetUInt64(field)
	case 60:
		prop.ChgParamVal[1] = r.GetUInt64(field)
	case 61:
		prop.ChgParamVal[2] = r.GetUInt64(field)
	case 62:
		prop.DestData1[0] = r.GetInt64(field)
	case 63:
		prop.DestData1[1] = r.GetInt64(field)
	case 64:
		prop.DestData1[2] = r.GetInt64(field)
	case 65:
		prop.ActiveSkill = r.GetUInt32(field)
	case 66:
		prop.ActiveSkillLv = r.GetUInt64(field)
	case 67:
		prop.ActiveSkillRate = r.GetUInt64(field)
	case 68:
		prop.ReqMp = r.GetUInt64(field)
	case 69:
		prop.ReqFp = r.GetUInt64(field)
	case 70:
		prop.ReqDisLV = r.GetUInt64(field)
	case 71:
		prop.ReSkill1 = r.GetUInt64(field)
	case 72:
		prop.ReSkillLevel1 = r.GetUInt64(field)
	case 73:
		prop.ReSkill2 = r.GetUInt64(field)
	case 74:
		prop.ReSkillLevel2 = r.GetUInt64(field)
	case 75:
		prop.SkillReadyType = r.GetUInt64(field)
	case 76:
		prop.SkillReady = r.GetUInt64(field)
	case 77:
		prop.SkillRange = r.GetUInt64(field)
	case 78:
		prop.SfxElemental = r.GetUInt64(field)
	case 79:
		prop.SfxObj = r.GetUInt32(field)
	case 80:
		prop.SfxObj2 = r.GetUInt32(field)
	case 81:
		prop.SfxObj3 = r.GetUInt32(field)
	case 82:
		prop.SfxObj4 = r.GetUInt32(field)
	case 83:
		prop.SfxObj5 = r.GetUInt32(field)
	case 84:
		prop.UseMotion = r.GetUInt32(field)
	case 85:
		prop.CircleTime = r.GetUInt64(field)
	case 86:
		prop.SkillTime = r.GetUInt64(field)
	case 87:
		prop.ExeTarget = r.GetUInt64(field)
	case 88:
		prop.UseChance = r.GetUInt32(field)
	case 89:
		prop.SpellRegion = r.GetUInt64(field)
	case 90:
		prop.SpellType = r.GetUInt64(field)
	case 91:
		prop.ReferStat1 = r.GetUInt32(field)
	case 92:
		prop.ReferStat2 = r.GetUInt32(field)
	case 93:
		prop.ReferTarget1 = r.GetUInt32(field)
	case 94:
		prop.ReferTarget2 = r.GetUInt32(field)
	case 95:
		prop.ReferValue1 = r.GetUInt64(field)
	case 96:
		prop.ReferValue2 = r.GetUInt64(field)
	case 97:
		prop.SkillType = r.GetUInt64(field)
	case 98:
		prop.ItemResistElecricity = r.GetInt32(field)
	case 99:
		prop.ItemResistFire = r.GetInt32(field)
	case 100:
		prop.ItemResistWind = r.GetInt32(field)
	case 101:
		prop.ItemResistWater = r.GetInt32(field)
	case 102:
		prop.ItemResistEarth = r.GetInt32(field)
	case 103:
		prop.Evildoing = r.GetInt64(field)
	case 104:
		prop.ExpertLV = r.GetUInt32(field)
	case 105:
		prop.ExpertMax = r.GetUInt64(field)
	case 106:
		prop.SubDefine = r.GetUInt32(field)
	case 107:
		prop.Exp = r.GetUInt64(field)
	case 108:
		prop.ComboStyle = r.GetUInt64(field)
	case 109:
		prop.FlightSpeed = r.GetFloat64(field)
	case 110:
		prop.FlightLRAngle = r.GetFloat64(field)
	case 111:
		prop.FlightTBAngle = r.GetFloat64(field)
	case 112:
		prop.FlightLimit = r.GetUInt64(field)
	case 113:
		prop.FFuelReMax = r.GetUInt64(field)
	case 114:
		prop.AFuelReMax = r.GetUInt64(field)
	case 115:
		prop.FuelRe = r.GetUInt64(field)
	case 116:
		prop.LimitLevel1 = r.GetUInt64(field)
	case 117:
		prop.Reflect = r.GetInt32(field)
	case 118:
		prop.SndAttack1 = r.GetUInt32(field)
	case 119:
		prop.SndAttack2 = r.GetUInt32(field)
	case 120:
		prop.QuestID = r.GetUInt64(field)
	case 121:
		prop.TextFileName = field
	}
}

func loadPropItem() (props map[int32]*ItemProp) {
	props = make(map[int32]*ItemProp)
	r := &reader.Reader{Filename: "propItem.txt"}
	if err := r.ReadAll(); err != nil {
		panic(err)
	}

	lines, err := scanLines(r)
	if err != nil {
		panic(err.Error())
	}

	for _, line := range lines {
		prop := new(ItemProp)
		for i, field := range strings.Fields(line) {
			fillField(prop, i, strings.TrimSpace(field), r)
		}
		props[prop.ID] = prop
		//props = append(props, prop)
	}

	return props
}

func scanLines(r *reader.Reader) ([]string, error) {
	scanner := bufio.NewScanner(r.BytesReader)

	scanner.Split(bufio.ScanLines)

	var lines []string
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "//") {
			continue
		}
		if len(strings.Fields(line)) != 124 {
			continue
		}

		lines = append(lines, line)
	}

	return lines, nil
}
