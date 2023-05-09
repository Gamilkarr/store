package repository

import (
	"context"
	"github.com/Gamilkarr/store/internal/models"
	"github.com/pashagolub/pgxmock/v2"
	"testing"
)

func TestItemReserved(t *testing.T) {
	t.Parallel()
	type args struct {
		storeID int64
		items   []models.Reserved
	}
	tests := []struct {
		name     string
		args     args
		mockInit func(pgxmock.PgxPoolIface)
		wantErr  string
	}{
		{
			name: "base",
			args: args{
				storeID: 1,
				items: []models.Reserved{
					{ItemID: 1, Quantity: 1},
				},
			},
			mockInit: func(mock pgxmock.PgxPoolIface) {
				mock.ExpectBegin()
				mock.ExpectQuery("SELECT availability").WithArgs(int64(1)).
					WillReturnRows(pgxmock.NewRows([]string{"availability"}).AddRow(true))
				mock.ExpectQuery("UPDATE available").WithArgs(int64(1), int64(1), int64(1)).
					WillReturnRows(pgxmock.NewRows([]string{"item_id"}).AddRow(int64(1)))
				mock.ExpectCommit()
			},
		},
		{
			name: "unavailable",
			args: args{
				storeID: 2,
				items: []models.Reserved{
					{ItemID: 2, Quantity: 2},
				},
			},
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
			if err := r.ItemReserved(context.Background(), tt.args.storeID, tt.args.items); (err != nil) &&
				err.Error() != tt.wantErr {
				t.Errorf("ItemReserved() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
