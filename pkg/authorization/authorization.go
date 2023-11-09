/*
 * @Author: huangpengju 15713716933@163.com
 * @Date: 2023-10-23 16:23:34
 * @LastEditors: huangpengju 15713716933@163.com
 * @LastEditTime: 2023-11-09 16:21:57
 * @FilePath: \chitchat4.0\pkg\authorization\authorization.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package authorization

import (
	"chitchat4.0/pkg/model"
	"chitchat4.0/pkg/repository"
	"chitchat4.0/pkg/utils/request"
)

var store repository.Repository

func Authorize(user *model.User, ri *request.RequestInfo) (bool, error) {
	if user == nil || ri == nil {
		return false, nil
	}

	// if user.ID == 0 {
	// 	group, err := store.Group().GetGroupByName(model.UnAuthenticatedGroup)
	// 	if err != nil {
	// 		return false, err
	// 	}
	// 	user.Groups = append(user.Groups, *group)
	// } else {
	// 	group, err := store.Group().GetGroupByName(model.AuthenticatedGroup)
	// 	if err != nil {
	// 		return false, err
	// 	}
	// 	user.Groups = append(user.Groups, *group)
	// }

	var err error
	if user.ID != 0 {
		// store 是 repository
		_, err = store.User().GetUserByID(user.ID)
	}

	if err != nil {
		return false, err
	}

	// roles := make([]model.Role, 0)
	// roles = append(roles, user.Roles...)
	// for _, g := range user.Groups {
	// 	roles = append(roles, g.Roles...)
	// }

	// for _, role := range roles {
	// 	if ri.Namespace == "" && role.Scope == model.NamespaceScope {
	// 		continue
	// 	}

	// 	if ri.Namespace != "" && (role.Scope == model.NamespaceScope && role.Namespace != ri.Namespace) {
	// 		continue
	// 	}

	// 	for _, rule := range role.Rules {
	// 		if (rule.Resource == model.All || rule.Resource == ri.Resource) && rule.Operation.Contain(ri.Verb) {
	// 			return true, nil
	// 		}
	// 	}
	// }
	return true, nil

	// return false, nil
}
