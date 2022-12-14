package service

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
	"github.com/mssola/user_agent"
	"github.com/it00021hot/gen-admin/api/v1/system"
	commonService "github.com/it00021hot/gen-admin/internal/app/common/service"
	"github.com/it00021hot/gen-admin/internal/app/system/consts"
	"github.com/it00021hot/gen-admin/internal/app/system/model"
	"github.com/it00021hot/gen-admin/internal/app/system/model/entity"
	"github.com/it00021hot/gen-admin/internal/app/system/service/internal/dao"
	"github.com/it00021hot/gen-admin/internal/app/system/service/internal/do"
	"github.com/it00021hot/gen-admin/library/libUtils"
	"github.com/it00021hot/gen-admin/library/liberr"
)

type IUser interface {
	GetAdminUserByUsernamePassword(ctx context.Context, req *system.UserLoginReq) (user *model.LoginUserRes, err error)
	LoginLog(ctx context.Context, params *model.LoginLogParams)
	UpdateLoginInfo(ctx context.Context, id uint64, ip string) (err error)
	NotCheckAuthAdminIds(ctx context.Context) *gset.Set
	GetAdminRules(ctx context.Context, userId uint64) (menuList []*model.UserMenus, permissions []string, err error)
	List(ctx context.Context, req *system.UserSearchReq) (total int, userList []*entity.SysUser, err error)
	GetUsersRoleDept(ctx context.Context, userList []*entity.SysUser) (users []*model.SysUserRoleDeptRes, err error)
	Add(ctx context.Context, req *system.UserAddReq) (err error)
	GetEditUser(ctx context.Context, id uint64) (res *system.UserGetEditRes, err error)
	Edit(ctx context.Context, req *system.UserEditReq) (err error)
	ResetUserPwd(ctx context.Context, req *system.UserResetPwdReq) (err error)
	ChangeUserStatus(ctx context.Context, req *system.UserStatusReq) (err error)
	Delete(ctx context.Context, ids []int) (err error)
}

type userImpl struct {
	CasBinUserPrefix string //CasBin ??????id??????
}

var (
	notCheckAuthAdminIds *gset.Set //???????????????????????????id
	userService          = userImpl{
		CasBinUserPrefix: "u_",
	}
)

func User() IUser {
	return &userService
}

func (s *userImpl) NotCheckAuthAdminIds(ctx context.Context) *gset.Set {
	ids := g.Cfg().MustGet(ctx, "system.notCheckAuthAdminIds")
	if !g.IsNil(ids) {
		notCheckAuthAdminIds = gset.NewFrom(ids)
	}
	return notCheckAuthAdminIds
}

func (s *userImpl) GetAdminUserByUsernamePassword(ctx context.Context, req *system.UserLoginReq) (user *model.LoginUserRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		user, err = s.GetUserByUsername(ctx, req.Username)
		liberr.ErrIsNil(ctx, err)
		liberr.ValueIsNil(user, "??????????????????")
		//????????????
		if libUtils.EncryptPassword(req.Password, user.UserSalt) != user.UserPassword {
			liberr.ErrIsNil(ctx, gerror.New("??????????????????"))
		}
		//????????????
		if user.UserStatus == 0 {
			liberr.ErrIsNil(ctx, gerror.New("??????????????????"))
		}
	})
	return
}

// GetUserByUsername ?????????????????????????????????
func (s *userImpl) GetUserByUsername(ctx context.Context, userName string) (user *model.LoginUserRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		user = &model.LoginUserRes{}
		err = dao.SysUser.Ctx(ctx).Fields(user).Where(dao.SysUser.Columns().UserName, userName).Scan(user)
		liberr.ErrIsNil(ctx, err, "??????????????????")
	})
	return
}

