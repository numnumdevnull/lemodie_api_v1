package handler

import (
	"database/sql"
	"net/http"
	"strconv"

	"lemodie_api_v1/internal/model"

	"github.com/gin-gonic/gin"
)

type DictionaryTypesHandler struct {
	DB *sql.DB
}

var allowedLimits = map[int]bool{10: true, 20: true, 50: true}

func (h *DictionaryTypesHandler) GetAll(c *gin.Context) {
	// --- парсимо page ---
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "невалідний page"})
		return
	}

	// --- парсимо limit ---
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil || !allowedLimits[limit] {
		c.JSON(http.StatusBadRequest, gin.H{"error": "limit має бути 10, 20 або 50"})
		return
	}

	offset := (page - 1) * limit

	// --- total ---
	var total int
	if err := h.DB.QueryRowContext(c.Request.Context(),
		`SELECT COUNT(*) FROM dictionary_types`,
	).Scan(&total); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "count failed"})
		return
	}

	// --- дані ---
	rows, err := h.DB.QueryContext(c.Request.Context(),
		`SELECT id, value, meta FROM dictionary_types ORDER BY id LIMIT ? OFFSET ?`,
		limit, offset,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "query failed"})
		return
	}
	defer rows.Close()

	items := make([]model.DictionaryTypes, 0)
	for rows.Next() {
		var item model.DictionaryTypes
		if err := rows.Scan(&item.ID, &item.Value, &item.Meta); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "scan failed"})
			return
		}
		items = append(items, item)
	}

	totalPages := (total + limit - 1) / limit

	c.JSON(http.StatusOK, model.PaginatedResponse[model.DictionaryTypes]{
		Data:       items,
		Total:      total,
		Page:       page,
		Limit:      limit,
		TotalPages: totalPages,
	})
}
