package main 
import (
	"fmt"
	"math/rand"
)

func main(){
	fmt.Println("Hello, choose a stat rolling method by typing 1-4 and then hit enter>");
	var method int;
	var myCharacter characterSheet;
	var sorted bool;
	var stats []int;
	fmt.Scan(&method);
	switch method {
	case 1 :
		stats, sorted = rollMethod1();
		break

	case 2 :
		method2();
	case 3 :
		method3();
	case 4 :
		method4();
	}
	//check for and run statpicker 
	 if sorted == false{
		 stats = statpicker(stats);
	 }

	 //writting stats to the myCharacter struc
	 myCharacter.Str = stats[0];
	 myCharacter.Dex = stats[1];
	 myCharacter.Con = stats[2];
	 myCharacter.Int = stats[3];
	 myCharacter.Wis = stats[4];
	 myCharacter.Char = stats[5];

	// later stuff to write
	fmt.Println("Choose your Class by typing it");
	fmt.Scan(&myCharacter.Class);
	fmt.Println("Choose your Race by typing it");
	fmt.Scan(&myCharacter.Race);

	fmt.Printf("Here's your character %v\n", myCharacter);

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

//taks unsorted stats and has the player pick the order
func statpicker( stats []int) []int{
	var sorted [6]int
	used := [6]bool{false,false,false,false,false,false}
	abilities := [6]string{"Str", "Dex", "Con", "Int", "Wis", "Char"};
	for i := 0; i < 6; i++{
		var index int;
		fmt.Printf("choose a stat for %s\n", abilities[i]);
		for loop := true; loop == true;{
			fmt.Printf("your stats %d\n", stats);
			fmt.Scan(&index);
			if used[index] == false{
				used[index] = true;
				sorted[i] = stats[index];
				loop = false
			}else{
				fmt.Println("You already used that one. Choose another.");
			}
		}
	}
	return sorted[:];
}


func rollMethod1()([]int, bool){
	fmt.Println("Method 1");
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

func method2(){
	fmt.Println("Method 2");
}

func method3(){
	fmt.Println("Method 3");
}

func method4(){
	fmt.Println("Method 4");
}

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


