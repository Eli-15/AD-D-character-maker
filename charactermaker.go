package main 
import (
	"fmt"
	"math/rand"
)

// Global variable section
// Race max scoures
var racemax = map[string][]int{
	"Dwarf" : {18,17,19,18,18,16},
	"Elf" : {18,19,18,18,18,18},
	"Gnome" : {18,18,18,18,18,18},
	"Half-Elf" : {18,18,18,18,18,18,},
	"Halfling" : {17,19,18,18,17,18},
	"Half-Orc" : {19,17,19,17,14,12},
	"Human" : {18,18,18,18,18,18},
}

var racemin = map[string][]int{
	"Dwarf" : {8,3,12,3,3,3},
	"Elf" : {3,7,6,8,3,8},
	"Gnome" : {6,3,8,7,3,3},
	"Half-Elf" : {3,6,6,4,3,3},
	"Halfling" : {6,8,10,6,3,3},
	"Half-Orc" : {6,3,13,3,3,3},
	"Human" : {3,3,3,3,3,3},
}

/*var fmaxstr = map[string]int{
	"Dwarf" : 17,
	"Elf" : 16,
	"Gnome" : 15,
	"Half-Elf" : 17,
	"Halfling" : 14,
	"Half-Orc" : 18,
	"Human" : 18,
}*/

//start of main !!!!!!!
func main(){
	fmt.Println("Hello, choose a stat rolling method by typing 1-4 and then hit enter>");
	var method int;
	var myCharacter characterSheet;
	var sorted bool;
	var stats []int;
	for loop := true; loop == true;{
		_, err := fmt.Scanf("%d\n", &method);
		if(err != nil || method <1 || method >4){
			fmt.Println("invaild input! please enter a number between 1-4.");
			loop = true;
		}else{
			loop = false;
		}
	}
	switch method {
	case 1 :
		stats, sorted = rollMethod1();
		break

	case 2 :
		stats, sorted = rollMethod2();
		break
	case 3 :
		stats, sorted = rollMethod3();
		break
	case 4 :
		stats, sorted = rollMethod4();
		break
	}
	//check for and run statpicker 
	 if sorted == false{
		 stats = statpicker(stats);
	 }

	 //writting stats to the myCharacter struc
	 myCharacter.Str  = stats[0];
	 myCharacter.Dex  = stats[1];
	 myCharacter.Con  = stats[2];
	 myCharacter.Int  = stats[3];
	 myCharacter.Wis  = stats[4];
	 myCharacter.Char = stats[5];

	validraces := checkrace(stats);
	fmt.Printf("%#v\n", validraces);

	// later stuff to write
	fmt.Println("Choose your Class by typing it");
	fmt.Scanf("%s", &myCharacter.Class);
	fmt.Println("Choose your Race by typing it");
	fmt.Scanf("%s", &myCharacter.Race);

	fmt.Printf("Here's your character %#v\n", myCharacter);

}

type characterSheet struct{
	// Ability Scores
	Str int;
	Dex int;
	Con int;
	Int int;
	Wis int;
	Char int;

	// Race and Class
	Class string;
	Race string;

}

//takes unsorted stats and has the player pick the order
func statpicker( stats []int) []int{
	var sorted [6]int
	used := [6]bool{false,false,false,false,false,false}
	abilities := [6]string{"Str", "Dex", "Con", "Int", "Wis", "Char"};
	for i := 0; i < 6; i++{
		var index int;
		fmt.Printf("choose a stat for %s\n", abilities[i]);
		for loop := true; loop == true;{
			fmt.Printf("your stats %d\n", stats);
			_, err := fmt.Scanf("%d\n",&index);
			if err == nil && index<7 && index>0{
				if used[index] == false{
					used[index] = true;
					sorted[i] = stats[index];
					loop = false
				}else{
				fmt.Println("You already used that one. Choose another.");
				}
			}else{
				fmt.Println("Invalid input. Please enter a number between 1-6.");
			}
		}
	}
	return sorted[:];
}

func rollMethod1()([]int, bool){
	var rolls [6]int;
	for i:=0; i<6; i++ {
		var roll [4]int;
		for j:=0; j<4; j++ {
			roll[j] = rand.Intn(5) +1;
		}
		sort(roll[:])
		sum := roll[1] + roll[2] +roll[3];
		rolls[i] = sum;
	}
	return rolls[:], false;
}

func rollMethod2()([]int, bool){
	var rolls [12]int;
	for i := 0; i <12; i++{
		rolls[i] = ((rand.Intn(5)+1)+(rand.Intn(5)+1)+(rand.Intn(5)+1));
	}
	sort(rolls[:]);
	return rolls[5:11], false;
}


func rollMethod3()([]int, bool){
	var rolls [6]int;
	for i := 0; i < 6; i++{
		var roller [6]int;
		for j := 0; j < 6; j++{
			roller[j] = ((rand.Intn(5)+1)+(rand.Intn(5)+1)+(rand.Intn(5)+1));
		}
		sort(roller[:]);
		rolls[i] = roller[5];
	}
	return rolls[:], true;
}

func rollMethod4()( []int, bool){
	fmt.Println("Method 4");
	var rmx [12][6]int;
	for i:=0;i<12;i++{
		for j:=0;j<6;j++{
			rmx[i][j] = ((rand.Intn(5)+1)+(rand.Intn(5)+1)+(rand.Intn(5)+1));
		}
	}
	fmt.Println("Choose one of the characters bellow by entering it's number.");
	for i:=0;i<12;i++{
		//the next line's arguments are spread over two lines
		fmt.Printf("character %2d  Str %2d  Dex %2d  Con %2d  Int %2d  Wis %2d  Char %2d\n", 
		i, rmx[i][0], rmx[i][1], rmx[i][2], rmx[i][3], rmx[i][4], rmx[i][5]);
	}
	var selection int;
	fmt.Scanf("%d", &selection);
	return rmx[selection][:], true;
}

//used by the rolling methods
func sort(arr []int) {
	var i int = 1
	for i < len(arr){
		var j int = i
		for j > 0 && arr[j-1] > arr[j]{
			arr[j], arr[j-1] = arr[j-1], arr[j]
			j = j-1
		}
		i = i+1;
	}
}

func checkrace(stats []int)(map[string]bool){
	fmt.Println("checking valid races");
	validrace := make(map[string]bool);

	for key := range racemax{
		legit := true;
		for i:=0; i<6; i++{
			if stats[i] > racemax[key][i] || stats[i] < racemin[key][i]{
				legit = false;
			}
		}
		if legit == true{
			validrace[key] = true;
		}
	}

	return validrace;

}
