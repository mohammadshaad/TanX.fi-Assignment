package services

import (
    "price-alert-app/cmd/internal/models"
    "price-alert-app/cmd/internal/repository"
)

func CreateAlert(alert models.Alert) error {
    return repository.CreateAlert(alert)
}

func DeleteAlert(id string) error {
    return repository.DeleteAlert(id)
}

func FetchAlerts(status string) ([]models.Alert, error) {
    return repository.FetchAlerts(status)
}
