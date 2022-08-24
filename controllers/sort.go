package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jinpikaFE/go_fiber/models"
	"github.com/jinpikaFE/go_fiber/pkg/app"
	"github.com/jinpikaFE/go_fiber/pkg/e"
	"github.com/jinpikaFE/go_fiber/pkg/logging"
	"github.com/jinpikaFE/go_fiber/pkg/valodates"
)

// 获取分类列表
// @Summary 获取分类列表
// @Description 获取分类列表
// @Tags Sort
// @Accept json
// @Produce json
// @Param shop_sort query models.ShopSort true "shop_sort"
// @Param page query e.PageStruct true "page"
// @Success 200 {object} ResponseHTTP{}
// @Failure 503 {object} ResponseHTTP{}
// @Router /v1/sorts [get]
func GetSorts(c *fiber.Ctx) error {
	appF := app.Fiber{C: c}
	maps := &models.ShopSort{}
	page := &e.PageStruct{}
	err2 := c.QueryParser(page)

	if err := c.QueryParser(maps); err != nil && err2 != nil {
		logging.Error(err)
		return appF.Response(fiber.StatusInternalServerError, fiber.StatusInternalServerError, "参数解析错误", err)
	}

	// 入参验证
	if errors := valodates.ValidateStruct(*page); errors != nil {
		return appF.Response(fiber.StatusBadRequest, fiber.StatusBadRequest, "检验参数错误", errors)
	}

	res, errs := models.GetSorts((page.Page-1)*page.PageSize, page.PageSize, maps)
	if errs != nil {
		return appF.Response(fiber.StatusInternalServerError, fiber.StatusInternalServerError, "查询失败", errs)
	}

	total, errTotal := models.GetSortTotal(maps)
	if errTotal != nil {
		return appF.Response(fiber.StatusInternalServerError, fiber.StatusInternalServerError, "查询失败", errTotal)
	}

	data := make(map[string]interface{})
	data["list"] = res
	data["total"] = total
	data["pageSize"] = page.PageSize
	data["page"] = page.Page

	return appF.Response(fiber.StatusOK, fiber.StatusOK, "SUCCESS", data)
}

// 获取分类详情
// @Summary 获取分类详情
// @Description 获取分类详情
// @Tags Sort
// @Accept json
// @Produce json
// @Param shop_sort query models.ShopSort true "shop_sort"
// @Success 200 {object} ResponseHTTP{}
// @Failure 503 {object} ResponseHTTP{}
// @Router /v1/sort [get]
func GetSort(c *fiber.Ctx) error {
	appF := app.Fiber{C: c}
	maps := &models.ShopSort{}

	if err := c.QueryParser(maps); err != nil {
		logging.Error(err)
		return appF.Response(fiber.StatusInternalServerError, fiber.StatusInternalServerError, "参数解析错误", err)
	}

	res, errs := models.GetSort(maps)

	if errs != nil {
		return appF.Response(fiber.StatusInternalServerError, fiber.StatusInternalServerError, "查询失败", errs)
	}

	if !(res.ID > 0) {
		return appF.Response(fiber.StatusBadRequest, fiber.StatusBadRequest, "分类不存在", nil)
	}

	return appF.Response(fiber.StatusOK, fiber.StatusOK, "SUCCESS", res)
}

// 添加sort
// @Summary 添加sort
// @Description 添加sort
// @Tags Sort
// @Accept json
// @Produce json
// @Param shop_sort body models.ShopSort true "shop_sort"
// @Success 200 {object} ResponseHTTP{}
// @Failure 503 {object} ResponseHTTP{}
// @Router /v1/sort [post]
func AddSort(c *fiber.Ctx) error {
	appF := app.Fiber{C: c}
	maps := &models.ShopSort{}

	if err := c.BodyParser(maps); err != nil {
		return appF.Response(fiber.StatusInternalServerError, fiber.StatusInternalServerError, "参数解析错误", err)
	}

	// 入参验证
	errors := valodates.ValidateStruct(*maps)

	if errors != nil {
		return appF.Response(fiber.StatusBadRequest, fiber.StatusBadRequest, "检验参数错误", errors)
	}

	err := models.AddSort(*maps)

	if err != nil {
		return appF.Response(fiber.StatusInternalServerError, fiber.StatusInternalServerError, "添加失败", err)
	}

	return appF.Response(fiber.StatusOK, fiber.StatusOK, "SUCCESS", nil)
}

// 编辑Sort
// @Summary 编辑Sort
// @Description 编辑Sort
// @Tags Sort
// @Accept json
// @Produce json
// @Param id path int true "ShopSort ID"
// @Param shop_sort body models.ShopSort true "shop_sort"
// @Success 200 {object} ResponseHTTP{}
// @Failure 503 {object} ResponseHTTP{}
// @Router /v1/sort/{id} [put]
func EditSort(c *fiber.Ctx) error {
	appF := app.Fiber{C: c}
	id, err := strconv.Atoi(c.Params("id"))
	maps := &models.ShopSort{}

	if err := c.BodyParser(maps); err != nil {
		logging.Error(err)
		return appF.Response(fiber.StatusInternalServerError, fiber.StatusInternalServerError, "参数解析错误", err)
	}

	if err != nil {
		logging.Error(err)
		return appF.Response(fiber.StatusInternalServerError, fiber.StatusInternalServerError, "Params： id 参数解析错误", err)
	}
	if models.ExistSortByID(id) {
		if err := models.EditSort(id, *maps); err != nil {
			return appF.Response(fiber.StatusInternalServerError, fiber.StatusInternalServerError, "编辑失败", err)
		}
		return appF.Response(fiber.StatusOK, fiber.StatusOK, "SUCCESS", nil)
	}

	return appF.Response(fiber.StatusBadRequest, fiber.StatusBadRequest, "id不存在", nil)
}

// 删除sort
// @Summary 删除sort
// @Description 删除sort
// @Tags Sort
// @Accept json
// @Produce json
// @Param id path int true "ShopSort ID"
// @Success 200 {object} ResponseHTTP{}
// @Failure 503 {object} ResponseHTTP{}
// @Router /v1/sort/{id} [delete]
func DelSort(c *fiber.Ctx) error {
	appF := app.Fiber{C: c}
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		logging.Error(err)
		return appF.Response(fiber.StatusInternalServerError, fiber.StatusInternalServerError, "Params： id 参数解析错误", err)
	}

	if models.ExistSortByID(id) {
		if err := models.DeleteSort(id); err != nil {
			return appF.Response(fiber.StatusInternalServerError, fiber.StatusInternalServerError, "删除失败", err)
		}
		return appF.Response(fiber.StatusOK, fiber.StatusOK, "SUCCESS", nil)
	}

	return appF.Response(fiber.StatusBadRequest, fiber.StatusBadRequest, "id不存在", nil)

}
