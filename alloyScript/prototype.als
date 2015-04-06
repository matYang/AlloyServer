//List of all courses that exist in this alloy prototype world
abstract sig Course {}
one sig SE101, ECE100, ECE459, CS444, PSYCH101 extends Course {}

//List of requirements to graduate
abstract sig List {}
one sig Core, ATE, Linkage extends List {}

//List of courses needed to full fill sig list above
//List of courses also have to be in sig Course above
fun courses [] : List -> Course {
	Core -> SE101 +
	ATE -> ECE459 +
	ATE -> CS444 +
	Linkage -> PSYCH101
}

// don't need to name student: implicit in generation
// abstract sig Student {}
// one sig Golson extends Student {}

//Hard coded transcript of a student with theoretical grades
//Alloy has a hard time with large numbers so grades on scale
//0 - 9
fun transcript [] : Course -> Int {
	ECE100 -> 9 +
	ECE459 -> 8 +
	CS444 -> 8 +
	PSYCH101 -> 5
}

//Making an exception for course substitution ECE100 for SE101
fun substitutions [] : List -> Course -> Course {
	Core -> ECE100 -> SE101
}

// plan is implicit in file generation
// abstract sig Plan {}
// one sig SE, SECogSci extends Plan {}

//Requests a result
one sig Result {
	rawresult : List -> Course,
}

fun result [] : List -> Course { Result.rawresult }
-- courses counting towards the major
fun major [] : Course { Core.result + ATE.result }
-- courses that fill some requirement but do not counting towards the major
fun nonmajor  [] : Course { List.result - major }

//You need to set a upperbound as well as a lower bound for your results
pred SE2011 {
	upperbound
	//Here is list of minimum requirments 
	-- all core courses (or substitutions)
	all c : Core.courses | some c' : c + (Core.substitutions).c | {
		(Core->c') in result and c' in transcript.Int} // and (c' != c => (c->c' in rsubst))}
	-- two ATEs
	#(ATE.result) = 2
	-- one Linkage
	#(Linkage.result) = 1
	-- average >= 60% for major courses
	-- sum major courses > 60 * #major
	-- N.B.: really 59.5 * #major
	mul[6, #major] <= (sum c : major | c.transcript)
	-- average >= 50% for non-major courses
	mul[5, #nonmajor] <= (sum c : nonmajor | c.transcript)
}
pred upperbound {
	-- result constrained by transcript
	List.result in transcript.Int	
	-- result constrained by list definitions + substitutions
	-- leftovers are results that are not on lists, and so must be covered by substitutions
	let leftovers = result-courses | leftovers in substitutions.Course
}

run SE2011 for 3 but 7 Int
