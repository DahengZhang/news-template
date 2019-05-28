package db

import (
	"dahengzhang/news/dto"
)

// GetNews 新闻列表
func GetNews() ([]dto.News, error) {
	stmtOut, err := dbConn.Prepare("SELECT nid, title, LEFT(content, 50), create_time, update_time FROM news")
	if err != nil {
		return []dto.News{}, err
	}
	defer stmtOut.Close()

	var result []dto.News

	rows, err := stmtOut.Query()
	for rows.Next() {
		var nid int
		var title, content, creatTime, updateTime string
		err := rows.Scan(&nid, &title, &content, &creatTime, &updateTime)
		if err != nil {
			continue
		}
		result = append(result, dto.News{
			Nid:        nid,
			Title:      title,
			Content:    content,
			CreateTime: creatTime,
			UpdateTime: updateTime,
		})
	}

	return result, nil
}

// CreateNews 创建新闻
func CreateNews(title, content string) error {
	stmtIns, err := dbConn.Prepare("INSERT INTO news (title, content) VALUES (?,?)")
	if err != nil {
		return err
	}
	defer stmtIns.Close()

	_, err = stmtIns.Exec(title, content)
	if err != nil {
		return err
	}

	return nil
}

// SearchNews 查询新闻
func SearchNews(id int) (dto.News, error) {
	stmtOut, err := dbConn.Prepare("SELECT nid, title, content, create_time, update_time FROM news WHERE nid=?")
	if err != nil {
		return dto.News{}, err
	}
	defer stmtOut.Close()

	var nid int
	var title, content, createTime, updateTime string
	err = stmtOut.QueryRow(id).Scan(&nid, &title, &content, &createTime, &updateTime)
	if err != nil {
		return dto.News{}, err
	}

	return dto.News{
		Nid:        nid,
		Title:      title,
		Content:    content,
		CreateTime: createTime,
		UpdateTime: updateTime,
	}, nil
}

// EditNews 编辑新闻
func EditNews(id int, title, content string) error {
	stmtIns, err := dbConn.Prepare("UPDATE news SET title=?, content=? WHERE nid=?")
	if err != nil {
		return err
	}
	defer stmtIns.Close()

	_, err = stmtIns.Exec(title, content, id)
	if err != nil {
		return err
	}

	return nil
}

// DeleteNews 删除新闻
func DeleteNews(id int) error {
	stmt, err := dbConn.Prepare("DELETE FROM news WHERE nid=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
