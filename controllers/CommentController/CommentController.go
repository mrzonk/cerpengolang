package CommentController

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
	var comments []models.Comment
	//var commentResponse []models.CommentResponse

	//if err := config.DB.Joins("Cerpen").Find(&comments).Find(&commentResponse).Error; err != nil {
	//if err := config.DB.Joins("Cerpen").Find(&comments).Error; err != nil {
	if err := config.DB.Find(&comments).Error; err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	helper.Response(w, 200, "List Comment", comments)
}

func Create(w http.ResponseWriter, r *http.Request) {
	var comment models.Comment

	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	defer r.Body.Close()

	// Cek Author
	var cerpen models.Cerpen
	if err := config.DB.First(&cerpen, comment.Cerpen_id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.Response(w, 404, "Cerpen Not Found", nil)
			return
		}

		helper.Response(w, 500, err.Error(), nil)
		return
	}

	if err := config.DB.Create(&comment).Error; err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	helper.Response(w, 201, "Success Create comment", nil)
}

func Detail(w http.ResponseWriter, r *http.Request) {
	idParams := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idParams)

	var comment models.Comment
	//var commentResponse models.CommentResponse

	//if err := config.DB.Joins("Cerpen").First(&comment, id).First(&commentResponse).Error; err != nil {
	if err := config.DB.First(&comment, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.Response(w, 404, "comment Not Found", nil)
			return
		}

		helper.Response(w, 500, err.Error(), nil)
		return
	}

	helper.Response(w, 200, "Detail comment", comment)
}

func Update(w http.ResponseWriter, r *http.Request) {
	idParams := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idParams)

	var comment models.Comment

	if err := config.DB.First(&comment, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.Response(w, 404, "Comment Not Found", nil)
			return
		}

		helper.Response(w, 500, err.Error(), nil)
		return
	}

	var commentPayload models.Comment
	if err := json.NewDecoder(r.Body).Decode(&commentPayload); err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	defer r.Body.Close()

	var cerpen models.Cerpen
	if commentPayload.Cerpen_id != 0 {
		if err := config.DB.First(&cerpen, commentPayload.Cerpen_id).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				helper.Response(w, 404, "Property Not Found", nil)
				return
			}

			helper.Response(w, 500, err.Error(), nil)
			return
		}
	}

	if err := config.DB.Where("id = ?", id).Updates(&commentPayload).Error; err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	helper.Response(w, 201, "Success Update Comment", nil)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(params)

	var comment models.Comment
	res := config.DB.Delete(&comment, id)

	if res.Error != nil {
		helper.Response(w, 500, res.Error.Error(), nil)
		return
	}

	if res.RowsAffected == 0 {
		helper.Response(w, 404, "Comment Not Found", nil)
		return
	}
	helper.Response(w, 200, "Success Delete Comment", nil)
}
