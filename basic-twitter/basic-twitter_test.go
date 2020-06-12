package basictwitter

import (
	"testing"
)

func TestFollow_Success(t *testing.T) {
	// add user to use that has no followers
	newUser := &user{
		userID: 1,
		following: []int{},
	}
	users = append(users, newUser)

	result, err := follow(1, 9)
	expected := true

	if result != expected {
		t.Errorf("When following a user with a valid user and follower that is not already followed, got: %v, want: %v.", result, expected)
	}

	if err != nil {
		t.Errorf("When following a user with a valid user and follower that is not already followed, no error should have been given, got: %v, want: %v.", err, nil)
	}
}

func TestFollow_InvalidUserId(t *testing.T) {
	result, err := follow(0, 9)
	expected := false

	if result != expected {
		t.Errorf("When following a user but user does not exist, got: %v, want: %v.", result, expected)
	}

	if err == nil {
		t.Errorf("When following a user but user does not exist, an error should have been given")
	}
}

func TestFollow_InvalidFollowerId(t *testing.T) {
	result, err := follow(1, 0)
	expected := false

	if result != expected {
		t.Errorf("When following a user but user does not exist, got: %v, want: %v.", result, expected)
	}

	if err == nil {
		t.Errorf("When following a user but user does not exist, an error should have been given")
	}
}

func TestFollow_InvalidSelfUserAndFollowerId(t *testing.T) {
	result, err := follow(1, 1)
	expected := false

	if result != expected {
		t.Errorf("When following a user but user does not exist, got: %v, want: %v.", result, expected)
	}

	if err == nil {
		t.Errorf("When following a user but user does not exist, an error should have been given")
	}
}

func TestFollow_BadUserId(t *testing.T) {
	result, err := follow(9999, 9)
	expected := false

	if result != expected {
		t.Errorf("When following a user but user does not exist, got: %v, want: %v.", result, expected)
	}

	if err == nil {
		t.Errorf("When following a user but user does not exist, an error should have been given")
	}
}

func TestFollow_InvalidFollowerUserIdAlreadyFollowing(t *testing.T) {
	// reset our users
	followerId := 9

	newUser := &user{
		userID: 2,
		following: []int{followerId},
	}
	users = append(users, newUser)

	result, err := follow(2, followerId)
	expected := false

	if result != expected {
		t.Errorf("When following a user but user does not exist, got: %v, want: %v.", result, expected)
	}

	if err == nil {
		t.Errorf("When following a user but user does not exist, an error should have been given")
	}
}
