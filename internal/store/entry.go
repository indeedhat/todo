package store

import "gorm.io/gorm"

type Entry struct {
	Model
	ListId uint   `gorm:"index" json:"-"`
	Text   string `gorm:"size:255" json:"text"`
	Done   bool
}

// CreateEntry on a TODO list
func CreateEntry(db *gorm.DB, listId uint, text string) *Entry {
	entry := Entry{
		ListId: listId,
		Text:   text,
	}

	if tx := db.Create(&entry); tx.Error != nil {
		return nil
	}

	return &entry
}

// FindEntriy by its id
func FindEntriy(db *gorm.DB, entryId uint) *Entry {
	var entry Entry

	if tx := db.First(&entry, entryId); tx.Error != nil {
		return nil
	}

	return &entry
}

// ListEntries on a todo list
func ListEntries(db *gorm.DB, listId uint) []Entry {
	var entries []Entry

	_ = db.Find(&entries, "list_id = ?", listId)

	return entries
}