// LoginLog ??????????????????
func (s *userImpl) LoginLog(ctx context.Context, params *model.LoginLogParams) {
	ua := user_agent.New(params.UserAgent)
	browser, _ := ua.Browser()
	loginData := &do.SysLoginLog{
		LoginName:     params.Username,
		Ipaddr:        params.Ip,
		LoginLocation: libUtils.GetCityByIp(params.Ip),
		Browser:       browser,
		Os:            ua.OS(),
		Status:        params.Status,
		Msg:           params.Msg,
		LoginTime:     gtime.Now(),
		Module:        params.Module,
	}
	_, err := dao.SysLoginLog.Ctx(ctx).Insert(loginData)
	if err != nil {
		g.Log().Error(ctx, err)
	}
}

func (s *userImpl) UpdateLoginInfo(ctx context.Context, id uint64, ip string) (err error) {
	g.Try(ctx, func(ctx context.Context) {
		_, err = dao.SysUser.Ctx(ctx).WherePri(id).Update(g.Map{
			dao.SysUser.Columns().LastLoginIp:   ip,
			dao.SysUser.Columns().LastLoginTime: gtime.Now(),
		})
		liberr.ErrIsNil(ctx, err, "??????????????????????????????")
	})
	return
}

// GetAdminRules ????????????????????????
func (s *userImpl) GetAdminRules(ctx context.Context, userId uint64) (menuList []*model.UserMenus, permissions []string, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		//????????????
		isSuperAdmin := false
		//?????????????????????????????????id
		s.NotCheckAuthAdminIds(ctx).Iterator(func(v interface{}) bool {
			if gconv.Uint64(v) == userId {
				isSuperAdmin = true
				return false
			}
			return true
		})
		//????????????????????????
		allRoles, err := Role().GetRoleList(ctx)
		liberr.ErrIsNil(ctx, err)
		roles, err := s.GetAdminRole(ctx, userId, allRoles)
		liberr.ErrIsNil(ctx, err)
		name := make([]string, len(roles))
		roleIds := make([]uint, len(roles))
		for k, v := range roles {
			name[k] = v.Name
			roleIds[k] = v.Id
		}
		//??????????????????
		if isSuperAdmin {
			//????????????????????????
			permissions = []string{"*/*/*"}
			menuList, err = s.GetAllMenus(ctx)
			liberr.ErrIsNil(ctx, err)
		} else {
			menuList, err = s.GetAdminMenusByRoleIds(ctx, roleIds)
			liberr.ErrIsNil(ctx, err)
			permissions, err = s.GetPermissions(ctx, roleIds)
			liberr.ErrIsNil(ctx, err)
		}
	})
	return
}

// GetAdminRole ??????????????????
func (s *userImpl) GetAdminRole(ctx context.Context, userId uint64, allRoleList []*entity.SysRole) (roles []*entity.SysRole, err error) {
	var roleIds []uint
	roleIds, err = s.GetAdminRoleIds(ctx, userId)
	if err != nil {
		return
	}
	roles = make([]*entity.SysRole, 0, len(allRoleList))
	for _, v := range allRoleList {
		for _, id := range roleIds {
			if id == v.Id {
				roles = append(roles, v)
			}
		}
		if len(roles) == len(roleIds) {
			break
		}
	}
	return
}

// GetAdminRoleIds ??????????????????ids
func (s *userImpl) GetAdminRoleIds(ctx context.Context, userId uint64) (roleIds []uint, err error) {
	enforcer, e := commonService.CasbinEnforcer(ctx)
	if e != nil {
		err = e
		return
	}
	//????????????????????????
	groupPolicy := enforcer.GetFilteredGroupingPolicy(0, fmt.Sprintf("%s%d", s.CasBinUserPrefix, userId))
	if len(groupPolicy) > 0 {
		roleIds = make([]uint, len(groupPolicy))
		//????????????id?????????
		for k, v := range groupPolicy {
			roleIds[k] = gconv.Uint(v[1])
		}
	}
	return
}

func (s *userImpl) GetAllMenus(ctx context.Context) (menus []*model.UserMenus, err error) {
	//???????????????????????????
	var allMenus []*model.SysAuthRuleInfoRes
	allMenus, err = Rule().GetIsMenuList(ctx)
	if err != nil {
		return
	}
	menus = make([]*model.UserMenus, len(allMenus))
	for k, v := range allMenus {
		var menu *model.UserMenu
		menu = s.setMenuData(menu, v)
		menus[k] = &model.UserMenus{UserMenu: menu}
	}
	menus = s.GetMenusTree(menus, 0)
	return
}

