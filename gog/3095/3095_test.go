package pull

import (
	"testing"

	"gopkg.in/check.v1"
)

func Test(t *testing.T) { check.TestingT(t) }

func init() {
	check.Suite(&PullSuite{})
}

type PullSuite struct{}

func (s *PullSuite) TestThatMyFriend(c *check.C) {
	c.Assert(true, check.Equals, true)
}
