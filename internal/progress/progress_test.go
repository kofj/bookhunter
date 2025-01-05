package progress

import (
	"fmt"
	"os"
	"path"
	"strconv"
	"testing"
	"time"
)

func tempFile() string {
	tempDir := os.TempDir()
	return path.Join(tempDir, "acquire-book-id_"+strconv.FormatInt(time.Now().Unix(), 10))
}

func TestProgress_AcquireBookID(t *testing.T) {
	file := tempFile()
	defer func() { _ = os.Remove(file) }()

	s, err := NewProgress(1, 10000, 1000000, file)
	if err != nil {
		t.Errorf("Error in creating Progress: %v", err)
	}

	now := time.Now().UnixMilli()
	for i := 0; i < 1000; i++ {
		bookID := s.AcquireBookID()
		if bookID != int64(i+1) {
			t.Errorf("The book id doesn't match the desired id.")
		}
	}
	fmt.Println("Total time for acquiring 1000 book IDs: ", time.Now().UnixMilli()-now, "ms")
}

func TestProgress_SaveBookID(t *testing.T) {
	file := tempFile()
	defer func() { _ = os.Remove(file) }()

	s, err := NewProgress(1, 1000, 1000000, file)
	if err != nil {
		t.Errorf("Error in creating Progress: %v", err)
	}

	now := time.Now().UnixMilli()
	for i := 0; i < 500; i++ {
		bookID := s.AcquireBookID()
		if bookID != int64(i+1) {
			t.Errorf("The book id doesn't match the desired id.")
		}

		err = s.SaveBookID(bookID)
		if err != nil {
			t.Errorf("Error in saving download book id: %v", err)
		}
	}
	fmt.Println("Total time for saving 500 book IDs: ", time.Now().UnixMilli()-now, "ms")

	s2, err := NewProgress(1, 1000, 10000, file)
	if err != nil {
		t.Errorf("Error in creating Progress: %v", err)
	}

	bookID := s2.AcquireBookID()
	if bookID != 501 {
		t.Errorf("Error in acquire book id from Progress file. Book id should be %d, but it's %d", 501, bookID)
	}
}
