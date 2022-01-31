package model

type Announcement struct {
	Id     int    `gorm:"column:id" json:"id"`
	CreatedAt     Time    `gorm:"column:created_at" json:"created_at"`
	Title    string    `gorm:"column:title" json:"title"`
	Content string `gorm:"column:content" json:"content"`
}

func (a Announcement) TableName() string {
	return "announcement"
}

func AddAnnouncement(a *Announcement) error{
	return DB.Create(a).Error
}

func DeleteAnnouncement(a *Announcement) error{
	return DB.Where("title = ?",a.Title).Delete(a).Error
}
