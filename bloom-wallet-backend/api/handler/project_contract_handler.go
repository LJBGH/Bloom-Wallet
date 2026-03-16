package handler

import (
	"bloom-wallet/api/request"
	"bloom-wallet/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProjectContractHandler struct {
	ProjectContractService *services.ProjectContractService
}

func NewProjectContractHandler(projectContractService *services.ProjectContractService) *ProjectContractHandler {
	return &ProjectContractHandler{
		ProjectContractService: projectContractService,
	}
}

// @Summary      根据ID获取项目合约
// @Description  根据ID获取项目合约
// @Tags         salecontract
// @Accept       json
// @Produce      json
// @Param        id  path      string  true  "项目合约ID"
// @Success      200      {object}  model.ProjectContract
// @Router       /project/{id} [get]
func (h *ProjectContractHandler) GetById(c *gin.Context) {
	idStr := c.Param("id")
	idUint64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	saleContract, err := h.ProjectContractService.GetById(uint(idUint64))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": saleContract})
}

// @Summary      获取所有项目合约
// @Description  获取所有项目合约
// @Tags         salecontract
// @Accept       json
// @Produce      json
// @Success      200      {object}  []model.ProjectContract
// @Router       /project/all [get]
func (h *ProjectContractHandler) GetAll(c *gin.Context) {
	saleContracts, err := h.ProjectContractService.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": saleContracts})
}

// @Summary      更新项目合约
// @Description  更新项目合约
// @Tags         salecontract
// @Accept       json
// @Produce      json
// @Param        id  path      string  true  "项目合约ID"
// @Param        request  body      request.ProjectContractUpdateRequest  true  "项目合约更新参数"
// @Success      200      {object}  string
// @Router       /project/{id} [put]
func (h *ProjectContractHandler) Update(c *gin.Context) {
	var request request.ProjectContractUpdateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.ProjectContractService.Update(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "success"})
}