func (s *userImpl) GetAdminMenusByRoleIds(ctx context.Context, roleIds []uint) (menus []*model.UserMenus, err error) {
	//???????????????????????????id
	err = g.Try(ctx, func(ctx context.Context) {
		enforcer, e := commonService.CasbinEnforcer(ctx)
		liberr.ErrIsNil(ctx, e)
		menuIds := map[int64]int64{}
		for _, roleId := range roleIds {
			//??????????????????
			gp := enforcer.GetFilteredPolicy(0, gconv.String(roleId))
			for _, p := range gp {
				mid := gconv.Int64(p[1])
				menuIds[mid] = mid
			}
		}
		//???????????????????????????
		allMenus, err := Rule().GetIsMenuList(ctx)
		liberr.ErrIsNil(ctx, err)
		menus = make([]*model.UserMenus, 0, len(allMenus))
		for _, v := range allMenus {
			if _, ok := menuIds[gconv.Int64(v.Id)]; gstr.Equal(v.Condition, "nocheck") || ok {
				var roleMenu *model.UserMenu
				roleMenu = s.setMenuData(roleMenu, v)
				menus = append(menus, &model.UserMenus{UserMenu: roleMenu})
			}
		}
		menus = s.GetMenusTree(menus, 0)
	})
	return
}

func (s *userImpl) GetMenusTree(menus []*model.UserMenus, pid uint) []*model.UserMenus {
	returnList := make([]*model.UserMenus, 0, len(menus))
	for _, menu := range menus {
		if menu.Pid == pid {
			menu.Children = s.GetMenusTree(menus, menu.Id)
			returnList = append(returnList, menu)
		}
	}
	return returnList
}

func (s *userImpl) setMenuData(menu *model.UserMenu, entity *model.SysAuthRuleInfoRes) *model.UserMenu {
	menu = &model.UserMenu{
		Id:        entity.Id,
		Pid:       entity.Pid,
		Name:      gstr.CaseCamelLower(gstr.Replace(entity.Name, "/", "_")),
		Component: entity.Component,
		Path:      entity.Path,
		MenuMeta: &model.MenuMeta{
			Icon:        entity.Icon,
			Title:       entity.Title,
			IsLink:      "",
			IsHide:      entity.IsHide == 1,
			IsKeepAlive: entity.IsCached == 1,
			IsAffix:     entity.IsAffix == 1,
			IsIframe:    entity.IsIframe == 1,
		},
	}
	if menu.MenuMeta.IsIframe || entity.IsLink == 1 {
		menu.MenuMeta.IsLink = entity.LinkUrl
	}
	return menu
}

func (s *userImpl) GetPermissions(ctx context.Context, roleIds []uint) (userButtons []string, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		//???????????????????????????id
		enforcer, err := commonService.CasbinEnforcer(ctx)
		liberr.ErrIsNil(ctx, err)
		menuIds := map[int64]int64{}
		for _, roleId := range roleIds {
			//??????????????????
			gp := enforcer.GetFilteredPolicy(0, gconv.String(roleId))
			for _, p := range gp {
				mid := gconv.Int64(p[1])
				menuIds[mid] = mid
			}
		}
		//???????????????????????????
		allButtons, err := Rule().GetIsButtonList(ctx)
		liberr.ErrIsNil(ctx, err)
		userButtons = make([]string, 0, len(allButtons))
		for _, button := range allButtons {
			if _, ok := menuIds[gconv.Int64(button.Id)]; gstr.Equal(button.Condition, "nocheck") || ok {
				userButtons = append(userButtons, button.Name)
			}
		}
	})
	return
}

