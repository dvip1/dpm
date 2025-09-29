package client

// Define structs to hold the route paths for different API groups.
// The struct fields (Login, Logout, etc.) are capitalized to be exported.
type authRoutes struct {
	Root           string
	Login          string
	Logout         string
	Register       string
	Me             string
	Refresh        string
	ChangePassword string
}

type packageRoutes struct {
	Root string
}

// The main 'routes' struct that holds all the nested route groups.
type routes struct {
	Root     string
	Auth     authRoutes
	Packages packageRoutes
}

// AppRoutes is a public, pre-configured variable that you can import and use anywhere.
// This is the single source of truth for all your backend routes.
var AppRoutes = routes{
	Root: "/",
	Auth: authRoutes{
		Root:           "/api/auth",
		Login:          "/api/auth/login",
		Logout:         "/api/auth/logout",
		Register:       "/api/auth/register",
		Me:             "/api/auth/me",
		Refresh:        "/api/auth/token/refresh",
		ChangePassword: "/api/auth/change-password",
	},
	Packages: packageRoutes{
		Root: "/api/packages",
	},
}
