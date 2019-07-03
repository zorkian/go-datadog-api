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

func TestAPPKeyCreateGetAndDelete(t *testing.T) {
	keyName := "client-test-key"
	expected, err := client.CreateAPPKey(keyName)
	if err != nil {
		t.Fatalf("Creating APP key failed when it shouldn't. (%s)", err)
	}
	defer cleanUpAPPKey(t, *expected.Hash)

	if *expected.Name != keyName {
		t.Errorf("Created key has wrong name. Got %s, want %s", *expected.Name, keyName)
	}

	// now try to fetch it freshly and compare it again
	actual, err := client.GetAPPKey(*expected.Hash)
	if err != nil {
		t.Fatalf("Retrieving APP key failed when it shouldn't. (%s)", err)
	}
	assertAPPKeyEquals(t, actual, expected)
}

func TestAPPKeyUpdateName(t *testing.T) {
	keyName := "client-test-key"
	newKeyName := "client-test-key-new"
	keyStruct, err := client.CreateAPPKey(keyName)
	if err != nil {
		t.Fatalf("Creating APP key failed when it shouldn't. (%s)", err)
	}
	defer cleanUpAPPKey(t, *keyStruct.Hash)

	*keyStruct.Name = newKeyName
	err = client.UpdateAPPKey(keyStruct)
	if err != nil {
		t.Fatalf("Updating APP key failed when it shouldn't. (%s)", err)
	}

	if *keyStruct.Name != newKeyName {
		t.Errorf("APP key name not updated. Got %s, want %s", *keyStruct.Name, newKeyName)
	}
}

func TestAPPKeyGetMultipleKeys(t *testing.T) {
	key1Name := "client-test-1"
	key2Name := "client-test-2"
	key1, err := client.CreateAPPKey(key1Name)
	if err != nil {
		t.Fatalf("Creating APP key failed when it shouldn't. (%s)", err)
	}
	defer cleanUpAPPKey(t, *key1.Hash)
	key2, err := client.CreateAPPKey(key2Name)
	if err != nil {
		t.Fatalf("Creating APP key failed when it shouldn't. (%s)", err)
	}
	defer cleanUpAPPKey(t, *key2.Hash)
	allKeys, err := client.GetAPPKeys()
	if err != nil {
		t.Fatalf("Getting all APP keys failed when it shouldn't (%s)", err)
	}
	key1Found, key2Found := false, false
	for _, key := range allKeys {
		switch *key.Name {
		case key1Name:
			assertAPPKeyEquals(t, &key, key1)
			key1Found = true
		case key2Name:
			assertAPPKeyEquals(t, &key, key2)
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

func assertAPPKeyEquals(t *testing.T, actual, expected *datadog.APPKey) {
	if *actual.Owner != *expected.Owner {
		t.Errorf("APPKey owner does not match: %s != %s", *actual.Owner, *expected.Owner)
	}
	if *actual.Hash != *expected.Hash {
		t.Errorf("APPKey hash does not match: %s != %s", *actual.Hash, *expected.Hash)
	}
	if *actual.Name != *expected.Name {
		t.Errorf("APPKey name does not match: %s != %s", *actual.Name, *expected.Name)
	}
}

func cleanUpAPPKey(t *testing.T, hash string) {
	if err := client.DeleteAPPKey(hash); err != nil {
		t.Fatalf("Deleting app key failed when it shouldn't. Manual cleanup needed. (%s)", err)
	}

	deletedKey, err := client.GetAPPKey(hash)
	if deletedKey != nil {
		t.Fatal("APP key hasn't been deleted when it should have been. Manual cleanup needed.")
	}

	if err == nil {
		t.Fatal("Fetching deleted APP key didn't lead to an error. Manual cleanup needed.")
	}
}
