package service

import "frame/internal/request"

func GetFinalData(req request.PageInfo, total int64, result interface{}) map[string]interface{} {
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}
	finalData := make(map[string]interface{})
	finalData["page"] = req.Page
	finalData["pageSize"] = req.PageSize
	finalData["total"] = total
	finalData["list"] = result
	return finalData
}
