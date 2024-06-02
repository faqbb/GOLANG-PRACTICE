package main

import (
	"fmt"
	"strconv"
)

type User struct {
	Usuario    string
	Contraseña string
}

func getInput() int {
	var input string
	fmt.Scanln(&input)
	command, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Ingrese un numero valido")
		return 0
	}
	return command
}

func getUserData() string {
	var data string
	fmt.Scanln(&data)
	return data
}

func login() User {
	fmt.Println("Ingrese usuario")
	var username string = getUserData()
	fmt.Println("Ingrese contraseña")
	var password string = getUserData()
	user := User{
		Usuario:    username,
		Contraseña: password,
	}
	return user
}

func adminMenu(entries []string, changes []int, laborers []string) int {
	fmt.Println("Ingrese 1 para agregar un laborer, 2 para eliminar un laborer, 3 para volver a iniciar sesion y 4 para cerrar el programa")
	var input int = getInput()
	switch input {
	case 1:
		laborers = append(laborers, "laborer")
		entries = append(entries, "laborer agregado")
		fmt.Println("Laborer creado")
		changes[0] += 1
		return 1
	case 2:
		laborers = laborers[:len(laborers)-1]
		entries = append(entries, "laborer eliminado")
		fmt.Println("Laborer eliminado")
		changes[1] += 1
		return 2
	case 3:
		return 10
	case 4:
		return 100
	default:
		fmt.Println("Entrada inválida")
		return 0
	}
}

func supervisorMenu(changes []int) int {
	fmt.Println("Ingrese 1 para crear registro de administrador, 2 para volver a iniciar sesion y 3 para cerrar el programa")
	var input int = getInput()
	switch input {
	case 1:
		log := fmt.Sprintf("Laborers agregados: %d Laborers eliminados: %d", changes[0], changes[1])
		fmt.Println(log)
		return 1
	case 2:
		return 10
	case 3:
		return 100
	default:
		fmt.Println("Entrada inválida")
		return 0
	}
}

func main() {
	admin := User{
		Usuario:    "admin",
		Contraseña: "root",
	}
	supervisor := User{
		Usuario:    "seeker223",
		Contraseña: "seekr",
	}
	laborers := []string{"laborer1", "laborer2", "laborer3", "laborer4"}
	adminEntries := []string{}
	adminChanges := []int{0, 0}
	var mainSwitch int = 1

	for mainSwitch != 0 {
		fmt.Println("Presione 1 para iniciar sesion o 0 para cerrar el programa")
		var off int = getInput()
		if off == 1 {
			fmt.Println("Inicie sesion")
			var user User = login()
			if user.Usuario == admin.Usuario && user.Contraseña == admin.Contraseña {
				var token int = 1
				for token != 0 {
					var menu int = adminMenu(adminEntries, adminChanges, laborers)
					switch menu {
					case 10:
						token = 0
					case 100:
						mainSwitch = 0
						token = 0
					}
				}
			} else if user.Usuario == supervisor.Usuario && user.Contraseña == user.Contraseña {
				var token int = 1
				for token != 0 {
					var menu int = supervisorMenu(adminChanges)
					switch menu {
					case 10:
						token = 0
					case 100:
						mainSwitch = 0
						token = 0
					}
				}
			} else {
				fmt.Println("Usuario o contraseña erroneos")
			}
		} else {
			mainSwitch = 0
		}
	}
}
