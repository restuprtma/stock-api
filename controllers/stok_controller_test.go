package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"stock-api/config"
	"stock-api/models"
	"stock-api/tests"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateStok_Success(t *testing.T) {
	// Setup
	tests.SetupTestDB()
	router := tests.SetupTestRouter()

	reqBody := models.CreateStokRequest{
		Tanggal:       "2026-07-10",
		JumlahStokTon: 1500.5,
		LokasiGudang:  "Gudang Test",
		Catatan:       "Catatan Test",
	}
	jsonValue, _ := json.Marshal(reqBody)

	// Action
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/stok", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusCreated, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, float64(1500.5), response["jumlah_stok_ton"])
	assert.Equal(t, "Gudang Test", response["lokasi_gudang"])
}

func TestCreateStok_InvalidPayload(t *testing.T) {
	// Setup
	tests.SetupTestDB()
	router := tests.SetupTestRouter()

	// Field tanggal salah (format bukan YYYY-MM-DD)
	reqBody := map[string]interface{}{
		"tanggal":         "10-07-2026",
		"jumlah_stok_ton": 1500.5,
		"lokasi_gudang":   "Gudang Test",
	}
	jsonValue, _ := json.Marshal(reqBody)

	// Action
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/stok", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetStok_Success(t *testing.T) {
	// Setup
	tests.SetupTestDB()
	router := tests.SetupTestRouter()

	// Insert data dummy secara manual ke dalam memory db
	config.DB.Create(&models.StokBeras{
		Tanggal:       time.Now(),
		JumlahStokTon: 120.0,
		LokasiGudang:  "Gudang A",
		Catatan:       "Test",
	})

	// Action
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/stok", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)

	var response []models.StokBeras
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Len(t, response, 1)
	assert.Equal(t, "Gudang A", response[0].LokasiGudang)
}
