/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service
 *
 * API version: 2.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package SubscriberDataManagement

import (
	"free5gc/lib/http_wrapper"
	"free5gc/src/udm/udm_handler"
	"free5gc/src/udm/udm_handler/udm_message"

	"github.com/gin-gonic/gin"
)

// GetSmfSelectData - retrieve a UE's SMF Selection Subscription Data
func GetSmfSelectData(c *gin.Context) {

	req := http_wrapper.NewRequest(c.Request, nil)
	req.Params["supi"] = c.Params.ByName("supi")
	req.Query.Set("plmn-id", c.Query("plmn-id"))

	handlerMsg := udm_message.NewHandlerMessage(udm_message.EventGetSmfSelectData, req)
	udm_handler.SendMessage(handlerMsg)

	rsp := <-handlerMsg.ResponseChan

	HTTPResponse := rsp.HTTPResponse

	c.JSON(HTTPResponse.Status, HTTPResponse.Body)
	return
}