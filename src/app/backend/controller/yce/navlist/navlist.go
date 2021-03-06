package navlist

import (
	"github.com/kataras/iris"
)

type NavListController struct {
	*iris.Context
}

var navList = `
{
    "list": [
        {"id": 1, "name": "Dashboard", "state": "main.dashboard","includeState": "main.dashboard","className":"fa-dashboard"},
        {"id": 2, "name": "应用管理", "state": "main.appManage","includeState": "main.appManage","className":"fa-adn",
            "item": [
                {"id": 21, "name": "发布", "state": "main.appManageDeployment", "includeState": "main.appManageDeployment"},
                {"id": 22, "name": "回滚", "state": "main.appManageRollback", "includeState": "main.appManageRollback"},
                {"id": 22, "name": "滚动升级", "state": "main.appManageRollup", "includeState": "main.appManageRollup"},
                {"id": 22, "name": "撤销", "state": "main.appManageCancel", "includeState": "main.appManageCancel"},
                {"id": 22, "name": "历史发布", "state": "main.appManageHistory", "includeState": "main.appManageHistory"}
            ]
        },
        {"id": 3, "name": "镜像管理", "state": "main.imageManage", "includeState": "main.imageManage","className":"fa-file-archive-o",
            "item": [
                {"id": 31, "name": "查找镜像", "state": "main.imageManageSearch", "includeState": "main.imageManageSearch"},
                {"id": 32, "name": "删除镜像", "state": "main.imageManageDelete", "includeState": "main.imageManageDelete"}
            ]
        },
        {"id": 4, "name": "云盘管理", "state": "main.rbdManage", "includeState": "main.rbdManage","className":"fa-cloud"},
        {"id": 5, "name": "扩展功能", "state": "main.extensions", "includeState": "main.extensions","className":"fa-arrows",
            "item": [
                {"id": 51, "name": "创建服务", "state": "main.extensionsService", "includeState": "main.extensionsService"},
                {"id": 52, "name": "创建访问点", "state": "main.extensionsEndpoint", "includeState": "main.extensionsEndpoint"}
            ]
        },
        {"id": 6, "name": "计费&充值", "state": "main.costManage", "includeState": "main.costManage","className":"fa-credit-card"}
    ]
}
`

func (nlc NavListController) Get() {
	nlc.Write(navList)
}