// List ????????????
func (s *userImpl) List(ctx context.Context, req *system.UserSearchReq) (total int, userList []*entity.SysUser, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.SysUser.Ctx(ctx)
		if req.KeyWords != "" {
			keyWords := "%" + req.KeyWords + "%"
			m = m.Where("user_name like ? or  user_nickname like ?", keyWords, keyWords)
		}
		if req.DeptId != "" {
			deptIds, e := s.getSearchDeptIds(ctx, gconv.Int64(req.DeptId))
			liberr.ErrIsNil(ctx, e)
			m = m.Where("dept_id in (?)", deptIds)
		}
		if req.Status != "" {
			m = m.Where("user_status", gconv.Int(req.Status))
		}
		if req.Mobile != "" {
			m = m.Where("mobile like ?", "%"+req.Mobile+"%")
		}
		if len(req.DateRange) > 0 {
			m = m.Where("created_at >=? AND created_at <=?", req.DateRange[0], req.DateRange[1])
		}
		if req.PageSize == 0 {
			req.PageSize = consts.PageSize
		}
		if req.PageNum == 0 {
			req.PageNum = 1
		}
		total, err = m.Count()
		liberr.ErrIsNil(ctx, err, "????????????????????????")
		err = m.FieldsEx(dao.SysUser.Columns().UserPassword, dao.SysUser.Columns().UserSalt).
			Page(req.PageNum, req.PageSize).Order("id asc").Scan(&userList)
		liberr.ErrIsNil(ctx, err, "????????????????????????")
	})
	return
}

// GetUsersRoleDept ???????????????????????? ????????????
func (s *userImpl) GetUsersRoleDept(ctx context.Context, userList []*entity.SysUser) (users []*model.SysUserRoleDeptRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		allRoles, e := Role().GetRoleList(ctx)
		liberr.ErrIsNil(ctx, e)
		depts, e := Dept().GetFromCache(ctx)
		liberr.ErrIsNil(ctx, e)
		users = make([]*model.SysUserRoleDeptRes, len(userList))
		for k, u := range userList {
			var dept *entity.SysDept
			users[k] = &model.SysUserRoleDeptRes{
				SysUser: u,
			}
			for _, d := range depts {
				if u.DeptId == uint64(d.DeptId) {
					dept = d
				}
			}
			users[k].Dept = dept
			roles, e := s.GetAdminRole(ctx, u.Id, allRoles)
			liberr.ErrIsNil(ctx, e)
			for _, r := range roles {
				users[k].RoleInfo = append(users[k].RoleInfo, &model.SysUserRoleInfoRes{RoleId: r.Id, Name: r.Name})
			}
		}
	})
	return
}

func (s *userImpl) getSearchDeptIds(ctx context.Context, deptId int64) (deptIds []int64, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		deptAll, e := Dept().GetFromCache(ctx)
		liberr.ErrIsNil(ctx, e)
		deptWithChildren := Dept().FindSonByParentId(deptAll, gconv.Int64(deptId))
		deptIds = make([]int64, len(deptWithChildren))
		for k, v := range deptWithChildren {
			deptIds[k] = v.DeptId
		}
		deptIds = append(deptIds, deptId)
	})
	return
}

func (s *userImpl) Add(ctx context.Context, req *system.UserAddReq) (err error) {
	err = s.userNameOrMobileExists(ctx, req.UserName, req.Mobile)
	if err != nil {
		return
	}
	req.UserSalt = grand.S(10)
	req.Password = libUtils.EncryptPassword(req.Password, req.UserSalt)
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			userId, e := dao.SysUser.Ctx(ctx).TX(tx).InsertAndGetId(do.SysUser{
				UserName:     req.UserName,
				Mobile:       req.Mobile,
				UserNickname: req.NickName,
				UserPassword: req.Password,
				UserSalt:     req.UserSalt,
				UserStatus:   req.Status,
				UserEmail:    req.Email,
				Sex:          req.Sex,
				DeptId:       req.DeptId,
				Remark:       req.Remark,
				IsAdmin:      req.IsAdmin,
			})
			liberr.ErrIsNil(ctx, e, "??????????????????")
			e = s.addUserRole(ctx, req.RoleIds, userId)
			liberr.ErrIsNil(ctx, e, "????????????????????????")
			e = s.AddUserPost(ctx, tx, req.PostIds, userId)
			liberr.ErrIsNil(ctx, e)
		})
		return err
	})
	return
}

