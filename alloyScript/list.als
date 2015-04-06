open course

//List of requirements to graduate
abstract sig List {}
one sig Core, ATE, Linkage, SocialIssue, Science, Language, Complementary extends List {}

//List of courses needed to full fill sig list above
//List of courses also have to be in sig Course above
fun courses [] : List -> Course {
	Core -> SE101 +
	ATE -> ECE459 +
	ATE -> CS444 +
	Linkage -> PSYCH101
}

fun lists [] : List -> List{
	Linkage -> (SocialIssue + Science + Language + Complementary)
}
