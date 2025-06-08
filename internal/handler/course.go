package handler

import (
	"devSystem/internal/usecase"
	"devSystem/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
)

// GetCourse godoc
// @Summary Получить курс по ID
// @Description Получение сведений о курсе по его ID.
// @Tags courses
// @Accept json
// @Produce json
// @Param id path int true "ID курса"
// @Success 200 {object} models.CourseResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /courses/{id} [get]
func (h *Handler) getCourse(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("course_id"))
	if err != nil {
		log.Printf("Invalid course ID from request: %v", c.Param("id"))
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid course ID"})
		return
	}

	course, err := h.usecases.GetCourse(id)
	if err != nil {
		log.Printf("Error fetching course with ID %d: %v", id, err)
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Error fetching course"})
		return
	}

	if course == nil {
		log.Printf("Course with ID %d not found", id)
		c.JSON(http.StatusNotFound, ErrorResponse{Error: "Course not found"})
		return
	}

	log.Printf("Returning course with ID %d: %+v", id, course)
	c.JSON(http.StatusOK, course)
}

// GetAllCourses godoc
// @Summary Получить все курсы
// @Description Получение списка курсов с учётом роли и организации.
// @Tags courses
// @Accept json
// @Produce json
// @Success 200 {array} models.CourseResponse
// @Failure 401 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /courses [get]
func (h *Handler) getAllCourses(c *gin.Context) {
	log.Println("Fetching all courses request received")

	claims, err := h.GetJWTClaims(c)
	if err != nil {
		log.Printf("Unauthorized: %v", err)
		c.JSON(http.StatusUnauthorized, ErrorResponse{Error: "Unauthorized"})
		return
	}

	courses, err := h.usecases.GetCoursesByClaims(claims)
	if err != nil {
		log.Printf("Error fetching courses: %v", err)
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Error fetching courses"})
		return
	}

	log.Printf("Returning %d courses for role=%d, orgID=%v", len(courses), claims.Role, claims.OrganizationID)
	c.JSON(http.StatusOK, courses)
}

// CreateCourse godoc
// @Summary Создать курс
// @Description Создание нового курса по входным данным с указанием материалов и компетенций.
// @Tags courses
// @Accept json
// @Produce json
// @Param course body models.Course true "Входные данные"
// @Success 201 {object} models.CourseResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /courses [post]
func (h *Handler) createCourse(c *gin.Context) {
	var input models.Course
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid input"})
		return
	}

	claims, err := h.GetJWTClaims(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, ErrorResponse{Error: "Unauthorized"})
		return
	}

	input.OrganizationID = *claims.OrganizationID
	input.CreatedBy = claims.Email
	if input.CreateDate.IsZero() {
		input.CreateDate = time.Now().UTC()
	}

	courseID, err := h.usecases.CreateCourse(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Error creating course"})
		return
	}

	course, err := h.usecases.GetCourse(courseID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Error fetching course"})
		return
	}

	c.JSON(http.StatusCreated, course)
}

// UpdateCourse godoc
// @Summary Обновить курс
// @Description Обновление курса по его ID.
// @Tags courses
// @Accept json
// @Produce json
// @Param id path int true "ID курса"
// @Param course body models.Course true "Входные данные"
// @Success 204 {object} nil
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /courses/{id} [put]
func (h *Handler) updateCourse(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("course_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid course ID"})
		return
	}

	var input models.Course
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid input"})
		return
	}
	input.CourseID = id

	claims, err := h.GetJWTClaims(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, ErrorResponse{Error: "Unauthorized"})
		return
	}

	existingCourse, err := h.usecases.GetCourse(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Failed to fetch existing course"})
		return
	}
	if existingCourse == nil {
		c.JSON(http.StatusNotFound, ErrorResponse{Error: "Course not found"})
		return
	}

	if claims.Role != usecase.RoleSuperAdmin {
		if claims.OrganizationID == nil || *claims.OrganizationID != existingCourse.OrganizationID {
			c.JSON(http.StatusForbidden, ErrorResponse{Error: "Access denied"})
			return
		}
	}

	input.OrganizationID = existingCourse.OrganizationID

	err = h.usecases.UpdateCourse(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Error updating course"})
		return
	}

	c.Status(http.StatusOK)
}

// DeleteCourse godoc
// @Summary Удалить курс
// @Description Удаление курса по его ID.
// @Tags courses
// @Accept json
// @Produce json
// @Param id path int true "ID курса"
// @Success 204 {object} nil
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /courses/{id} [delete]
func (h *Handler) deleteCourse(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("course_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid course ID"})
		return
	}

	claims, err := h.GetJWTClaims(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, ErrorResponse{Error: "Unauthorized"})
		return
	}

	course, err := h.usecases.GetCourse(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Failed to fetch course"})
		return
	}
	if course == nil {
		c.JSON(http.StatusNotFound, ErrorResponse{Error: "Course not found"})
		return
	}

	if claims.Role != usecase.RoleSuperAdmin {
		if claims.OrganizationID == nil || *claims.OrganizationID != course.OrganizationID {
			c.JSON(http.StatusForbidden, ErrorResponse{Error: "Access denied"})
			return
		}
	}

	if err := h.usecases.DeleteCourse(id); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Error deleting course"})
		return
	}

	c.Status(http.StatusNoContent)
}

