package probabilities

import (
	"github.com/dlsniper/u/tmp/service"
)

type Probs struct{}

func (p *Probs) Get() {}

func NewProbabilities(service service.SkyService) *Probs {
	return *Probs{

	}
}