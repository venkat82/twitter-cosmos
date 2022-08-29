Commands used - 

#
ignite scaffold chain github.com/venkat/twitter
#
ignite scaffold map userData tweets:array.string followers:array.string  --module twitter –no-message
#
ignite scaffold map userFeed feed:array.string  --module twitter –no-message 
#
ignite scaffold message addTweet content --module twitter
#
ignite scaffold message addFollower followerId --module twitter
#
ignite scaffold message fetchFeed --module twitter --response followerTweets:array.string

#
twitterd tx twitter add-tweet "first tweet of alice" --from $alice
#
twitterd tx twitter add-follower $alice --from $bob
#
twitterd tx twitter fetch-feed --from $bob
#
