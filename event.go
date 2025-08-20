package sdk

type Event struct {
	Type string                 `json:"type"`
	Data map[string]interface{} `json:"data"`
}

type Bus interface {
	Publish(ev Event)
	Subscribe() chan Event
	Unsubscribe(ch chan Event)
}
