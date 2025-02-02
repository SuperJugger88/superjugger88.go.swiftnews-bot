package models

type State uint

const (
	StateDefault State = iota
	StateAskName
	StateShowKeyboard
)
