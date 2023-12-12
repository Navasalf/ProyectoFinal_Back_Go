package main

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

var db *gorm.DB

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}

func init() {
	var err error
	db, err = gorm.Open("postgres", "user=postgres password=Lar1ss0n dbname=snicky port=5432 sslmode=disable")
	if err != nil {
		panic(err)
	}
	db.LogMode(true)
}

func main() {
	defer db.Close()
	db.AutoMigrate(&User{})

	tableName := "users"
	db.Table(tableName).AutoMigrate(&User{})

	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://127.0.0.1:5173"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	r.Use(cors.New(config))

	r.POST("/crear-usuario", crearUsuario)
	r.POST("/iniciar-sesion", iniciarSesion)
	r.GET("/consultar-usuario/:email", consultarUsuario)
	r.DELETE("/eliminar-usuario/:id", eliminarUsuario)
	r.Run(":8080")
}

func crearUsuario(c *gin.Context) {
	var user User
	c.Bind(&user)
	if user.Username == "" || user.Password == "" || user.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Faltan datos"})
		return
	}
	var existingUser User
	if err := db.Where("email = ?", user.Email).Or("username = ?", user.Username).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "El usuario o email ya existe"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	user.Password = string(hashedPassword)
	db.Create(&user)
	c.JSON(http.StatusOK, gin.H{"message": "Usuario creado exitosamente"})
}

func iniciarSesion(c *gin.Context) {
	var loginRequest struct {
		Password string `json:"password"`
		Email    string `json:"email"`
	}
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user User
	if err := db.Where("email = ?", loginRequest.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no encontrado"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Contraseña incorrecta"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Inicio de sesión exitoso", "user": user})
}

func consultarUsuario(c *gin.Context) {
	email := c.Param("email")
	var user User
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func eliminarUsuario(c *gin.Context) {
	id := c.Param("id")
	if err := db.Where("id = ?", id).Delete(&User{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar usuario"})
		return
	}
	c.JSON(http.StatusOK, fmt.Sprintf("Usuario con ID %s ha sido eliminado", id))
}
