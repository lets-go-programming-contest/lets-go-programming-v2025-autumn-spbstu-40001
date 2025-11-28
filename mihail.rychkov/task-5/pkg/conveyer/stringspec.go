package conveyer

import "errors"

type StringConveyer struct {
	Conveyer[string]
}

func New(channelCapacity int) StringConveyer {
	return StringConveyer{NewConveyer[string](channelCapacity)}
}

func (obj *StringConveyer) Recv(outChName string) (string, error) {
	res, err := obj.Conveyer.Recv(outChName)
	if errors.Is(err, ErrClosedChanelEmpty) {
		return "undefined", nil
	}

	return res, err
}
