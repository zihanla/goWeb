package model

import "time"

// Article 新闻结构体
type Article struct {
	Id      int64     `json:"id"`
	Title   string    `json:"title"`
	Author  string    `json:"author"`
	Content string    `json:"content"`
	Hits    int       `json:"hits"`  // 点击量
	Utime   time.Time `json:"utime"` // 修改时间
}

// ArticleGet 查询一条数据
func ArticleGet(id int64) (Article, error) {
	// 创建一个容器 接收数据
	mod := Article{}
	// Unsafe() 不检测数据库对于关系
	err := Db.Unsafe().Get(&mod, "select * from article where id = ? limit 1", id)
	return mod, err
}

// ArticleList 返回数据列表
func ArticleList() ([]Article, error) {
	// 创建一个切片接收集合数据
	mods := make([]Article, 0, 10)
	// order by id desc 为了方便演示 倒叙排列数据
	err := Db.Unsafe().Select(&mods, "select * from article order by id desc limit 10")
	return mods, err
}

// ArticlePage 分页
func ArticlePage(pi, ps int) ([]Article, error) {
	mods := make([]Article, 0, 10)
	// 数据库 分页查询实现方法 (pi -1)*ps,ps
	err := Db.Unsafe().Select(&mods, "select * from article order by id desc limit ?,?", (pi-1)*ps, ps)
	return mods, err
}

// ArticlePageCount 数据总数
func ArticlePageCount() int {
	count := 0
	Db.Get(&count, "select count(id) as count from article\n")
	return count
}

// ArticleDel 删除数据
func ArticleDel(id int64) bool {
	res, _ := Db.Exec("delete from article where id = ?", id)
	if res == nil {
		return false
	}
	// 受影响行数
	rows, _ := res.RowsAffected()
	if rows >= 1 {
		return true
	}
	return false
}

// ArticleAdd 添加数据
func ArticleAdd(article *Article) error {
	_, err := Db.Exec("insert into article (title,author,content,hits,utime) values (?,?,?,?,?)", article.Title, article.Author, article.Content, article.Hits, article.Utime)
	return err
}

// ArticleEdit 修改数据
func ArticleEdit(article *Article) error {
	_, err := Db.Exec("update article set title=?,author=?,content=?,hits=? where id=? ", article.Title, article.Author, article.Content, article.Hits, article.Id)
	return err
}
