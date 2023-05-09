package repository

import (
	"context"
	"reflect"
	"testing"

	"github.com/Gamilkarr/store/internal/models"
	"github.com/pashagolub/pgxmock/v2"
)

func TestGetItemsQuantityOnStore(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		storeID  int64
		mockInit func(pgxmock.PgxPoolIface)
		want     []models.Item
		wantErr  string
	}{
		{
			name:    "base",
			storeID: 1,
			mockInit: func(mock pgxmock.PgxPoolIface) {
				mock.ExpectBegin()
				mock.ExpectQuery("SELECT availability").WithArgs(int64(1)).
					WillReturnRows(pgxmock.NewRows([]string{"availability"}).AddRow(true))
				mock.ExpectQuery("SELECT item_id, name, reserved_item, item_quantity - reserved_item AS availability").WithArgs(int64(1)).
					WillReturnRows(pgxmock.NewRows([]string{"item_id", "name", "reserved_item", "availability"}).
						AddRow(int64(1), "книга", int64(10), int64(3)).
						AddRow(int64(2), "стол", int64(5), int64(1)))
				mock.ExpectCommit()
			},
			want: []models.Item{
				{
					ItemID:            1,
					Name:              "книга",
					AvailableQuantity: 3,
					ReservedQuantity:  10,
				},
				{
					ItemID:            2,
					Name:              "стол",
					AvailableQuantity: 1,
					ReservedQuantity:  5,
				},
			},
		},
		{
			name:    "unavailable",
			storeID: 2,
			mockInit: func(mock pgxmock.PgxPoolIface) {
				mock.ExpectBegin()
				mock.ExpectQuery("SELECT availability").WithArgs(int64(2)).
					WillReturnRows(pgxmock.NewRows([]string{"availability"}).AddRow(false))
				mock.ExpectCommit()
			},
			wantErr: "store is not available",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mock, err := pgxmock.NewPool()
			if err != nil {
				t.Fatal(err)
			}
			defer mock.Close()
			if tt.mockInit != nil {
				tt.mockInit(mock)
			}
			r := &Repository{Conn: mock}
			got, err := r.GetItemsQuantityOnStore(context.Background(), tt.storeID)
			if (err != nil) && err.Error() != tt.wantErr {
				t.Errorf("GetItemsQuantityOnStore() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetItemsQuantityOnStore() got = %v, want %v", got, tt.want)
			}
		})
	}
}
