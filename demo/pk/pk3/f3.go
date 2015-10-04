/**
 * @author Florin Patan <florinpatan@gmail.com>
 */

// Package pk3 ...
package pk3

type (
	De interface {
		GetMessage() string
		UpdateMessage(string) De
	}

	demo struct {
		message string
	}
)

func (d demo) GetMessage() string {
	return d.message
}

func (d demo) UpdateMessage(message string) De {
	d.message = message
	return d
}

func NewDemo() De {
	d := demo{message: "newDemo"}
	return d
}