func (s *userImpl) Edit(ctx context.Context, req *system.UserEditReq) (err error) {
	err = s.userNameOrMobileExists(ctx, "", req.Mobile, req.UserId)
	if err != nil {
		return
	}
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			_, err = dao.SysUser.Ctx(ctx).TX(tx).WherePri(req.UserId).Update(do.SysUser{
				Mobile:       req.Mobile,
				UserNickname: req.NickName,
				UserStatus:   req.Status,
				UserEmail:    req.Email,
				Sex:          req.Sex,
				DeptId:       req.DeptId,
				Remark:       req.Remark,
				IsAdmin:      req.IsAdmin,
			})
			liberr.ErrIsNil(ctx, err, "????????????????????????")
			//??????????????????????????????
			err = s.EditUserRole(ctx, req.RoleIds, req.UserId)
			liberr.ErrIsNil(ctx, err, "????????????????????????")
			err = s.AddUserPost(ctx, tx, req.PostIds, req.UserId)
			liberr.ErrIsNil(ctx, err)
		})
		return err
	})
	return
}

// AddUserPost ????????????????????????
func (s *userImpl) AddUserPost(ctx context.Context, tx *gdb.TX, postIds []int64, userId int64) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		//?????????????????????
		_, err = dao.SysUserPost.Ctx(ctx).TX(tx).Where(dao.SysUserPost.Columns().UserId, userId).Delete()
		liberr.ErrIsNil(ctx, err, "????????????????????????")
		if len(postIds) == 0 {
			return
		}
		//????????????????????????
		data := g.List{}
		for _, v := range postIds {
			data = append(data, g.Map{
				dao.SysUserPost.Columns().UserId: userId,
				dao.SysUserPost.Columns().PostId: v,
			})
		}
		_, err = dao.SysUserPost.Ctx(ctx).TX(tx).Data(data).Insert()
		liberr.ErrIsNil(ctx, err, "????????????????????????")
	})
	return
}

// AddUserRole ????????????????????????
func (s *userImpl) addUserRole(ctx context.Context, roleIds []int64, userId int64) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		enforcer, e := commonService.CasbinEnforcer(ctx)
		liberr.ErrIsNil(ctx, e)
		for _, v := range roleIds {
			_, e = enforcer.AddGroupingPolicy(fmt.Sprintf("%s%d", s.CasBinUserPrefix, userId), gconv.String(v))
			liberr.ErrIsNil(ctx, e)
		}
	})
	return
}

// EditUserRole ????????????????????????
func (s *userImpl) EditUserRole(ctx context.Context, roleIds []int64, userId int64) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		enforcer, e := commonService.CasbinEnforcer(ctx)
		liberr.ErrIsNil(ctx, e)

		//???????????????????????????
		enforcer.RemoveFilteredGroupingPolicy(0, fmt.Sprintf("%s%d", s.CasBinUserPrefix, userId))
		for _, v := range roleIds {
			_, err = enforcer.AddGroupingPolicy(fmt.Sprintf("%s%d", s.CasBinUserPrefix, userId), gconv.String(v))
			liberr.ErrIsNil(ctx, err)
		}
	})
	return
}

func (s *userImpl) userNameOrMobileExists(ctx context.Context, userName, mobile string, id ...int64) error {
	user := (*entity.SysUser)(nil)
	err := g.Try(ctx, func(ctx context.Context) {
		m := dao.SysUser.Ctx(ctx)
		if len(id) > 0 {
			m = m.Where(dao.SysUser.Columns().Id+" != ", id)
		}
		m = m.Where(fmt.Sprintf("%s='%s' OR %s='%s'",
			dao.SysUser.Columns().UserName,
			userName,
			dao.SysUser.Columns().Mobile,
			mobile))
		err := m.Limit(1).Scan(&user)
		liberr.ErrIsNil(ctx, err, "????????????????????????")
		if user == nil {
			return
		}
		if user.UserName == userName {
			liberr.ErrIsNil(ctx, gerror.New("??????????????????"))
		}
		if user.Mobile == mobile {
			liberr.ErrIsNil(ctx, gerror.New("??????????????????"))
		}
	})
	return err
}

