/*
 * Datadog API for Go
 *
 * Please see the included LICENSE file for licensing information.
 *
 * Copyright 2019 by authors and contributors.
 */
package integration

import (
	"testing"

	"github.com/zorkian/go-datadog-api"
)

func TestAPIKeyCreateGetAndDelete(t *testing.T) {
	keyName := "client-test-key"
	expected, err := client.CreateAPIKey(keyName)
	if err != nil {
		t.Fatalf("Creating API key failed when it shouldn't. (%s)", err)
	}
	defer cleanUpAPIKey(t, *expected.Key)

	if *expected.Name != keyName {
		t.Errorf("Created key has wrong name. Got %s, want %s", *expected.Name, keyName)
	}

	// now try to fetch it freshly and compare it again
	actual, err := client.GetAPIKey(*expected.Key)
	if err != nil {
		t.Fatalf("Retrieving API key failed when it shouldn't. (%s)", err)
	}
	assertAPIKeyEquals(t, actual, expected)
}

func TestAPIKeyUpdateName(t *testing.T) {
	keyName := "client-test-key"
	newKeyName := "client-test-key-new"
	keyStruct, err := client.CreateAPIKey(keyName)
	if err != nil {
		t.Fatalf("Creating API key failed when it shouldn't. (%s)", err)
	}
	defer cleanUpAPIKey(t, *keyStruct.Key)

	*keyStruct.Name = newKeyName
	err = client.UpdateAPIKey(keyStruct)
	if err != nil {
		t.Fatalf("Updating API key failed when it shouldn't. (%s)", err)
	}

	if *keyStruct.Name != newKeyName {
		t.Errorf("API key name not updated. Got %s, want %s", *keyStruct.Name, newKeyName)
	}
}

func TestAPIKeyGetMultipleKeys(t *testing.T) {
	key1Name := "client-test-1"
	key2Name := "client-test-2"
	key1, err := client.CreateAPIKey(key1Name)
	if err != nil {
		t.Fatalf("Creating API key failed when it shouldn't. (%s)", err)
	}
	defer cleanUpAPIKey(t, *key1.Key)
	key2, err := client.CreateAPIKey(key2Name)
	if err != nil {
		t.Fatalf("Creating API key failed when it shouldn't. (%s)", err)
	}
	defer cleanUpAPIKey(t, *key2.Key)
	allKeys, err := client.GetAPIKeys()
	if err != nil {
		t.Fatalf("Getting all API keys failed when it shouldn't (%s)", err)
	}
	key1Found, key2Found := false, false
	for _, key := range allKeys {
		switch *key.Name {
		case key1Name:
			assertAPIKeyEquals(t, &key, key1)
			key1Found = true
		case key2Name:
			assertAPIKeyEquals(t, &key, key2)
			key2Found = true
		}
	}
	if key1Found == false {
		t.Errorf("Key 1 not found while getting multiple keys.")
	}
	if key2Found == false {
		t.Errorf("Key 2 not found while getting multiple keys.")
	}
}

func assertAPIKeyEquals(t *testing.T, actual, expected *datadog.APIKey) {
	if *actual.Created != *expected.Created {
		t.Errorf("APIKey created does not match: %s != %s", *actual.Created, *expected.Created)
	}
	if *actual.CreatedBy != *expected.CreatedBy {
		t.Errorf("APIKey created_by does not match: %s != %s", *actual.CreatedBy, *expected.CreatedBy)
	}
	if *actual.Key != *expected.Key {
		t.Errorf("APIKey key does not match: %s != %s", *actual.Key, *expected.Key)
	}
	if *actual.Name != *expected.Name {
		t.Errorf("APIKey name does not match: %s != %s", *actual.Name, *expected.Name)
	}
}

func cleanUpAPIKey(t *testing.T, key string) {
	if err := client.DeleteAPIKey(key); err != nil {
		t.Fatalf("Deleting api key failed when it shouldn't. Manual cleanup needed. (%s)", err)
	}

	deletedKey, err := client.GetAPIKey(key)
	if deletedKey != nil {
		t.Fatal("API key hasn't been deleted when it should have been. Manual cleanup needed.")
	}

	if err == nil {
		t.Fatal("Fetching deleted API key didn't lead to an error. Manual cleanup needed.")
	}
}
