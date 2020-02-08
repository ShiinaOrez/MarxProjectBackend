package model

type NewsModel struct {
	BaseModel
	Kind    string `json:"kind" gorm:"column:kind;not null"`
	Title   string `json:"title" gorm:"column:title;not null"`
	Author  string `json:"author" gorm:"column:author;not null"`
	Content string `json:"content" gorm:"column:content;not null"`
	Images  string `json:"images" gorm:"column:images;not null"`
}

func (c *NewsModel) TableName() string {
	return "news"
}

// Create creates a new News account.
func (u *NewsModel) Create() error {
	return DB.Self.Create(&u).Error
}

// DeleteNews deletes the News by the News identifier.
func DeleteNews(id uint64) error {
	News := NewsModel{}
	News.BaseModel.Id = id
	return DB.Self.Delete(&News).Error
}

// Update updates an News account information.
func (u *NewsModel) Update() error {
	return DB.Self.Save(u).Error
}

// GetNews gets an News by the News identifier.
func GetHistoryNewById(id int) (*NewsModel, error) {
	u := &NewsModel{}
	d := DB.Self.Where("id = ? AND kind = 'history'", id).First(&u)
	return u, d.Error
}

func GetStudyNewById(id int) (*NewsModel, error) {
	u := &NewsModel{}
	d := DB.Self.Where("id = ? AND kind = 'study'", id).First(&u)
	return u, d.Error
}

func GetHistoryNews(page, limit int) ([]NewsModel, uint64, error) {
	news := make([]NewsModel, 0)
	var count uint64
	where := "kind = 'history'"
	if err := DB.Self.Model(&NewsModel{}).Where(where).Count(&count).Error; err != nil {
		return news, count, err
	}
	if err := DB.Self.Where(where).Offset((page - 1) * limit).Limit(limit).Find(&news).Error; err != nil {
		return news, count, err
	}
	return news, count, nil
}

func GetStudyNews(page, limit int) ([]NewsModel, uint64, error) {
	news := make([]NewsModel, 0)
	var count uint64
	where := "kind = 'study'"
	if err := DB.Self.Model(&NewsModel{}).Where(where).Count(&count).Error; err != nil {
		return news, count, err
	}
	if err := DB.Self.Where(where).Offset((page - 1) * limit).Limit(limit).Find(&news).Error; err != nil {
		return news, count, err
	}
	return news, count, nil
}
