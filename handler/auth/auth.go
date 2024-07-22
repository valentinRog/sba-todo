package auth

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/valentinRog/sba-todo/auth"
	"github.com/valentinRog/sba-todo/store"
	userstore "github.com/valentinRog/sba-todo/store/user"
	"net/http"
)

type Handlers struct {
	q   *store.Queries
	ctx context.Context
}

func New(ctx context.Context, q *store.Queries) *Handlers {
	return &Handlers{
		ctx: ctx,
		q:   q,
	}
}

func (h *Handlers) PostSignup(c echo.Context) error {
	name := c.FormValue("username")
	password := c.FormValue("password")
	userID, _ := h.q.User.CreateUser(h.ctx, userstore.CreateUserParams{Username: name, Password: password})
	token := auth.CreateSession(userID)
	cookie := auth.CreateCookie(token)
	c.SetCookie(cookie)
	return c.Redirect(http.StatusMovedPermanently, "/")
}

func (h *Handlers) PostSignin(c echo.Context) error {
	name := c.FormValue("username")
	// password := c.FormValue("password")
	user, err := h.q.User.GetUserByName(h.ctx, name)
	if err != nil {
		return c.Redirect(http.StatusMovedPermanently, "/")
	}
	token := auth.CreateSession(user.ID)
	cookie := auth.CreateCookie(token)
	c.SetCookie(cookie)
	return c.Redirect(http.StatusMovedPermanently, "/")
}

func (h *Handlers) PostLogout(c echo.Context) error {
	c.SetCookie(auth.DeleteCookie())
	return c.Redirect(http.StatusMovedPermanently, "/login")
}
