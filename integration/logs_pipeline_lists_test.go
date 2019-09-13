package integration

import (
	"github.com/stretchr/testify/assert"
	"github.com/zorkian/go-datadog-api"
	"testing"
)

func TestLogsPipelineListGetAndUpdate(t *testing.T) {
	createdPipeline1, err := client.CreateLogsPipeline(
		&datadog.LogsPipeline{
			Name:      datadog.String("my first pipeline"),
			IsEnabled: datadog.Bool(true),
			Filter: &datadog.FilterConfiguration{
				Query: datadog.String("source:redis"),
			},
		})
	assert.Nil(t, err)
	defer assertPipelineDelete(t, *createdPipeline1.Id)

	createdPipeline2, err := client.CreateLogsPipeline(
		&datadog.LogsPipeline{
			Name:      datadog.String("my second pipeline"),
			IsEnabled: datadog.Bool(true),
			Filter: &datadog.FilterConfiguration{
				Query: datadog.String("source:redis"),
			},
		})
	assert.Nil(t, err)
	defer assertPipelineDelete(t, *createdPipeline2.Id)

	pipelineList, err := client.GetLogsPipelineList()
	assert.Nil(t, err)
	size := len(pipelineList.PipelineIds)
	assert.True(t, size >= 2)
	assert.Equal(t, *createdPipeline1.Id, pipelineList.PipelineIds[size-2])
	assert.Equal(t, *createdPipeline2.Id, pipelineList.PipelineIds[size-1])

	pipelineList.PipelineIds[size-2], pipelineList.PipelineIds[size-1] =
		pipelineList.PipelineIds[size-1], pipelineList.PipelineIds[size-2]
	updatedList, err := client.UpdateLogsPipelineList(pipelineList)
	assert.Nil(t, err)
	size = len(updatedList.PipelineIds)
	assert.True(t, size >= 2)
	assert.Equal(t, *createdPipeline1.Id, updatedList.PipelineIds[size-1])
	assert.Equal(t, *createdPipeline2.Id, updatedList.PipelineIds[size-2])

}
