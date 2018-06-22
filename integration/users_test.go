package integration

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zorkian/go-datadog-api"
)

func TestUserCreateAndDelete(t *testing.T) {
	handle := "test@example.com"
	name := "tester"

	user, err := client.CreateUser(datadog.String(handle), datadog.String(name))
	assert.NotNil(t, user)
	assert.Nil(t, err)
	// Users aren't actually deleted; they're disabled
	// As a result, this doesn't really create a new user. The existing user is disabled
	// at the end of the test, so we enable it here so that we can test deletion (disabling).
	user.Disabled = datadog.Bool(false)
	err = client.UpdateUser(*user)
	assert.Nil(t, err)

	defer func() {
		err := client.DeleteUser(handle)
		if err != nil {
			t.Fatalf("Failed to delete user: %s", err)
		}
	}()

	assert.Equal(t, *user.Handle, handle)
	assert.Equal(t, *user.Name, name)

	newUser, err := client.GetUser(handle)
	assert.Nil(t, err)
	assert.Equal(t, *newUser.Handle, handle)
	assert.Equal(t, *newUser.Name, name)
}
