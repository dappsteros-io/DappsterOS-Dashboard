package controllers

import (
	"github.com/dappster-io/DappsterOS-Dashboard/backend/src/models"
	"github.com/kataras/iris/v12"
)

type ProfileController struct {
	/* dependencies */
}

func GetMe(ctx iris.Context) {
	user := ctx.Value("user")
	ctx.JSON(iris.Map{"user": user})
}

// GET: /api/profile
func (c *ProfileController) Get() ([]models.Profile, error) {
	p := &models.Profile{}
	res, err := p.Get()
	if err != nil {
		return nil, err
	}
	return res, nil
}

// POST: /api/profile
func (c *ProfileController) Post(p models.Profile) (*models.Profile, error) {
	res, err := p.Post()

	if err != nil {
		return nil, err
	}

	return res, nil
}
