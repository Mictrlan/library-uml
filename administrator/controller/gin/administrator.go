package gin

import (
	"database/sql"
	"errors"
	"log"
	"net/http"

	adminmysql "github.com/Mictrlan/library/administrator/model/mysql"
	"github.com/gin-gonic/gin"
)

var (
	errInvalidServer = errors.New("[RegisterRouter]: server is nil")
)

// AdministratorController -
type AdministratorController struct {
	db *sql.DB
}

// New create new AdministratorController
func New(db *sql.DB) *AdministratorController {
	return &AdministratorController{
		db: db,
	}
}

// Register register administrator router
func (ac *AdministratorController) Register(r gin.IRouter) error {

	if r == nil {
		log.Fatal(errInvalidServer)
	}

	err := adminmysql.AdministratorCreatetable
	if err != nil {
		log.Fatal(err)
	}

	r.POST("api/v1/admin/addNewAdmin", ac.adminAddNew)

	return nil
}

func (ac *AdministratorController) adminAddNew(ctx *gin.Context) {
	var (
		req struct {
			Name string `json: name`
			Pwd  string `json:pwd`
		}
	)

	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadGateway})
		return
	}

	err = adminmysql.AdministratorAdd(ac.db, req.Name, req.Pwd)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusPreconditionFailed, gin.H{"status": http.StatusPreconditionFailed})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
}

func (ac *AdministratorController) adminModifyPwd(ctx *gin.Context) {
	var (
		req struct {
			Name     string `json : name`
			Password string `json: pwd`
		}
	)

	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadGateway})
		return
	}

	err = adminmysql.AdministratorModify(ac.db, req.Name, req.Password)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusPreconditionFailed, gin.H{"status": http.StatusPreconditionFailed})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
}
