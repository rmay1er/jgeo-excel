/*
Copyright © 2025 Ruslan Mayer
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd представляет базовую команду приложения
var rootCmd = &cobra.Command{
	Use:   "jgeo-excel",
	Short: "Работа с Excel и GeoJSON",
	Long: `jgeo-excel - это CLI инструмент для работы с координатами и файлами Excel.

Приложение позволяет:
  • Читать координаты и данные из Excel файла
  • Добавлять точки в существующий GeoJSON файл (FeatureCollection)
  • Конфигурировать столбцы Excel и параметры через YAML файл
  • Поддерживает несколько листов в Excel

Использование:
  jgeo-excel to-geojson --config config.yaml

Для примера конфигурационного файла смотрите config.example.yaml`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
		}
	},
}

// Execute добавляет все подкоманды к корневой команде и устанавливает флаги.
// Это вызывается из main.main() один раз.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка: %v\n", err)
		os.Exit(1)
	}
}

func init() {
	// Здесь определяются глобальные флаги и конфигурация.
	// Cobra поддерживает persistent флаги, которые, если определены здесь,
	// будут глобальными для всего приложения.
}
