
type Difficulty string

import (
	"fmt"
	"strconv"
)

const (
	easy Difficulty "easy"
	normal Difficulty "normal" 
	hard Difficulty "hard"
	boss Difficulty "boss"
)

type Player struct {
	Name string
	Level int
	HP int
	Damage int
	Defense int
}

type Enemy struct {
	Difficulty Difficulty
	Level int
	HP int
	Damage int
	Defense int
	Stage Stage
}

type Stage struct {
	minLevel int
	maxLevel int
}

type Stages struct {
	stages []Stage
}

type Level struct {
	KillCondition int
	isBossLevel bool
}

type Levels struct {
	levels []Level
}


func createNewPlayer(name string) Player {
	newPLayer := Player{
		Name: name,
		Level: 1,
		HP: 10,
		Damage: 5,
		Defense: 5
	}
	return newPLayer;
}

func createNewEnemy(name string, level int, hp int, damage int, defense int) Enemy {
	newEnemy := Enemy{
		Name: name,
		Level: level,
		HP: hp,
		Damage: damage,
		Defense: defense
	}
	return newPLayer;
}

func createNewLevel(isBoss bool, killCondition int) {
	newLevel := Level{
		KillCondition: killCondition,
		isBossLevel: isBoss
	}
	return newLevel
} 

func createNewStage(min int, max int) Stage{
	newStage := Stage{
		minLevel: min,
		maxLevel: max
	}
	return newStage
}


func fightResult(player Player, enemy Enemy) bool {
	var playerAttack int = player.Attack * randomNum(1, 10)
	var enemyAttack int = enemy.Attack * randomNum(1, 10)
	var isEnemyDead bool = false
	for player.HP <= 0 || enemy.HP <= 0 {
		enemy.HP -= playerAttack - enemy.Defense
		fmt.Printf("Atacas con %d de poder!", playerAttack )
		if(enemy.HP <= 0) {
			isEnemyDead = true
			fmt.Println("El enemigo murio!!!")
		}
		player.HP -= enemyAttack - player.defense
	}
	return isEnemyDead	
}

func getPlayerName() string {
	var data string
	fmt.Scanln(&data)
	return data
}

func getEnemyName() string {
	names := {"hongo", "perro", "lobo", "oso", "dragon", "satan"}
	return names[randomNum(0, len(names) - 1)]
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

func randomNum(min int, max int) int {
	return rand.Intn(min, max)
}

func main() 
	fmt.Println("Ingrese 1 para iniciar el juego")
	var input string = getInput()
	if(input == 1) {
		fmt.Println("Ingrese su nombre")
		var isBossDead bool = false
		var cantLevels int = 1
		var name string = getPlayerName()
		var player Player = createNewPlayer(name)
		for !isBossDead {
			var min int = 1 * cantLevels
			var max int = 2 * cantLevels
			var isBossLevel bool = (cantLevels == 4)
			var level Level = createNewLevel(isBossLevel, randomNum(min, max))
			var stage Stage = createNewStage(min, max)
			for i := 0; i < level.KillCondition {
				var enemy Enemy = createNewEnemy(getEnemyName(), randomNum(min,max), randomNum(min, max), randomNum(min, max), randomNum(min, max))
				fmt.Printf("Ha aparecido el enemigo %s que tiene %d cantidad de vida", enemy.Name, enemy.HP)
				if(fightResult(player, enemy)) {
					i++
				}

			} 
			cantLevels++
		}
	}

}