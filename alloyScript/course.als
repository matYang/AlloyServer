//List of all courses that we concern
abstract sig Course {}

abstract sig 
    CoreCourse, 
    ATECourseECE, 
    ATECourseCS, 
    ATECourseMore, 
    ATECourseExtra, 
    LinkageCourse, 
    ScienceCourse,
    CommunicationCourse, 
    ComplementaryCourseA,
    //ComplementaryCourseB,
    ComplementaryCourseC
    //ComplementaryCourseD,
    extends Course {}

//core:
one sig
//1A
CS137,
ECE105,
ECE140,
MATH115,
MATH117,
SE101, 
//1B
CS138,
ECE106,
ECE124,
MATH119,
MATH135,
//2A
CHE102,
CS241,
ECE222,
SE212,
STAT206,
//2B
CS240,
CS247,
MSCI261,
MATH213,
MATH239,
//3A
CS341,
CS349,
SE350,
SE380,
//3B
CS343,
CS348,
SE390,
SE463,
//4A
ECE358,
SE464,
SE490,
//4B
SE465,
SE491 
extends CoreCourse {}

//ATE
//CS
one sig 
CS442,
CS444,
CS448,
CS449,
CS450,
CS452,
CS454,
CS457,
CS458,
CS473,
CS484,
CS485,
CS486,
CS488
extends ATECourseCS {}

//ECE
one sig
ECE409,
ECE416,
ECE417,
ECE418,
ECE419,
ECE423,
ECE429,
ECE454,
ECE455,
ECE457A,
ECE457B,
ECE458,
ECE459,
ECE481,
ECE486,
ECE488
extends ATECourseECE {}

//More
one sig
CO487,
CS462,
CS466,
CS467,
CS475,
CS487,
CS489,
CS490,
ECE406,
SE498,
SE499,
SYDE433,
SYDE475,
SYDE558
extends ATECourseMore {}

//Extra
one sig
CS360,
CS365,
CS370,
CS371,
ECE207,
ECE224,
ECE327,
SYDE348 
extends ATECourseExtra {}


/*
Linkages
*/
//Social Issue
one sig CS492 extends LinkageCourse {}

//Science
one sig 
BIOL130, BIOL130L,
BIOL239, 
BIOL240, 
BIOL273, 
CHE161, 
CHEM262, CHE262L,
EARTH121, EARTH121L,
PHYS234,
PHYS246, 
PHYS275, 
PHYS334, 
PHYS375,
ECE209,
SCI238,
SCI215
extends ScienceCourse {}

//Communications
one sig 
ENGL109,
ENGL119, 
ENGL140R, 
ENGL210F, 
ENGL210G, 
SPCOM100, 
SPCOM223, 
SPCOM324, 
SPCOM432 
extends CommunicationCourse {}

//Complementary Studies List A
one sig 
ECE390,
ENVS105,
ERS215,
ERS315,
GENE22A,
GEOG203,
GEOG368,
MSCI422,
MSCI442,
PHIL226,
SOC232,
STV100,
STV202,
STV203,
STV205,
STV210,
STV302,
STV404,
SYDE261,
WS205
extends ComplementaryCourseA {}


//List C
one sig
ECON102,
ECON202,
MSCI211,
MSCI311,
PSCI260,
PSYCH101,
SOC101,
HIST253,
HIST254,
HIST275,
PHIL250A,
PHIL250B,
PHIL315 //(GENE412) - (cross-listed)
extends ComplementaryCourseC {}
