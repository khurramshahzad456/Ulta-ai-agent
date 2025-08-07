package api

import (
	"net/http"
	"ultahost-ai-gateway/internal/agents"
	"ultahost-ai-gateway/internal/ai"
	"ultahost-ai-gateway/internal/pkg/models"

	"github.com/gin-gonic/gin"
)

func HandleChat(c *gin.Context) {
	var req *models.ChatRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	req.UserToken = c.GetString("user_token")

	category, err := ai.ClassifyPromptCategory(&models.CategoryRequest{
		Query: req.Message,
		Categories: []string{
			"billing",
			"vps",
			"domain",
			"products",
			"support",
			"server_metrics",
			"vm_command",
			"unknown",
			"wordpress",
		},
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "AI classifier failed", "details": err.Error()})
		return
	}

	switch category {
	case "vps", "vm_command", "server_metrics", "wordpress":
		resp, err := agents.HandleVPS(req, agents.VPSFunctionList)
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"response": resp})

	case "billing":
		resp, err := agents.HandleBilling(req)
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"response": resp})

	case "domain":
		resp, err := agents.HandleDomain(req)
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"response": resp})

	case "products", "product_info", "hosting_plans":
		resp, err := agents.HandleProducts(req, agents.ProductsFunctionList)
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"response": resp})

	default:
		c.JSON(http.StatusNotImplemented, gin.H{"response": "I couldn’t process this request. Please rephrase or try again."})
	}
}
