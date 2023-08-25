package main

import (
	"context"
	feed "douyin/kitex_gen/feed"
)

// FeedServiceImpl implements the last service interface defined in the IDL.
type FeedServiceImpl struct{}

// ListVideos implements the FeedServiceImpl interface.
func (s *FeedServiceImpl) ListVideos(ctx context.Context, request *feed.DouyinFeedRequest) (resp *feed.DouyinFeedResponse, err error) {
	// TODO: Your code here...
	return
}

// QueryVideos implements the FeedServiceImpl interface.
func (s *FeedServiceImpl) QueryVideos(ctx context.Context, video *feed.Video) (resp *feed.VideoIdRequest, err error) {
	// TODO: Your code here...
	return
}
