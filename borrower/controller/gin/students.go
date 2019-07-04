package gin

import (
	"database/sql"
	"errors"
	"log"
	"net/http"

	mysql "github.com/Mictrlan/library/students/model/mysql"

	"github.com/gin-gonic/gin"
)

var (
	errInvalidServer = errors.New("[RegisterRouter]: server is nil")
)

// StudnetsController -
type StudnetsController struct {
	db *sql.DB
}

// New create new StudentsController
func New(db *sql.DB) *StudnetsController {
	return &StudnetsController{
		db: db,
	}
}

// Register -
func (sc *StudnetsController) Register(r gin.IRouter) error {
	if r == nil {
		log.Fatal(errInvalidServer)
	}

	err := mysql.StudentsCreateTable(sc.db)
	if err != nil {
		log.Fatal(err)
	}

	r.POST("api/v2/students/add")
	r.POST("api/v2/students/modifypassword")
	r.POST("api/v2/students/delete")
	r.POST("api/v2/students/")
	r.POST("api/v2/students/")
	r.POST("api/v2/students/")
	r.POST("api/v2/students/")
	r.POST("api/v2/students/")

	return nil
}

func (sc *StudnetsController) studentsadd(ctx *gin.Context) {
	var (
		req struct {
			Name       string `json:"name"`
			Pwd        string `json:"pwd"`
			Mobile     string `json:"mobile"`
			Department string `json:"department"`
		}
	)

	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	id, err := mysql.StudnetsAddInfo(sc.db, req.Name, req.Pwd, req.Mobile, req.Department)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusPreconditionFailed, gin.H{"status": http.StatusPreconditionFailed})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"id":     id,
	})
}
