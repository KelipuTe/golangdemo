package test

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type xxxSuite struct {
	suite.Suite
}

func (s *xxxSuite) SetupSuite() {
	s.T().Log("整个测试执行之前执行一次")
}

func (s *xxxSuite) TearDownSuite() {
	s.T().Log("整个测试执行之后执行一次")
}

func (s *xxxSuite) SetupTest() {
	s.T().Log("每个测试执行之前执行一次")
}

func (s *xxxSuite) TearDownTest() {
	s.T().Log("每个测试执行之后执行一次")
}

func (s *xxxSuite) TestXxxSuite01() {
	t := s.T()
	t.Log("执行测试01")
}

func (s *xxxSuite) TestXxxSuite02() {
	t := s.T()
	t.Log("执行测试02")
}

func TestXxx(t *testing.T) {
	suite.Run(t, &xxxSuite{})
}
