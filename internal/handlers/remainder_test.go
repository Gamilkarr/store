package handlers

import (
	"github.com/Gamilkarr/store/internal/models"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
)

func TestRemainder(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repository := NewMockRepository(ctrl)

	tests := []struct {
		name      string
		req       RemainderRequest
		setupRepo func(*MockRepository)
		want      *RemainderResponse
		wantErr   string
	}{
		{
			name: "base",
			req:  RemainderRequest{StoreID: 1},
			setupRepo: func(repository *MockRepository) {
				repository.EXPECT().GetItemsQuantityOnStore(gomock.Any(), int64(1)).Return(
					[]models.Item{{ItemID: 1, Name: "книга", AvailableQuantity: 1, ReservedQuantity: 1}}, nil)
			},
			want: &RemainderResponse{Items: []models.Item{
				{ItemID: 1, Name: "книга", AvailableQuantity: 1, ReservedQuantity: 1},
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Store{Repository: repository}
			if tt.setupRepo != nil {
				tt.setupRepo(repository)
			}
			got := new(RemainderResponse)
			if err := s.Remainder(tt.req, got); (err != nil) && err.Error() != tt.wantErr {
				t.Errorf("Remainder() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Remainder() got = %v, want %v", got, tt.want)
			}
		})
	}
}
