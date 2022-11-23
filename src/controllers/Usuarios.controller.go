package controllers

import (
	"github.com/albinovejar/api_templateV1/src/config"
	"github.com/albinovejar/api_templateV1/src/models"
	"github.com/gofiber/fiber/v2"
)

func Usuarios_GetAll(context *fiber.Ctx) error {
	db := config.Connection()
	var usuarios []models.Usuario
	db.Raw("SELECT * FROM Usuarios;").Scan(&usuarios)
	return context.JSON(fiber.Map{
		"data":   usuarios,
		"status": fiber.StatusOK,
	})
}

func Usuarios_GetOne(context *fiber.Ctx) error {
	id, _ := context.ParamsInt("id")
	db := config.Connection()
	var usuario models.Usuario
	db.Raw("SELECT * FROM Usuarios WHERE id = (?);", id).Scan(&usuario)
	return context.JSON(fiber.Map{
		"data":   usuario,
		"status": fiber.StatusOK,
	})
}

func Usuarios_Edit(context *fiber.Ctx) error {
	id, _ := context.ParamsInt("id")
	db := config.Connection()
	var usuario models.Usuario
	context.BodyParser(&usuario)
	db.Raw("UPDATE Usuarios SET nombre = (?), apellidos = (?), direccion = (?), ciudad = (?), estado = (?), cp = (?), FechaNacimiento = (?) WHERE id = (?)",
		usuario.Nombre,
		usuario.Apellidos,
		usuario.Direccion,
		usuario.Ciudad,
		usuario.Estado,
		usuario.Cp,
		usuario.FechaNacimiento,
		id).Scan(&usuario.Id)
	return context.JSON(fiber.Map{
		"data":   usuario,
		"status": fiber.StatusOK,
	})
}

func Usuarios_Delete(context *fiber.Ctx) error {
	id, _ := context.ParamsInt("id")
	db := config.Connection()
	db.Exec("DELETE FROM Usuarios WHERE id = (?);", id)
	return context.JSON(fiber.Map{
		"status": fiber.StatusOK,
	})
}

func Usuarios_Add(context *fiber.Ctx) error {
	db := config.Connection()
	var usuario models.Usuario
	context.BodyParser(&usuario)
	db.Exec("INSERT INTO Usuarios(nombre, apellidos, direccion, ciudad, estado, cp, fechaNacimiento) VALUES(?,?,?,?,?,?,?)",
		usuario.Nombre,
		usuario.Apellidos,
		usuario.Direccion,
		usuario.Ciudad,
		usuario.Estado,
		usuario.Cp,
		usuario.FechaNacimiento)
	return context.JSON(fiber.Map{
		"status": fiber.StatusOK,
	})
}
