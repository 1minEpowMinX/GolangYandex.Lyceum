package main

import (
	"fmt"
	"net/http"
	"slices"
)

func RoleBasedAuthMiddleware(allowedRoles []string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user_role := r.Header.Get("X-User-Role")
		if !slices.Contains(allowedRoles, user_role) {
			http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

func AdminHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Admin Resource")
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "User Resource")
}

func main() {
	allowedAdminRoles := []string{"admin", "superadmin"}
	allowedUserRoles := []string{"user"}

	// Создание маршрута и применение Middleware для пути "/admin".
	adminHandler := RoleBasedAuthMiddleware(allowedAdminRoles, http.HandlerFunc(AdminHandler))
	http.Handle("/admin", adminHandler)

	// Создание маршрута и применение Middleware для пути "/user".
	userHandler := RoleBasedAuthMiddleware(allowedUserRoles, http.HandlerFunc(UserHandler))
	http.Handle("/user", userHandler)

	// Запуск веб-сервера на порту 8080.
	http.ListenAndServe(":8080", nil)
}
