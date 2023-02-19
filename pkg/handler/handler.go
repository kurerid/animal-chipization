package handler

import (
	"animal-chipization/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	router.POST("/registration", h.createAccount)

	accounts := router.Group("/accounts")
	{
		//Получение информации об аккаунте пользователя
		accounts.GET("/:id")
		//Поиск аккаунтов пользователей по параметрам
		accounts.GET("/search")
		//Обновление данных аккаунта пользователя
		accounts.PUT(":id")
		//Удаление аккаунта пользователя
		accounts.DELETE(":id")
	}

	animals := router.Group("/animals")
	{
		//Получение информации о животном
		animals.GET("/:animalId")
		//Поиск животных по параметрам
		animals.GET("/search")
		//Добавление нового животного
		animals.POST("")
		//Обновление информации о животном
		animals.PUT("/:animalId")
		//Удаление животного
		animals.DELETE("/:animalId")

		//Добавление типа конкретному животному
		animals.POST(":animalId/types/:typeId")
		//Изменить тип у конкретного животного
		animals.PUT("/:animalId/types")
		//Удаление типа у конкретного животного
		animals.DELETE("/:animalId/types/:typeId")

		//Просмотр локаций, посещенных животным
		animals.GET("/:animalId/locations")
		//Добавление локации, посещенной животным
		animals.POST("/:animalId/locations/:pointId")
		//Изменение локации, посещенной животным
		animals.PUT("/:animalId/locations")
		//Удаление локации, посещенной животным
		animals.DELETE("/:animalId/locations/:visitedPointId")

		types := animals.Group("/types")
		{
			//Получение информации о типе животного
			types.GET("/:typeId")
			//Добавление типа животного
			types.POST("")
			//Изменение типа животного
			types.PUT("/:typeId")
			//Удаление типа животнгого
			types.DELETE("/:typeId")
		}
	}

	locations := router.Group("/locations")
	{
		//Получение информации о точке локации животных
		locations.GET("/:pointId")
		//Добавление точки локации животных
		locations.POST("")
		//Изменение точки локации животных
		locations.PUT("/:pointId")
		//Удаление точки локации животных
		locations.DELETE("/:pointId")
	}
	return router
}
