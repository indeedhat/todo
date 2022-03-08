package store

import "gorm.io/gorm"

type List struct {
	Model
	Slug  string `gorm:"uniqueIndex;size:100" json:"slug"`
	Title string `json:"title"`
}

// CreateList create a new TODO list
func CreateList(db *gorm.DB, title, slug string) *List {
	list := List{
		Title: title,
		Slug:  slug,
	}

	if tx := db.Create(&list); tx.Error != nil {
		return nil
	}

	return &list
}

// FindListBySlug
func FindListBySlug(db *gorm.DB, slug string) *List {
	var list List

	if tx := db.First(&list, "slug = ?", slug); tx.Error != nil {
		return nil
	}

	return &list
}

// ListLists gets a list of all todo lists
//
// This i a great name for a function
func ListLists(db *gorm.DB) []List {
	var lists []List

	_ = db.Find(&lists)

	return lists
}
