package readers

import "github.com/rmay1er/jgeo-excel/internal/models"

// Reader определяет интерфейс для чтения координат из различных источников
type Reader interface {
	// Read читает координаты из источника
	Read() (*[]models.CordsData, error)
	// Close закрывает соединение с источником
	Close() error
}
