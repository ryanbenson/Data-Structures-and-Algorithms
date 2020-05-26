package basictwitter

import "time"

type tweet struct {
  id int
  userId int
  content string
  createdAt time.Time
}
