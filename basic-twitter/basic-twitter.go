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
  userID int
  following []int
}

var tweets []*tweet
var users []*user

func getNewsFeed(userID int) ([]*tweet, error) {
  if userID == 0 {
    return nil, errors.New("A user is required to get a news feed")
  }

  feed := []*tweet{}
  
  for _, tweet := range tweets {
    if tweet.userID == userID {
      feed = append(feed, tweet)
    }
  }

  return feed, nil
}

func postTweet(userID int, post string) (bool, error) {
  if userID == 0 {
    return false, errors.New("A user is required to post a tweet")
  }

  if post == "" {
    return false, errors.New("A post must have some content")
  }

  t := &tweet{
    userID:     userID,
    content:    post,
    createdAt:  time.Now(),
  }
  
  tweets = append(tweets, t)

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
	
  // do we have a user?
  var userMatch *user
  for _, user := range users {
    if user.userID == userID {
      userMatch = user
      break
    }
  }

  if userMatch == nil {
    newUser := &user{
      userID: userID,
      following: []int{followerUserID},
    }
    users = append(users, newUser)
    return true, nil
  }

  // are we already following that user?
  for _, followingUser := range userMatch.following {
    if followingUser == followerUserID {
      return false, errors.New("Already following this user")
    }
  }

  // add the new follower to our list
  _ = append(userMatch.following, followerUserID)
  return true, nil
}

func unfollow(userID int, followerUserID int) (bool, error) {
  if userID == followerUserID {
    return false, errors.New("Unable to unfollow yourself")
  }

  return true, nil
}
