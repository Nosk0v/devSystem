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
	id, err := strconv.Atoi(c.Param("id"))
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
	id, err := strconv.Atoi(c.Param("id"))
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
	id, err := strconv.Atoi(c.Param("id"))
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
