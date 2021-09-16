package dao

import (
	"frame/internal/model"
	"frame/pkg/app"
)

func (d *Dao) GetTagList(name string, state int, page, pageSize int) (*[]model.Tag, error) {
	tag := model.Tag{Name: name, State: state}
	pageOffset := app.GetPageOffset(page, pageSize)
	return tag.List(d.engine, pageOffset, pageSize)
}

//func (d *Dao) CountTag(name string, state int) (int, error) {
//	tag := model.Tag{Name: name, State: state}
//	return tag.Count(d.engine)
//}
//
//func (d *Dao) CreateTag(name string, state int, createBy string) error {
//	tag := model.Tag{
//		Name: name,
//		State: state,
//		Model: &model.Model{CreateBy: createBy},
//	}
//	return tag.Create(d.engine)
//}
//
//func (d *Dao) UpdateTag(id int, name string, state int, modifiedBy string) error {
//	tag := model.Tag{
//		Name: name,
//		State: state,
//		Model: &model.Model{ModifyedBy: modifiedBy},
//	}
//	return tag.Update(d.engine)
//}
//
//func (d *Dao) DeleteTag(id int) error {
//	tag := model.Tag{ID: id, Model: &model.Model{ID: id}}
//	return tag.Delete(d.engine)
//}