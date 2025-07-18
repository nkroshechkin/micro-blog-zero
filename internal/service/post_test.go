package service

import (
	"testing"

	"github.com/google/uuid"
	"github.com/nkroshechkin/micro-blog-zero/internal/models"
	"github.com/stretchr/testify/assert"
)

func Test_GetAllPost(t *testing.T) {
	tests := []struct {
		name    string
		posts   []models.Post
		want    []models.Post
		wantErr bool
	}{
		{
			name: "success - get all posts",
			posts: []models.Post{
				{Id: "1", AuthorId: "user1", Text: "Post 1"},
				{Id: "2", AuthorId: "user2", Text: "Post 2"},
			},
			want: []models.Post{
				{Id: "1", AuthorId: "user1", Text: "Post 1"},
				{Id: "2", AuthorId: "user2", Text: "Post 2"},
			},
			wantErr: false,
		},
		{
			name:    "success - empty posts list",
			posts:   []models.Post{},
			want:    []models.Post{},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := &models.DataStructures{Posts: tt.posts}
			service := &postService{ds: ds}

			got, err := service.GetAllPost()

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func Test_GetPost(t *testing.T) {
	basePosts := []models.Post{
		{Id: "1", AuthorId: "user1", Text: "Post 1"},
		{Id: "2", AuthorId: "user2", Text: "Post 2"},
	}

	tests := []struct {
		name    string
		id      string
		want    models.Post
		wantErr bool
		errMsg  string
	}{
		{
			name:    "success - get existing post",
			id:      "1",
			want:    models.Post{Id: "1", AuthorId: "user1", Text: "Post 1"},
			wantErr: false,
		},
		{
			name:    "fail - empty id",
			id:      "",
			want:    models.Post{},
			wantErr: true,
			errMsg:  "id пустой",
		},
		{
			name:    "fail - post not found",
			id:      "999",
			want:    models.Post{},
			wantErr: true,
			errMsg:  "пост не найден",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := &models.DataStructures{Posts: basePosts}
			service := &postService{ds: ds}

			got, err := service.GetPost(tt.id)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, tt.errMsg, err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func Test_CreatePost(t *testing.T) {
	baseUsers := []models.User{
		{Id: "user1", Username: "user1"},
		{Id: "user2", Username: "user2"},
	}

	tests := []struct {
		name      string
		authorId  string
		text      string
		wantErr   bool
		errMsg    string
		wantPosts int
	}{
		{
			name:      "success - create new post",
			authorId:  "user1",
			text:      "New post",
			wantErr:   false,
			wantPosts: 1,
		},
		{
			name:      "fail - invalid author",
			authorId:  "invalid_user",
			text:      "New post",
			wantErr:   true,
			errMsg:    "некоректный пользователь",
			wantPosts: 0,
		},
		{
			name:      "fail - empty text",
			authorId:  "user1",
			text:      "",
			wantErr:   false,
			wantPosts: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := &models.DataStructures{
				Users: baseUsers,
				Posts: []models.Post{},
			}
			service := &postService{ds: ds}

			got, err := service.CreatePost(tt.authorId, tt.text)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, tt.errMsg, err.Error())
				assert.Equal(t, "", got)
			} else {
				assert.NoError(t, err)
				_, err := uuid.Parse(got)
				assert.NoError(t, err)
				assert.Len(t, ds.Posts, tt.wantPosts)
				if tt.wantPosts > 0 {
					assert.Equal(t, tt.authorId, ds.Posts[0].AuthorId)
					assert.Equal(t, tt.text, ds.Posts[0].Text)
				}
			}
		})
	}
}

func Test_LikePost(t *testing.T) {
	baseUsers := []models.User{
		{Id: "user1", Username: "user1", Likes: []string{}},
		{Id: "user2", Username: "user2", Likes: []string{}},
	}

	basePosts := []models.Post{
		{Id: "post1", AuthorId: "user1", Text: "Post 1", LikeList: []string{}},
		{Id: "post2", AuthorId: "user2", Text: "Post 2", LikeList: []string{"user1"}},
	}

	tests := []struct {
		name          string
		userId        string
		postId        string
		wantErr       bool
		errMsg        string
		wantUserLikes []string
		wantPostLikes []string
	}{
		{
			name:          "success - like post",
			userId:        "user1",
			postId:        "post1",
			wantErr:       false,
			wantUserLikes: []string{"post1"},
			wantPostLikes: []string{"user1"},
		},
		{
			name:          "success - unlike post",
			userId:        "user1",
			postId:        "post2",
			wantErr:       false,
			wantUserLikes: []string{},
			wantPostLikes: []string{},
		},
		{
			name:    "fail - invalid user",
			userId:  "invalid_user",
			postId:  "post1",
			wantErr: true,
			errMsg:  "некоректный пользователь",
		},
		{
			name:    "fail - invalid post",
			userId:  "user1",
			postId:  "invalid_post",
			wantErr: true,
			errMsg:  "некоректный пост",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			users := make([]models.User, len(baseUsers))
			copy(users, baseUsers)

			posts := make([]models.Post, len(basePosts))
			copy(posts, basePosts)

			ds := &models.DataStructures{
				Users: users,
				Posts: posts,
			}
			service := &postService{ds: ds}

			_, err := service.LikePost(tt.userId, tt.postId)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, tt.errMsg, err.Error())
			} else {
				assert.NoError(t, err)

				var user *models.User
				for i := range ds.Users {
					if ds.Users[i].Id == tt.userId {
						user = &ds.Users[i]
						break
					}
				}

				var post *models.Post
				for i := range ds.Posts {
					if ds.Posts[i].Id == tt.postId {
						post = &ds.Posts[i]
						break
					}
				}

				assert.Equal(t, tt.wantUserLikes, user.Likes)
				assert.Equal(t, tt.wantPostLikes, post.LikeList)
			}
		})
	}
}
