package main 
import (
	"fmt"
	"math/rand"
	"os"
)

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
	Sex string;

}

// Global variable section
var racemod = map[string][]int{
	"Dwarf"    : { 0, 0, 1, 0, 0,-1},
	"Elf"      : { 0, 1,-1, 0, 0, 0},
	"Gnome"    : { 0, 0, 0, 0, 0, 0},
	"Half-Elf" : { 0, 0, 0, 0, 0, 0},
	"Halfling" : {-1, 1, 0, 0, 0, 0},
	"Half-Orc" : { 1, 0, 1, 0, 0,-2},
	"Human"    : { 0, 0, 0, 0, 0, 0},
}

//start of main !!!!!!!
func main(){
	fmt.Println("Hello, choose a stat rolling method by typing 1-4 and then hit enter>");
	var method int;
	var myCharacter characterSheet;
	var sorted bool;
	var stats []int;
	for true{
		_, err := fmt.Scanf("%d\n", &method);
		if(err != nil || method <1 || method >4){
			fmt.Fprintf(os.Stderr,"invaild input! please enter a number between 1-4.\n");
		}else{
			break
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

	myCharacter = raceselect(myCharacter,stats);
	myCharacter = pickClass(myCharacter);

	// later stuff to write
	fmt.Printf("Here's your character %#v\n", myCharacter);
}

//takes unsorted stats and has the player pick the order
func statpicker( stats []int) []int{
	var sorted [6]int
	used := [6]bool{false,false,false,false,false,false}
	abilities := [6]string{"Strength", "Dexterity", "Constitution", "Intelligence", "Wisdom", "Charisma"};
	for abl := 0; abl < 6; abl++{
		fmt.Printf("Choose a roll for %v by typing it.\nAvalible rolls:", abilities[abl])
		for i, roll := range stats{
			if used[i] == false{
				fmt.Printf(" %v,", roll);
			}
		}
		fmt.Printf("\n");
		var choice int;
		for loop := true; loop == true;{
			_, err := fmt.Scanf("%d\n", &choice);
			if err != nil {
				fmt.Println("Invalid input! Please choose one of the avalible rolls.");
			}else{
				 matched := false;
				 for index, roll := range stats{
					if choice == roll && used[index] == false{
						sorted[index] = roll;
						used[index] = true;
						matched = true;
						break;
					}
				}
				if matched == true{
					break;
				}else{
					fmt.Println("Invalid input! Please used a avalible roll.");
				}
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
			i+1, rmx[i][0], rmx[i][1], rmx[i][2], rmx[i][3], rmx[i][4], rmx[i][5]);
	}
	var selection int;
	for 1 == 1{
		_, err := fmt.Scanf("%d", &selection);
		if err == nil && selection > 0 && selection < 13{
			break
		}else{
			fmt.Fprintf(os.Stderr, "Invalid input! Please enter a number bwtween 1-12.\n");
		}
	}
	return rmx[selection-1][:], true;
}

//used by the rolling methods 1 & 2
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

func checkrace(stats []int)(map[string]bool, map[string]bool){
	var racemax = map[string][]int{
		"Dwarf"    : {18,17,19,18,18,16},
		"Elf"      : {18,19,18,18,18,18},
		"Gnome"    : {18,18,18,18,18,18},
		"Half-Elf" : {18,18,18,18,18,18,},
		"Halfling" : {17,19,18,18,17,18},
		"Half-Orc" : {19,17,19,17,14,12},
		"Human"    : {18,18,18,18,18,18},
	}
	var racemin = map[string][]int{
		"Dwarf"    : { 8, 3,12, 3, 3, 3},
		"Elf"      : { 3, 7, 6, 8, 3, 8},
		"Gnome"    : { 6, 3, 8, 7, 3, 3},
		"Half-Elf" : { 3, 6, 6, 4, 3, 3},
		"Halfling" : { 6, 8,10, 6, 3, 3},
		"Half-Orc" : { 6, 3,13, 3, 3, 3},
		"Human"    : { 3, 3, 3, 3, 3, 3},
	}
	var fmaxstr = map[string]int{
		"Dwarf"    : 17,
		"Elf"      : 16,
		"Gnome"    : 15,
		"Half-Elf" : 17,
		"Halfling" : 14,
		"Half-Orc" : 18,
		"Human"    : 18,
	}
	maleraces := make(map[string]bool);
	femraces := make(map[string]bool);
	for key := range racemax{
		legit := true;
		for i:=0; i<6; i++{
			if (stats[i] + racemod[key][i])  > racemax[key][i] || (stats[i] + racemod[key][i]) < racemin[key][i]{
				legit = false;
			}
		}
		if legit == true{
			maleraces[key] = true;
		}else{ maleraces[key] = false;}
		if maleraces[key] == true && stats[0]<fmaxstr[key]{
			femraces[key] = true;
		}else{ femraces[key] = false;}
	}

	return maleraces, femraces;
}

// Race and Sex selector
func raceselect(player characterSheet, stats []int) characterSheet{
	maleraces, femraces := checkrace(stats);
	//split print args
	fmt.Printf("Your stats are: Str %v, Dex %v, Con %v, Int %v, Wis %v, Char %v.\n", 
		stats[0], stats[1], stats[2], stats[3], stats[4], stats[5]);
	fmt.Printf("Here are your valid race for male characters:");
	for key, value := range  maleraces{
		if value == true{
			fmt.Printf(" %v;", key);
		}
	}
	fmt.Printf("\nAnd valid races for female characters:");
	for key, value := range femraces{
		if value == true{
			fmt.Printf(" %v;", key);
		}
	}
	fmt.Printf("\nPease Choose either 'Male' or 'Female' by typing and entering it.\n");
	var choice string;
	for true{
		_, err := fmt.Scanf("%s\n", &choice);
		if ((choice != "Male" && choice != "Female") || err != nil){
			fmt.Fprintf(os.Stderr, "Invalid input! Please enter either 'Male' or 'Female'.\n");
		}else{
			break
		}
	}
	player.Sex = choice;
	switch player.Sex{
	case "Male":
		fmt.Println("Choose one of the valid male race options");
		fmt.Printf("Here are your valid race for male characters:");
		for key, value := range  maleraces{
			if value == true{
				fmt.Printf(" %v;", key);
			}
		}
		fmt.Printf("\n");
		var choice string;
		for true {
			fmt.Scanf("%s\n", &choice);
			if maleraces[choice] != true{
				fmt.Fprintf(os.Stderr, "Invalid input! Please choose a valid option.\n");
			}else{
				player.Race = choice;
				break
			}
		}
		break
	case "Female":
		fmt.Println("Choose one of the valid female race options");
		fmt.Printf("Here are your valid race for female characters:");
		for key, value := range  femraces{
			if value == true{
				fmt.Printf(" %v;", key);
			}
		}
		fmt.Printf("\n");
		var choice string;
		for true{
			fmt.Scanf("%s\n", &choice);
			if femraces[choice] != true{
				fmt.Fprintf(os.Stderr, "Invalid input! Please choose a valid option.\n");
			}else{
				player.Race = choice;
				break
			}
		}
		break
	}
	player.Str  = stats[0] + racemod[player.Race][0];
	player.Dex  = stats[1] + racemod[player.Race][1];
	player.Con  = stats[2] + racemod[player.Race][2];
	player.Int  = stats[3] + racemod[player.Race][3];
	player.Wis  = stats[4] + racemod[player.Race][4];
	player.Char = stats[5] + racemod[player.Race][5];

	return player;
}

func checkclass(player characterSheet)[]string{
	var classes = [10]string{"Cleric","Druid","Fighter","Paladin","Ranger","Magic-User","Illusionist","Theif","Assassin","Monk"};
	var classesPerRace = map[string][]bool{
		"Dwarf"    :{false,false, true,false,false,false,false, true, true,false},
		"Elf"      :{false,false, true,false,false, true,false, true, true,false},
		"Gnome"    :{false,false, true,false,false,false, true, true, true,false},
		"Half-Elf" :{ true, true, true,false, true, true,false, true, true,false},
		"Halfling" :{false,false, true,false,false,false,false, true,false,false},
		"Half-Orc" :{ true,false, true,false,false,false,false, true, true,false},
		"Human"    :{ true, true, true, true, true, true, true, true, true, true},
	}
	minStr  := [10]int{ 0, 0, 9,12,13, 0, 0, 0,12,15};
	minDex  := [10]int{ 0, 0, 0, 0, 0, 6,16, 9,12,15};
	minCon  := [10]int{ 0, 0, 7, 9,14, 0, 0, 0, 0,11};
	minInt  := [10]int{ 0, 0, 0, 9,13, 9,15, 0,11, 0};
	minWis  := [10]int{ 9,12, 0,13,14, 0, 0, 0, 0,15};
	minChar := [10]int{ 0,15, 0,17, 0, 0, 0, 0, 0, 0};
	var validclasses [10]string;
	var numvalidclasses int;
	numvalidclasses =0;
	for i := 0; i<10; i++{
		if (classesPerRace[player.Race][i] == true && player.Str >= minStr[i] && player.Dex >= minDex[i] &&
		player.Con >= minCon[i] && player.Int >= minInt[i] && player.Wis >= minWis[i] && player.Char >= minChar[i]){
			validclasses[numvalidclasses] = classes[i];
			numvalidclasses++
		}
	}
	return validclasses[0:numvalidclasses];
}

func pickClass(player characterSheet) characterSheet{
	validclasses := checkclass(player);
	fmt.Printf("Choose a valid class by typing its name and pressing 'enter'.\nValid classes:");
	for class := 0; class< len(validclasses); class++{
		fmt.Printf(" %v;", validclasses[class]);
	}
	fmt.Printf("\n");
	var choice string;
	for true{
		_, err := fmt.Scanf("%s\n", &choice);
		if err != nil{
			fmt.Fprintln(os.Stderr, "Invalid input! Please type a class name.");
		}else{
			match := false
			for _, class := range validclasses {
				if choice == class{
					match = true 
				}
			}
			if match == true{
				break;
			}else{
				fmt.Fprintln(os.Stderr, "Invalid input! Please type a class name.");
			}
		}
	}
	player.Class = choice;
	return player;
}
