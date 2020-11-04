package stubs

var CreateChannel = "Broker.CreateChannel"
var Publish = "Broker.Publish"
var Subscribe = "Broker.Subscribe"

type Pair struct {
	X int
	Y int
}

type PublishRequest struct {
	Topic string
	Pair Pair
}


type ChannelRequest struct {
	Topic string
	Buffer int
}

type Subscription struct {
	Topic string
	FactoryAddress string
	Callback string
}

type JobReport struct {
	Result int
}

type StatusReport struct {
	Message string
}
