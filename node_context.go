package kdtree

type NodeContext interface {
	IsValid(NodeContext) bool
	IsEqual(NodeContext) bool
}

type DummyNodeContext struct{}

func (context *DummyNodeContext) IsValid(other NodeContext) bool {
	return true
}
func (context *DummyNodeContext) IsEqual(other NodeContext) bool {
	return context == other
}

var DummyNodeContextInstance = &DummyNodeContext{}
