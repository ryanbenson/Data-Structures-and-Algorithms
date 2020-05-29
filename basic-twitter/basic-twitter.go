package basictwitter

import "time"

type tweet struct {
  id int
  userID int
  content string
  createdAt time.Time
}

func getNewsFeed(userID int) []tweet {
}
