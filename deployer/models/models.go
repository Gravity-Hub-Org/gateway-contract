package models

type States []State
type Status int
type RqType int

type State struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
	Type  string      `json:"type"`
}

const (
	None Status = iota
	New
	Rejected
	Success
	Returned
)

const (
	Lock RqType = iota
	Unlock
	Mint
	Burn
)

func (states States) Map() map[string]State {
	stateMap := make(map[string]State)
	for _, v := range states {
		stateMap[v.Key] = v
	}
	return stateMap
}
