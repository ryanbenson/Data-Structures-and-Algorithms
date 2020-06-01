package basictwitter

import (
  "errors"
  "time"
)

type tweet struct {
  id int
  userID int
  content string
  createdAt time.Time
}

type user struct {
  id int
  username string
  following []int
}

func getNewsFeed(userID int) []tweet {
  return []tweet{}
}

func postTweet(userID int, post tweet) (bool, error) {
  if userID == 0 {
    return false, errors.New("A user is required to post a tweet")
  }

  return true
}

func follow(userID int, followerUserID int) (bool, error) {
  if userID == 0 {
    return false, errors.New("A user is follow another user")
  }

  if followerUserID == 0 {
    return false, errors.New("A user to follow is follow another user")
  }

  if userID == followerUserID {
    return false, errors.New("Unable to follow yourself")
  }

  return true
}

func unfollow(userID int, followerUserID int) (bool, error) {
  if userID == followerUserID {
    return false, errors.New("Unable to unfollow yourself")
  }

  return true
}
