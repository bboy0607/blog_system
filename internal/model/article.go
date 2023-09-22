package model

import (
	"github.com/jinzhu/gorm"
)

type Article struct {
	*Model
	Title         string
	Desc          string
	CoverImageUrl string
	Content       string
	State         uint8
}

type ArticleRow struct {
	ArticleID     uint32
	Title         string
	Desc          string
	CoverImageUrl string
	Content       string
	State         uint8
	TagID         uint32
	TagName       string
}

func (a Article) TableName() string {
	return "blog_article"
}

func (a Article) Create(db *gorm.DB) (*Article, error) {
	if err := db.Create(&a).Error; err != nil {
		return nil, err
	}

	return &a, nil
}

func (a Article) CreateInTransaction(tx *gorm.DB) (*Article, error) {
	if err := tx.Create(&a).Error; err != nil {
		return nil, err
	}

	return &a, nil
}

func (a Article) Count(db *gorm.DB) (int, error) {
	var count int
	if a.Title != "" {
		db = db.Where("title = ?", a.Title)
	}
	db = db.Where("state = ?", a.State)

	err := db.Model(&a).Where("is_del = ?", 0).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (a Article) ListByTagID(db *gorm.DB, tagID uint32, pageOffset int, pageSize int) ([]*ArticleRow, error) {
	fields := []string{"ar.id AS article_id", "ar.title", "ar.desc", "ar.cover_image_url", "ar.content", "ar.state"}
	//加入blog_article_tag表中的tag_id
	fields = append(fields, "at.tag_id")
	//加入blog_tag表中所對應的tag_name
	fields = append(fields, "t.name AS tag_name")

	rows, err := db.Select(fields).Table(Article{}.TableName()+" AS ar").
		Joins("LEFT JOIN `"+ArticleTag{}.TableName()+"` AS at ON ar.id = at.article_id").
		Joins("LEFT JOIN `"+Tag{}.TableName()+"` AS t ON at.tag_id = t.id").
		Where("at.tag_id = ? AND ar.state = ? AND ar.is_del = ?", tagID, a.State, 0).
		Rows()

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articles []*ArticleRow
	for rows.Next() {
		r := &ArticleRow{}
		err := rows.Scan(&r.ArticleID, &r.Title, &r.Desc, &r.CoverImageUrl, &r.Content, &r.State, &r.TagID, &r.TagName)
		if err != nil {
			return nil, err
		}

		articles = append(articles, r)
	}

	return articles, nil
}

func (a Article) List(db *gorm.DB, pageOffset int, pageSize int) ([]*Article, error) {
	var articles []*Article
	query := db.Model(&Article{})

	//如果標題不為空，則在搜索條件內新增查詢標題
	if a.Title != "" {
		query = db.Where("title = ?", a.Title)
	}

	//加入查詢狀態條件與分頁查詢
	query = query.Where("state = ?", a.State).Offset(pageOffset).Limit(pageSize)

	err := query.Where("is_del = ?", 0).Find(&articles).Error
	if err != nil {
		return nil, err
	}

	return articles, nil
}

func (a Article) Get(db *gorm.DB) (*Article, error) {
	article := &Article{}
	query := db.Model(&Article{})
	if a.ID != 0 {
		query = query.Where("id = ?", a.ID)
	}

	err := query.Where("is_del = ?", 0).Find(&article).Error
	if err != nil {
		return nil, err
	}

	return article, nil
}

func (a Article) Update(db *gorm.DB, values interface{}) error {
	query := db.Model(&Article{})
	if a.ID != 0 {
		query = query.Where("id = ?", a.ID)
	}

	return query.Where("is_del = ?", 0).Updates(values).Error
}

func (a Article) Delete(db *gorm.DB) error {
	return db.Where("id = ? AND is_del = ?", a.Model.ID, 0).Delete(&a).Error
}

func (a Article) CountByTagID(db *gorm.DB, tagID uint32) (int, error) {
	var count int

	err := db.Table(ArticleTag{}.TableName()+" AS at").
		Joins("LEFT JOIN `"+Tag{}.TableName()+"` AS t ON at.tag_id = t.id").
		Joins("LEFT JOIN `"+Article{}.TableName()+"` AS ar ON at.article_id = ar.id").
		Where("at.tag_id = ? AND ar.state = ? AND ar.is_del = ?", tagID, a.State, 0).
		Count(&count).Error

	if err != nil {
		return 0, err
	}

	return count, nil
}
