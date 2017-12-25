package integration

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func Jsonify(v interface{}) (string, error) {
	b, err := json.MarshalIndent(v, "", " ")
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func init() {
	client = initTest()
}

func TestSnapshot(t *testing.T) {

	end := time.Now().Unix()
	start := end - 3600

	q := "avg:system.mem.used{*}"
	url, err := client.Snapshot(q, time.Unix(start, 0), time.Unix(end, 0), "")

	if err != nil {
		t.Fatalf("Couldn't create snapshot for query(%s): %s", q, err)
	}

	fmt.Printf("query snapshot url: %s\n", url)
}

func TestSnapshotGeneric(t *testing.T) {

	end := time.Now().Unix()
	start := end - 3600

	// create new graph def
	graphs := createCustomGraph()

	g, err := Jsonify(graphs[0].Definition)
	if err != nil {
		t.Fatalf("Couldn't create graphDef: %s", err)
	}

	options := map[string]string{"graphDef": g}

	url, err := client.SnapshotGeneric(options, time.Unix(start, 0), time.Unix(end, 0))

	if err != nil {
		t.Fatalf("Couldn't create snapshot from graphDef(%s): %s", g, err)
	}

	fmt.Printf("Graph def snapshot url: %s\n", url)
}
