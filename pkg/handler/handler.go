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

	router.POST("/registration", h.signUp)

	accounts := router.Group("/accounts")
	{
		//Получение информации об аккаунте пользователя
		accounts.GET("/:accountId", h.accountGetById)
		//Поиск аккаунтов пользователей по параметрам
		accounts.GET("/search", h.searchAccount)
		//Обновление данных аккаунта пользователя
		accounts.PUT(":id", h.updateAccount)
		//Удаление аккаунта пользователя
		accounts.DELETE(":id", h.removeAccount)
	}

	animals := router.Group("/animals")
	{
		//Получение информации о животном
		animals.GET("/:animalId", h.animalGetById)
		//Поиск животных по параметрам
		animals.GET("/search", h.searchAnimal)
		//Добавление нового животного
		animals.POST("", h.createAnimal)
		//Обновление информации о животном
		animals.PUT("/:animalId", h.updateAnimal)
		//Удаление животного
		animals.DELETE("/:animalId", h.removeAnimal)

		//Добавление типа конкретному животному
		animals.POST(":animalId/types/:typeId", h.addTypeToAnimal)
		//Изменить тип у конкретного животного
		animals.PUT("/:animalId/types", h.changeAnimalType)
		//Удаление типа у конкретного животного
		animals.DELETE("/:animalId/types/:typeId", h.removeTypeFromAnimal)

		//Просмотр локаций, посещенных животным
		animals.GET("/:animalId/locations", h.getVisitedLocations)
		//Добавление локации, посещенной животным
		animals.POST("/:animalId/locations/:pointId", h.createVisitedLocation)
		//Изменение локации, посещенной животным
		animals.PUT("/:animalId/locations", h.changeVisitedLocation)
		//Удаление локации, посещенной животным
		animals.DELETE("/:animalId/locations/:visitedPointId", h.removeVisitedLocation)

		types := animals.Group("/types")
		{
			//Получение информации о типе животного
			types.GET("/:typeId", h.getAnimalType)
			//Добавление типа животного
			types.POST("", h.createAnimalType)
			//Изменение типа животного
			types.PUT("/:typeId", h.modifyAnimalType)
			//Удаление типа животнгого
			types.DELETE("/:typeId", h.removeAnimalType)
		}
	}

	locations := router.Group("/locations")
	{
		//Получение информации о точке локации животных
		locations.GET("/:pointId", h.getLocationById)
		//Добавление точки локации животных
		locations.POST("", h.createLocation)
		//Изменение точки локации животных
		locations.PUT("/:pointId", h.modifyLocation)
		//Удаление точки локации животных
		locations.DELETE("/:pointId", h.removeLocation)
	}
	return router
}
