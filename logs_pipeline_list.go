package datadog

const (
	logsPipelineListPath = "/v1/logs/config/pipeline-order"
)

type LogsPipelineList struct {
	PipelineIds []string `json:"pipeline_ids"`
}

func (client *Client) GetLogsPipelineList() (*LogsPipelineList, error) {
	var pipelineList LogsPipelineList
	if err := client.doJsonRequest("GET", logsPipelineListPath, nil, &pipelineList); err != nil {
		return nil, err
	}
	return &pipelineList, nil
}

func (client *Client) UpdateLogsPipelineList(pipelineList *LogsPipelineList) (*LogsPipelineList, error) {
	var updatedPipelineList = &LogsPipelineList{}
	if err := client.doJsonRequest("PUT", logsPipelineListPath, pipelineList, updatedPipelineList); err != nil {
		return nil, err
	}
	return updatedPipelineList, nil
}
