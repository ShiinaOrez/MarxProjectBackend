package model

type ArticlesModel struct {
	BaseModel
	Title   string `json:"title" gorm:"column:title;not null"`
	Author  string `json:"author" gorm:"column:author;not null"`
	Content string `json:"content" gorm:"column:content;not null"`
	Images  string `json:"images" gorm:"column:images;not null"`
}

func (c *ArticlesModel) TableName() string {
	return "articles"
}

// Create creates a new Articles account.
func (u *ArticlesModel) Create() error {
	return DB.Self.Create(&u).Error
}

// DeleteArticles deletes the Articles by the Articles identifier.
func DeleteArticles(id uint64) error {
	Articles := ArticlesModel{}
	Articles.BaseModel.Id = id
	return DB.Self.Delete(&Articles).Error
}

// Update updates an Articles account information.
func (u *ArticlesModel) Update() error {
	return DB.Self.Save(u).Error
}

// GetArticles gets an Articles by the Articles identifier.
func GetArticleById(id int) (*ArticlesModel, error) {
	u := &ArticlesModel{}
	d := DB.Self.Where("id = ?", id).First(&u)
	return u, d.Error
}

func GetArticles(page, limit int) ([]ArticlesModel, uint64, error) {
	articles := make([]ArticlesModel, 0)
	var count uint64
	if err := DB.Self.Model(&ArticlesModel{}).Count(&count).Error; err != nil {
		return articles, count, err
	}
	if err := DB.Self.Offset((page - 1) * limit).Limit(limit).Find(&articles).Error; err != nil {
		return articles, count, err
	}
	return articles, count, nil
}
