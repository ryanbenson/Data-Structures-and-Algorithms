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

// getNewsFeed returns the content that the users posted based on the user's followers list
func getNewsFeed(userID int) ([]*tweet, error) {
  if userID == 0 {
    return nil, errors.New("A user is required to get a news feed")
  }

  feed := []*tweet{}

  // do we have a user?
  var userMatch *user
  for _, user := range users {
    if user.userID == userID {
      userMatch = user
      break
    }
  }

  if userMatch == nil {
    return nil, errors.New("Invalid user")
  }

  followerLen := len(userMatch.following)
  if followerLen == 0 {
    return nil, errors.New("Not following anyone")
  }
  
  // find our followers and their posts
  for _, followingUser := range userMatch.following {
    for _, tweet := range tweets {
      if tweet.userID == followingUser {
        feed = append(feed, tweet)
      }
    }
  }

  return feed, nil
}

// postTweet adds a new tweet to the system
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

// follow adds a follower to the list of followers for the user
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
    return false, errors.New("Unable to find user")

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

// unfollow removes a follower from a user's follower list
func unfollow(userID int, followerUserID int) (bool, error) {
  if userID == followerUserID {
    return false, errors.New("Unable to unfollow yourself")
  }

  // get our user
  var userMatch *user
  for _, user := range users {
    if user.userID == userID {
      userMatch = user
      break
    }
  }

  if userMatch == nil {
    return false, errors.New("User does not exist")
  }

  // find the user we are to unfollow
  for i, followingUser := range userMatch.following {
    // if we find the match, remove it and return early
    if followingUser == followerUserID {
      userMatch.following = removeIndex(userMatch.following, i)
      return true, nil
    }
  }

  // if we get here, the user was never following that user
  return false, errors.New("Unable to unfollow a user that is not followed")
}

// removeIndex removes an item from an array at the given index
func removeIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}
