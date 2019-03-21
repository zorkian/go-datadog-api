package integration

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zorkian/go-datadog-api"
)

func TestScreenboardCreateAndDelete(t *testing.T) {
	expected := getTestScreenboard()
	// create the screenboard and compare it
	actual, err := client.CreateScreenboard(expected)
	if err != nil {
		t.Fatalf("Creating a screenboard failed when it shouldn't. (%s)", err)
	}

	defer cleanUpScreenboard(t, *actual.Id)

	assertScreenboardEquals(t, actual, expected)

	// now try to fetch it freshly and compare it again
	actual, err = client.GetScreenboard(*actual.Id)
	if err != nil {
		t.Fatalf("Retrieving a screenboard failed when it shouldn't. (%s)", err)
	}

	assertScreenboardEquals(t, actual, expected)

}

func TestScreenboardShareAndRevoke(t *testing.T) {
	expected := getTestScreenboard()
	// create the screenboard
	actual, err := client.CreateScreenboard(expected)
	if err != nil {
		t.Fatalf("Creating a screenboard failed when it shouldn't: %s", err)
	}

	defer cleanUpScreenboard(t, *actual.Id)

	// share screenboard and verify it was shared
	var response datadog.ScreenShareResponse
	err = client.ShareScreenboard(*actual.Id, &response)
	if err != nil {
		t.Fatalf("Failed to share screenboard: %s", err)
	}

	// revoke screenboard
	err = client.RevokeScreenboard(*actual.Id)
	if err != nil {
		t.Fatalf("Failed to revoke sharing of screenboard: %s", err)
	}
}

func TestScreenboardUpdate(t *testing.T) {
	board := createTestScreenboard(t)
	defer cleanUpScreenboard(t, *board.Id)

	board.Title = datadog.String("___New-Test-Board___")
	if err := client.UpdateScreenboard(board); err != nil {
		t.Fatalf("Updating a screenboard failed when it shouldn't: %s", err)
	}

	actual, err := client.GetScreenboard(*board.Id)
	if err != nil {
		t.Fatalf("Retrieving a screenboard failed when it shouldn't: %s", err)
	}

	assertScreenboardEquals(t, actual, board)

}

func TestScreenboardGet(t *testing.T) {
	boards, err := client.GetScreenboards()
	if err != nil {
		t.Fatalf("Retrieving screenboards failed when it shouldn't: %s", err)
	}
	num := len(boards)

	board := createTestScreenboard(t)
	defer cleanUpScreenboard(t, *board.Id)

	boards, err = client.GetScreenboards()
	if err != nil {
		t.Fatalf("Retrieving screenboards failed when it shouldn't: %s", err)
	}

	if num+1 != len(boards) {
		t.Fatalf("Number of screenboards didn't match expected: %d != %d", len(boards), num+1)
	}
}

func TestScreenboardGetWithNewId(t *testing.T) {
	expected := getTestScreenboard()
	// create the screenboard and compare it
	actual, err := client.CreateScreenboard(expected)
	if err != nil {
		t.Fatalf("Creating a screenboard failed when it shouldn't. (%s)", err)
	}

	defer cleanUpScreenboard(t, *actual.Id)

	assertScreenboardEquals(t, actual, expected)

	// try to fetch it freshly and compare it again
	actual, err = client.GetScreenboard(*actual.Id)
	if err != nil {
		t.Fatalf("Retrieving a screenboard failed when it shouldn't. (%s)", err)
	}

	assertScreenboardEquals(t, actual, expected)

	// try to fetch it freshly using the new id format and compare it again
	actualWithNewId, err := client.GetScreenboard(*actual.NewId)
	if err != nil {
		t.Fatalf("Retrieving a screenboard failed when it shouldn't. (%s)", err)
	}
	assertScreenboardEquals(t, actualWithNewId, expected)

	// the ids are equal whether fetching using the old or the new id
	assert.Equal(t, *actualWithNewId.Id, *actual.Id)

	// try to fetch it freshly using a string, but with a wrong value
	actual, err = client.GetScreenboard("random_string")
	if assert.NotNil(t, err) {
		// it should not fail because of the id format
		assert.NotContains(t, err.Error(), "unsupported id type")
		assert.Contains(t, err.Error(), "404")
	}

	// try to fetch it freshly using a boolean
	actual, err = client.GetScreenboard(true)
	if assert.NotNil(t, err) {
		// it should fail because of the id format
		assert.Contains(t, err.Error(), "unsupported id type")
	}

	// try to fetch it freshly using a float64
	actual, err = client.GetScreenboard(5.5)
	if assert.NotNil(t, err) {
		// it should fail because of the id format
		assert.Contains(t, err.Error(), "unsupported id type")
	}

}

func getTestScreenboard() *datadog.Screenboard {
	return &datadog.Screenboard{
		Title:   datadog.String("___Test-Board___"),
		Height:  datadog.Int(600),
		Width:   datadog.Int(800),
		Widgets: []datadog.Widget{},
	}
}

func createTestScreenboard(t *testing.T) *datadog.Screenboard {
	board := getTestScreenboard()
	board, err := client.CreateScreenboard(board)
	if err != nil {
		t.Fatalf("Creating a screenboard failed when it shouldn't: %s", err)
	}

	return board
}

func cleanUpScreenboard(t *testing.T, id int) {
	if err := client.DeleteScreenboard(id); err != nil {
		t.Fatalf("Deleting a screenboard failed when it shouldn't. Manual cleanup needed. (%s)", err)
	}

	deletedBoard, err := client.GetScreenboard(id)
	if deletedBoard != nil {
		t.Fatal("Screenboard hasn't been deleted when it should have been. Manual cleanup needed.")
	}

	if err == nil {
		t.Fatal("Fetching deleted screenboard didn't lead to an error. Manual cleanup needed.")
	}
}

func assertScreenboardEquals(t *testing.T, actual, expected *datadog.Screenboard) {
	if *actual.Title != *expected.Title {
		t.Errorf("Screenboard title does not match: %s != %s", *actual.Title, *expected.Title)
	}
	if *actual.Width != *expected.Width {
		t.Errorf("Screenboard width does not match: %d != %d", *actual.Width, *expected.Width)
	}
	if *actual.Height != *expected.Height {
		t.Errorf("Screenboard width does not match: %d != %d", *actual.Height, *expected.Height)
	}
	if len(actual.Widgets) != len(expected.Widgets) {
		t.Errorf("Number of Screenboard widgets does not match: %d != %d", len(actual.Widgets), len(expected.Widgets))
	}
}
