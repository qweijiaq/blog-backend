package flag

import "backend/models"

func EsCreateIndex() {
	models.ArticleModel{}.CreateIndex()
	models.FullTextModel{}.CreateIndex()
}
