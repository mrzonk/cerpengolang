package CerpenController

import (
	"cerpengolang/config"
	"cerpengolang/helper"
	"cerpengolang/models"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func Index(w http.ResponseWriter, r *http.Request) {
	var cerpen []models.Cerpen

	if err := config.DB.Find(&cerpen).Error; err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	helper.Response(w, 200, "Data Cerpen's", cerpen)
}

func Create(w http.ResponseWriter, r *http.Request) {
	var cerpen models.Cerpen

	if err := json.NewDecoder(r.Body).Decode(&cerpen); err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	defer r.Body.Close()

	if err := config.DB.Create(&cerpen).Error; err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}
	helper.Response(w, 201, "Success create cerpen", nil)

}

func Detail(w http.ResponseWriter, r *http.Request) {
	idParams := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idParams)

	var cerpen models.Cerpen

	if err := config.DB.First(&cerpen, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.Response(w, 404, "Cerpen Not Found", nil)
			return
		}

		helper.Response(w, 500, err.Error(), nil)
		return
	}

	helper.Response(w, 200, "Detail cerpen", cerpen)
}

func Update(w http.ResponseWriter, r *http.Request) {
	idParams := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idParams)

	var cerpen models.Cerpen

	if err := config.DB.First(&cerpen, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.Response(w, 404, "cerpen Not Found", nil)
			return
		}

		helper.Response(w, 500, err.Error(), nil)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&cerpen); err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	defer r.Body.Close()

	if err := config.DB.Where("id = ?", id).Updates(&cerpen).Error; err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	helper.Response(w, 201, "Success Update cerpen", nil)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(params)

	var cerpen models.Cerpen
	res := config.DB.Delete(&cerpen, id)

	if res.Error != nil {
		helper.Response(w, 500, res.Error.Error(), nil)
		return
	}

	if res.RowsAffected == 0 {
		helper.Response(w, 404, "cerpen Not Found", nil)
		return
	}
	helper.Response(w, 200, "Success Delete cerpen", nil)
}
