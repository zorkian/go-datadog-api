package integration

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func Jsonify(v interface{}) (string, error) {
	if b, err := json.MarshalIndent(v, "", " "); err != nil {
		return "", err
	} else {
		return string(b), nil
	}
}

func TestSnapshot(t *testing.T) {

	end := time.Now().Unix()
	start := end - 3600

	query_s := "avg:system.mem.used{*}"
	url, err := client.Snapshot(query_s, time.Unix(start, 0), time.Unix(end, 0), "")

	if err != nil {
		t.Fatalf("Couldn't create snapshot for query(%s): %s", query_s, err)
	}

	fmt.Printf("query snapshot url: %s\n", url)
}

func TestSnapshotGeneric(t *testing.T) {

	end := time.Now().Unix()
	start := end - 3600

	// create new graph def
	graphs := createCustomGraph()

	graph_def, err := Jsonify(graphs[0].Definition)
	if err != nil {
		t.Fatalf("Couldn't create graph_def: %s", err)
	}

	options := map[string]string{"graph_def": graph_def}

	url, err := client.SnapshotGeneric(options, time.Unix(start, 0), time.Unix(end, 0))

	if err != nil {
		t.Fatalf("Couldn't create snapshot from graph_def(%s): %s", graph_def, err)
	}

	fmt.Printf("Graph def snapshot url: %s\n", url)
}
