package service

import (
	"frame/internal/model"
	"frame/internal/request"
)

func (s *Service) GetTagList(param *request.TagListRequest) (*[]model.Tag, error) {
	return s.dao.GetTagList(param.Name, param.State, param.Page, param.PageSize)
}

//func (s *Service) CountTag(param *request.CountTagRequest) (int, error) {
//	return s.dao.CountTag(param.Name, param.State)
//}
//
//func (s *Service) CreateTag(param *request.CreateTagRequest) error {
//	return s.dao.CreateTag(param.Name, param.State, param.CreateBy)
//}
//
//func (s *Service) UpdateTag(param *request.UpdateTagRequest) error {
//	return s.dao.UpdateTag(param.ID, param.Name, param.State, param.ModifiedBy)
//}
//
//func (s *Service) DeleteTag(param *request.DeleteTagRequest) error {
//	return s.dao.DeleteTag(param.ID)
//}