// GetEditUser ????????????????????????
func (s *userImpl) GetEditUser(ctx context.Context, id uint64) (res *system.UserGetEditRes, err error) {
	res = new(system.UserGetEditRes)
	err = g.Try(ctx, func(ctx context.Context) {
		//??????????????????
		res.User, err = s.GetUserInfoById(ctx, id)
		liberr.ErrIsNil(ctx, err)
		//??????????????????????????????
		res.CheckedRoleIds, err = s.GetAdminRoleIds(ctx, id)
		liberr.ErrIsNil(ctx, err)
		res.CheckedPosts, err = s.GetUserPostIds(ctx, id)
		liberr.ErrIsNil(ctx, err)
	})
	return
}

// GetUserInfoById ??????Id??????????????????
func (s *userImpl) GetUserInfoById(ctx context.Context, id uint64, withPwd ...bool) (user *entity.SysUser, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		if len(withPwd) > 0 && withPwd[0] {
			//??????????????????
			err = dao.SysUser.Ctx(ctx).Where(dao.SysUser.Columns().Id, id).Scan(&user)
		} else {
			//??????????????????
			err = dao.SysUser.Ctx(ctx).Where(dao.SysUser.Columns().Id, id).
				FieldsEx(dao.SysUser.Columns().UserPassword, dao.SysUser.Columns().UserSalt).Scan(&user)
		}
		liberr.ErrIsNil(ctx, err, "????????????????????????")
	})
	return
}

// GetUserPostIds ??????????????????
func (s *userImpl) GetUserPostIds(ctx context.Context, userId uint64) (postIds []int64, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		var list []*entity.SysUserPost
		err = dao.SysUserPost.Ctx(ctx).Where(dao.SysUserPost.Columns().UserId, userId).Scan(&list)
		liberr.ErrIsNil(ctx, err, "??????????????????????????????")
		postIds = make([]int64, 0)
		for _, entity := range list {
			postIds = append(postIds, entity.PostId)
		}
	})
	return
}

// ResetUserPwd ??????????????????
func (s *userImpl) ResetUserPwd(ctx context.Context, req *system.UserResetPwdReq) (err error) {
	salt := grand.S(10)
	password := libUtils.EncryptPassword(req.Password, salt)
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.SysUser.Ctx(ctx).WherePri(req.Id).Update(g.Map{
			dao.SysUser.Columns().UserSalt:     salt,
			dao.SysUser.Columns().UserPassword: password,
		})
		liberr.ErrIsNil(ctx, err, "????????????????????????")
	})
	return
}

func (s *userImpl) ChangeUserStatus(ctx context.Context, req *system.UserStatusReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.SysUser.Ctx(ctx).WherePri(req.Id).Update(do.SysUser{UserStatus: req.UserStatus})
		liberr.ErrIsNil(ctx, err, "????????????????????????")
	})
	return
}

// Delete ????????????
func (s *userImpl) Delete(ctx context.Context, ids []int) (err error) {
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			_, err = dao.SysUser.Ctx(ctx).TX(tx).Where(dao.SysUser.Columns().Id+" in(?)", ids).Delete()
			liberr.ErrIsNil(ctx, err, "??????????????????")
			//??????????????????
			enforcer, e := commonService.CasbinEnforcer(ctx)
			liberr.ErrIsNil(ctx, e)
			for _, v := range ids {
				enforcer.RemoveFilteredGroupingPolicy(0, fmt.Sprintf("%s%d", s.CasBinUserPrefix, v))
			}
			//???????????????????????????
			_, err = dao.SysUserPost.Ctx(ctx).TX(tx).Delete(dao.SysUserPost.Columns().UserId+" in (?)", ids)
			liberr.ErrIsNil(ctx, err, "???????????????????????????")
		})
		return err
	})
	return
}
