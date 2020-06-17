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
	followerID := 9

	newUser := &user{
		userID: 2,
		following: []int{followerID},
	}
	users = append(users, newUser)

	result, err := follow(2, followerID)
	expected := false

	if result != expected {
		t.Errorf("When following a user but user does not exist, got: %v, want: %v.", result, expected)
	}

	if err == nil {
		t.Errorf("When following a user but user does not exist, an error should have been given")
	}
}

func TestUnfollow_InvalidSelfIDs(t *testing.T) {
	followerID := 1
	userID := 1

	result, err := unfollow(userID, followerID)
	expected := false

	if result != expected {
		t.Errorf("When unfollowing a user that is yourself, got: %v, want: %v.", result, expected)
	}

	if err == nil {
		t.Errorf("When unfollowing a user that is yourself , an error should have been given, got: %v, want: %v.", err, nil)
	}
}

func TestUnfollow_NoUserIDMatch(t *testing.T) {
	followerID := 1
	userID := 890

	result, err := unfollow(userID, followerID)
	expected := false

	if result != expected {
		t.Errorf("When unfollowing a user when the user does not exist, got: %v, want: %v.", result, expected)
	}

	if err == nil {
		t.Errorf("When unfollowing a user when the user does not exist, an error should have been given, got: %v, want: %v.", err, nil)
	}
}

func TestUnfollow_NotFollowingUserID(t *testing.T) {
	followerID := 1
	userID := 3

	newUser := &user{
		userID: userID,
		following: []int{2},
	}
	users = append(users, newUser)
	
	result, err := unfollow(userID, followerID)
	expected := false

	if result != expected {
		t.Errorf("When unfollowing a user but you are not following that user, got: %v, want: %v.", result, expected)
	}

	if err == nil {
		t.Errorf("When unfollowing a user but you are not following that user, an error should have been given, got: %v, want: %v.", err, nil)
	}
}

func TestUnfollow(t *testing.T) {
	followerID := 9
	userID := 4

	newUser := &user{
		userID: userID,
		following: []int{followerID},
	}
	users = append(users, newUser)
	
	result, err := unfollow(userID, followerID)
	expected := true

	if result != expected {
		t.Errorf("When unfollowing a user, got: %v, want: %v.", result, expected)
	}

	if err != nil {
		t.Errorf("When unfollowing a user, an error should not have been given, got: %v, want: %v.", err, nil)
	}
}
