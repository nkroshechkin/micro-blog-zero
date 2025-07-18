package service

import (
	"testing"

	"github.com/nkroshechkin/micro-blog-zero/internal/models"
	"github.com/stretchr/testify/assert"
)

func Test_GetAllUser(t *testing.T) {
	tests := []struct {
		name    string
		users   []models.User
		want    []models.User
		wantErr bool
	}{
		{
			name: "success - get all users",
			users: []models.User{
				{Id: "1", Username: "user1"},
				{Id: "2", Username: "user2"},
			},
			want: []models.User{
				{Id: "1", Username: "user1"},
				{Id: "2", Username: "user2"},
			},
			wantErr: false,
		},
		{
			name:    "success - empty users list",
			users:   []models.User{},
			want:    []models.User{},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := &models.DataStructures{Users: tt.users}
			service := &userService{ds: ds}

			got, err := service.GetAllUser()

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func Test_GetUser(t *testing.T) {
	baseUsers := []models.User{
		{Id: "1", Username: "user1"},
		{Id: "2", Username: "user2"},
	}

	tests := []struct {
		name    string
		id      string
		users   []models.User
		want    models.User
		wantErr bool
		errMsg  string
	}{
		{
			name:    "success - get existing user",
			id:      "1",
			users:   baseUsers,
			want:    models.User{Id: "1", Username: "user1"},
			wantErr: false,
		},
		{
			name:    "fail - empty id",
			id:      "",
			users:   baseUsers,
			want:    models.User{},
			wantErr: true,
			errMsg:  "id пустой",
		},
		{
			name:    "fail - user not found",
			id:      "999",
			users:   baseUsers,
			want:    models.User{},
			wantErr: true,
			errMsg:  "пользователь не найден",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := &models.DataStructures{Users: tt.users}
			service := &userService{ds: ds}

			got, err := service.GetUser(tt.id)

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

func Test_CreateUser(t *testing.T) {
	tests := []struct {
		name     string
		username string
		users    []models.User
		wantErr  bool
		errMsg   string
	}{
		{
			name:     "success - create new user",
			username: "newuser",
			users:    []models.User{},
			wantErr:  false,
		},
		{
			name:     "fail - empty username",
			username: "",
			users:    []models.User{},
			wantErr:  true,
			errMsg:   "имя пользователя пустое",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := &models.DataStructures{Users: tt.users}
			service := &userService{ds: ds}

			got, err := service.CreateUser(tt.username)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, tt.errMsg, err.Error())
				assert.Equal(t, "", got)
			} else {
				assert.NoError(t, err)
				assert.Len(t, ds.Users, len(tt.users)+1)
				assert.Equal(t, tt.username, ds.Users[len(ds.Users)-1].Username)
			}
		})
	}
}
