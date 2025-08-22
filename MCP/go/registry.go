package main

import (
	"github.com/apimanagementclient/mcp-server/config"
	"github.com/apimanagementclient/mcp-server/models"
	tools_operationsbytags "github.com/apimanagementclient/mcp-server/tools/operationsbytags"
	tools_apitags "github.com/apimanagementclient/mcp-server/tools/apitags"
	tools_tags "github.com/apimanagementclient/mcp-server/tools/tags"
	tools_operationtags "github.com/apimanagementclient/mcp-server/tools/operationtags"
	tools_apitagdescriptions "github.com/apimanagementclient/mcp-server/tools/apitagdescriptions"
	tools_producttags "github.com/apimanagementclient/mcp-server/tools/producttags"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_operationsbytags.CreateOperation_listbytagsTool(cfg),
		tools_apitags.CreateTag_listbyapiTool(cfg),
		tools_apitags.CreateTag_detachfromapiTool(cfg),
		tools_apitags.CreateTag_getbyapiTool(cfg),
		tools_apitags.CreateTag_getentitystatebyapiTool(cfg),
		tools_apitags.CreateTag_assigntoapiTool(cfg),
		tools_tags.CreateTag_listbyserviceTool(cfg),
		tools_tags.CreateTag_deleteTool(cfg),
		tools_tags.CreateTag_getTool(cfg),
		tools_tags.CreateTag_getentitystateTool(cfg),
		tools_tags.CreateTag_updateTool(cfg),
		tools_tags.CreateTag_createorupdateTool(cfg),
		tools_operationtags.CreateTag_listbyoperationTool(cfg),
		tools_apitagdescriptions.CreateTagdescription_listbyapiTool(cfg),
		tools_apitagdescriptions.CreateTagdescription_deleteTool(cfg),
		tools_apitagdescriptions.CreateTagdescription_getTool(cfg),
		tools_apitagdescriptions.CreateTagdescription_getentitystateTool(cfg),
		tools_apitagdescriptions.CreateTagdescription_createorupdateTool(cfg),
		tools_producttags.CreateTag_listbyproductTool(cfg),
		tools_operationtags.CreateTag_detachfromoperationTool(cfg),
		tools_operationtags.CreateTag_getbyoperationTool(cfg),
		tools_operationtags.CreateTag_getentitystatebyoperationTool(cfg),
		tools_operationtags.CreateTag_assigntooperationTool(cfg),
		tools_producttags.CreateTag_detachfromproductTool(cfg),
		tools_producttags.CreateTag_getbyproductTool(cfg),
		tools_producttags.CreateTag_getentitystatebyproductTool(cfg),
		tools_producttags.CreateTag_assigntoproductTool(cfg),
	}
}
