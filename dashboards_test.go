package datadog

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type Alias Yaxis

type YAxisTestSuite struct {
	suite.Suite
	yJSON                     []byte
	yMarshalledJSON           []byte
	y                         Yaxis
	yAutoMinMaxJSON           []byte
	yAutoMinMaxMarshalledJSON []byte
	yAutoMinMax               Yaxis
}

func (suite *YAxisTestSuite) SetupTest() {
	// Custom Y.Min, Y.Max
	suite.yJSON = []byte(`{"min":0,"max":1,"scale":"linear"}`)
	suite.yMarshalledJSON = suite.yJSON
	yMinFloat := float64(0)
	yMaxFloat := float64(1)
	yScale := "linear"
	suite.y = Yaxis{
		Min:     &yMinFloat,
		AutoMin: false,
		Max:     &yMaxFloat,
		AutoMax: false,
		Scale:   &yScale,
	}
	// Auto Y.Min, Y.Max
	suite.yAutoMinMaxJSON = []byte(`{"min":"auto","max":"auto","scale":"linear"}`)
	suite.yAutoMinMaxMarshalledJSON = []byte(`{"scale":"linear"}`)
	suite.yAutoMinMax = Yaxis{
		Min:     nil,
		AutoMin: true,
		Max:     nil,
		AutoMax: true,
		Scale:   &yScale,
	}
}

func (suite *YAxisTestSuite) TestYAxisJSONMarshalCustomMinMax() {
	jsonData, err := json.Marshal(suite.y)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), suite.yMarshalledJSON, jsonData)
}

func (suite *YAxisTestSuite) TestYAxisJSONUnmarshalCustomMinMax() {
	var res Yaxis
	err := json.Unmarshal(suite.yJSON, &res)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), suite.y, res)
}

func (suite *YAxisTestSuite) TestYAxisJSONMarshalAutoMinMax() {
	jsonData, err := json.Marshal(suite.yAutoMinMax)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), suite.yAutoMinMaxMarshalledJSON, jsonData)
}

func (suite *YAxisTestSuite) TestYAxisJSONUnmarshalAutoMinMax() {
	var res Yaxis
	err := json.Unmarshal(suite.yAutoMinMaxJSON, &res)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), suite.yAutoMinMax, res)
}

func TestYAxisTestSuite(t *testing.T) {
	suite.Run(t, new(YAxisTestSuite))
}
