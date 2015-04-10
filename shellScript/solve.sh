#!/bin/bash

SOURCEJSON="transcript.json"
INTERMEDIATEFILE="transcript.als"
TARGETPARSER="transcript_parser.py"

TARGETALS="match.als"
JARNAME="alloy4.2.jar "
MAINCLASS="edu.mit.csail.sdg.alloy4whole.ExampleUsingTheCompiler"

#parse json, if fail then do not run alloy
python $TARGETPARSER --from_json_file  $SOURCEJSON
if [[ $? -ne 0 ]] ; then
    exit 1
fi
java -cp $JARNAME $MAINCLASS $TARGETALS
