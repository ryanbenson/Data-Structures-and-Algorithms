package basictwitter

import "time"

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

func postTweet(userID int, post tweet) bool {
  return true
}

func follow(userID int, followerUserID int) bool {
  return true
}

func unfollow(userID int, followerUserID int) bool {
  return true
}