package controllers

import (
	"go-clean/models"
	"go-clean/request"
	"go-clean/responses"

	"github.com/gin-gonic/gin"
)

var MachineModelCall = new(models.MachineModels)

type ControlMachine struct{}

func (Machine *ControlMachine) GetDeviceCmd(c *gin.Context) {
	sn := c.Query("sn")
	_ = c.Param("kategori")
	data, err := MachineModelCall.GetMachine(sn)
	if err != nil {
		c.JSON(404, responses.MachineDefault{
			Msg:    err.Error(),
			Status: "ERROR",
		})
		c.Abort()
	} else {
		c.JSON(200, responses.MachineData{
			Status: "SUCCESS",
			Data:   data,
		})
		c.Abort()
	}
}
func (Machine *ControlMachine) PostData(c *gin.Context) {
	sn := c.Param("sn")
	var data request.AddMachine
	if c.BindJSON(&data) != nil {
		c.JSON(406, responses.MachineDefault{
			Msg:    "Invalid Form",
			Status: "ERROR",
		})
	} else {
		_, err := MachineModelCall.GetMachine(sn)
		if err != nil {
			c.JSON(402, gin.H{
				"status": "ERROR",
				"msg":    err.Error(),
			})
		} else {
			c.JSON(200, gin.H{
				"status": "SUCCESS",
				"msg":    "Created Data Success",
			})
		}
	}
}
