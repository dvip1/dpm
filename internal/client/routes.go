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
	Version  string
	BaseUrl string
	Auth     authRoutes
	Packages packageRoutes
}

// AppRoutes is a public, pre-configured variable that you can import and use anywhere.
// This is the single source of truth for all your backend routes.
var AppRoutes = routes{
	Root:    "/",
	Version: "v1",
	BaseUrl: "http://localhost:8000",
	Auth: authRoutes{
		Root:     "/auth",
		Login:    "/auth/login",
		Logout:   "/auth/logout",
		Register: "/auth/register",
		Me:       "/auth/me",
		Refresh:  "/auth/token/refresh",
	},
	Packages: packageRoutes{
		Root: "api/v1/packages",
	},
}