// CompleteCourse godoc
// @Summary Завершить курс
// @Description Отмечает курс как завершённый для пользователя
// @Tags courses
// @Accept json
// @Produce json
// @Param id path int true "ID курса"
// @Success 200 {string} string "Course marked as completed"
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /courses/{id}/complete [post]
func (h *Handler) completeCourse(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("course_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid course ID"})
		return
	}

	claims, err := h.GetJWTClaims(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, ErrorResponse{Error: "Unauthorized"})
		return
	}

	if err := h.usecases.CompleteCourse(claims.Email, id); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Failed to mark course as completed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Course marked as completed"})
}

// MarkMaterialAsCompleted godoc
// @Summary Отметить материал как пройденный
// @Description Отмечает материал в курсе как пройденный пользователем
// @Tags courses
// @Accept json
// @Produce json
// @Param course_id path int true "ID курса"
// @Param material_id path int true "ID материала"
// @Success 200 {string} string "Material marked as completed"
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /courses/{course_id}/materials/{material_id}/complete [post]
func (h *Handler) markMaterialAsCompleted(c *gin.Context) {
	courseID, err := strconv.Atoi(c.Param("course_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid course ID"})
		return
	}

	materialID, err := strconv.Atoi(c.Param("material_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid material ID"})
		return
	}

	claims, err := h.GetJWTClaims(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, ErrorResponse{Error: "Unauthorized"})
		return
	}

	if err := h.usecases.MarkMaterialAsCompleted(claims.Email, courseID, materialID); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Failed to mark material as completed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Material marked as completed"})
}

// GetCourseProgress godoc
// @Summary Получить прогресс курса
// @Description Получение списка ID завершённых материалов по курсу
// @Tags courses
// @Accept json
// @Produce json
// @Param id path int true "ID курса"
// @Success 200 {object} models.CourseProgressResponse
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /courses/{id}/progress [get]
func (h *Handler) getCourseProgress(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("course_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid course ID"})
		return
	}

	claims, err := h.GetJWTClaims(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, ErrorResponse{Error: "Unauthorized"})
		return
	}

	progress, err := h.usecases.GetUserCourseProgress(claims.Email, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Failed to fetch course progress"})
		return
	}

	c.JSON(http.StatusOK, models.CourseProgressResponse{CompletedMaterials: progress})
}

// IsCourseCompleted godoc
// @Summary Проверка завершения курса
// @Description Проверяет, завершён ли курс пользователем
// @Tags courses
// @Accept json
// @Produce json
// @Param id path int true "ID курса"
// @Success 200 {object} models.CourseProgressResponse
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /courses/{id}/completed [get]
func (h *Handler) isCourseCompleted(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("course_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid course ID"})
		return
	}

	claims, err := h.GetJWTClaims(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, ErrorResponse{Error: "Unauthorized"})
		return
	}

	isCompleted, err := h.usecases.IsCourseCompleted(claims.Email, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Failed to check course completion"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"completed": isCompleted})
}

// GetCompletedCourses godoc
// @Summary Получить завершённые курсы
// @Description Возвращает список завершённых курсов для текущего пользователя
// @Tags courses
// @Accept json
// @Produce json
// @Success 200 {array} models.CourseResponse
// @Failure 401 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /courses/completed [get]
func (h *Handler) getCompletedCourses(c *gin.Context) {
	claims, err := h.GetJWTClaims(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, ErrorResponse{Error: "Unauthorized"})
		return
	}

	courses, err := h.usecases.GetCompletedCourses(claims.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Failed to fetch completed courses"})
		return
	}

	c.JSON(http.StatusOK, courses)
}

// GetOrganizationCourseProgress godoc
// @Summary Прогресс всех пользователей организации
// @Description Получает список пользователей и их завершённые курсы по организации
// @Tags courses
// @Accept json
// @Produce json
// @Param org_id query int false "ID организации (только для SuperAdmin)"
// @Success 200 {array} models.UserCourseProgress
// @Failure 401 {object} ErrorResponse
// @Failure 403 {object} ErrorResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /courses/progress/organization [get]
func (h *Handler) getOrganizationCourseProgress(c *gin.Context) {
	claims, err := h.GetJWTClaims(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, ErrorResponse{Error: "Unauthorized"})
		return
	}

	var orgID int

	switch claims.Role {
	case usecase.RoleAdmin:
		if claims.OrganizationID == nil || *claims.OrganizationID <= 0 {
			c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Missing organization ID in token"})
			return
		}
		orgID = *claims.OrganizationID
	case usecase.RoleSuperAdmin:
		rawOrgID := c.Query("org_id")
		if rawOrgID == "" {
			c.JSON(http.StatusBadRequest, ErrorResponse{Error: "org_id query parameter is required"})
			return
		}
		orgID, err = strconv.Atoi(rawOrgID)
		if err != nil || orgID <= 0 {
			c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid org_id"})
			return
		}
	default:
		c.JSON(http.StatusForbidden, ErrorResponse{Error: "Forbidden"})
		return
	}

	result, err := h.usecases.GetOrganizationCourseProgress(orgID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Failed to fetch organization course progress"})
		return
	}

	c.JSON(http.StatusOK, result)
}
