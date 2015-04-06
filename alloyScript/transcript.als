open course

//Hard coded transcript of a student
fun transcript [] : set Course {
    CoreCourse +
    CS492 +
    ENGL109 +
    CS360 +
	ECE327 +
    CO487 +
    ECE409 +
    CS442 +
    ECON102 +
    ECON202 +
    SCI238 +
    SCI215 
}

//Making an exception for course substitution ECE100 for SE101
/*
fun substitutions [] : List -> Course -> Course {
	Core -> ECE100 -> SE101
}
*/
