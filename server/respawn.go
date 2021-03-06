package main

import (
	"github.com/yenkeia/mirgo/common"
)

type Respawn struct {
	Map          *Map
	Info         common.RespawnInfo
	AliveMonster []*Monster
	DeadMonster  []*Monster
}

func NewRespawn(m *Map, ri common.RespawnInfo) (r *Respawn, err error) {
	r = &Respawn{
		Map:          m,
		Info:         ri,
		AliveMonster: []*Monster{},
		DeadMonster:  []*Monster{},
	}
	for i := 0; i < r.Info.Count; i++ {
		m, err := NewMonster(r)
		if err != nil {
			return nil, err
		}
		r.AliveMonster = append(r.AliveMonster, m)
	}
	return r, nil
}
