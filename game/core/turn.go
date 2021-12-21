package core

type Turn struct {
	Steps []*Step
}

type Step struct {
	Actions []Action
	Results []Result
}

type Result interface {
	accountInto(state *GameState)
}

type Action interface {
	performWith(state *GameState)
}
