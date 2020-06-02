package basictwitter

import (
  "errors"
  "time"
)

type tweet struct {
  userID int
  content string
  createdAt time.Time
}

type user struct {
  username string
  following []int
}

var tweets []tweet

func getNewsFeed(userID int) []tweet {
  return []tweet{}
}

func postTweet(userID int, post string) (bool, error) {
  if userID == 0 {
    return false, errors.New("A user is required to post a tweet")
  }

  if post == "" {
    return false, errors.New("A post must have some content")
  }

  t := tweet{
		userID:     userID,
		content:    post,
		createdAt:  time.Now(),
  }
  
  append(tweets, t)

  return true, nil
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

  return true, nil
}

func unfollow(userID int, followerUserID int) (bool, error) {
  if userID == followerUserID {
    return false, errors.New("Unable to unfollow yourself")
  }

  return true, nil
}
