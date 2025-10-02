package book_test

// import (
// 	"myapp/api/resource/book"
// 	mockDB "myapp/mock/db"
// 	testUtil "myapp/util/test"
// 	"testing"

// 	"github.com/DATA-DOG/go-sqlmock"
// 	"github.com/google/uuid"
// )

// func TestRepository_List(t *testing.T) {
// 	t.Parallel()

// 	db, mock, err := mockDB.NewMockDB()
// 	testUtil.NoError(t, err)

// 	repo := book.NewRepository(db)

// 	mockRows := sqlmock.NewRows([]string{"id", "title", "author"}).
// 		AddRow(uuid.New(), "Book1", "Author1").
// 		AddRow(uuid.New(), "Book2", "Author2")

// 	mock.ExpectQuery("")

// }
