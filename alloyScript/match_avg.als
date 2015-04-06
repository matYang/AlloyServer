open course
open list
open transcript

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
