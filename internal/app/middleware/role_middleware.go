package middleware

import (
	"developer-orientenergy-golang/internal/app/factory"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"

	. "developer-orientenergy-golang/internal/app/init_global"
	router "developer-orientenergy-golang/internal/app/router/router_defined"
	. "developer-orientenergy-golang/internal/pkg/util"
)

func RoleMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlerName := "UNKNOWN"
		if route := mux.CurrentRoute(r); route != nil {
			routeName := route.GetName()
			if routeName != "" {
				handlerName = routeName
			}
		}
		var isPublicApi = router.IsPublicAPI(handlerName)
		if isPublicApi {
			next.ServeHTTP(w, r)
			return
		}
		var authService = factory.GetAuthService()
		var currentUser, err = strconv.Atoi(context.Get(r, "userInfo").(string))
		if err != nil {
			return
		}
		var userInfo, _ = authService.GetUserInfoByID(currentUser)
		if userInfo.UserAccountId == 0 {
			RespondJSONError(w, http.StatusForbidden, "UnAuthenticate")
			return
		}
		if userInfo.IsActive != true {
			RespondJSONError(w, http.StatusForbidden, "UnAuthenticate")
			return
		}
		var permission = make(map[string][]string, 0)
		if userInfo.Permissions != nil {
			_ = json.Unmarshal(userInfo.Permissions, &permission)
		}
		var currentRoute = mux.CurrentRoute(r)
		for _, val := range RoutersArr {
			for _, valR := range val.Data {
				if valR.Handler == currentRoute.GetName() {
					var key = valR.Key
					if key == "" {
						next.ServeHTTP(w, r)
						return
					}
					if permission[key] != nil {
						var yourPermissions = permission[key]
						for _, yourPermission := range yourPermissions {
							var routerMethods = valR.Permission.(map[string]interface{})
							for method, routerMethod := range routerMethods {
								if r.Method == method {
									for _, routerPermission := range routerMethod.([]interface{}) {
										if routerPermission == yourPermission {
											next.ServeHTTP(w, r)
											return
										}
									}
								}
							}
						}
					}
				}
			}
		}
		RespondJSONError(w, http.StatusForbidden, "Access denied")
		return
	})
}
