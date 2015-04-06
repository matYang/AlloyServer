open course
open transcript

// plan is implicit in file generation
// abstract sig Plan {}
// one sig SE, SECogSci extends Plan {}

//Requests a result
one sig Result {
    rawresult : set Course
}
fun result : set Course { Result.rawresult }

//You need to set a upperbound as well as a lower bound for your results
pred SEPlan {
	upperbound
	//Here is list of minimum requirments 
	-- all core courses (or substitutions)
    allOf[CoreCourse]
	-- other courses
    oneOf[ATECourseCS]
    oneOf[ATECourseECE]
    fourOf[ATECourseCS + ATECourseECE + ATECourseMore + ATECourseExtra]
    // this is too strict: a student can have two extra courses,
    // but they just cannot count both as ATEs.
	// in that case, the extra course will be excluded from result
    no disjoint x, y : ATECourseExtra & result | x != y
    twoOf[ScienceCourse]
    -- Linkage Electives
    oneOf[CommunicationCourse]
    oneOf[CS492 + ComplementaryCourseA]
    twoOf[ComplementaryCourseC]
    -- Substitutions
	//all c : Core.courses | some c' : c + (Core.substitutions).c | {
	//	(Core->c') in result and c' in transcript.Int} // and (c' != c => (c->c' in rsubst))}
}

pred allOf[ s : set Course ] {
    s in result
}

pred oneOf[ s : set Course ] {
	// some s & result
	// introduce quantifier/variable so we know what it is
    some c : s | c in result
}

pred twoOf[ s : set Course ] {
    some disjoint x, y : s & result | x != y
}

pred threeOf[ s : set Course ] {
    some disjoint x, y, z : s & result | x != y and x != z
}

pred fourOf[ s : set Course ] {
    some disjoint w, x, y, z : s & result | x != y and x != z and x != w
}



pred upperbound {
	-- result constrained by transcript
	result in transcript
	-- result constrained by list definitions + substitutions
	-- leftovers are results that are not on lists, and so must be covered by substitutions
	//let leftovers = result-courses | leftovers in substitutions.Course
}

run SEPlan
